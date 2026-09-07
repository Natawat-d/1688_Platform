import json,glob,os,html as H
OUT='/Users/natawatd/Desktop/temp/1688-api-docs'
T={}
for fn in glob.glob('translations/*.json'): T.update(json.load(open(fn,encoding='utf-8')))
src=open('translate.py',encoding='utf-8').read().split('# ---------------------------------------------------------------- coverage check')[0]
ns={}; exec(src,ns); CATS=ns['CATS']; API=ns['API']
def tr(s): s=(s or '').strip(); return T.get(s,s)
e=H.escape
def ex(s): return H.unescape(s or '').strip()
LST=json.load(open('apilist.json',encoding='utf-8'))
def rows(fields,depth,out,arr=False):
    for f in fields or []:
        if '_cycle' in f: out.append((depth,'…','recursive',False,'same structure as the enclosing model','',False)); continue
        typ=f.get('type') or ''
        if typ.startswith('message:'): typ=typ[len('message:'):].split('.')[-1]
        isobj=bool(f.get('fields'))
        out.append((depth,f.get('name') or '',typ,bool(f.get('required')),tr(f.get('description')),ex(f.get('exampleValue')),isobj))
        if isobj: rows(f['fields'],depth+1,out)
def table(fields,cls=''):
    r=[]; rows(fields,0,r)
    if not r: return '<p class="none">No parameters.</p>'
    h=[f'<div class="tw"><table class="{cls}"><thead><tr><th>Field</th><th>Type</th><th>Req</th><th>Description</th><th>Example</th></tr></thead><tbody>']
    for depth,name,typ,req,desc,exv,isobj in r:
        h.append(f'<tr class="d{min(depth,6)}{" obj" if isobj else ""}"><td class="f"><span class="tree"></span><code>{e(name)}</code></td><td class="t">{e(typ)}</td><td class="r">{"<b>yes</b>" if req else "no"}</td><td class="desc">{e(desc)}</td><td class="ex">{e(exv)}</td></tr>')
    h.append('</tbody></table></div>'); return ''.join(h)
nav=[]; body=[]; n=0
for cat in LST['apis']:
    cn=cat['categoryFamilyName']; en=CATS.get(cn,cn); cid='c-'+''.join(ch if ch.isalnum() else '-' for ch in en.lower())
    nav.append(f'<div class="grp"><div class="gname" data-cat="{cid}">{e(en)} <span class="cn">{e(cn)}</span></div>')
    body.append(f'<h2 class="cat" id="{cid}">{e(en)} <span class="cn">{e(cn)}</span> <span class="n">{len(cat["modules"])} APIs</span></h2>')
    for m in cat['modules']:
        n+=1; aid='a-'+m['name'].replace('.','-').lower()+'-'+str(m['version']); name_en=API[m['fullName']][0]
        nav.append(f'<a href="#{aid}" data-cat="{cid}">{e(name_en)}</a>')
        fn=f"{OUT}/raw/{m['namespace']}.{m['name']}-{m['version']}.json"
        head=f'<section class="api" id="{aid}" data-cat="{cid}"><header><div class="num">{n}</div><div class="hd"><h3>{e(name_en)} <span class="cn">{e(m["displayName"].strip())}</span></h3><div class="id"><code>{e(m["fullName"])}</code><button class="copy" data-copy="{e(m["fullName"])}" title="Copy API id">copy</button></div>'
        if not os.path.exists(fn):
            body.append(head+'</div></header><p class="none">Detail doc unavailable on open.1688.com (server answers "error finding API").</p></section>'); continue
        d=json.load(open(fn,encoding='utf-8'))
        url=f"https://gw.open.1688.com/openapi/param2/{d['version']}/{d['namespace']}/{d['name']}/{{appKey}}"
        chips=''.join(f'<span class="chip {"req" if p.get("required") else "opt"}">{e(p["name"])}<i>{"required" if p.get("required") else "optional"}</i></span>' for p in d.get('apiSystemParamVOList') or [])
        chips+=f'<span class="chip flag">{"auth required" if d.get("needAuth") else "no user auth"}</span><span class="chip flag">{"signed" if d.get("needSignature") else "unsigned"}</span>'
        body.append(head+f'<div class="url"><span class="method">POST</span><code>{e(url)}</code><button class="copy" data-copy="{e(url)}" title="Copy URL">copy</button></div><div class="chips">{chips}</div></div></header>'
            f'<p class="desc">{e(API[m["fullName"]][1])}</p>'
            f'<h4>Request body</h4>{table(d.get("apiAppParamVOList"),"req")}'
            f'<details><summary>Response fields</summary>{table(d.get("apiReturnParamVOList"),"resp")}</details>')
        ecs=d.get('apiErrorCodeVOList') or []
        if ecs:
            body.append('<details><summary>Error codes ('+str(len(ecs))+')</summary><div class="tw"><table class="err"><thead><tr><th>Error</th><th>Symptom</th><th>How to fix</th></tr></thead><tbody>'+''.join(f'<tr><td><code>{e(tr(x.get("code")))}</code></td><td>{e(tr(x.get("desc")))}</td><td>{e(tr(x.get("howToFix")))}</td></tr>' for x in ecs)+'</tbody></table></div></details>')
        sm=d.get('apiDocSampleVOList') or []
        if sm:
            body.append('<details><summary>Samples ('+str(len(sm))+')</summary>'+''.join(f'<p class="sn">{e(tr(x.get("name")))}</p><pre>{e(ex(x.get("sample")))}</pre>' for x in sm)+'</details>')
        body.append('</section>')
    nav.append('</div>')
css='''
:root{--bg:#f5f6f8;--sf:#fff;--ink:#1b1f26;--mut:#5f6875;--ln:#dde1e7;--ln2:#c3c9d2;--ac:#d85a00;--acs:#fff1e6;--code:#eef0f3;--ok:#1e7a4a;--oks:#e4f4ea;--sel:#fff7ed}
@media(prefers-color-scheme:dark){:root{--bg:#121519;--sf:#1a1e24;--ink:#e7e9ed;--mut:#98a2af;--ln:#2a303a;--ln2:#3a424e;--ac:#ff8b45;--acs:#2a1e14;--code:#22272f;--ok:#63c48f;--oks:#17302a;--sel:#2a2119}}
*{box-sizing:border-box}html{scroll-behavior:smooth;scroll-padding-top:70px}
body{margin:0;background:var(--bg);color:var(--ink);font:14.5px/1.5 -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,"PingFang SC","Microsoft YaHei",sans-serif}
code,pre{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,monospace;font-size:.86em}
code{background:var(--code);padding:.1em .35em;border-radius:3px;word-break:break-all}
a{color:var(--ac);text-decoration:none}a:hover{text-decoration:underline}
.cn{color:var(--mut);font-weight:400;font-size:.9em}
.top{position:sticky;top:0;z-index:5;background:var(--sf);border-bottom:1px solid var(--ln);display:flex;gap:16px;align-items:center;padding:10px 20px}
.top h1{font-size:1.05rem;margin:0;white-space:nowrap}
.top input{flex:1;max-width:560px;padding:8px 12px;border:1px solid var(--ln2);border-radius:5px;background:var(--bg);color:var(--ink);font:inherit}
.top .cnt{color:var(--mut);font-size:.85rem;white-space:nowrap}
.wrap{display:grid;grid-template-columns:290px minmax(0,1fr);gap:0}
nav{position:sticky;top:53px;height:calc(100vh - 53px);overflow:auto;border-right:1px solid var(--ln);padding:14px 10px 40px;font-size:.86rem;background:var(--sf)}
nav .gname{font-weight:600;margin:12px 6px 4px;font-size:.8rem;text-transform:uppercase;letter-spacing:.05em;color:var(--mut)}
nav a{display:block;color:var(--ink);padding:3px 8px;border-radius:4px;border-left:2px solid transparent}
nav a:hover{background:var(--acs);text-decoration:none}nav a.on{border-left-color:var(--ac);color:var(--ac)}
nav a.hide,nav .grp.hide{display:none}
main{padding:20px 28px 80px;max-width:1300px}
.intro{color:var(--mut);max-width:80ch;margin:0 0 20px}
h2.cat{margin:34px 0 12px;font-size:1.35rem;padding-bottom:6px;border-bottom:2px solid var(--ln2)}h2.cat .n{font-size:.8rem;color:var(--mut);font-weight:500;margin-left:8px}
section.api{background:var(--sf);border:1px solid var(--ln);border-radius:6px;padding:16px 18px;margin:0 0 14px}
section.api.hide,h2.cat.hide{display:none}
section.api header{display:flex;gap:14px;align-items:flex-start}
.num{flex:none;font-variant-numeric:tabular-nums;color:var(--mut);font-size:.8rem;padding-top:6px;min-width:2.2em;text-align:right}
.hd{min-width:0;flex:1}
h3{margin:0 0 6px;font-size:1.12rem}
.id,.url{display:flex;align-items:center;gap:8px;flex-wrap:wrap;margin:4px 0;font-size:.9rem}
.method{font-weight:700;font-size:.72rem;letter-spacing:.06em;color:var(--ok);background:var(--oks);padding:1px 7px;border-radius:3px}
.copy{font:inherit;font-size:.72rem;color:var(--mut);background:transparent;border:1px solid var(--ln2);border-radius:3px;padding:1px 7px;cursor:pointer}
.copy:hover{color:var(--ac);border-color:var(--ac)}.copy.done{color:var(--ok);border-color:var(--ok)}
.chips{display:flex;flex-wrap:wrap;gap:6px;margin:8px 0 2px}
.chip{font-size:.76rem;font-family:ui-monospace,Menlo,monospace;border:1px solid var(--ln2);border-radius:999px;padding:2px 9px;color:var(--ink)}
.chip i{font-style:normal;color:var(--mut);margin-left:6px;font-family:-apple-system,sans-serif}
.chip.req{border-color:var(--ac);background:var(--acs)}.chip.flag{background:var(--code);border-color:transparent;color:var(--mut);font-family:-apple-system,sans-serif}
p.desc{margin:12px 0 10px;max-width:90ch}
h4{margin:14px 0 6px;font-size:.78rem;letter-spacing:.07em;text-transform:uppercase;color:var(--mut)}
details{margin:10px 0 0}summary{cursor:pointer;font-size:.86rem;color:var(--ac);font-weight:500;padding:4px 0}
table{border-collapse:collapse;width:100%;font-size:.86rem}.tw{overflow-x:auto;border:1px solid var(--ln);border-radius:4px}
th{text-align:left;font-size:.7rem;letter-spacing:.06em;text-transform:uppercase;color:var(--mut);font-weight:600;padding:6px 10px;border-bottom:1px solid var(--ln2);background:var(--code)}
td{padding:6px 10px;border-bottom:1px solid var(--ln);vertical-align:top}
td.f{white-space:nowrap}td.t{color:var(--mut);font-family:ui-monospace,Menlo,monospace;font-size:.8rem;white-space:nowrap}td.r{white-space:nowrap}td.r b{color:var(--ac)}
td.desc{min-width:260px;max-width:560px}td.ex{color:var(--mut);font-family:ui-monospace,Menlo,monospace;font-size:.78rem;max-width:320px;word-break:break-all}
tr.obj td.f code{background:var(--acs);color:var(--ac)}
tr.d1 td.f{padding-left:28px}tr.d2 td.f{padding-left:46px}tr.d3 td.f{padding-left:64px}tr.d4 td.f{padding-left:82px}tr.d5 td.f{padding-left:100px}tr.d6 td.f{padding-left:118px}
tr[class*="d"]:not(.d0) td.f .tree::before{content:"↳ ";color:var(--mut)}
tr.hit td{background:var(--sel)}
pre{background:var(--code);padding:10px 12px;border-radius:5px;overflow-x:auto;max-height:340px;white-space:pre-wrap;word-break:break-all}
p.sn{margin:8px 0 4px;font-size:.85rem;font-weight:600}p.none{color:var(--mut);font-size:.9rem}
@media(max-width:900px){.wrap{grid-template-columns:1fr}nav{position:static;height:auto;max-height:40vh;border-right:0;border-bottom:1px solid var(--ln)}main{padding:14px}}
'''
js='''
(function(){
var q=document.getElementById('q'),cnt=document.getElementById('cnt');
var secs=[].slice.call(document.querySelectorAll('section.api')),cats=[].slice.call(document.querySelectorAll('h2.cat'));
var links=[].slice.call(document.querySelectorAll('nav a')),grps=[].slice.call(document.querySelectorAll('nav .grp'));
var total=secs.length;
function apply(){var t=q.value.trim().toLowerCase(),shown=0,vis={};
 secs.forEach(function(s){[].forEach.call(s.querySelectorAll('tr.hit'),function(r){r.classList.remove('hit')});
  var ok=!t||s.textContent.toLowerCase().indexOf(t)>-1;s.classList.toggle('hide',!ok);
  if(ok){shown++;vis[s.dataset.cat]=1;if(t){[].forEach.call(s.querySelectorAll('tbody tr'),function(r){if(r.textContent.toLowerCase().indexOf(t)>-1)r.classList.add('hit')})}}
  var l=document.querySelector('nav a[href="#'+s.id+'"]');if(l)l.classList.toggle('hide',!ok);});
 cats.forEach(function(c){c.classList.toggle('hide',!vis[c.id])});grps.forEach(function(g){g.classList.toggle('hide',!vis[g.querySelector('.gname').dataset.cat])});
 cnt.textContent=t?shown+' of '+total+' APIs':total+' APIs';}
q.addEventListener('input',apply);
document.addEventListener('click',function(ev){var b=ev.target.closest('.copy');if(!b)return;var v=b.dataset.copy;
 (navigator.clipboard?navigator.clipboard.writeText(v):Promise.reject()).then(function(){b.classList.add('done');b.textContent='copied';setTimeout(function(){b.classList.remove('done');b.textContent='copy'},1200)},function(){window.prompt('Copy:',v)});});
if('IntersectionObserver' in window){var io=new IntersectionObserver(function(es){es.forEach(function(en){if(en.isIntersecting){links.forEach(function(l){l.classList.toggle('on',l.getAttribute('href')==='#'+en.target.id)})}})},{rootMargin:'-15% 0px -75% 0px'});secs.forEach(function(s){io.observe(s)});}
})();
'''
page=f'''<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>1688 XunYuanTong APIs</title><style>{css}</style></head><body>
<div class="top"><h1>1688 XunYuanTong APIs</h1><input id="q" type="search" placeholder="Search API name, id, field, description…" aria-label="Search"><span class="cnt" id="cnt">{n} APIs</span></div>
<div class="wrap"><nav>{''.join(nav)}</nav><main>
<p class="intro">Solution 1703167397970 · pulled from open.1688.com on 2026-09-03. Every call is an HTTP POST (form-encoded) to <code>https://gw.open.1688.com/openapi/param2/{{version}}/{{namespace}}/{{name}}/{{appKey}}</code>. System parameters travel with the body: <code>access_token</code> (from the OAuth authorization of the admin account), <code>_aop_signature</code> (HMAC-SHA1 request signature), <code>_aop_timestamp</code> (optional). Object-typed fields are sent as JSON strings; nested fields are indented under their parent (highlighted names are objects).</p>
{''.join(body)}
</main></div><script>{js}</script></body></html>'''
open(f'{OUT}/ALL-APIS-en.html','w',encoding='utf-8').write(page)
print('apis',n,'bytes',os.path.getsize(f'{OUT}/ALL-APIS-en.html'))
