package main

import (
	"fmt"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

const dag_rules = `
rule "1"
begin
	print(1, ",")
end
rule "2"
	begin
print(2, ",")
end
rule "3"
begin
	print(3, ",")
end
rule "4"
begin
	print(4, ",")
end
rule "5"
begin
	print(5, ",")
end
rule "6"
begin
	print(6, ",")
end
rule "7"
	begin
print(7, ",")
end
rule "8"
begin
	print(8, ",")
end
rule "9"
begin
	print(9, ",")
end
rule "10"
begin
	print(10, ",")
end
rule "11"
begin
	print(11, ",")
end
rule "12"
begin
	print(12, ",")
end
`

func main() {

	dataContext := context.NewDataContext()
	dataContext.Add("print", hello)

	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(dag_rules)
	if e1 != nil {
		panic(e1)
	}

	gengine := engine.NewGengine()
	names := makeDAG()

	e := gengine.ExecuteDAGModel(ruleBuilder, names)
	if e != nil {
		panic(e)
	}
}

func makeDAG() [][]string {
	//base
	names := make([][]string, 5)

	//第1列(层)
	namesCol1 := make([]string, 3)
	namesCol1[0] = "1"
	namesCol1[1] = "2"
	namesCol1[2] = "3"
	names[1] = namesCol1

	//第2列(层)
	namesCol2 := make([]string, 1)
	namesCol2[0] = "4"
	names[2] = namesCol2

	//第3列(层)
	namesCol3 := make([]string, 5)
	namesCol3[0] = "5"
	namesCol3[1] = "6"
	namesCol3[2] = "7"

	//add the rules not exist
	namesCol3[3] = "100"
	namesCol3[4] = "200"
	names[3] = namesCol3

	//第4列(层)
	namesCol4 := make([]string, 5)
	namesCol4[0] = "8"
	namesCol4[1] = "9"
	namesCol4[2] = "10"
	namesCol4[3] = "11"
	namesCol4[4] = "12"
	names[4] = namesCol4
	return names
}

func hello(x int, y string) {
	fmt.Printf("workflow %v \n", x)
}
