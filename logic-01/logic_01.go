package logic_01

import "strconv"

func Logic01Number1(num int) (result []int) {
	result = make([]int, num)
	value := 1
	for i := 0; i < num; i++ { // fill the value, forwards
		result[i] = value
		value += 2
	}
	return result
}

func Logic01Number2(num int) (result []int) {
	result = make([]int, num)
	value := 2
	for i := 0; i < num; i++ { // fill the value, forwards
		result[i] = value
		value += 2
	}
	return result
}

func Logic01Number3(num int) (result []int) {
	result = make([]int, num)
	value := 3
	for i := 0; i < num; i++ { // fill the value, forwards
		result[i] = value
		value += 3
	}
	return result
}

func Logic01Number4(num int) (result []int) {
	result = make([]int, num)
	value := 1
	for i := num - 1; i >= 0; i-- { // fill with the value, backwards
		result[i] = value
		value += 2
	}
	return result
}

func Logic01Number5(num int) (result []int) {
	result = make([]int, num)
	value := 2
	for i := num - 1; i >= 0; i-- { // fill with the value, backwards
		result[i] = value
		value += 2
	}
	return result
}

func Logic01Number6(num int) (result []int) {
	result = make([]int, num)
	value := 3
	for i := num - 1; i >= 0; i-- { // fill with the value, backwards
		result[i] = value
		value += 3
	}
	return result
}

func Logic01Number7(num int) (result []int) {
	result = make([]int, num)
	value := 1
	midPoint := ((num + 1) / 2) - 1
	for i := 0; i < num; i++ {
		result[i] = value
		if i == midPoint && num%2 == 0 { // if the index is in the midpoint and the series is even, do nothing to the value
			continue
		} else if i >= midPoint { // if the index is equal (only if odd) or more than the midpoint, minus the value
			value -= 2
		} else { // if the index is before the midpoint, add the value
			value += 2
		}
	}
	return result
}

func Logic01Number8(num int) (result []int) {
	result = make([]int, num)
	value := 2
	midPoint := ((num + 1) / 2) - 1
	for i := 0; i < num; i++ {
		result[i] = value
		if i == midPoint && num%2 == 0 { // if the index is in the midpoint and the series is even, do nothing to the value
			continue
		} else if i >= midPoint { // if the index is equal (only if odd) or more than the midpoint, minus the value
			value -= 2
		} else { // if the index is before the midpoint, add the value
			value += 2
		}
	}
	return result
}

func Logic01Number9(num int) (result []int) {
	result = make([]int, num)
	value := 3
	midPoint := ((num + 1) / 2) - 1
	for i := 0; i < num; i++ {
		result[i] = value
		if i == midPoint && num%2 == 0 { // if the index is in the midpoint and the series is even, do nothing to the value
			continue
		} else if i >= midPoint { // if the index is equal (only if odd) or more than the midpoint, minus the value
			value -= 3
		} else { // if the index is before the midpoint, add the value
			value += 3
		}
	}
	return result
}

func Logic01Number10(num int) (result []string) {
	result = make([]string, num)
	value := 2
	for i := 0; i < num; i++ {
		if (i+1)%2 == 0 { // if index is even, fill with "fizz"
			result[i] = "fizz"
		} else { // if index is odd, fill with the value
			result[i] = strconv.Itoa(value)
			value *= 2
		}
	}
	return result
}

func Logic01Number11(num int) (result []string) {
	result = make([]string, num)
	value := 3
	for i := 0; i < num; i++ {
		if i == 1 { // if the index is 1, fill with plain "1"
			result[i] = "1"
		} else if (i+1)%2 == 1 { // if the index is odd, fill with "buzz"
			result[i] = "buzz"
		} else { // if the index is even (except the '1' index), fill with the value
			result[i] = strconv.Itoa(value)
			value *= 2
		}
	}
	return result
}

func Logic01Number12(num int) (result []int) {
	numSlice := Logic01Number1(4)
	result = make([]int, num)
	pointer := 0
	for i := 0; i < num; i++ {
		result[i] = numSlice[pointer]
		if (i+1)%4 == 0 { // if the index is divisible by 4, reset the value pointer
			pointer = 0
		} else { // if not, add 1 to the value pointer
			pointer++
		}
	}
	return result
}
