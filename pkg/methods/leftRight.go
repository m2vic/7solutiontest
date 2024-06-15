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

func LeftRight(alp string) string {

	var arr []int
	curr := 0
	val := 0
	alphabet := string(alp)

	if string(alphabet[0]) == "=" {
		if string(alphabet[1]) == "L" {
			v := checkL(alphabet[1:])
			arr = append(arr, v)
		} else {
			arr = append(arr, val)
		}
	}

	if string(alphabet[0]) == "R" {
		arr = append(arr, val)
	}

	if string(alphabet[0]) == "L" {
		v := checkL(alphabet)
		arr = append(arr, v)
	}

	for curr < len(alphabet) {
		if string(alphabet[curr]) == "R" {
			//			if curr+1 < len(alp) && alp[curr+1] == 'L' {

			if curr+1 < len(alp) && string(alphabet[curr+1]) == "L" {
				v := checkL(alphabet[curr+1:])
				if v == 1 {
					val++
					arr = append(arr, val)
				} else {
					arr = append(arr, v)
				}
			} else {
				val++
				arr = append(arr, val)
			}
		}

		if string(alphabet[curr]) == "L" {
			if curr+1 < len(alp) && curr > 0 && string(alphabet[curr-1]) == "R" && string(alphabet[curr+1]) != "L" {
				val = 0
			} else {
				val = arr[curr] - 1
			}
			//fmt.Println(val)
			arr = append(arr, val)
		}

		if string(alphabet[curr]) == "=" {
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
