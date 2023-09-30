# Micro Session Service
Aplikasi ini merupakan service untuk mengolah data sessian dan generate token.

## GRPC Service
```proto3
syntax = "proto3";

package session;

import "proto/session/type/session.proto";

option go_package = 
  "github.com/fbriansyah/micro-payment-proto/protogen/go/session";

message Session {
  string id=1;
  string user_id=2;
  string access_token=3;
  string refresh_token=4;
  google.type.DateTime access_token_expires_at=5;
  google.type.DateTime refresh_token_expires_at=6;
}

message UserID {
  string user_id=1;
}

message SessionID {
  string session_id=1;
}

message Token {
  string access_token=1;
}

message Payload {
  string user_id=1;
}


service SessionService {
  rpc CreateSession(UserID) returns (Session) {}
  rpc RefreshToken(SessionID) returns (Session) {}
  rpc DeleteSession(SessionID) returns (SessionID) {}
  rpc GetPayloadFromToken(Token) returns (Payload) {}
}
```