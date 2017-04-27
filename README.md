# Web API Base

This is a base API for webservices. It uses gorilla/mux to deliver calls.
Includes basic logging of every request/response issued by the server to STDOUT.

### Execution

```sh
go build && ./web-api-base
```

### Usage

#### Adding a Route
  To add a new route, simply add the 4 pieces of data needed to [routes.go](./routes.go)
	 - Name of the route
	 - Path pattern that matches the route
	 - The HTTP Method the route returns true on
	 - The Handler Function that executes for this route.


#### Adding a Handler Function
  To add a new handler, please place all your code in its own file, and the handler that executes it in handler.go
