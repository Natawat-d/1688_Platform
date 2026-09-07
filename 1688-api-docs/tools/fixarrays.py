import json,glob,urllib.request,urllib.parse,time,concurrent.futures,threading,os
BASE='https://open.1688.com'; lock=threading.Lock(); stats={'req':0,'filled':0,'empty':0}
def get(path,params):
    url=BASE+path+'?'+urllib.parse.urlencode(params); last=None
    for i in range(5):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':'Mozilla/5.0','Referer':'https://open.1688.com/api/apidocdetail.htm'})
            with urllib.request.urlopen(req,timeout=40) as r:
                with lock: stats['req']+=1
                return json.loads(r.read().decode('utf-8'))
        except Exception as e: last=e; time.sleep(1.5*(i+1))
    raise last
def clean(tn): return tn[:-2] if tn.endswith('[]') else tn
def resolve(ns,name,ver,typ,tn,cache,stack):
    c=clean(tn)
    if c in stack: return [{'_cycle':c}]
    if (typ,c) in cache: return cache[(typ,c)]
    r=get('/api/data/getModelInfo.json',{'namespace':ns,'apiname':name,'version':ver,'type':typ,'typeName':c,'_input_charset':'UTF-8'})
    fields=r.get('result') or []
    with lock: stats['filled' if fields else 'empty']+=1
    cache[(typ,c)]=fields
    for f in fields:
        if f.get('complexTypeFlag') and f.get('typeName'):
            f['fields']=resolve(ns,name,ver,typ,f['typeName'],cache,stack|{c})
    return fields
def fix(fields,ns,name,ver,typ,cache,stack):
    n=0
    for f in fields or []:
        if '_cycle' in f: continue
        tn=f.get('typeName')
        if f.get('complexTypeFlag') and tn and tn.endswith('[]') and not f.get('fields'):
            f['fields']=resolve(ns,name,ver,typ,tn,cache,stack); n+=1
        elif f.get('fields'):
            n+=fix(f['fields'],ns,name,ver,typ,cache,stack|{clean(tn or '')})
    return n
def one(fn):
    d=json.load(open(fn,encoding='utf-8')); cache={}; n=0
    for typ,k in ((1,'apiAppParamVOList'),(2,'apiReturnParamVOList')):
        n+=fix(d.get(k),d['namespace'],d['name'],d['version'],typ,cache,set())
    if n: json.dump(d,open(fn,'w',encoding='utf-8'),ensure_ascii=False,indent=1)
    return n
files=sorted(glob.glob('/Users/natawatd/Desktop/temp/1688-api-docs/raw/*.json'))
tot=0
with concurrent.futures.ThreadPoolExecutor(max_workers=6) as ex:
    for fn,n in zip(files,ex.map(one,files)): tot+=n
print('array nodes re-resolved',tot,'requests',stats['req'],'filled',stats['filled'],'still empty',stats['empty'])
