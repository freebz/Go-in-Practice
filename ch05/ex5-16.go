// Listing 5.16  Asserting a writer

func main() {
	m := map[string]interface(){
		"w": &MyWriter(),
	}
}

func doSomething(m map[string]interface()) {
	w := m["w"].(io.Writer)
}
