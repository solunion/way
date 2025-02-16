# Database

## Create migration script
```sh
  migrate -database 'postgres://way:way@postgres-127.0.0.1.nip.io:5432/way?sslmode=disable&search_path=way' create -dir=db/migrations -seq -ext=sql my-script
```

## Run migration script
```sh
  migrate -database 'postgres://way:way@postgres-127.0.0.1.nip.io:5432/way?sslmode=disable&search_path=way' -path db/migrations up
```
