package fourth

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type AssignmentPair struct {
	AreaOne Area
	AreaTwo Area
}

type Area struct {
	Start int
	End   int
}

func (a *Area) String() string {
	return fmt.Sprintf("Area start %d, end %d", a.Start, a.End)
}

func (a *AssignmentPair) String() string {
	return fmt.Sprintf("Area one %v, Area two %v", a.AreaOne, a.AreaTwo)
}

func (a *AssignmentPair) ContainsDuplicates() bool {
	for n1 := a.AreaOne.Start; n1 <= a.AreaOne.End; n1++ {
		for n2 := a.AreaTwo.Start; n2 <= a.AreaTwo.End; n2++ {
			if n1 == n2 {
				return true
			}
		}
	}
	return false
}

func (a *AssignmentPair) ContainsDuplicateEntire() bool {
	areaOne := "-"
	areaTwo := "-"

	for n1 := a.AreaOne.Start; n1 <= a.AreaOne.End; n1++ {
		areaOne += strconv.Itoa(n1) + "-"
	}
	for n1 := a.AreaTwo.Start; n1 <= a.AreaTwo.End; n1++ {
		areaTwo += strconv.Itoa(n1) + "-"
	}

	if strings.Contains(areaOne, areaTwo) {
		fmt.Printf("1 %s %s is complete duplicate\n", areaOne, areaTwo)
		return true
	}

	if strings.Contains(areaTwo, areaOne) {
		fmt.Printf("2 %s %s is complete duplicate\n", areaOne, areaTwo)
		return true
	}
	fmt.Printf("%s %s is not complete duplicate\n", areaOne, areaTwo)
	return false
}

const fileInput = "fourth/input.txt"

func GetAssignmentPairs() []*AssignmentPair {
	assignmentPairs := make([]*AssignmentPair, 0)

	file, err := os.Open(fileInput)
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

	for _, p := range strings.Split(string(content), "\n") {
		if p != "" {
			areaArray := strings.Split(p, ",")
			areaOne := strings.Split(areaArray[0], "-")
			areaTwo := strings.Split(areaArray[1], "-")

			areaOneStart, err := strconv.Atoi(areaOne[0])
			if err != nil {
				log.Fatal(err)
			}
			areaOneEnd, err := strconv.Atoi(areaOne[1])
			if err != nil {
				log.Fatal(err)
			}

			areaTwoStart, err := strconv.Atoi(areaTwo[0])
			if err != nil {
				log.Fatal(err)
			}
			areaTwoEnd, err := strconv.Atoi(areaTwo[1])
			if err != nil {
				log.Fatal(err)
			}

			newAssignmentPair := AssignmentPair{
				AreaOne: Area{
					Start: areaOneStart,
					End:   areaOneEnd,
				},
				AreaTwo: Area{
					Start: areaTwoStart,
					End:   areaTwoEnd,
				},
			}

			assignmentPairs = append(assignmentPairs, &newAssignmentPair)
		}
	}

	return assignmentPairs
}

func printAnswerPartOne() {
	fmt.Println("--- Part One ---")

	assignmentPairs := GetAssignmentPairs()
	duplicates := 0
	for _, a := range assignmentPairs {
		if a.ContainsDuplicateEntire() {
			duplicates++
			fmt.Printf("Area %s contains duplicates\n", a)
		} else {
			fmt.Printf("Area %s does not contain duplicates\n", a)
		}

	}
	//fmt.Println(assignmentPairs)
	fmt.Printf("Duplicate area assignments: %d\n\n", duplicates)
}

func printAnswerPartTwo() {
	fmt.Println("--- Part Two ---")

	assignmentPairs := GetAssignmentPairs()
	duplicates := 0
	for _, a := range assignmentPairs {
		if a.ContainsDuplicates() {
			duplicates++
			fmt.Printf("Area %s contains duplicates\n", a)
		} else {
			fmt.Printf("Area %s does not contain duplicates\n", a)
		}

	}
	//fmt.Println(assignmentPairs)
	fmt.Printf("Duplicate area assignments: %d\n\n", duplicates)
}

func PrintAnswer() {
	fmt.Println("--- Day 4: Camp Cleanup ---")
	printAnswerPartOne()
	printAnswerPartTwo()
}
