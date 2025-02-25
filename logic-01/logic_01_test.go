package logic_01

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogic01Number1(t *testing.T) {
	result := Logic01Number1(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number2(t *testing.T) {
	result := Logic01Number2(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number3(t *testing.T) {
	result := Logic01Number3(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number4(t *testing.T) {
	result := Logic01Number4(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{19, 17, 15, 13, 11, 9, 7, 5, 3, 1}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number5(t *testing.T) {
	result := Logic01Number5(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 2}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number6(t *testing.T) {
	result := Logic01Number6(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{30, 27, 24, 21, 18, 15, 12, 9, 6, 3}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number7(t *testing.T) {
	result := Logic01Number7(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{1, 3, 5, 7, 9, 9, 7, 5, 3, 1}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)

	result = Logic01Number7(11)
	assert.Equal(t, 11, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected = []int{1, 3, 5, 7, 9, 11, 9, 7, 5, 3, 1}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number8(t *testing.T) {
	result := Logic01Number8(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{2, 4, 6, 8, 10, 10, 8, 6, 4, 2}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)

	result = Logic01Number8(11)
	assert.Equal(t, 11, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected = []int{2, 4, 6, 8, 10, 12, 10, 8, 6, 4, 2}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number9(t *testing.T) {
	result := Logic01Number9(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{3, 6, 9, 12, 15, 15, 12, 9, 6, 3}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)

	result = Logic01Number9(11)
	assert.Equal(t, 11, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected = []int{3, 6, 9, 12, 15, 18, 15, 12, 9, 6, 3}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number10(t *testing.T) {
	result := Logic01Number10(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []string{"2", "fizz", "4", "fizz", "8", "fizz", "16", "fizz", "32", "fizz"}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number11(t *testing.T) {
	result := Logic01Number11(10)
	assert.Equal(t, 10, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []string{"buzz", "1", "buzz", "3", "buzz", "6", "buzz", "12", "buzz", "24"}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}

func TestLogic01Number12(t *testing.T) {
	result := Logic01Number12(12)
	assert.Equal(t, 12, len(result))
	fmt.Println("Slice Length: ", len(result))

	expected := []int{1, 3, 5, 7, 1, 3, 5, 7, 1, 3, 5, 7}
	assert.Equal(t, expected, result)
	fmt.Println("Slice Result:\n", result)
}
