package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
)

// Entry point of diceroll CLI application
func main() {
	dice := flag.String("d", "d6", "The type of the dice to roll. Format: dX, where X is an integer")
	numRoll := flag.Int("n", 1, "The number dice to roll")
	sum := flag.Bool("s", false, "Get the sum of all the dice rolls")
	advantage := flag.Bool("adv", false, "Roll the dice with advantage")
	disadvantage := flag.Bool("dis", false, "Roll the dice with disadvantage")
	flag.Parse()

	matched, _ := regexp.Match("d\\d+", []byte(*dice))
	if matched {
		rolls := rollDice(dice, numRoll)
		printDice(rolls)
		if *sum {
			diceSum := sumDice(rolls)
			fmt.Printf("The sum of the dice was %d\n", diceSum)
		}
		if *advantage {
			roll := rollWithAdvantage(rolls)
			fmt.Printf("The roll with advantage was %d\n", roll)
		}
		if *disadvantage {
			roll := rollWithDisadvantage(rolls)
			fmt.Printf("The roll with disadvantage was %d\n", roll)
		}
	} else {
		log.Fatal("Improper format for dice. Format should be dX where X is an integer")
	}
}

func rollDice(dice *string, times *int) []int {
	var rolls []int
	diceSlots := (*dice)[1:]
	d, err := strconv.Atoi(diceSlots)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *times; i++ {
		roll := rand.Intn(d) + 1
		rolls = append(rolls, roll)
	}
	return rolls
}

func printDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0
	for _, dice := range rolls {
		sum += dice
	}
	return sum
}

func rollWithAdvantage(rolls []int) int {
	return getRoll(rolls, len(rolls)-1)
}

func rollWithDisadvantage(rolls []int) int {
	return getRoll(rolls, 0)
}

func getRoll(rolls []int, index int) int {
	sort.Ints(rolls)
	return rolls[index]
}
