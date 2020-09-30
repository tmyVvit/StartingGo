package leetcode

import (
	"regexp"
	"strings"
)

// func main() {
// 	res := letterCasePermutation("a1b4")
// 	fmt.Println(res)
// }

func letterCasePermutation(S string) (res []string) {
	if S == "" || len(S) == 0 {
		return res
	}

	var letterCase func(letters string) []string
	letterCase = func(letters string) []string {
		if len(letters) > 1 {
			letter := letters[:1]
			sub := letters[1:]
			return appendStrs(letter, letterCase(sub))
		}
		return appendStrs(letters, nil)
	}
	return letterCase(S)
}

func appendStrs(letter string, subs []string) []string {
	twoCases := getAllCases(letter)
	if subs == nil {
		return twoCases
	}
	var res []string
	for _, val := range subs {
		for _, cases := range twoCases {
			res = append(res, cases+val)
		}
	}
	return res
}

func getAllCases(letter string) (res []string) {
	res = append(res, letter)
	if isUppercase(letter) {
		res = append(res, strings.ToLower(letter))
	} else if isLowercase(letter) {
		res = append(res, strings.ToUpper(letter))
	}
	return res
}

func isUppercase(s string) bool {
	upperPattern := "[A-Z]"
	result, _ := regexp.MatchString(upperPattern, s)
	return result
}

func isLowercase(s string) bool {
	lowerPattern := "[a-z]"
	result, _ := regexp.MatchString(lowerPattern, s)
	return result
}
