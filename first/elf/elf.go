package elf

type Elf struct {
	Number   int
	Calories []int
}

func (e *Elf) AddCalorie(calorie int) {
	if e.Calories == nil {
		e.Calories = make([]int, 0)
	}

	e.Calories = append(e.Calories, calorie)
}

func (e *Elf) GetTotalCalories() int {
	sum := 0
	for _, c := range e.Calories {
		sum += c
	}
	return sum
}

func New(number int) *Elf {
	elf := Elf{Number: number}
	return &elf
}
