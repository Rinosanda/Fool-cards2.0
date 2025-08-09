package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

var suits = []string{"Пики", "Червы", "Бубны", "Трефы"}
var values = []string{"6", "7", "8", "9", "10", "Валет", "Дама", "Король", "Туз"}

func createDeck() []Card {
	var deck []Card
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

func shuffleDeck(deck []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func dealCards(deck []Card) ([]Card, []Card, []Card) {
	playerHand := deck[:6]
	computerHand := deck[6:12]
	remainingDeck := deck[12:]
	return playerHand, computerHand, remainingDeck
}

func printHand(name string, hand []Card) {
	fmt.Printf("%s: ", name)
	for i, card := range hand {
		fmt.Printf("[%d: %s %s] ", i, card.Value, card.Suit)
	}
	fmt.Println()
}

func playerTurn(hand []Card) ([]Card, Card) {
	printHand("Ваши карты", hand)
	var index int
	for {
		fmt.Print("Выберите номер карты для хода: ")
		_, err := fmt.Scanf("%d\n", &index)
		if err == nil && index >= 0 && index < len(hand) {
			break
		}
		fmt.Println("Неверный ввод, попробуйте снова.")
	}
	card := hand[index]
	hand = append(hand[:index], hand[index+1:]...)
	return hand, card
}

func computerTurn(hand []Card) ([]Card, Card) {
	index := rand.Intn(len(hand))
	card := hand[index]
	hand = append(hand[:index], hand[index+1:]...)
	fmt.Printf("Компьютер ходит картой: %s %s\n", card.Value, card.Suit)
	return hand, card
}

func main() {
	deck := shuffleDeck(createDeck())

	playerHand, computerHand, remainingDeck := dealCards(deck)
	trump := remainingDeck[len(remainingDeck)-1].Suit
	fmt.Println("Козырь:", trump)

	for len(playerHand) > 0 && len(computerHand) > 0 {
		// Игрок ходит
		var playedCard Card
		playerHand, playedCard = playerTurn(playerHand)
		fmt.Printf("Вы выложили карту: %s %s\n", playedCard.Value, playedCard.Suit)

		// Компьютер отвечает
		computerHand, _ = computerTurn(computerHand)

		fmt.Println("Раунд завершён.")
		fmt.Println()
	}

	fmt.Println("Игра окончена.")
	if len(playerHand) == 0 {
		fmt.Println("Вы победили!")
	} else {
		fmt.Println("Компьютер победил!")
	}
}
