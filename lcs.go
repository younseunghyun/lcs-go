package main

import "fmt"

func lcs(s1 string, s2 string) (int) {
	result := make([][]int, len(s1)+1)
	for i := range result {
		result[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				result[i][j] = result[i-1][j-1] + 1
			} else {
				result[i][j] = max(result[i-1][j], result[i][j-1])
			}
		}
	}

	return result[len(s1)][len(s2)]
}

func max(i1 int, i2 int) (int) {
	if i1 >= i2 {
		return i1
	} else {
		return i2
	}
}

func main() {
	s1 := "ABCDAA"
	s2 := "BACDAA"

	fmt.Println(lcs(s1, s2))
}
