🔹 1. What is gRPC?

    gRPC (Google Remote Procedure Call) is a high-performance, open-source RPC framework.
    
    It uses Protocol Buffers (Protobuf) as the interface definition language.
    
    Communication is over HTTP/2 (faster, multiplexing, streaming).
    
    Supports many languages (Go, Java, Python, etc.).

🔹 2. Why use gRPC?

    Strongly typed contracts (proto files).
    
    Faster than REST (binary over HTTP/2 instead of JSON over HTTP/1.1).
    
    Built-in support for streaming.
    
    Easy code generation (client & server stubs auto-generated).
    
    Great for microservices communication.

🔹 3. gRPC Communication Types

    gRPC supports 4 types of RPC calls:
    
    Unary RPC → simple request/response.
    Example: GetUser(id) → User.
    
    Server Streaming RPC → client sends 1 request, server sends multiple responses.
    Example: ListUsers() → stream of User.
    
    Client Streaming RPC → client sends multiple requests, server responds once.
    Example: UploadLogs(stream Log) → UploadStatus.
    
    Bidirectional Streaming RPC → client & server both stream messages.
    Example: Chat(stream Message) ↔ stream Message.




    
<img width="729" height="377" alt="image" src="https://github.com/user-attachments/assets/60d02776-38c0-45fa-afcd-70aa9b19319d" />
