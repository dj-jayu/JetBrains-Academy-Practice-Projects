package Auxiliary

import (
	"log"
	"math"
	"strconv"
)

type NumberOrOperator int

const (
	Number NumberOrOperator = iota
	Operator
)

var calculator Calculator = CalculatorConstructor()

type Calculator struct {
	Variables map[string]string
}

func CalculatorConstructor() Calculator {
	return Calculator{Variables: make(map[string]string, 0)}
}

func (c *Calculator) EvaluateExpression(expression []string) float64 {
	expressionPostFix := c.ToPostFix(expression)
	result := c.EvaluatePostFix(expressionPostFix)
	return result
}

func (c *Calculator) GetType(s string) NumberOrOperator {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return Operator
	}
	return Number
}

func (c *Calculator) EvaluatePostFix(postFixSlice []string) float64 {
	var stack Stack[string]
	for _, symbol := range postFixSlice {
		if specifications.IsNumber(symbol) {
			stack.Push(symbol)
		} else {
			operand2, _ := stack.Pop()
			operand1, _ := stack.Pop()
			result := c.EvaluateSimpleOperation(operand1, symbol, operand2)
			stack.Push(result)
		}
	}
	resultAsString := stack.Peck()
	resultAsFloat, _ := strconv.ParseFloat(resultAsString, 64)
	return resultAsFloat
}

func (c *Calculator) EvaluateSimpleOperation(operand1 string, operator string, operand2 string) string {
	operand1Float, err := strconv.ParseFloat(operand1, 64)
	if err != nil {
		log.Fatalln(err)
	}
	operand2Float, err := strconv.ParseFloat(operand2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	var total float64
	switch operator {
	case "+":
		total = operand1Float + operand2Float
	case "-":
		total = operand1Float - operand2Float
	case "*":
		total = operand1Float * operand2Float
	case "/":
		total = operand1Float / operand2Float
	case "^":
		total = math.Pow(operand1Float, operand2Float)
	}
	totalString := strconv.FormatFloat(total, 'f', -1, 64)
	return totalString
}
