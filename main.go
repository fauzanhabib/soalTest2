package main

import (
	"awesomeCodingTest/anagram"
	"awesomeCodingTest/rotate"
)

func main() {
	// Group Anagram soal test nomor 1
	input := []string{"kasur","sukar","rakus","siap","kelas","selak"}
	anagram.GroupAnagrams(input)

	// Rotate soal test nomor 2
	inputArray := []int{1,2,3,4,5,6,7}
	rotationAmount := 3
	rotate.Rotate(inputArray, rotationAmount)
}
