import json,glob,os,html as H,sys
OUT='/Users/natawatd/Desktop/temp/1688-api-docs'
LANG=sys.argv[1] if len(sys.argv)>1 else 'zh'
T={}
if LANG=='en':
    for fn in glob.glob('translations/*.json'):
        T.update(json.load(open(fn,encoding='utf-8')))
def tr(s):
    s=(s or '').strip()
    if LANG=='zh' or not s: return s
    return T.get(s,s)
def cell(s): return (s or '').replace('|','\\|').replace('\n','<br>')
def unq(s): return H.unescape(s or '')
os.makedirs(f'{OUT}/{LANG}',exist_ok=True)
LST=json.load(open('apilist.json',encoding='utf-8'))
CATS={}
if LANG=='en':
    import importlib.util
    spec=importlib.util.spec_from_file_location('tl','translate.py'); 
    src=open('translate.py',encoding='utf-8').read().split('# ---------------------------------------------------------------- coverage check')[0]
    ns={}; exec(src,ns); CATS=ns['CATS']; API=ns['API']
def table(fields,seen,out,depth=0):
    out.append('| Field | Type | Required | Description | Example |' if LANG=='en' else '| 字段 | 类型 | 必填 | 描述 | 示例 |')
    out.append('|---|---|---|---|---|')
    subs=[]
    for f in fields or []:
        if '_cycle' in f: continue
        typ=f.get('type') or ''
        if f.get('complexTypeFlag') and f.get('typeName'):
            typ=f"[{typ}](#{anchor(f['typeName'])})"
            if f['typeName'] not in seen:
                seen.add(f['typeName']); subs.append(f)
        req=('yes' if f.get('required') else 'no') if LANG=='en' else ('是' if f.get('required') else '否')
        out.append(f"| `{f.get('name')}` | {cell(typ)} | {req} | {cell(tr(f.get('description')))} | {cell(unq(f.get('exampleValue')))} |")
    out.append('')
    for f in subs:
        out.append(f"<a id=\"{anchor(f['typeName'])}\"></a>")
        out.append(f"#### {f['typeName']}\n")
        table(f.get('fields'),seen,out,depth+1)
def anchor(tn): return 'm-'+tn.replace('.','-').lower()
index=[]
for cat in LST['apis']:
    cn=cat['categoryFamilyName']; index.append(f"\n## {CATS.get(cn,cn) if LANG=='en' else cn}\n")
    for m in cat['modules']:
        base=f"{m['namespace']}.{m['name']}-{m['version']}"
        fn=f"{OUT}/raw/{base}.json"
        if not os.path.exists(fn):
            index.append(f"- {m['displayName']} `{m['fullName']}` — detail page unavailable on open.1688.com"); continue
        d=json.load(open(fn,encoding='utf-8'))
        name_en=API[m['fullName']][0] if LANG=='en' else d['displayName']
        o=[]
        o.append(f"# {name_en}\n")
        if LANG=='en': o.append(f"Original name: {d['displayName']}  ")
        o.append(f"API: `{d['namespace']}:{d['name']}:{d['version']}` · Category: {CATS.get(cn,cn) if LANG=='en' else cn}  ")
        o.append(f"Doc page: https://open.1688.com/api/apidocdetail.htm?id={d['namespace']}:{d['name']}-{d['version']}  ")
        o.append(f"Request URL: `https://gw.open.1688.com/openapi/param2/{d['version']}/{d['namespace']}/{d['name']}/{{appKey}}`  ")
        flags=[]
        flags.append(('Requires user authorization (access_token)' if LANG=='en' else '需要授权 (access_token)') if d.get('needAuth') else ('No user authorization' if LANG=='en' else '无需授权'))
        flags.append(('Requires signature' if LANG=='en' else '需要签名') if d.get('needSignature') else ('No signature' if LANG=='en' else '无需签名'))
        o.append(' · '.join(flags)+'\n')
        o.append(tr(d.get('description'))+'\n')
        o.append('## System parameters\n' if LANG=='en' else '## 系统级参数\n')
        table(d.get('apiSystemParamVOList'),set(),o)
        o.append('## Request parameters\n' if LANG=='en' else '## 应用级参数\n')
        table(d.get('apiAppParamVOList'),set(),o)
        o.append('## Response\n' if LANG=='en' else '## 返回结果\n')
        table(d.get('apiReturnParamVOList'),set(),o)
        if d.get('apiErrorCodeVOList'):
            o.append('## Error codes\n' if LANG=='en' else '## 错误码\n')
            o.append('| Error | Symptom | How to fix |' if LANG=='en' else '| 错误码 | 现象 | 解决方案 |'); o.append('|---|---|---|')
            for ec in d['apiErrorCodeVOList']:
                o.append(f"| {cell(tr(ec.get('code') or ''))} | {cell(tr(ec.get('desc') or ''))} | {cell(tr(ec.get('howToFix') or ''))} |")
            o.append('')
        if d.get('apiDocSampleVOList'):
            o.append('## Samples\n' if LANG=='en' else '## 示例\n')
            for s in d['apiDocSampleVOList']:
                o.append(f"**{tr(s.get('name') or '')}**\n")
                for k in ('sample',):
                    if s.get(k): o.append('```\n'+unq(str(s[k]))+'\n```\n')
        if d.get('returnExample') and d['returnExample'] not in ('[]','{}',''):
            o.append('## Response example\n' if LANG=='en' else '## 返回示例\n'); o.append('```json\n'+unq(d['returnExample'])+'\n```\n')
        open(f"{OUT}/{LANG}/{base}.md",'w',encoding='utf-8').write('\n'.join(o))
        index.append(f"- [{name_en}]({LANG}/{base}.md) `{m['fullName']}`" + (f" — {d['displayName']}" if LANG=='en' else ''))
title='# 1688 API docs (English)' if LANG=='en' else '# 1688 API 文档（中文原文）'
open(f"{OUT}/INDEX-{LANG}.md",'w',encoding='utf-8').write(title+'\n\nSolution 1703167397970 · pulled from open.1688.com on 2026-09-03\n'+'\n'.join(index)+'\n')
print('rendered',len(glob.glob(f'{OUT}/{LANG}/*.md')),'files')
