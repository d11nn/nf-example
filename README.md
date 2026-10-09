# NF-Example

This repository is a sample NF for a simple HTTP service!
Try adding more services and learn how to collaborate using GitHub.

## Compile & Run

```sh
make
./bin/nf -c config/nfcfg.yaml
```

## Try Service

```sh
> curl -X GET http://127.0.0.163:8000/default/
"Hello free5GC!"

> curl -X GET http://127.0.0.163:8000/spyfamily/
"Hello SPYxFAMILY!"

> curl -X GET http://127.0.0.163:8000/spyfamily/character/Loid
"Character: Loid Forger"
```

## Go Test

```sh
> go test -v ./...
```

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
