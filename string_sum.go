package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {

	input = strings.Join(strings.Fields(input), "")
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	var a = regexp.MustCompile(`([+-]?[^+-]+)`)
	b := a.FindAllStringSubmatch(input, -1)
	var x []int

	for _, v := range b {
		value, err := strconv.Atoi(v[0])
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}
		x = append(x, value)
	}

	if len(x) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	return strconv.Itoa(x[0] + x[1]), nil
}
