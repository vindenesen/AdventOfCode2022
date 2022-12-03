package third

import (
	"AdventOfCode2022/types/rucksack"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "third/input"
)

func GetRuckSacks() []*rucksack.Rucksack {
	rucksacks := make([]*rucksack.Rucksack, 0)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	for id, rucksackContent := range strings.Split(string(content), "\n") {
		if rucksackContent != "" {
			newRucksack := rucksack.New(id)
			contentSlice := strings.Split(rucksackContent, "")
			for _, i := range contentSlice[0:(len(contentSlice) / 2)] {
				newItem := rucksack.NewItem(i)
				newRucksack.CompartmentOne.AddItem(newItem)
			}
			for _, i := range contentSlice[len(contentSlice)/2:] {
				newRucksack.CompartmentTwo.AddItem(&rucksack.Item{Value: i})
			}

			rucksacks = append(rucksacks, newRucksack)
		}
	}

	return rucksacks
}

func printAnswerPartOne() {
	fmt.Println("--- Part One ---")
	rucksacks := GetRuckSacks()
	sum := 0
	for _, r := range rucksacks {
		duplicates := r.GetDuplicateItems()

		for _, d := range duplicates {
			fmt.Printf("Rucksack id %d, item %s, priority %d\n", r.Number, d, d.Priority())
			sum += int(d.Priority())
		}
	}

	fmt.Println(sum)
}

func PrintAnswer() {
	fmt.Println("--- Day 3: Rucksack Reorganization ---")
	printAnswerPartOne()
}
