package main

import (
	"calculator/Auxiliary"
	"fmt"
	"strconv"
	"strings"
)

var calculator Auxiliary.Calculator = Auxiliary.CalculatorConstructor()
var inputGetter Auxiliary.Specifications = Auxiliary.SpecificationsConstructor()

func main() {
	var answerType Auxiliary.Answer
	var expression []string

	// loop until user types /exit
	for answerType != Auxiliary.Exit {
		answerType, expression = inputGetter.GetValidInput()
		if answerType == Auxiliary.Empty {
			continue
		}
		if answerType == Auxiliary.Help {
			fmt.Println("The program calculates the sum and subtraction of expression")
		}
		if answerType == Auxiliary.OneEquals {
			equalsIndex := Auxiliary.GetIndexOfValue(expression, "=")
			expressionLeft := expression[:equalsIndex]
			expressionRight := expression[equalsIndex+1:]
			variableName := strings.Join(expressionLeft, "")
			leftOK := Auxiliary.CheckValidVariable(variableName)
			if !leftOK {
				fmt.Println("Invalid Identifier")
				continue
			}
			var result float64
			expressionWithVariablesReplaced, err := Auxiliary.ReplaceVariables(expressionRight, &calculator)
			if err == Auxiliary.UnknownVariable {
				fmt.Println("UnknownVariable")
				continue
			} else if err == Auxiliary.InvalidName {
				fmt.Println("Invalid identifier")
				continue
			} else {
				result = calculator.EvaluateExpression(expressionWithVariablesReplaced)
				calculator.Variables[variableName] = strconv.FormatFloat(result, 'f', -1, 64)
			}

			// process first:
			//   check if all chars are letters
			//	 if yes: create variable
			// process second:
			//   replace variables, compute value, assign to variable
		}
		if answerType == Auxiliary.NoEquals {
			expressionWithVariablesReplaced, err := Auxiliary.ReplaceVariables(expression, &calculator)
			if err == Auxiliary.UnknownVariable {
				fmt.Println("UnknownVariable")
			} else if err == Auxiliary.InvalidName {
				fmt.Println("Invalid identifier")
			} else {
				result := calculator.EvaluateExpression(expressionWithVariablesReplaced)
				fmt.Println(result)
			}
		}
	}
	fmt.Println("Bye!")
}
