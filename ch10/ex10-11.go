// Listing 10.11  Define messages and RPC calls with protocol buffer

syntax = "proto3";

package chapter10;

service Hello {
	rpc Say (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
	string name = 1;
}

message HelloResponse {
	string message = 1;
}
