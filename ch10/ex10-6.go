// Listing 10.6  Decode JSON into an instance of a type

var u2 user.User
err = codec.NewDecoderBytes(out, jh).Decode(&u2)
if err != nil {
	...
}

fmt.Println(u2)
