# Middleware payment

### Use examples

Request:

`api.exampleheadwaydomain.com/v1/third_party_pay?product_id=16e64fb&service=paypal`

Headers:
```
HTTP/2 200
content-type: application/json;charset=UTF-8
content-length: 94
date: Fri, 01 Oct 2021 16:01:10 GMT
server: nginx
x-response-time: 19ms
x-server-name: api.exampleheadwaydomain.com
```
Body:
```json
{
  "status": "ok",
  "href": "https://api-m.paypal.com/v2/payments/authorizations/0VF52814937998046"
}
```

Request:

`api.exampleheadwaydomain.com/v1/third_party_pay?product_id=16e64fb&service=INVALIDSERVICE`

Headers:
```
HTTP/2 404
content-type: application/json;charset=UTF-8
content-length: 87
date: Fri, 01 Oct 2021 16:01:10 GMT
server: nginx
x-response-time: 0ms
x-server-name: api.exampleheadwaydomain.com
```
Body:
```json
{
  "status": "fail",
  "error": "unknown payment service",
  "href": "https://headway.onelink.me/8zSH/playstore"
}
```