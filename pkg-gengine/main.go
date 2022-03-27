package main

import (
	"fmt"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
	"ruleng/define"
)
// 文档 https://github.com/bilibili/gengine/wiki/
//定义规则
const rule1 = `
rule "TotalAmount"
begin
		displayTotal(req.GetTotalAmount())
		if req.GetTotalAmount() < 0{
			return false
		}
        return true
end
`
func displayTotal(total int64)string{
	return fmt.Sprintf("total is %v \n",total)
}
func main()  {
	req := &define.CreateOrderRequest{
		TotalAmount: 1,
		FinalPrice: 0,
		AdminFee: -1,
	}
	dataContext := context.NewDataContext()
	//注入初始化的结构体
	dataContext.Add("req", req)
	// 注入函数
	dataContext.Add("displayTotal",displayTotal)
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	err := ruleBuilder.BuildRuleFromString(rule1)
	if err != nil{
		fmt.Println("err:%s ", err)
	}else{
		eng := engine.NewGengine()
		_ = eng.Execute(ruleBuilder,true)
		result,_ := eng.GetRulesResultMap()
		r := result["TotalAmount"]
		if r == nil{
			fmt.Println("r is nil")
		}else {
			fmt.Printf("Total amount check result is %v\n",r.(bool))
		}

	}
}

