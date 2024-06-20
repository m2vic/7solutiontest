package pkg

import (
	"strconv"
	"strings"
)

func checkL(n string) int {
	count := 0

	for _, val := range n {
		if string(val) != "L" {
			return count
		}
		count++
	}
	return count
}

func LeftRight(alphabet string) string {

	arr := []int{}
	curr := 0
	val := 0

	// init and push the first value
	firstAlphabet := string(alphabet[0])
	secondAlphabet := string(alphabet[1])

	if firstAlphabet == "=" {
		if secondAlphabet == "L" {
			v := checkL(alphabet[1:])
			arr = append(arr, v)
		} else {
			arr = append(arr, val)
		}
	} else if firstAlphabet == "L" {
		v := checkL(alphabet)
		arr = append(arr, v)
	} else {
		arr = append(arr, val)
	}

	// iterate through input
	for curr < len(alphabet) {
		currentAlphabet := string(alphabet[curr])

		prevAlphabet := ""
		if curr > 0 {
			prevAlphabet = string(alphabet[curr-1])
		}

		nextAlphabet := ""
		if curr+1 < len(alphabet) {
			nextAlphabet = string(alphabet[curr+1])
		}

		if currentAlphabet == "R" {
			if nextAlphabet == "L" {
				v := checkL(alphabet[curr+1:])
				if v == 1 {
					// If the next alphabet is L, We have to check consecutive amount of L , due to make sure that the number is more than the previous
					//and able to decrease if there is more than a single L
					val++
					arr = append(arr, val)
					// we use val , to make sure the placed val is correct
				} else {
					arr = append(arr, v)
				}
			} else {
				val++
				arr = append(arr, val)
			}
		}
		if currentAlphabet == "L" {
			if prevAlphabet == "R" && nextAlphabet != "L" {
				val = 0
			} else {
				val = arr[curr] - 1
			}
			arr = append(arr, val)
		}

		if currentAlphabet == "=" {
			arr = append(arr, arr[curr])
		}
		curr++

	}
	var sb strings.Builder

	for _, num := range arr {
		sb.WriteString(strconv.Itoa(num))
	}
	result := sb.String()

	return result
}
