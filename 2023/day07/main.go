package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	OnePair int = iota + 1
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
	ScoreCount
)

var cardStrength = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

var cardStrengthPartTwo = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1,
}
var cardRankingPartTwo = []rune{
	'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J',
}

func sortByStrength(hands []string) []string {
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})
	return hands
}

func compareHands(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {
		if cardStrength[rune(hand1[i])] != cardStrength[rune(hand2[i])] {
			return cardStrength[rune(hand1[i])] < cardStrength[rune(hand2[i])]
		}
	}
	return false
}
func sortByStrengthPartTwo(hands []string) []string {
	sort.Slice(hands, func(i, j int) bool {
		return compareHandsPartTwo(hands[i], hands[j])
	})
	return hands
}

func compareHandsPartTwo(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {
		if cardStrengthPartTwo[rune(hand1[i])] != cardStrengthPartTwo[rune(hand2[i])] {
			return cardStrengthPartTwo[rune(hand1[i])] < cardStrengthPartTwo[rune(hand2[i])]
		}
	}
	return false
}

func replaceJokers(cardsToCount *map[string]int, countToCards *map[int][]string, jokerCount int) {
	for count := 5; count > 0; count-- {
		maxCards, found := (*countToCards)[count]
		if !found {
			continue
		}
		if len(maxCards) == 1 {
			if maxCards[0] == string("J") && count == 5 {
				(*cardsToCount)["A"] += jokerCount
			} else if maxCards[0] == string("J") {
				continue
			} else {
				(*cardsToCount)[maxCards[0]] += jokerCount
			}
			delete(*cardsToCount, "J")
			return
		}
		// iterate over maxCards, and find strongest card there
		for _, card1 := range cardRankingPartTwo {
			for _, card2 := range maxCards {
				if string(card1) == card2 {
					(*cardsToCount)[card2] += jokerCount
					delete(*cardsToCount, "J")
					return
				}
			}
		}
	}

}

func calculateHandScore(hand *string, replaceJoker bool) int {
	cardsToCount := make(map[string]int)
	countToCards := make(map[int][]string)
	for _, card := range *hand {
		c := string(card)
		cardsToCount[c] = cardsToCount[c] + 1
		countToCards[cardsToCount[c]] = append(countToCards[cardsToCount[c]], c)
	}
	jokerCount := cardsToCount["J"]
	if replaceJoker && jokerCount > 0 {
		replaceJokers(&cardsToCount, &countToCards, jokerCount)
	}
	fullHouseCount := 0
	twoPairCandidate := false
	score := 0
	for _, count := range cardsToCount {
		// five of a kind
		if count == 5 && score < FiveOfAKind {
			score = FiveOfAKind
		}
		// four of a kind
		if count == 4 && score < FourOfAKind {
			score = FourOfAKind
		}
		// full house
		if count == 3 || count == 2 {
			if fullHouseCount != 0 && fullHouseCount != count && score < FullHouse {
				score = FullHouse
			}
			if fullHouseCount == 0 {
				fullHouseCount = count
			}
		}
		// three of a kind
		if count == 3 && score < ThreeOfAKind {
			score = ThreeOfAKind
		}
		// two pair
		if count == 2 {
			if twoPairCandidate && score < TwoPair {
				score = TwoPair
			}
			if !twoPairCandidate {
				twoPairCandidate = true
			}
		}
		// one pair
		if count == 2 && score < OnePair {
			score = OnePair
		}
	}
	return score
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	scoreToHands := make(map[int][]string)
	handToBid := make(map[string]int)
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) != 2 {
			panic("malformed line " + line)
		}
		hand := lineSplit[0]
		if len(hand) != 5 {
			panic("invalid assumption about hand size")
		}
		score := calculateHandScore(&hand, false)
		scoreToHands[score] = append(scoreToHands[score], hand)
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			panic(err)
		}
		handToBid[hand] = bid
	}
	allSortedHands := make([]string, len(lines))
	count := 0
	for i := 0; i < ScoreCount; i++ {
		hands, found := scoreToHands[i]
		if !found {
			continue
		}
		if len(hands) == 1 {
			allSortedHands[count] = hands[0]
			count += 1
		} else {
			sortedHands := sortByStrength(hands)
			for _, hand := range sortedHands {
				allSortedHands[count] = hand
				count += 1
			}
		}
	}
	ans := 0
	for i := 0; i < len(allSortedHands); i++ {
		bid := handToBid[allSortedHands[i]]
		ans += bid * (i + 1)
	}
	fmt.Println("Part 1 Answer: ", ans)
}

func secondPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	scoreToHands := make(map[int][]string)
	handToBid := make(map[string]int)
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) != 2 {
			panic("malformed line " + line)
		}
		hand := lineSplit[0]
		if len(hand) != 5 {
			panic("invalid assumption about hand size")
		}
		score := calculateHandScore(&hand, true)
		scoreToHands[score] = append(scoreToHands[score], hand)
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			panic(err)
		}
		handToBid[hand] = bid
	}
	allSortedHands := make([]string, len(lines))
	count := 0
	for i := 0; i < ScoreCount; i++ {
		hands, found := scoreToHands[i]
		if !found {
			continue
		}
		if len(hands) == 1 {
			allSortedHands[count] = hands[0]
			count += 1
		} else {
			sortedHands := sortByStrengthPartTwo(hands)
			for _, hand := range sortedHands {
				allSortedHands[count] = hand
				count += 1
			}
		}
	}
	ans := 0
	for i := 0; i < len(allSortedHands); i++ {
		bid := handToBid[allSortedHands[i]]
		ans += bid * (i + 1)
	}
	fmt.Println("Part 2 Answer: ", ans)
}

func main() {
	firstPart()
	secondPart()
}
