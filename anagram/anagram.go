package anagram

import (
	"fmt"
	"sort"
)

func isAnagram(str1, str2 string) bool {
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	sort.Slice(runeStr1, func(i, j int) bool { return runeStr1[i] < runeStr1[j] })
	sort.Slice(runeStr2, func(i, j int) bool { return runeStr2[i] < runeStr2[j] })

	return string(runeStr1) == string(runeStr2)
}

// return multiple array string.
func GroupAnagrams(strInput []string) [][]string {
	var result [][]string

	for _, str := range strInput {
		// Mengecek apakah string sudah termasuk dalam grup yang ada.
		var grouped bool
		for j, group := range result {
			//cek jika string tersebut adalah anagram
			if isAnagram(group[0], str) {
				result[j] = append(result[j], str)
				grouped = true
				break
			}
		}

		// Jika tidak termasuk dalam grup yang ada, buat grup baru.
		if !grouped {
			result = append(result, []string{str})
		}
	}
	// looping hasil
	for _, group := range result {
		fmt.Println(group)
	}

	return result
}
