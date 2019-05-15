// Listing 10.10  Protocol buffer client

res, err := http.Get("http://localhost:8080")
if err != nil {
	...
}
defer res.Body.Close()

b, err = ioutil.ReadAll(res.Body)
if err != nil {
	...
}

var u pb.User
err = proto.Unmarshal(b, &u)
if err != nil {
	...
}

fmt.Println(u.GetName())
fmt.Println(u.GetId())
fmt.Printlnu.GetEmail())
