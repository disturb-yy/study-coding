# Introduction to gRPC

## Overview

在 gRPC 中，客户端应用程序可以直接调用不同机器上的服务器应用程序上的方法，就好像它是本地对象一样，这样可以更容易地创建分布式应用程序和服务。与许多RPC系统一样，gRPC 基于定义服务的思想，指定可以通过参数和返回类型远程调用的方法。在服务器端，服务器实现了这个接口，并运行一个 gRPC 服务器来处理客户端调用。在客户端，客户端有一个 stub （在某些语言中称为客户端），它提供与服务器相同的方法。

![Concept Diagram](https://grpc.io/img/landing-2.svg)

**RPC 调用具体过程**

- Client 通过本地调用，调用 Client Stub
- Client Stub 将参数打包 (Marshalling) 成一个消息，然后发送这个消息
- Client 所在的操作系统将消息发送给 Server
- Server 接收到消息后，将消息传递给 Server Stub
- Server Stub 将消息解包 (Unmarshalling) 得到参数
- Server Stub 调用服务端的子程序（函数），处理后，将最终结果按照相反的步骤返回给 Client



### Working with Protocol Buffers

gRPC 使用 Protocol Buffers 作为**数据通信协议** （类似 JSON），其主要特性有：

- 更快的传输速度：使用 HTTP2
- 跨平台多语言
- 具有良好的扩展性和兼容性
- 基于 IDL (Interface Definition Language) 文件定义服务