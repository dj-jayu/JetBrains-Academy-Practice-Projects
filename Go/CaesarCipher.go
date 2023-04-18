package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func calculate(base, mod, times int) int {
	// g**b % p
	var c int = 1
	for i := 0; i < times; i++ {
		c = (c * base) % mod
	}
	return c
}
func getBrokenInput() []string {
	var input string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	// var broken_input []string
	broken_input := strings.Fields(input)
	return broken_input
}
func getGP() (int, int) {
	var broken_input []string = getBrokenInput()
	g := broken_input[2]
	p := broken_input[6]
	gInt, _ := strconv.Atoi(g)
	pInt, _ := strconv.Atoi(p)
	return gInt, pInt
}
func getA() int {
	var broken_input []string = getBrokenInput()
	var AStr string = broken_input[2]
	A, _ := strconv.Atoi(AStr)
	return A
}
func encript(str string, secret int) string {
	var builder strings.Builder
	var minAscii int
	validRange := 26
	for _, c := range str {
		if !unicode.IsLetter(c) {
			builder.WriteRune(c)
			continue
		}
		if unicode.IsUpper(c) {
			minAscii = 65
		} else {
			minAscii = 97
		}
		// fmt.Println(string(rune(minAscii)), "c:", c, "int(c):", int(c), "secret:", secret, "secret%validRange:", secret%validRange, "int(c)+secret", int(c)+secret, "(int(c)+secret)%validRange:", (int(c)+secret)%validRange)
		encriptedC := minAscii + (-minAscii+int(c)+int(secret))%(validRange)
		builder.WriteRune(rune(encriptedC))
	}
	return builder.String()
}
func decript(str string, secret int) string {
	var builder strings.Builder
	var minAscii int
	validRange := 26
	secret = -(secret%validRange - validRange)
	for _, c := range str {
		if !unicode.IsLetter(c) {
			builder.WriteRune(c)
			continue
		}
		if unicode.IsUpper(c) {
			minAscii = 65
		} else {
			minAscii = 97
		}
		// fmt.Println(string(rune(minAscii)), "c:", c, "int(c):", int(c), "int(secret):", int(secret), "int(c)-secret", int(c)-secret, "(int(c)-secret)%validRange:", (int(c)-secret)%validRange)
		decriptedC := minAscii + (-minAscii+int(c)+int(secret))%(validRange)
		builder.WriteRune(rune(decriptedC))
	}
	return builder.String()
}
func sendProposalGetAnswer(firstEncriptedString string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(firstEncriptedString)
	scanner.Scan()
	return scanner.Text()
}
func getFinalAnswer(decriptedAnswer string) string {
	switch decriptedAnswer {
	case "Yeah, okay!":
		return "Great!"
	case "Let's be friends.":
		return "What a pity!"
	}
	return "noAnswer"
}
func main() {
	var base1, mod int = getGP()
	fmt.Println("OK")
	var A int = getA()
	// var b int = rand.Intn(mod)
	var b = 7
	var B int = calculate(base1, mod, b)
	var S int = calculate(A, mod, b)
	fmt.Printf("B is %d\n", B)
	proposal := "Will you marry me?"
	encriptedProposal := encript(proposal, S)
	aliceEncriptedAnswer := sendProposalGetAnswer(encriptedProposal)
	aliceDecriptedAnswer := decript(aliceEncriptedAnswer, S)
	finalAnswer := getFinalAnswer(aliceDecriptedAnswer)
	if finalAnswer != "noAnswer" {
		finalEncriptedAnswer := encript(finalAnswer, S)
		fmt.Println(finalEncriptedAnswer)
	}

	// var input string
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// var broken []string = strings.Split(scanner.Text(), ",")
	// secretInt, _ := strconv.Atoi(broken[1])
	// input = broken[0]
	// fmt.Println(decript(input, secretInt))
}
