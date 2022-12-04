package first

import (
	"AdventOfCode2022/types/elf"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// File is saved using browser save as
const listFile = "first/input"

func GetElfList() []*elf.Elf {
	// Make a slice containing our elf's
	elves := make([]*elf.Elf, 0)

	// Open input data and read it
	file, err := os.Open(listFile)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Err:", err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)

	// Looping through the input data, and assigning calories to elf's
	elfCalories := strings.Split(content, "\n\n")
	for elfId := 0; elfId < len(elfCalories); elfId++ {
		newElf := elf.New(elfId)
		elves = append(elves, newElf)

		for _, c := range strings.Split(elfCalories[elfId], "\n") {
			if c != "" {
				calorie, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal("Encountered calorie entry which isn't a number", err)
				}
				newElf.AddCalorie(calorie)
			}
		}
	}

	// Sort list of elf's from the largest calorie count to the smallest
	sort.Slice(elves, func(i, j int) bool { return elves[i].GetTotalCalories() > elves[j].GetTotalCalories() })

	return elves
}

func printPartOne() {
	fmt.Println("--- Part One ---")
	elves := GetElfList()

	// Print the elf with the most calories
	fmt.Printf("The elf with the most calories is elf number %d with %d calories\n\n", elves[0].Number, elves[0].GetTotalCalories())
}

func printPartTwo() {
	fmt.Println("--- Part Two ---")
	elves := GetElfList()

	sum := elves[0].GetTotalCalories() + elves[1].GetTotalCalories() + elves[2].GetTotalCalories()
	fmt.Printf("The sum of calories which the three elf's with the most calories have is %d\n\n", sum)
}

func PrintAnswer() {
	fmt.Println("--- Day 1: Calorie Counting ---")
	printPartOne()
	printPartTwo()
}
