# 1688 Open Platform API docs (solution 1703167397970)

Local, offline copy of the API documentation for the 1688 "XunYuanTong cross-border
key-account sourcing" solution (跨境大客户寻源通解决方案), pulled from open.1688.com on
2026-09-03 and translated to English.

Source page: https://open.1688.com/solution/detail?key=1703167397970

## Layout

| Path | What it is |
|---|---|
| `OVERVIEW-en.md` | The solution page itself: all 143 APIs and 39 message topics with names and descriptions, change log, SDK notes |
| `INDEX-en.md` / `INDEX-zh.md` | Index of the per-API detail docs (English / Chinese original) |
| `en/<namespace>.<name>-<version>.md` | English detail doc per API: request URL, system/request/response parameters (nested models expanded), error codes, samples |
| `zh/...` | Same docs, Chinese original |
| `raw/*.json` | Raw JSON as returned by open.1688.com (`getApiDetail.json` + resolved `getModelInfo.json` models) |
| `INDEX-messages-en.md` / `INDEX-messages-zh.md` | Index of push-message topic docs |
| `messages-en/<TOPIC>.md`, `messages-zh/`, `messages-raw/` | Per-topic payload field docs (English / Chinese / raw JSON) |
| `tools/` | Scripts used to pull, resolve and render everything (see below) |

Field names, enum values, identifiers, URLs and sample payloads are kept verbatim; only
Chinese prose was translated. Every English doc shows the Chinese original name under the
title so entries can be matched with the live site.

## Known gaps

* `com.alibaba.trade:alibaba.trade.queryOrderByInsure:1` has no detail doc: the 1688 detail
  endpoint answers "error finding API" for it, although it is listed in the solution.
* One nested model type returned no fields from 1688 (shown as an empty table).
* Request-URL format and signing rules are summarised in `../PLATFORM-PLAN.md`; the
  authoritative text is https://open.1688.com/doc/apiInvoke.htm and
  https://open.1688.com/doc/apiAuth.htm.

## Refreshing the docs

The pages on open.1688.com are JavaScript apps; the data comes from JSON endpoints:

```
Solution:     https://open.1688.com/solution/data/getSolutionDetail.jsonp?solutionKey=1703167397970&callback=cb
API detail:   https://open.1688.com/api/data/getApiDetail.json?namespace=..&name=..&version=..
Model fields: https://open.1688.com/api/data/getModelInfo.json?namespace=..&apiname=..&version=..&type=1|2&typeName=..
Topic:        https://open.1688.com/msg/dataNew/getTopic.json?topicId=<TOPIC_NAME>
```

`tools/crawl.py` pulls all API details (needs `apilist.json` next to it), `tools/fixarrays.py`
re-resolves array model types (the `[]` suffix must be stripped for `getModelInfo`),
`tools/render.py zh|en` renders the markdown, `tools/render_msgs.py zh|en` renders topics.
English rendering expects a `translations/*.json` folder of `{chinese: english}` maps.
