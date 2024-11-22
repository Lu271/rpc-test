namespace go hello

struct HelloRequest {
    1: string message
}

struct HelloResponse {
    1: string message
}

service HelloService {
    HelloResponse SayHello(HelloRequest req)
}