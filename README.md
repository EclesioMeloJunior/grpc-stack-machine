# grpc-stack-machine

The intention for this project is to learn about gRPC communication and its implementation. Here we built a stack machine that execute a set of instructions and then returns the output.

To execute the gRPC server:

```sh
go run main.go -p 9000
```

Then the server will starts at port 9000. To setup the client and send the instructions to the stack machine you need to run:

```sh
go run client/main.go -serverAddress localhost:9000
```

Immediately the client will send RPC calls to the server runnig on port 9000 and will get the output result.

To run unit test: 

```sh
go test ./... -v
```