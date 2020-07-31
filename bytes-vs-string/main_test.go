package bytes_vs_string

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
)

func Benchmark_minifyByBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minifyByBytes()
	}
}

func Benchmark_minifyStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minifyStr()
	}
}

func Benchmark_minifyPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minifyPackage()
	}
}
func minifyByBytes() string {
	b := []byte(`
{
   "message" : {
      "query" : "this is a test",
      "operator":"and"
   }
}`)

	insideStr := false
	cnt := 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == '"' {
			insideStr = !insideStr
		}

		if insideStr {
			continue
		}

		switch b[i] {
		case '\t', '\n', ' ':
			copy(b[i:], b[i+1:])
			cnt++
		}
	}
	b = b[:len(b)-cnt]
	return string(b)
}
func minifyStr() string {
	b := []byte(`
{
   "message" : {
      "query" : "this is a test",
      "operator":"and"
   }
}`)

	newb := make([]byte, len(b))
	copy(newb, b)

	for _, s := range []string{" ", "\t", "\n"} {
		newb = bytes.ReplaceAll(newb, []byte(s), []byte(""))
	}
	return string(newb)
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
