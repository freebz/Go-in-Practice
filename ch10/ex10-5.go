// Listing 10.5  Encode an instance to JSON with codec

jh := new(codec.JsonHandle)
u := &user.User{
	Name:  "Inigo Montoya",
	Email: "inigo@montoya.example.com",
}

var out []byte
err := codec.NewEncoderBytes(&out, jh).Encode(&u)
if err != nil {
	...
}

fmt.Println(string(out))
