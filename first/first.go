package first

import (
	"AdventOfCode2022/first/elf"
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
	elfs := make([]*elf.Elf, 0)

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
		elfs = append(elfs, newElf)

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
	sort.Slice(elfs, func(i, j int) bool { return elfs[i].GetTotalCalories() > elfs[j].GetTotalCalories() })

	return elfs
}

func PrintAnswerOne() {
	fmt.Println("--- Part One ---")
	elfs := GetElfList()

	// Print the elf with the most calories
	fmt.Printf("The elf with the most calories is elf number %d with %d calories\n\n", elfs[0].Number, elfs[0].GetTotalCalories())
}

func PrintAnswerTwo() {
	fmt.Println("--- Part Two ---")
	elfs := GetElfList()

	sum := elfs[0].GetTotalCalories() + elfs[1].GetTotalCalories() + elfs[2].GetTotalCalories()
	fmt.Printf("The sum of calories which the three elf's with the most calories have is %d\n\n", sum)
}

func PrintAnswer() {
	fmt.Println("--- Day 1: Calorie Counting ---")
	PrintAnswerOne()
	PrintAnswerTwo()
}
