# protoc-gen-gosvc

This repo is a bit of an experiment into generating service interfaces with 
protoc. The goal is to be able to generate as much of the http hendler code
as possible, including: request parsing, error handling, response writing etc. 

## Run the examle

As a proof of concept I've implemented a small Hello World service with the 
following proto definition. 
```
syntax = "proto3";

package helloworld;

import "google/api/annotations.proto";

service HelloWorld {
    rpc SayHello(SayHelloRequest) returns (SayHelloResponse) {
        option (google.api.http).post = "/hello";
    }

    rpc SayGoodbye(SayGoodbyeRequest) returns (SayGoodbyeResponse) {
        option (google.api.http).get = "/goodbye";
    }
}

message SayHelloRequest {
    string name = 1;
}

message SayHelloResponse {
    string message = 1;
}

message SayGoodbyeRequest {}

message SayGoodbyeResponse {
    string message = 1;
}
```

The example will generate code using this proto definition and then compile it
alongside the handler implementations. 
Running the resulting binary will spin up an HTTP server on `localhost:8080` 
that handles the `/hello` and `/goodbye` endpoints.
```
[user@host]$ make example
```

You can hit these with `httpie`, for example. 

```
[user@host]$ http POST localhost:8080/hello name=dave
HTTP/1.1 200 OK
Content-Length: 29
Content-Type: application/json
Date: Sun, 23 May 2021 21:29:02 GMT

{
    "message": "Hello dave!"
}

[user@host]$ http localhost:8080/goodbye
HTTP/1.1 200 OK
Content-Length: 22
Content-Type: application/json
Date: Sun, 23 May 2021 21:29:17 GMT

{
    "message": "👋"
}
```

## Dependencies

```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

