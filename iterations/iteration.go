package iteration

import "strings"

func Repeat(letter string, repetitions int) string {
	var repeated string

	for i := 0; i < repetitions; i++ {
		repeated += letter
	}
	return repeated
}

func RepeatWithFunction(letter string, repetitions int) string {
	return strings.Repeat(letter, repetitions)
}
