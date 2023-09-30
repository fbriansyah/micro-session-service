# Micro Auth Service
Aplikasi ini merupakan service untuk melakukan authentication.

## GRPC Service
```proto3
syntax = "proto3";

package payment;


import "proto/session/type/session.proto";

option go_package = 
  "github.com/fbriansyah/micro-payment-proto/protogen/go/auth";

message LoginRequest {
  string username=1 [json_name="username"];
  string password=2 [json_name="password"];
}

message LoginResponse {
  string userid=1 [json_name="userid"];
  string name=2 [json_name="name"];
  session.Session session=3 [json_name="session"];
}

message CreateUserRequest {
  string username=1 [json_name="username"];
  string password=2 [json_name="password"];
  string name=3 [json_name="name"];
}

message CreateUserResponse {
  string userid=1 [json_name="userid"];
  string username=2 [json_name="username"];
  string name=3 [json_name="name"];
}

service AuthService {
  rpc Login (LoginRequest) returns (LoginResponse);
  rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
}
```