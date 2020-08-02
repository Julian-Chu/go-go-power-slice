package bytes_vs_string

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
)




func Benchmark_removeByCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeByCopy(str)
	}
}

func Benchmark_removeByBytesReplaceAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeByBytesReplaceAll(str)
	}
}
func Benchmark_removeByStringsReplaceAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeByStringsReplaceAll(str)
	}
}


func Benchmark_minifyPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minifyPackage()
	}
}


var str = `
{
   "message" : {
      "query" : "abc",
      "operator":"and"
   }
}`
//==>  `{"message":{"query":"abc","operator":"and"}}`

func removeByStringsReplaceAll(str string) string {
	for _, s := range []string{" ", "\t", "\n"} {
		str = strings.ReplaceAll(str, s, "")
	}
	return str
}

func removeByBytesReplaceAll(str string) string {
	b := []byte(str)
	for _, s := range []string{" ", "\t", "\n"} {
		b = bytes.ReplaceAll(b, []byte(s), nil)
	}
	return string(b)
}

func removeByCopy(str string) string {
	b:= []byte(str)
	cnt := 0
	for i := len(b) - 1; i >= 0; i-- {
		switch b[i] {
		case '\t', '\n', ' ':
			copy(b[i:], b[i+1:])
			cnt++
		}
	}
	b = b[:len(b)-cnt]
	return string(b)
}

func minifyPackage() string {
	b := []byte(`
{
   "message" : {
      "query" : "this is a test",
      "operator":"and"
   }
}`)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	b, _ = m.Bytes("application/json", b)

	return string(b)
}
/*
Benchmark_removeByCopy-8                3653143   323 ns/op    128 B/op   2 allocs/op
Benchmark_removeByBytesReplaceAll-8     1378482   851 ns/op    304 B/op   5 allocs/op
Benchmark_removeByStringsReplaceAll-8   1620444   743 ns/op    224 B/op   4 allocs/op
*/
