# links
Golang URL Shortener

## Starting up
### Default postgresql version:
```
docker-compose up
```
### Dev version with no database:
```
make build-inmemory
make run
```

## REST API interface
### Create link
Request: `POST`

URL: `/create`

Body:

    {
        "url": < url >
    }

Response: `< link >`

### Create link #2
Request: `GET`

URL: `/create/< url >`

Response: `< link >`

### Use link
Request: `/< link >`

Response: redirect to the stored url.

## Example of use
- GET Request to the `http://localhost:8080/create/google.com`

   Response: "R"


- GET Request to the `http://localhost:8080/R`

   Response: redirect to the `http://google.com`

