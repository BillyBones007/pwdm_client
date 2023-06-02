package validcard

import (
	"testing"
)

func TestCnValidator(t *testing.T) {
	// Valid card number
	if err := CnValidator("4111 1111 1111 1111"); err != nil {
		t.Errorf("expected nil error but got %v", err)
	}
	// Invalid card number - too long
	if err := CnValidator("4111 1111 1111 1111 111"); err == nil {
		t.Error("expected an error but got nil")
	}
	// Invalid card number - last digit not a number or space
	if err := CnValidator("4111 1111 1111 111x"); err == nil {
		t.Error("expected an error but got nil")
	}
	// Invalid card number - must separate groups with spaces
	if err := CnValidator("4111 1111111 1111"); err == nil {
		t.Error("expected an error but got nil")
	}
	// Invalid card number - remaining digits not integers
	if err := CnValidator("4111 1111 1x11 1111"); err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestExpValidator(t *testing.T) {
	// Valid EXP
	if err := ExpValidator("12/22"); err != nil {
		t.Errorf("expected nil error but got %v", err)
	}
	// Invalid EXP - not all numbers
	if err := ExpValidator("1x/22"); err == nil {
		t.Error("expected an error but got nil")
	}
	// Invalid EXP - slash not in the 2nd index
	if err := ExpValidator("1/2/22"); err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestCVCValidator(t *testing.T) {
	// Valid CVC
	if err := CVCValidator("123"); err != nil {
		t.Errorf("expected nil error but got %v", err)
	}
	// Invalid CVC - not a number
	if err := CVCValidator("1x3"); err == nil {
		t.Error("expected an error but got nil")
	}
	// Invalid CVC - too long
	if err := CVCValidator("1234"); err == nil {
		t.Error("expected an error but got nil")
	}
}
