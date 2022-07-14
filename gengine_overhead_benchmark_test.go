package go_rule_eng_research

import (
	"testing"

	"github.com/bilibili/gengine/engine"
)

// 引入规则引擎造成的性能损耗

/*
✦2 ❯ go test -bench=. -benchmem -test.bench BenchmarkX -cpuprofile cpuprofile.out
goos: darwin
goarch: amd64
pkg: ruleng
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkXGo-16                 554308380                2.050 ns/op           0 B/op          0 allocs/op
BenchmarkXGengine-16              346998              3194 ns/op             816 B/op         18 allocs/op
BenchmarkXGengineInGoFnc-16       599295              1876 ns/op             560 B/op          9 allocs/op
PASS
ok      ruleng  5.343s

*/

type ReqStruct struct {
	A string
	B string
	C *ReqStruct
}

func newReq() *ReqStruct {
	req := &ReqStruct{
		A: "A",
		B: "B",
		C: &ReqStruct{"A", "B", nil},
	}
	return req
}

func BenchmarkXGo(b *testing.B) {
	req := newReq()
	b.ResetTimer() // 重置计时器，忽略前面的准备时间
	for n := 0; n < b.N; n++ {
		ValidateByGo(req)
	}
}
func ValidateByGo(req *ReqStruct) bool {
	if req.A != "A" {
		return false
	}
	if req.B != "B" {
		return false
	}
	if req.C.A != "A" {
		return false
	}
	if req.C.B != "B" {
		return false
	}
	return true
}

const rulereq = `
rule "rule1"
begin
      if req.A != "A"{return "A"}
      return ""
end

rule "rule2"
begin
      if req.B != "B"{return "B"}
      return ""
end

rule "rule3"
begin
      if req.C.A != "A"{return "CA"}
      return ""
end

rule "rule4"
begin
      if req.C.B != "B"{return "CB"}
      return ""
end
`

var pool = &engine.GenginePool{}

func BenchmarkXGengine(b *testing.B) {
	req := newReq()
	var err error
	pool, err = engine.NewGenginePool(10, 100, 2, rulereq, make(map[string]interface{}))
	if err != nil {
		b.Fatal(err)
	}
	datamap := map[string]interface{}{"req": req}

	b.ResetTimer() // 重置计时器，忽略前面的准备时间
	for n := 0; n < b.N; n++ {
		_ = ValidateByGengine(datamap, b)
	}
}

const rulereq2 = `
rule "rule1"
begin
    ValidateByGo(req)
    return ""
end
`

func ValidateByGengine(datamap map[string]interface{}, b *testing.B) bool {
	_, res := pool.Execute(datamap, true)
	for _, v := range res {
		if v != "" {
			return false
		}
	}
	return true
}

func ValidateByGengineInGoFnc(datamap map[string]interface{}, b *testing.B) bool {
	_, res := pool.Execute(datamap, true)
	for _, v := range res {
		if v != false {
			return false
		}
	}
	return true
}

func BenchmarkXGengineInGoFnc(b *testing.B) {
	req := newReq()
	var err error
	pool, err = engine.NewGenginePool(10, 100, 2, rulereq2, make(map[string]interface{}))
	if err != nil {
		b.Fatal(err)
	}
	datamap := map[string]interface{}{
		"req":          req,
		"ValidateByGo": ValidateByGo,
	}

	b.ResetTimer() // 重置计时器，忽略前面的准备时间
	for n := 0; n < b.N; n++ {
		_ = ValidateByGengineInGoFnc(datamap, b)
	}
}
