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




<img width="646" height="432" alt="image" src="https://github.com/user-attachments/assets/d92e22cb-aa1a-4f04-8cc5-67081098dd13" />

<img width="714" height="311" alt="image" src="https://github.com/user-attachments/assets/bf9f1bf3-42ea-4ffd-bc1b-f1b9774b526b" />
<img width="624" height="605" alt="image" src="https://github.com/user-attachments/assets/4949afbd-276c-4ca4-b530-0c0cc2bb23a8" />
<img width="592" height="553" alt="image" src="https://github.com/user-attachments/assets/1d7557c6-5c28-4277-9b48-63d666af8834" />

🔹 7. Advanced gRPC Concepts

        Deadlines & Timeouts → control request time.
        
        Error Handling → gRPC status codes (codes.NotFound, codes.Unauthenticated, etc.).
        
        Interceptors → middleware (logging, authentication).
        
        TLS/SSL → secure communication.
        
        Load Balancing → gRPC supports client-side LB.
        
        Reflection → debugging with tools like grpcurl.

🔹 8. Summary : 
        UnaryHello → single request & response.
        
        ServerStreamHello → one request, multiple responses.
        
        ClientStreamHello → multiple requests, one response.
        
        BidiHello → continuous two-way communication.
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Interceptors (middleware in gRPC) next (for logging, authentication, retry etc.) or should we first go deeper into deadlines, timeouts & error handling

🔹 1. Interceptors in gRPC (Middleware)

👉 Interceptors in gRPC are similar to middlewares in Gin.
    They allow you to:

        Log requests & responses
        
        Add authentication / authorization
        
        Retry policies
        
        Metrics & monitoring

    There are two types:

        Unary Interceptor → For Unary RPCs (one request, one response).
        
        Stream Interceptor → For Streaming RPCs (client, server, or bidi streaming).


<img width="628" height="451" alt="image" src="https://github.com/user-attachments/assets/20bc3112-ff50-4a7b-af69-b36ad2226b28" />
<img width="631" height="531" alt="image" src="https://github.com/user-attachments/assets/42fc03b8-7cf6-4ce2-bb4a-8faffad6ff63" />

🔹 2. Deadlines & Timeouts

👉 gRPC lets clients set deadlines so that requests don’t hang forever.
On the server, you can check context to stop work early.
<img width="653" height="456" alt="image" src="https://github.com/user-attachments/assets/7db18ef3-eabc-4ce5-b93a-cf07640b13a5" />

🔹 3. Error Handling in gRPC

    gRPC has structured error codes (like HTTP status codes).

    Common ones:

        codes.OK → Success
        
        codes.NotFound → Resource not found
        
        codes.InvalidArgument → Bad input
        
        codes.Unauthenticated → Not logged in
        
        codes.PermissionDenied → Unauthorized
        
        codes.DeadlineExceeded → Timeout
        
        codes.Unavailable → Service down
    


