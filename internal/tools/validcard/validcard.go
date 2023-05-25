package validcard

import (
	"fmt"
	"strconv"
	"strings"
)

// Validator functions to ensure valid input
func CnValidator(s string) error {
	// Card Number should a string less than 20 digits
	// It should include 16 integers and 3 spaces
	if len(s) > 16+3 {
		return fmt.Errorf("card number is too long")
	}

	// The last digit should be a number unless it is a multiple of 4 in which
	// case it should be a space
	if len(s)%5 == 0 && s[len(s)-1] != ' ' {
		return fmt.Errorf("card number must separate groups with spaces")
	}
	if len(s)%5 != 0 && (s[len(s)-1] < '0' || s[len(s)-1] > '9') {
		return fmt.Errorf("card number is invalid")
	}

	// The remaining digits should be integers
	c := strings.ReplaceAll(s, " ", "")
	_, err := strconv.ParseInt(c, 10, 64)
	if err != nil {
		return fmt.Errorf("the remaining digits should be integers")
	}

	return nil
}

func ExpValidator(s string) error {
	// The 3 character should be a slash (/)
	// The rest should be numbers
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("EXP is invalid")
	}

	// There should be only one slash and it should be in the 2nd index (3rd character)
	if len(s) >= 3 && (strings.Index(s, "/") != 2 || strings.LastIndex(s, "/") != 2) {
		return fmt.Errorf("EXP is invalid")
	}

	return nil
}

func CVCValidator(s string) error {
	// The CVC should be a number of 3 digits
	// Since the input will already ensure that the CVC is a string of length 3,
	// All we need to do is check that it is a number
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("CVC - digital code with a length of no more than 3 characters")
	}
	return nil
}
