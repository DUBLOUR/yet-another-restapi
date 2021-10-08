# Middleware payment

##### Setup:

`git clone https://github.com/DUBLOUR/yet-another-restapi.git && cd yet-another-restapi`

`go build -o bin/main cmd/main.go && ./bin/main`

##### Usage examples:

Query:

`http://localhost:8999/?product_id=2021bestsellers&service=vp`

Headers:
```
HTTP/1.1 200 OK
Content-Length: 92
Content-Type: text/json
X-Response-Time: 172us
X-Server-Name: api-headway.com
Date: Fri, 08 Oct 2021 03:06:58 GMT
```
Body:
```json
{
  "status":"ok",
  "href":"https://virtualpay.dev/?money=3.99\u0026id=5jnoe5r8apn5pztsgsvupau2"
}
```
Trace-log:
```
(II) Init server...
(--) Load file data/products.json
(--) Successfully loaded all products from data/products.json
(--) Load payment services...
(--) Load VP
(--) Load PayPal
(--) Load QIWI
(II) Initialize is successful
(II) Start server at port :8999
(II) Query: map[product_id:[2021bestsellers] service:[vp]]
(II) Respond with status 200
(--) {"status":"ok","href":"https://virtualpay.dev/?money=3.99\u0026id=5r3u589g8a05hz6zk94neuh2"}
(II) Query: map[product_id:[syllabus_classics] service:[binance]]
(II) Respond with status 500
(--) {"status":"fail","error":"unknown payment service","href":"https://headway.onelink.me/8zSH/playstore"}
...
```

#### Another requests

Unknown payment gateway

`http://localhost:8999/?product_id=syllabus_classics&service=binance`
```
HTTP/1.1 500 Internal Server Error
Content-Length: 102
Content-Type: text/json
X-Response-Time: 116us
X-Server-Name: api-headway.com
Date: Fri, 08 Oct 2021 03:18:06 GMT
```

```json
{
  "status":"fail",
  "error":"unknown payment service",
  "href":"https://headway.onelink.me/8zSH/playstore"
}
```

[comment]: <> (Unavailable product:)

[comment]: <> (`localhost:8999/?product_id=2020bestsellers&service=vp`)

[comment]: <> (```)

[comment]: <> (HTTP/1.1 500 Internal Server Error)

[comment]: <> (Content-Length: 112)

[comment]: <> (Content-Type: text/json)

[comment]: <> (X-Response-Time: 215us)

[comment]: <> (X-Server-Name: api-headway.com)

[comment]: <> (Date: Fri, 08 Oct 2021 02:57:38 GMT)

[comment]: <> (```)

[comment]: <> (```json)

[comment]: <> ({)

[comment]: <> (  "status":"fail",)

[comment]: <> (  "error":"product is not available for sale",)

[comment]: <> (  "href":"https://headway.onelink.me/8zSH/playstore")

[comment]: <> (} )

[comment]: <> (```)

Missed parameters

`localhost:8999/?product_id=weekly_digest_relationship`

```
HTTP/1.1 400 Bad Request
Content-Length: 113
Content-Type: text/json
X-Response-Time: 115us
X-Server-Name: api-headway.com
Date: Fri, 08 Oct 2021 03:29:52 GMT
```
```json
{
  "status":"fail",
  "error":"`product_id` or `service` is empty",
  "href":"https://headway.onelink.me/8zSH/playstore"
}
```