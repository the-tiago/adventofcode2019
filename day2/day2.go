package main

import "fmt"

// Day2 takes a program and returns the program processed
func Day2(program []int) []int {
	if len(program) == 0 {
		return program
	}

	for i := 0; i < len(program); {
		switch program[i] {
		case 1:
			program = sum(program, i)
		case 2:
			program = mul(program, i)
		case 99:
			return program
		}
		i += 4
	}

	return program
}

// Day2Part2 takes a program a finds the second and third items so the final program has the value of 19690720 in its first position
func Day2Part2(program []int) (noun, verb int) {
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			fmt.Printf("trying noun = %d and verb = %d\n", noun, verb)
			newProgram := copyProgram(program)
			newProgram[1] = noun
			newProgram[2] = verb
			processed := Day2(newProgram)
			if processed[0] == 19690720 {
				return
			}
		}
	}
	return -1, -1
}

func sum(program []int, index int) []int {
	a := program[program[index+1]]
	b := program[program[index+2]]
	program[program[index+3]] = a + b
	return program
}

func mul(program []int, index int) []int {
	a := program[program[index+1]]
	b := program[program[index+2]]
	program[program[index+3]] = a * b
	return program
}

func copyProgram(program []int) []int {
	programCopy := make([]int, len(program))
	copy(programCopy, program)
	return programCopy
}

func main() {
	/*
		processed := Day2([]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3,
			1, 5, 0, 3, 2, 1, 10, 19, 2, 9, 19, 23, 1, 9, 23, 27, 2, 27, 9, 31,
			1, 31, 5, 35, 2, 35, 9, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 47, 6, 51,
			2, 51, 10, 55, 1, 9, 55, 59, 2, 6, 59, 63, 1, 63, 6, 67, 1, 67, 10, 71,
			1, 71, 10, 75, 2, 9, 75, 79, 1, 5, 79, 83, 2, 9, 83, 87, 1, 87, 9, 91,
			2, 91, 13, 95, 1, 95, 9, 99, 1, 99, 6, 103, 2, 103, 6, 107, 1, 107, 5, 111,
			1, 13, 111, 115, 2, 115, 6, 119, 1, 119, 5, 123, 1, 2, 123, 127, 1, 6, 127, 0, 99,
			2, 14, 0, 0})
	*/
	input := []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3,
		1, 5, 0, 3, 2, 1, 10, 19, 2, 9, 19, 23, 1, 9, 23, 27, 2, 27, 9, 31,
		1, 31, 5, 35, 2, 35, 9, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 47, 6, 51,
		2, 51, 10, 55, 1, 9, 55, 59, 2, 6, 59, 63, 1, 63, 6, 67, 1, 67, 10, 71,
		1, 71, 10, 75, 2, 9, 75, 79, 1, 5, 79, 83, 2, 9, 83, 87, 1, 87, 9, 91,
		2, 91, 13, 95, 1, 95, 9, 99, 1, 99, 6, 103, 2, 103, 6, 107, 1, 107, 5, 111,
		1, 13, 111, 115, 2, 115, 6, 119, 1, 119, 5, 123, 1, 2, 123, 127, 1, 6, 127, 0, 99,
		2, 14, 0, 0}
	processed := Day2(copyProgram(input))
	fmt.Println(processed)

	fmt.Println(Day2Part2(input))
}
