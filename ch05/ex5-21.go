// Listing 5.21  Benchmark template compile and run

package main

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tpl)
		t.Execute(&buf, data)
		buf.Reset()
	}
}
