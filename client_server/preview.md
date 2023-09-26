

GRPC
    One component (the client) calls or invokes specific functions in another software component (the server).
    Default format for gRPC payloads is Protobuf,
RPC
HTTP
REST

gRPC and RPC    
    RPC is a generic protocol for remote procedure calls, while gRPC is a specific implementation of RPC that uses the HTTP/2 protocol for communication. 
    Also, RPC uses a binary encoding format to transmit data, while gRPC supports several serialization formats, including Protocol Buffers, JSON, and XML.

gRPC vs HTTP
    HTTP API requests are sent as text and can be read and created by humans. gRPC messages are encoded with Protobuf by default. While Protobuf is efficient to send and receive, its binary format isn't human readable

RPC vs HTTP
    In Remote Procedure Call (RPC), the client uses HTTP POST to call a specific function by name. 
    Client-side developers must know the function name and parameters in advance for RPC to work. In REST, clients and servers use HTTP verbs like GET, POST, PATCH, PUT, DELETE, and OPTIONS to perform options.