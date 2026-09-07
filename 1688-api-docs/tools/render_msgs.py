import json,glob,os,html as H,sys
OUT='/Users/natawatd/Desktop/temp/1688-api-docs'
LANG=sys.argv[1] if len(sys.argv)>1 else 'zh'
T={}
if LANG=='en':
    for fn in glob.glob('translations/*.json'): T.update(json.load(open(fn,encoding='utf-8')))
    src=open('translate.py',encoding='utf-8').read().split('# ---------------------------------------------------------------- coverage check')[0]
    ns={}; exec(src,ns); CATS=ns['CATS']; MSG=ns['MSG']
def tr(s):
    s=(s or '').strip()
    return s if LANG=='zh' or not s else T.get(s,s)
def cell(s): return (s or '').replace('|','\\|').replace('\n','<br>')
os.makedirs(f'{OUT}/messages-{LANG}',exist_ok=True)
LST=json.load(open('apilist.json',encoding='utf-8'))
def emit(nodes,out,depth=0):
    for n in nodes or []:
        ind='&nbsp;&nbsp;'*depth+('↳ ' if depth else '')
        req=('yes' if n.get('required') else 'no') if LANG=='en' else ('是' if n.get('required') else '否')
        out.append(f"| {ind}`{n.get('name')}` | {cell(n.get('type'))} | {req} | {cell(tr(n.get('desc')))} | {cell(H.unescape((n.get('sample') or '').strip()))} |")
        emit(n.get('children'),out,depth+1)
index=[]
for fam in LST['messages']:
    g=fam['categoryFamilyName']; index.append(f"\n## {CATS.get(g,g) if LANG=='en' else g}\n")
    for m in fam['modules']:
        tn=m['typeName']; fn=f'{OUT}/messages-raw/{tn}.json'
        if not os.path.exists(fn): continue
        d=json.load(open(fn,encoding='utf-8'))
        title=MSG[tn] if LANG=='en' else d.get('topicDisplayName') or tn
        o=[f"# {title}\n"]
        if LANG=='en': o.append(f"Original name: {d.get('topicDisplayName')}  ")
        o.append(f"Topic: `{tn}` · Group: {d.get('topicGroupName')} ({CATS.get(d.get('topicGroupDisplayName'),d.get('topicGroupDisplayName')) if LANG=='en' else d.get('topicGroupDisplayName')})  ")
        o.append(f"Doc page: https://open.1688.com/doc/topicDetail.htm?id={tn}\n")
        o.append(tr(d.get('desc'))+'\n')
        o.append('## Payload fields\n' if LANG=='en' else '## 消息字段\n')
        o.append('| Field | Type | Required | Description | Example |' if LANG=='en' else '| 字段 | 类型 | 必填 | 描述 | 示例 |'); o.append('|---|---|---|---|---|')
        emit(d.get('messageDocs'),o); o.append('')
        if d.get('sample') and d['sample'].strip():
            o.append('## Sample message\n' if LANG=='en' else '## 消息示例\n')
            s=H.unescape(d['sample'])
            try: s=json.dumps(json.loads(s),ensure_ascii=False,indent=2)
            except Exception: pass
            o.append('```json\n'+s+'\n```\n')
        open(f'{OUT}/messages-{LANG}/{tn}.md','w',encoding='utf-8').write('\n'.join(o))
        index.append(f"- [{title}](messages-{LANG}/{tn}.md) `{tn}`"+(f" — {d.get('topicDisplayName')}" if LANG=='en' else ''))
title='# 1688 message topics (English)' if LANG=='en' else '# 1688 消息主题（中文原文）'
open(f'{OUT}/INDEX-messages-{LANG}.md','w',encoding='utf-8').write(title+'\n\nSolution 1703167397970 · pulled from open.1688.com on 2026-09-03\n'+'\n'.join(index)+'\n')
print('rendered',len(glob.glob(f'{OUT}/messages-{LANG}/*.md')),'message docs')
