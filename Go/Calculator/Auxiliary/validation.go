package Auxiliary

import (
	"strconv"
	"strings"
	"unicode"
)

type CommandOrExpression int

const (
	Command = iota
	Assignment
	Expression
)

func SpecificationsConstructor() Specifications {
	inputGetter := Specifications{
		ValidCommands: map[string]bool{
			"/help": true,
			"/exit": true,
		},
		ValidSymbols: map[string]bool{
			"/": true,
			"*": true,
			"^": true,
			"=": true,
			"(": true,
			")": true,
			"+": true,
			"-": true,
			"0": true,
			"1": true,
			"2": true,
			"3": true,
			"4": true,
			"5": true,
			"6": true,
			"7": true,
			"8": true,
			"9": true,
		},
		Operators: map[string]bool{
			"+": true,
			"-": true,
			"=": true,
			"*": true,
			"/": true,
			"^": true,
		},
		OperatorsPriority: map[string]int{
			"+": 1,
			"-": 1,
			"*": 2,
			"/": 2,
			"^": 3,
		},
	}
	return inputGetter
}

func (specifications *Specifications) ValidateInput(input string) (CommandOrExpression, bool) {
	// empty string is valid
	if input == "" {
		return Expression, true
	}
	input = RemoveAllSpaces(input)
	inputFirstRune, inputLastRune := GetFirstAndLastRuneFromString(input)

	// is command
	if inputFirstRune == '/' {
		_, ok := specifications.ValidCommands[input]
		if ok {
			return Command, true
		} else {
			return Command, false
		}
	}

	// if input ends with operator, it's invalid
	_, ok := specifications.Operators[string(inputLastRune)]
	if ok {
		return Expression, false
	}

	numberOfEquals := strings.Count(input, "=")

	// more than one '='
	if numberOfEquals > 1 {
		return Assignment, false
	}

	// validate parenthesis
	if specifications.InvalidParenthesis(input) {
		return Expression, false
	}
	// two operands in sequence that are not '+' and '-' is invalid
	if specifications.HasTwoOperandsNextToEachOther(input) {
		return Expression, false
	}
	// is expression (consists of many symbols)
	inputAsRunes := []rune(input)
	for _, ch := range inputAsRunes {
		ok := specifications.ValidSymbols[string(ch)] || unicode.IsLetter(ch)
		if !ok {
			return Expression, false
		}
	}
	return Expression, true
}

func IsSumSubSimplified(s string) bool {
	for _, duplicated := range []string{"+-", "-+", "++", "--"} {
		if strings.Contains(s, duplicated) {
			return false
		}
	}
	return true
}

func CheckIfSymbolHasLetters(symbol string) bool {
	for _, ch := range symbol {
		if unicode.IsLetter(ch) {
			return true
		}
	}
	return false
}

func CheckValidVariable(symbol string) bool {
	for _, ch := range symbol {
		if !unicode.IsLetter(ch) {
			return false
		}
	}
	return true
}

func GetIndexOfValue(stringSlice []string, value string) int {
	for i, curr := range stringSlice {
		if curr == value {
			return i
		}
	}
	return -1
}

func (specifications *Specifications) InvalidParenthesis(input string) bool {
	numberOfOpening := strings.Count(input, "(")
	numberOfClosing := strings.Count(input, ")")
	inputAsRunes := []rune(input)
	// if it has no parenthesis, it's valid
	if numberOfOpening == 0 && numberOfClosing == 0 {
		return false
	}
	// if the number of "(" differs from ")", it's invalid
	if numberOfOpening != numberOfClosing {
		return true
	}
	// if there are "[operand])" or "[operand](", it's invalid
	for index := 0; index < len(inputAsRunes)-1; index++ {
		firstCh := inputAsRunes[index]
		secondCh := inputAsRunes[index+1]
		if firstCh == '=' {
			if secondCh == '(' {
				continue
			} else {
				return true
			}
		}
		_, firstChIsOperator := specifications.Operators[string(firstCh)]
		secondChIsParenthesis := secondCh == ')'
		if firstChIsOperator && secondChIsParenthesis {
			return true
		}
	}
	// check the if they are in the correct order
	stack := Stack[rune]{}
	for _, ch := range inputAsRunes {
		if ch == '(' {
			stack.Push(ch)
		}
		if ch == ')' {
			if stack.IsEmpty() {
				return true
			} else {
				topElement := stack.Peck()
				if topElement != '(' {
					return true
				} else if topElement == '(' {
					stack.Pop()
				}
			}
		}
	}
	if stack.IsEmpty() {
		return false
	}
	return true
}

func (specifications *Specifications) HasHigherPrecedence(operator1 string, operator2 string) bool {
	return specifications.OperatorsPriority[operator1] > specifications.OperatorsPriority[operator2]
}

// HasTwoOperandsNextToEachOther checks for invalid input like "5+*6" or "10//4", but allows "5++--4"
func (specifications *Specifications) HasTwoOperandsNextToEachOther(input string) bool {
	stringAsRunes := []rune(input)
	if len(stringAsRunes) == 1 {
		return false
	}
	for index := 0; index < len(stringAsRunes)-1; index++ {
		firstCh := stringAsRunes[index]
		secondCh := stringAsRunes[index+1]
		_, firstIsOperand := specifications.Operators[string(firstCh)]
		_, secondIsOperand := specifications.Operators[string(secondCh)]
		firstIsPlusOrMinus := firstCh == '+' || firstCh == '-'
		secondIsPlusOrMinus := secondCh == '+' || secondCh == '-'
		bothAreOperands := firstIsOperand && secondIsOperand
		bothArePlusOrMinus := firstIsPlusOrMinus && secondIsPlusOrMinus
		if bothAreOperands && !bothArePlusOrMinus {
			return true
		}
	}
	return false
}

func (specifications *Specifications) IsNumber(s string) bool {
	_, isNotNumber := strconv.ParseFloat(s, 64)
	if isNotNumber != nil {
		return false
	}
	return true
}
