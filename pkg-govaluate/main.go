package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"ruleng/define"
)

func main()  {
	req := define.CreateOrderRequest{
		TotalAmount: 1,
		FinalPrice: 0,
		AdminFee: -1,
	}
	exp1, _ := govaluate.NewEvaluableExpression("TotalAmount< 0");
	parameters := make(map[string]interface{}, 8)
	parameters["TotalAmount"] = req.TotalAmount;

	/*
		不支持
		parameters["req"] = req
		exp1, _ := govaluate.NewEvaluableExpression("req.TotalAmount< 0");
	*/


	result, _ := exp1.Evaluate(parameters);
	if result.(bool){
		fmt.Println("exp1 err")
	}
}