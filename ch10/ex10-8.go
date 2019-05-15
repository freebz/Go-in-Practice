// Listing 10.8  Protocol buffer server setup

import (
	"net/http"

	pb "github.com/Masterminds/go-in-practice/chapter10/userpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
