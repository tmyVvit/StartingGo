package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for  {
		fmt.Print("please input expression: ")
		var input, _, err = inputReader.ReadLine()
		var line = string(input)
		if err == nil && line != "q" {
			if str, err := valid(line); err == nil {
				err := cal(str)
				if err != nil {
					fmt.Println("error: ", err)
				}
			} else {
				fmt.Println("error: ", err)
			}
		} else {
			fmt.Println("quit...")
			break
		}
	}
}

func valid(line string) (result string, err error) {
	result = strings.NewReplacer(" ", "").Replace(line)
	reg, err := regexp.Compile("(-?[0-9]+)[+*/-](-?[0-9]+)=$")
	if err == nil && reg.MatchString(result) {
		return strings.TrimSuffix(result, "="), nil
	}
	return result, errors.New("invalid input")
}

func cal(exp string) (err error){
	negative := strings.Index(exp, "-") == 0
	if negative {
		exp = exp[1:]
	}
	var nums []string
	switch {
		case strings.Index(exp, "+") > -1:{
			nums = strings.Split(exp, "+")
			fmt.Println("result: ", neg(negative) * convert(nums[0]) + convert(nums[1]))
		}
		case strings.Index(exp, "*") > -1:{
			nums = strings.Split(exp, "*")
			fmt.Println("result: ", float32(neg(negative) * convert(nums[0])) * float32(convert(nums[1])))
		}
		case strings.Index(exp, "/") > -1:{
			nums = strings.Split(exp, "/")
			if  convert(nums[1]) == 0{
				err = errors.New("divisor cannot be 0")
				break
			}
			fmt.Println("result: ", float32(neg(negative) * convert(nums[0])) / float32(convert(nums[1])))
		}
		case strings.Index(exp, "-") > -1:{
			nums = strings.Split(exp, "-")
			fmt.Println("result: ", neg(negative) * convert(nums[0]) - convert(nums[1]))
		}
	}
	return err
}

func convert(str string) (data int) {
	reg, err := regexp.Compile("[0-9]*")
	if err == nil && reg.MatchString(str) {
		for _, value := range str {
			data = data * 10 + int(value) - int('0')
		}
	}
	return data
}

func neg(negative bool) int {
	if negative {
		return -1
	}
	return 1
}