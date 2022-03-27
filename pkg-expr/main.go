package main

import (
	"fmt"

	"github.com/antonmedv/expr"
	"ruleng/define"
)
func displayTotal(total int64)string{
	return fmt.Sprintf("total is %v \n",total)
}

const rule = `
       req.TotalAmount>0
    && req.FinalPrice>0
    && req.AdminFee>0`
func main() {
	req := &define.CreateOrderRequest{
		TotalAmount: 1,
		FinalPrice: 0,
		AdminFee: -1,
	}
	env := map[string]interface{}{
		"req":   req,
		"displayTotal": displayTotal, // You can pass any functions.
	}


	// Compile code into bytecode. This step can be done once and program may be reused.
	// Specify environment for type check.
	program, err := expr.Compile(rule, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}