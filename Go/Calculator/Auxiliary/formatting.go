package Auxiliary

import (
	"log"
	"strings"
)

type InvalidNameOrUnknownVariable int

const (
	InvalidName InvalidNameOrUnknownVariable = iota
	UnknownVariable
	OK
)

type Specifications struct {
	ValidCommands     map[string]bool
	ValidSymbols      map[string]bool
	Operators         map[string]bool
	OperatorsPriority map[string]int
}

var specifications Specifications = SpecificationsConstructor()

func RemoveAllSpaces(inputString string) string {
	return strings.ReplaceAll(inputString, " ", "")
}

// RemoveSpacesAndSplitOnOperatorsAndParenthesis removes duplicated operators and spaces
// return the input split on the operators and parenthesis
func (specifications *Specifications) RemoveSpacesAndSplitOnOperatorsAndParenthesis(inputString string) []string {
	inputStringWithoutSpaces := RemoveAllSpaces(inputString)
	for !IsSumSubSimplified(inputStringWithoutSpaces) {
		inputStringWithoutSpaces = strings.ReplaceAll(inputStringWithoutSpaces, "+-", "-")
		inputStringWithoutSpaces = strings.ReplaceAll(inputStringWithoutSpaces, "-+", "-")
		inputStringWithoutSpaces = strings.ReplaceAll(inputStringWithoutSpaces, "--", "+")
		inputStringWithoutSpaces = strings.ReplaceAll(inputStringWithoutSpaces, "++", "+")
	}
	simplifiedReadyToSplit := specifications.AddSpacesToSplitBy(inputStringWithoutSpaces)
	splitInput := strings.Split(simplifiedReadyToSplit, " ")
	return splitInput
}

func (specifications *Specifications) AddSpacesToSplitBy(s string) string {
	operatorAndParenthesis := make(map[string]bool)
	for key, value := range specifications.Operators {
		operatorAndParenthesis[key] = value
	}
	operatorAndParenthesis[")"] = true
	operatorAndParenthesis["("] = true
	for operator, _ := range operatorAndParenthesis {
		s = strings.ReplaceAll(s, operator, " "+operator+" ")
	}
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.TrimSpace(s)
	return s
}

func ReplaceVariables(expression []string, c *Calculator) ([]string, InvalidNameOrUnknownVariable) {
	var symbol string

	for i := 0; i < len(expression); i++ {
		symbol = expression[i]
		// is the symbol a variable?
		if CheckIfSymbolHasLetters(symbol) {
			// does it have a valid variable name?
			if CheckValidVariable(symbol) {
				// is it a known variable?
				value, ok := c.Variables[symbol]
				if !ok {
					return []string{}, UnknownVariable
				} else {
					expression[i] = value
				}
			} else {
				return []string{}, InvalidName
			}
		}
	}
	return expression, OK
}

func (c *Calculator) ToPostFix(expression []string) []string {
	var stack Stack[string]
	var resultSlice []string
	for i := 0; i < len(expression); i++ {
		symbol := expression[i]
		if symbol != "(" && symbol != ")" {
			if _, ok := specifications.Operators[symbol]; !ok && symbol != "(" && symbol != ")" {
				resultSlice = append(resultSlice, symbol)
			} else if stack.IsEmpty() || stack.Peck() == "(" {
				stack.Push(symbol)
			} else if specifications.HasHigherPrecedence(symbol, stack.Peck()) {
				stack.Push(symbol)
			} else {
				for !stack.IsEmpty() && stack.Peck() != "(" && !specifications.HasHigherPrecedence(symbol, stack.Peck()) {
					nextElement, err := stack.Pop()
					if err != nil {
						log.Fatalln(err)
					}
					resultSlice = append(resultSlice, nextElement)
				}
				stack.Push(symbol)
			}
		}
		if symbol == "(" {
			stack.Push(symbol)
		} else if symbol == ")" {
			for stack.Peck() != "(" {
				nextElement, err := stack.Pop()
				if err != nil {
					log.Fatalln(err)
				}
				resultSlice = append(resultSlice, nextElement)
			}
			stack.Pop()
		}
	}
	for !stack.IsEmpty() {
		nextElement, err := stack.Pop()
		if err != nil {
			log.Fatalln(err)
		}
		resultSlice = append(resultSlice, nextElement)
	}
	return resultSlice
}
