package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Flashcard struct {
	Front string
	Back  string
}
type FlashcardDeck struct {
	flashCardSlice []Flashcard
	size           int
	frontMap       map[string]*Flashcard
	backMap        map[string]*Flashcard
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

func (d *FlashcardDeck) addFlashCardToMaps(flashcard *Flashcard) {
	d.addFrontToMap(flashcard)
	d.addBackToMap(flashcard)
}

func (d *FlashcardDeck) addFrontToMap(flashcard *Flashcard) {
	d.frontMap[flashcard.getFront()] = flashcard
}

func (d *FlashcardDeck) addBackToMap(flashcard *Flashcard) {
	d.backMap[flashcard.getBack()] = flashcard
}

func (d *FlashcardDeck) addCard(flashcard *Flashcard) {
	d.flashCardSlice = append(d.flashCardSlice, *flashcard)
	d.size++
}

//	func (f *FlashcardDeck) getDeck() []Flashcard {
//		return f.flashCardSlice
//	}

func (d *FlashcardDeck) getFlashCard(i int) Flashcard {
	return d.flashCardSlice[i]
}

func (d *FlashcardDeck) hasFront(front string) bool {
	return d.frontMap[front] != nil
}

func (d *FlashcardDeck) hasBack(back string) bool {
	_, ok := d.backMap[back]
	return ok
}

func FlashcardCreator(front, back string) *Flashcard {
	return &Flashcard{Front: front, Back: back}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func (d *FlashcardDeck) getFrontFromBack(answer string) string {
	var cardPointer *Flashcard = d.backMap[answer]
	return (*cardPointer).getFront()
}

func (f *Flashcard) printFeedback(d *FlashcardDeck, answer string) {
	if f.checkAnswer(answer) {
		fmt.Println("Correct!")
	} else if d.hasBack(answer) {
		fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", f.getBack(), d.getFrontFromBack(answer))
	} else {
		fmt.Printf("Wrong. The right answer is \"%s\".\n", f.getBack())
	}
}

func getNumberOfCardsToCreate() int {
	var numberOfCardsToCreate int
	fmt.Println("Input the number of cards:")
	fmt.Scan(&numberOfCardsToCreate)
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
		newFront, _ = reader.ReadString('\n')
		newFront = strings.TrimSpace(newFront)
	}
	return newFront
}

func getOriginalBack(d *FlashcardDeck, back *string) string {
	if !d.hasBack(*back) {
		return *back
	}
	var reader = bufio.NewReader(os.Stdin)
	var newBack string = *back
	for d.hasBack(newBack) {
		fmt.Printf("The definition \"%s\" already exists. Try again:\n", newBack)
		newBack, _ = reader.ReadString('\n')
		newBack = strings.TrimSpace(newBack)
	}
	return newBack
}

func (d *FlashcardDeck) createNCards(numberOfCardsToCreate int) {
	var front, back string
	var reader = bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	for cardNumber := 1; cardNumber <= numberOfCardsToCreate; cardNumber++ {
		fmt.Printf("The term for card #%d:\n", cardNumber)
		front, _ = reader.ReadString('\n')
		front = strings.TrimSpace(front)
		front = getOriginalFront(d, &front)
		fmt.Printf("The definition for card #%d:\n", cardNumber)
		back, _ = reader.ReadString('\n')
		back = strings.TrimSpace(back)
		back = getOriginalBack(d, &back)
		newCard := FlashcardCreator(front, back)
		d.addCard(newCard)
		d.addFlashCardToMaps(newCard)
	}
}
func (d *FlashcardDeck) practiceAllCards() {
	var answer string
	var reader = bufio.NewReader(os.Stdin)
	var currentFlashCard Flashcard
	for i := 0; i < d.size; i++ {
		currentFlashCard = d.getFlashCard(i)
		fmt.Printf("Print the definition of \"%s\":\n", currentFlashCard.getFront())
		answer, _ = reader.ReadString('\n')
		currentFlashCard.printFeedback(d, strings.TrimSpace(answer))
	}
}
func FlashcardDeckCreator() FlashcardDeck {
	newDeck := FlashcardDeck{}
	newDeck.frontMap = make(map[string]*Flashcard)
	newDeck.backMap = make(map[string]*Flashcard)
	return newDeck
}

func main() {
	flashcardsDeck := FlashcardDeckCreator()
	numberOfCardsToCreate := getNumberOfCardsToCreate()
	flashcardsDeck.createNCards(numberOfCardsToCreate)
	flashcardsDeck.practiceAllCards()
}
