```
$ curl -XPOST -H "Content-Type: application/json" \
     -d '{"uid":"joske","app":"mysecureapp","fingerprint":"blahblah", "auth":"1factor","ip":"1.2.3.4","country":"GB"}' \ 
     http://localhost:7755/collect/zending
```

returns
```
{"uid":"joske","timestamp":"2017-05-16T16:00:34.762458174+02:00","ip":"1.2.3.4","app":"mysecureapp","fingerprint":"blahblah","auth":"1factor","country":"GB","id":"zending"}
```
