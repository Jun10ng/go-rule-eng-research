package go_rule_eng_research

import (
	"testing"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

// 引入规则引擎造成的性能损耗

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

func BenchmarkGo(b *testing.B) {
	req := newReq()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ValidateByGo(req)
	}
	b.StopTimer()
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
      if req.B != "B"{return "B"}
      if req.C.A != "A"{return "CA"}
      if req.C.B != "B"{return "CB"}
      return ""
end
`

var eng = engine.NewGengine()

func BenchmarkGengine(b *testing.B) {
	req := newReq()

	dataContext := context.NewDataContext()
	dataContext.Add("req", req)

	ruleBuilder := builder.NewRuleBuilder(dataContext)

	err := ruleBuilder.BuildRuleFromString(rulereq)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = ValidateByGengine(ruleBuilder)
	}
	b.StopTimer()
}

func ValidateByGengine(ruleBuilder *builder.RuleBuilder) bool {
	err := eng.Execute(ruleBuilder, false)
	if err != nil {
		return false
	}
	res, _ := eng.GetRulesResultMap()
	for _, v := range res {
		if v != "" {
			return false
		}
	}
	return true
}
