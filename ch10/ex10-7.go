// Listing 10.7  Protocol buffer file

package chapter10;

message User {
	required string name = 1;
	required int32 id = 2;
	optional string email = 3;
}
