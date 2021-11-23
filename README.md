# go-mongo-rest

Example of how to make REST endpoints with Go and Mongodb.

## How to run

Copy/paste this into your terminal:

```sh
git clone git@github.com:jakecoffman/go-mongo-rest.git
cd go-mongo-rest
go run cmd/server/server.go
```

## Dependencies

- [gin](https://github.com/gin-gonic/gin)
  - This is a good router that has been around for a while. Feel free to pick your own. The builtin http.ServeMux doesn't support path variables and method routing so doing a deep path with many variables (/authors/{authorId}/books/{bookId}) is unruly.
- [crud](https://github.com/jakecoffman/crud)
  - Provides an easy way to get OpenAPI docs and validation middleware.
- [mongo-driver](https://github.com/mongodb/mongo-go-driver)
  - Official Go Mongodb driver.

# Project layout

- cmd
  - Typical pattern in Go to store the main packages
- lib
  - It's common to separate the rest of the code into another package. I use lib, you can use pkg, it doesn't matter.
- lib/db
  - The mongo.Client and collections are exported from this package since they are thread safe. Makes things easier to manage than passing it down through Request Contexts.
- lib/endpoonts
  - All of our handlers and route definitions under this directory.
- lib/models
  - Models live in the same package since they often depend on each other, but we've separated concerns from the handlers. 

## TODO

- With generics around the corner, we can remove all the boilerplate in author_handlers.go
