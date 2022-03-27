package go_rule_eng_research

import (
	"testing"

	"github.com/Knetic/govaluate"
	"github.com/antonmedv/expr"
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)
func createParams() map[string]interface{} {
	params := make(map[string]interface{})
	params["Origin"] = "MOW"
	params["Country"] = "RU"
	params["Adults"] = 1
	params["Value"] = 100
	return params
}
type Params struct {
	Origin  string
	Country string
	Value   int
	Adults  int
}
const rule = `(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`



func Benchmark_govaluate(b *testing.B) {
	params := createParams()

	expression, err := govaluate.NewEvaluableExpression(rule)

	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expression.Evaluate(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
const rule1 = `
rule "rule1"
begin
      return (Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)
end
`
func Benchmark_gengine(b *testing.B)  {
	params := createParams()
	dataContext := context.NewDataContext()
	for k,v := range params{
		dataContext.Add(k,v)
	}
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	err := ruleBuilder.BuildRuleFromString(rule1)
	if err != nil {
		b.Fatal(err)
	}
	eng := engine.NewGengine()
	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = eng.Execute(ruleBuilder,true)

	}
	b.StopTimer()
	res,_ := eng.GetRulesResultMap()
	out = res["rule1"]
	if !out.(bool) {
		b.Fail()
	}
}
func Benchmark_expr(b *testing.B) {
	params := createParams()

	program, err := expr.Compile(rule, expr.Env(params))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

