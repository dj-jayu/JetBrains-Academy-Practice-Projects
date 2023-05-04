package Auxiliary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Answer int

const (
	OneEquals Answer = iota
	NoEquals
	Empty
	Exit
	Help
)

func (specifications *Specifications) GetValidInput() (Answer, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()
	CommandOrExpression, isInputValid := specifications.ValidateInput(inputString)

	// validate input
	for !isInputValid {
		switch CommandOrExpression {
		case Command:
			fmt.Println("Unknown command")
		case Expression:
			fmt.Println("Invalid Expression")
		case Assignment:
			fmt.Println("Invalid assignment")
		}

		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputString = scanner.Text()
		CommandOrExpression, isInputValid = specifications.ValidateInput(inputString)
	}
	// here, the input is valid
	if inputString == "" {
		return Empty, nil
	}
	if inputString == "/exit" {
		return Exit, nil
	}
	if inputString == "/help" {
		return Help, nil
	}
	equalsNumber := strings.Count(inputString, "=")
	splitInputString := specifications.RemoveSpacesAndSplitOnOperatorsAndParenthesis(inputString)
	if equalsNumber == 0 {
		return NoEquals, splitInputString
	}
	return OneEquals, splitInputString
}

func GetFirstAndLastRuneFromString(s string) (rune, rune) {
	if len(s) == 0 {
		return 0, 0
	}
	var sFirstRune, sLastRune rune
	// get first rune of input to compare with '/'
	sAsRuneSlice := []rune(s)
	sFirstRune = sAsRuneSlice[0]
	sLastRune = sAsRuneSlice[len(sAsRuneSlice)-1]
	return sFirstRune, sLastRune
}
