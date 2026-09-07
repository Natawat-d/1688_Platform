import json, os, time, urllib.request, urllib.parse, concurrent.futures, html as H, sys, threading
BASE='https://open.1688.com'
OUT='/Users/natawatd/Desktop/temp/1688-api-docs'
os.makedirs(f'{OUT}/raw',exist_ok=True); os.makedirs(f'{OUT}/zh',exist_ok=True)
LST=json.load(open('apilist.json',encoding='utf-8'))
apis=[(f['categoryFamilyName'],m) for f in LST['apis'] for m in f['modules']]
lock=threading.Lock(); stats={'req':0,'err':0}
def get(path,params):
    url=BASE+path+'?'+urllib.parse.urlencode(params)
    last=None
    for i in range(5):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':'Mozilla/5.0','Referer':'https://open.1688.com/api/apidocdetail.htm'})
            with urllib.request.urlopen(req,timeout=40) as r:
                with lock: stats['req']+=1
                return json.loads(r.read().decode('utf-8'))
        except Exception as e:
            last=e; time.sleep(1.5*(i+1))
    with lock: stats['err']+=1
    raise last
def resolve(ns,name,ver,typ,tn,cache,stack):
    if tn in stack: return [{'_cycle':tn}]
    key=(typ,tn)
    if key in cache: return cache[key]
    r=get('/api/data/getModelInfo.json',{'namespace':ns,'apiname':name,'version':ver,'type':typ,'typeName':tn,'_input_charset':'UTF-8'})
    fields=r.get('result') or []
    cache[key]=fields
    for f in fields:
        if f.get('complexTypeFlag') and f.get('typeName'):
            f['fields']=resolve(ns,name,ver,typ,f['typeName'],cache,stack|{tn})
    return fields
def one(item):
    cat,m=item
    ns,name,ver=m['namespace'],m['name'],m['version']
    fn=f"{OUT}/raw/{ns}.{name}-{ver}.json"
    if os.path.exists(fn):
        return fn
    d=get('/api/data/getApiDetail.json',{'namespace':ns,'name':name,'version':ver,'_input_charset':'UTF-8'})
    res=d.get('result')
    if not res:
        print('NO RESULT',ns,name,ver,str(d)[:200]); return None
    cache={}
    for typ,k in ((1,'apiAppParamVOList'),(2,'apiReturnParamVOList')):
        for p in res.get(k) or []:
            if p.get('complexTypeFlag') and p.get('typeName'):
                p['fields']=resolve(ns,name,ver,typ,p['typeName'],cache,set())
    res['_category']=cat
    res.pop('bizSolutionDTOList',None)
    json.dump(res,open(fn,'w',encoding='utf-8'),ensure_ascii=False,indent=1)
    return fn
with concurrent.futures.ThreadPoolExecutor(max_workers=6) as ex:
    futs={ex.submit(one,it):it for it in apis}
    done=0
    for f in concurrent.futures.as_completed(futs):
        done+=1
        try: f.result()
        except Exception as e: print('FAIL',futs[f][1]['fullName'],e)
        if done%20==0: print('progress',done,'/',len(apis),'requests',stats['req'],flush=True)
print('DONE',len(os.listdir(f'{OUT}/raw')),'raw files; requests',stats['req'],'errors',stats['err'])
