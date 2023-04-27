package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type Flashcard struct {
	Front  string
	Back   string
	Errors int
}
type FlashcardDeck struct {
	FlashCardSlice []Flashcard
	Size           int
	FrontMap       map[string]int
	BackMap        map[string]int
	difficultCards []string
	maxRepeats     int
}

func (d *FlashcardDeck) getIndexFromFront(card Flashcard) int {
	return d.FrontMap[card.getFront()]
}

func (d *FlashcardDeck) addAndReplace(card Flashcard) {
	var index int = d.getIndexFromFront(card)
	if d.hasFront(card.getFront()) {
		d.FlashCardSlice[index] = card
		delete(d.BackMap, card.getBack())
		d.addBackToMap(card)
	} else {
		d.addCard(card)
	}
}

func (f *Flashcard) getFront() string {
	return f.Front
}

func (f *Flashcard) getBack() string {
	return f.Back
}

func (f *Flashcard) checkAnswer(answer string) bool {
	return answer == f.getBack()
}

func (f *Flashcard) printFeedback(d *FlashcardDeck, answer string) {
	if f.checkAnswer(answer) {
		fmt.Println("Correct!")
		log = append(log, fmt.Sprintln("Correct!"))
	} else if d.hasBack(answer) {
		fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", f.getBack(), d.getFrontFromBack(answer))
		log = append(log, fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", f.getBack(), d.getFrontFromBack(answer)))
		f.Errors++
	} else {
		fmt.Printf("Wrong. The right answer is \"%s\".\n", f.getBack())
		log = append(log, fmt.Sprintf("Wrong. The right answer is \"%s\".\n", f.getBack()))
		f.Errors++
	}
}

func (d *FlashcardDeck) createNCards(numberOfCardsToCreate int) {
	var front, back string
	var reader = bufio.NewReader(os.Stdin)
	for cardNumber := 1; cardNumber <= numberOfCardsToCreate; cardNumber++ {
		fmt.Println("The card:")
		log = append(log, fmt.Sprintln("The card:"))
		front, _ = reader.ReadString('\n')
		log = append(log, fmt.Sprint(front))
		front = strings.TrimSpace(front)
		front = getOriginalFront(d, &front)
		fmt.Println("The definition of the card:")
		log = append(log, fmt.Sprintln("The definition of the card:"))
		back, _ = reader.ReadString('\n')
		log = append(log, fmt.Sprint(back))
		back = strings.TrimSpace(back)
		back = getOriginalBack(d, &back)
		newCard := FlashcardCreator(front, back)
		d.addCard(newCard)
		fmt.Printf("The pair (\"%s\":\"%s\") has been added.\n", front, back)
		log = append(log, fmt.Sprintf("The pair (\"%s\":\"%s\") has been added.\n", front, back))
	}
}
func (d *FlashcardDeck) practiceNCards(numberOfQuestions int) {
	var answer string
	var reader = bufio.NewReader(os.Stdin)
	var currentFlashCard *Flashcard
	for i := 0; i < numberOfQuestions; i++ {
		currentFlashCard = &d.FlashCardSlice[i%d.Size]
		fmt.Printf("Print the definition of \"%s\":\n", currentFlashCard.getFront())
		log = append(log, fmt.Sprintf("Print the definition of \"%s\":\n", currentFlashCard.getFront()))
		answer, _ = reader.ReadString('\n')
		log = append(log, answer)
		currentFlashCard.printFeedback(d, strings.TrimSpace(answer))
	}
}

func (d *FlashcardDeck) removeCard(front string) {
	index := d.FrontMap[front]
	delete(d.FrontMap, front)
	card := d.getFlashCard(index)
	delete(d.BackMap, card.getBack())
	d.FlashCardSlice = append(d.FlashCardSlice[:index], d.FlashCardSlice[index+1:]...)
	d.updateIndexes()
}
func (d *FlashcardDeck) updateIndexes() {
	for i, card := range d.FlashCardSlice {
		d.FrontMap[card.getFront()] = i
		d.BackMap[card.getBack()] = i
	}
}
func (d *FlashcardDeck) getFrontFromBack(answer string) string {
	cardIndex := d.BackMap[answer]
	var card Flashcard = d.getFlashCard(cardIndex)
	return card.getFront()
}

func (d *FlashcardDeck) addFlashCardToMaps(flashcard Flashcard) {
	d.addFrontToMap(flashcard)
	d.addBackToMap(flashcard)
}

func (d *FlashcardDeck) addFrontToMap(flashcard Flashcard) {
	d.FrontMap[flashcard.getFront()] = d.Size
}

func (d *FlashcardDeck) addBackToMap(flashcard Flashcard) {
	d.BackMap[flashcard.getBack()] = d.Size
}

func (d *FlashcardDeck) addCard(flashcard Flashcard) {
	d.FlashCardSlice = append(d.FlashCardSlice, flashcard)
	d.addFlashCardToMaps(flashcard)
	d.Size += 1
}

func (d *FlashcardDeck) getFlashCard(i int) Flashcard {
	return d.FlashCardSlice[i]
}

func (d *FlashcardDeck) hasFront(front string) bool {
	_, ok := d.FrontMap[front]
	return ok
}

func (d *FlashcardDeck) hasBack(back string) bool {
	_, ok := d.BackMap[back]
	return ok
}

func getOriginalBack(d *FlashcardDeck, back *string) string {
	if !d.hasBack(*back) {
		return *back
	}
	var reader = bufio.NewReader(os.Stdin)
	var newBack string = *back
	for d.hasBack(newBack) {
		fmt.Printf("The definition \"%s\" already exists. Try again:\n", newBack)
		log = append(log, fmt.Sprintf("The definition \"%s\" already exists. Try again:\n", newBack))
		newBack, _ = reader.ReadString('\n')
		log = append(log, fmt.Sprintln(newBack))
		newBack = strings.TrimSpace(newBack)
	}
	return newBack
}

func FlashcardCreator(front, back string) Flashcard {
	return Flashcard{Front: front, Back: back}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getNumberOfCardsToCreate() int {
	var numberOfCardsToCreate int
	fmt.Println("Input the number of cards:")
	log = append(log, fmt.Sprintln("Input the number of cards:"))
	fmt.Scan(&numberOfCardsToCreate)
	log = append(log, fmt.Sprintln(numberOfCardsToCreate))
	return numberOfCardsToCreate
}

func getOriginalFront(d *FlashcardDeck, front *string) string {
	if !d.hasFront(*front) {
		return *front
	}
	var reader = bufio.NewReader(os.Stdin)
	var newFront string = *front
	for d.hasFront(newFront) {
		fmt.Printf("The term \"%s\" already exists. Try again:\n", newFront)
		log = append(log, fmt.Sprintf("The term \"%s\" already exists. Try again:\n", newFront))
		newFront, _ = reader.ReadString('\n')
		log = append(log, fmt.Sprint(newFront))
		newFront = strings.TrimSpace(newFront)
	}
	return newFront
}

func FlashcardDeckCreator() *FlashcardDeck {
	newDeck := FlashcardDeck{}
	newDeck.FrontMap = make(map[string]int)
	newDeck.BackMap = make(map[string]int)
	return &newDeck
}

func addCard(flashcardsDeck *FlashcardDeck) {
	flashcardsDeck.createNCards(1)
}
func removeCard(flashcardsDeck *FlashcardDeck) {
	var front string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which card?")
	log = append(log, fmt.Sprintln("Which card?"))
	scanner.Scan()
	front = scanner.Text()
	log = append(log, fmt.Sprintln(front))
	ok := flashcardsDeck.hasFront(front)
	if ok {
		flashcardsDeck.removeCard(front)
		flashcardsDeck.Size--
		fmt.Println("The card has been removed.")
		log = append(log, fmt.Sprintln("The card has been removed."))
	} else {
		fmt.Printf("Can't remove \"%s\": there is no such card.\n", front)
		log = append(log, fmt.Sprintf("Can't remove \"%s\": there is no such card.\n", front))
	}
}

func importCards(deck *FlashcardDeck, givenFileName string, getFileName bool) {
	var fileName string
	if getFileName {
		fmt.Println("File name:")
		log = append(log, fmt.Sprintln("File name:"))
		fmt.Scanln(&fileName)
		log = append(log, fmt.Sprintln(fileName))
	} else {
		fileName = givenFileName
	}
	var newDeck FlashcardDeck
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("File not found.")
		log = append(log, fmt.Sprintln("File not found."))
		return
	}
	json.Unmarshal(data, &newDeck)
	n := 0
	for _, card := range newDeck.FlashCardSlice {
		n++
		deck.addAndReplace(card)
	}
	fmt.Printf("%d cards have been loaded.\n", n)
	log = append(log, fmt.Sprintf("%d cards have been loaded.\n", n))
}

func exportCards(deck *FlashcardDeck, givenFileName string, getFileName bool) {
	var fileName string
	if getFileName {
		fmt.Println("File name:")
		log = append(log, fmt.Sprintln("File name:"))
		fmt.Scanln(&fileName)
		log = append(log, fmt.Sprintln(fileName))
	} else {
		fileName = givenFileName
	}
	data, _ := json.Marshal(*deck)
	os.WriteFile(fileName, data, 0664)
	fmt.Printf("%d cards have been saved.\n", deck.Size)
	log = append(log, fmt.Sprintf("%d cards have been saved.\n", deck.Size))
}

func ask(deck *FlashcardDeck) {
	var numberOfQuestions int
	if deck.Size == 0 {
		fmt.Println("Please add a card before studying.")
		log = append(log, fmt.Sprintln("Please add a card before studying."))
		return
	}
	fmt.Println("How many times to ask?")
	log = append(log, fmt.Sprintln("How many times to ask?"))
	fmt.Scanln(&numberOfQuestions)
	log = append(log, fmt.Sprintln(numberOfQuestions))
	deck.practiceNCards(numberOfQuestions)
}
func logToFile(deck *FlashcardDeck) {
	var fileName string
	fmt.Println("File name:")
	log = append(log, fmt.Sprintln("File name:"))
	fmt.Scanln(&fileName)
	log = append(log, fmt.Sprintln(fileName))
	file, _ := os.Create(fileName)
	defer file.Close()
	for _, s := range log {
		file.WriteString(s)
	}
	fmt.Println("The log has been saved.")
	log = append(log, fmt.Sprintln("The log has been saved."))
}
func (d *FlashcardDeck) updateMaxDifficultieRepeats() {
	d.maxRepeats = 0
	for _, card := range d.FlashCardSlice {
		d.maxRepeats = int(math.Max(float64(card.Errors), float64(d.maxRepeats)))
	}
}
func (d *FlashcardDeck) updateMaxDifficultieCards() {
	var maxRepeats []string
	for _, card := range d.FlashCardSlice {
		if d.maxRepeats == card.Errors {
			maxRepeats = append(maxRepeats, card.getFront())
		}
	}
	d.difficultCards = maxRepeats
}
func hardestCard(deck *FlashcardDeck) {
	deck.updateMaxDifficultieRepeats()
	deck.updateMaxDifficultieCards()
	var cardName string
	difficultCardsNumber := len(deck.difficultCards)
	if deck.maxRepeats == 0 {
		fmt.Println("There are no cards with errors.")
		log = append(log, fmt.Sprintln("There are no cards with errors."))
	} else if difficultCardsNumber == 1 {
		cardName = deck.difficultCards[0]
		cardIndex := deck.FrontMap[cardName]
		fmt.Printf("The hardest card is \"%s\". You have %d errors answering it\n", cardName, deck.getFlashCard(cardIndex).Errors)
		log = append(log, fmt.Sprintf("The hardest card is \"%s\". You have %d errors answering it\n", cardName, deck.getFlashCard(cardIndex).Errors))
	} else {
		var builder strings.Builder
		builder.WriteString("The hardest cards are ")
		for i := 0; i < len(deck.difficultCards); i++ {
			cardName = deck.difficultCards[i]
			builder.WriteString("\"" + cardName + "\"" + ", ")
		}
		fmt.Println(builder.String()[:len(builder.String())-2])
		log = append(log, builder.String()[:len(builder.String())-2])

	}
}
func resetStats(deck *FlashcardDeck) {
	var card *Flashcard
	for i := 0; i < deck.Size; i++ {
		card = &deck.FlashCardSlice[i]
		card.Errors = 0
	}
	fmt.Println("Card statistics have been reset.")
	log = append(log, fmt.Sprintln("Card statistics have been reset."))
}

func exit(deck *FlashcardDeck, exportTo string) {
	exportCards(deck, exportTo, false)
}

var log []string = make([]string, 100)

func main() {
	var scanner *bufio.Scanner
	var answer string
	importFrom := flag.String("import_from", "", "Set the name of the file to import from.")
	exportTo := flag.String("export_to", "", "Set the name of the file to export to when exiting.")
	flag.Parse()
	var importDefined, exportDefined bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "import_from" {
			importDefined = true
		}
		if f.Name == "export_to" {
			exportDefined = true
		}
	})

	flashcardsDeck := FlashcardDeckCreator()
	if importDefined {
		importCards(flashcardsDeck, *importFrom, false)
	}
	for answer != "exit" {
		scanner = bufio.NewScanner(os.Stdin)
		fmt.Println("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats): 	")
		log = append(log, fmt.Sprintln("Input the action (add, remove, import, export, ask, exit):"))
		scanner.Scan()
		answer = scanner.Text()
		log = append(log, fmt.Sprintln(scanner.Text()))
		switch answer {
		case "exit":
			if exportDefined {
				exit(flashcardsDeck, *exportTo)
			}
		case "add":
			addCard(flashcardsDeck)
		case "remove":
			removeCard(flashcardsDeck)
		case "import":
			importCards(flashcardsDeck, "", true)
		case "export":
			exportCards(flashcardsDeck, "", true)
		case "ask":
			ask(flashcardsDeck)
		case "log":
			logToFile(flashcardsDeck)
		case "hardest card":
			hardestCard(flashcardsDeck)
		case "reset stats":
			resetStats(flashcardsDeck)
		}
		fmt.Println()
		log = append(log, "\n")
	}
	fmt.Println("bye bye")
	log = append(log, fmt.Sprintln("bye bye"))
}
