# candidate-matcher
A simple `candidate-matcher` Go Fiber app. It uses Facebook `ent` ORM framework backed by `sqlite3`.

## Running locally

### Using `go`
```sh
export GO111MODULE=on
git clone git@github.com:zeako/candidate-matcher.git && cd candidate-matcher
export DB_FILE_PATH=<db-file-path>  # db file location
go run cmd/app/main.go # exposes service on port 8080
```

### Using `docker` on port 8080
Without persistency:
```sh
docker run --rm -it -p 8080:8080 zeako/candidate-matcher # latest
```

With persistent sqlite3 db:
```sh
docker run --rm -it -p 8080:8080 -e DB_FILE_PATH=/tmp/db.sqlite3 --mount type=bind,source=<dna-file-path>,target=/tmp/db.sqlite3 zeako/candidate-matcher # latest
```

## Sample
Create candidates:
```sh
curl --location --request POST 'http://localhost:8080/jobs/' --header 'Content-Type: application/json' --data-raw '{
    "title": "Architect",
    "skill": "C"
}'

curl --location --request POST 'http://localhost:8080/jobs/' --header 'Content-Type: application/json' --data-raw '{
    "title": "Junior Architect",
    "skill": "C"
}'
```

Create job:
```sh
curl --location --request POST 'http://localhost:8080/jobs/' --header 'Content-Type: application/json' --data-raw '{
    "title": "Architect",
    "skill": "C"
}'
```

Find matching candidates:
```sh
curl --location --request GET 'http://localhost:8080/jobs/1/candidates'
```

\* Input is case-insensitive.
