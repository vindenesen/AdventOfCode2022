package main

import (
	"AdventOfCode2022/01/elf"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// File is saved using browser save as
const listFile = "input"

func main() {
	fmt.Println("Day 1: Calorie Counting")

	elfs := make([]*elf.Elf, 0)

	file, err := os.Open(listFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)

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

	sort.Slice(elfs, func(i, j int) bool { return elfs[i].GetTotalCalories() < elfs[j].GetTotalCalories() })
	fmt.Println(elfs[len(elfs)-1].GetTotalCalories())
}
