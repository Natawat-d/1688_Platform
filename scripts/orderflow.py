#!/usr/bin/env python3
"""Drive one order through the whole pipeline against a running api + stub.

    python3 scripts/orderflow.py [offerId ...]

    BASE_URL=https://d111.cloudfront.net ADMIN_TOKEN=... python3 scripts/orderflow.py

Adds one SKU from each given offer (defaults pick two different suppliers), checks
out, pays on our side, waits for the relay, pays the 1688-side parcels through the
stub's control endpoint, runs the simulator fast, then prints the timeline, the
tracking events and the message log. Exit code is non-zero if the order never
reaches the warehouse, so this doubles as an end-to-end check.

Environment:
    BASE_URL         where the api is (default http://localhost:8787)
    ADMIN_TOKEN      bearer token for the admin routes (default dev)
    CONTROL_TOKEN    forwarded to the stub's control plane when it is guarded
"""
import http.cookiejar
import json
import os
import sys
import time
import urllib.error
import urllib.request

API = os.environ.get("BASE_URL", "http://localhost:8787").rstrip("/")
ADMIN = {
    "Authorization": "Bearer " + os.environ.get("ADMIN_TOKEN", "dev"),
    "Content-Type": "application/json",
}
if os.environ.get("CONTROL_TOKEN"):
    ADMIN["X-Control-Token"] = os.environ["CONTROL_TOKEN"]

jar = http.cookiejar.CookieJar()
opener = urllib.request.build_opener(urllib.request.HTTPCookieProcessor(jar))


def call(method, path, body=None, headers=None):
    data = json.dumps(body).encode() if body is not None else None
    req = urllib.request.Request(API + path, data=data, method=method)
    req.add_header("Content-Type", "application/json")
    for k, v in (headers or {}).items():
        req.add_header(k, v)
    try:
        with opener.open(req, timeout=30) as r:
            return r.status, json.loads(r.read() or b"{}")
    except urllib.error.HTTPError as e:
        raw = e.read() or b"{}"
        try:
            return e.code, json.loads(raw)
        except ValueError:
            return e.code, {"error": raw[:200].decode("utf-8", "replace")}


def main():
    offers = sys.argv[1:] or ["900000000000", "900000000175"]
    print(f"== target {API} ==")

    print("== cart ==")
    cart = None
    for oid in offers:
        code, p = call("GET", f"/api/products/{oid}")
        if code != 200:
            print(f"  ! product {oid}: {code} {p.get('message') or p.get('error')}")
            sys.exit(2)
        sku = next(s for s in p["skus"] if s["available"])
        qty = max(p["moq"], 1)
        _, cart = call("POST", "/api/cart/items", {"offerId": oid, "skuId": sku["skuId"], "quantity": qty})
        print(f"  + {p['title'][:48]:48} seller={p['seller']['name'][:22]:22} x{qty} {sku['price']['text']}")
    print(f"  groups={len(cart['groups'])} checkoutable={cart['checkoutable']} total={cart['totals']['total']['text']}")
    for g in cart["groups"]:
        for i in g["issues"]:
            print(f"  ! {g['sellerName']}: {i['code']} {i['message']}")

    print("== checkout (1688 preview runs before any charge) ==")
    code, co = call("POST", "/api/checkout", {
        "email": "buyer@example.com", "name": "Test Buyer", "phone": "0812345678",
        "address": {"line1": "1 Sukhumvit Rd", "city": "Bangkok", "province": "Bangkok", "postcode": "10110"},
    })
    if code != 200:
        print(f"  refused {code}: {co.get('message')}")
        for i in co.get("issues", []):
            print(f"    {i['code']} on {i.get('field')}: {i['message']}")
        sys.exit(2)
    oid, tok = co["orderId"], co["token"]
    print(f"  order {oid} parcels={co['parcels']} total={co['total']['text']}")

    print("== pay on our side ==")
    _, o = call("POST", f"/api/orders/{oid}/pay?t={tok}")
    print(f"  {o['status']} / {o['statusLabel']}")

    def show(o):
        for p in o["parcels"]:
            print(f"  parcel {p['sellerName'][:26]:26} {p['status']:18} cbu={p.get('cbuOrderId') or '-':20} "
                  f"track={p.get('trackingNo') or '-':14} events={len(p['events'])}")

    print("== relay ==")
    for _ in range(15):
        time.sleep(2)
        _, o = call("GET", f"/api/orders/{oid}?t={tok}")
        if all(p.get("cbuOrderId") for p in o["parcels"]):
            break
    show(o)

    print("== pay the 1688 parcels on the stub cashier, run the simulator at 60x ==")
    for p in o["parcels"]:
        if p.get("cbuOrderId"):
            code, r = call("POST", f"/api/admin/stub/orders/{p['cbuOrderId']}/advance", {"to": "waitsellersend"}, ADMIN)
            print(f"  cashier {p['cbuOrderId']}: {code} {json.dumps(r)[:80]}")
    call("POST", "/api/admin/stub/clock", {"running": True, "speedX": 60}, ADMIN)

    final = None
    for _ in range(40):
        time.sleep(3)
        _, o = call("GET", f"/api/orders/{oid}?t={tok}")
        if o["status"] in ("ARRIVED_WAREHOUSE", "CANCELLED", "RELAY_FAILED"):
            final = o
            break
    o = final or o

    print(f"== final: {o['status']} / {o['statusLabel']} ==")
    for s in o["timeline"]:
        print(f"  [{'x' if s['done'] else ' '}] {s['label']} {s.get('detail', '')}")
    show(o)
    for p in o["parcels"]:
        for e in p["events"][:8]:
            print(f"      {e['at'][:19]} {e['source']:7} {e['text']}")

    print("== message log ==")
    code, msgs = call("GET", "/api/admin/messages?limit=12", headers=ADMIN)
    if code == 200:
        for m in msgs:
            print(f"  {m['type']:46} {'ERR ' + m['error'] if m['error'] else 'ok'}")
    else:
        print(f"  admin log unavailable: {code}")

    print(f"ORDER={oid} TOKEN={tok}")
    sys.exit(0 if o["status"] == "ARRIVED_WAREHOUSE" else 1)


if __name__ == "__main__":
    main()
