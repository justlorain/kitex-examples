namespace go api

struct EchoRequest {
  1: string message (vt.min_size = "0") // length of Message should be greater than or equal to 0
}

struct EchoResponse {
  1: string message
}

service EchoService {
    EchoResponse echo(1: EchoRequest req)
}
