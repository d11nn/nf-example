## Greeting Service

GET reads the name from a query parameter.
If omitted, the name defaults to free5GC.

```sh
curl --noproxy '*' 'http://127.0.0.163:8000/greeting/?name=Anya'
# {"message":"Hello Anya!"}
```

POST reads the name from a JSON body.
A missing or empty name returns HTTP 400.

```sh
curl --noproxy '*' http://127.0.0.163:8000/greeting/ \
  -H 'Content-Type: application/json' \
  -d '{"name":"Loid"}'
# {"message":"Hello Loid!"}
```