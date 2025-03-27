package domain_test

import (
	"strings"
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
)

func TestCanCreateValidBoilerplateName(t *testing.T) {
	validNames := []string{
		"ValidName",
		"AnotherValidName",
		"Valid Name With Spaces",
		"Short",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"}
	for _, name := range validNames {
		if _, err := domain.NewBoilerplateName(name); err != nil {
			t.Errorf("Expected valid name, but got invalid: %s", name)
		}
	}

	invalidNames := []string{
		" InvalidName",
		"InvalidName ",
		"Invalid Name ",
		" Invalid Name",
		"Invalid@Name",
		"Invalid#Name",
		"Invalid123Name",
		"Invalid Name With Spaces ",
		" Invalid Name With Spaces",
		"Invalid Name With Spaces And Symbols!",
	}
	for _, name := range invalidNames {
		if _, err := domain.NewBoilerplateName(name); err == nil {
			t.Errorf("Expected valid name, but got invalid: %s", name)
		}
	}

	invalidLengthNames := []string{
		strings.Repeat("a", domain.BOILERPLATE_NAME_MAX_LENGTH+1),
		strings.Repeat("b", domain.BOILERPLATE_NAME_MIN_LENGTH-1),
		"",
	}
	for _, name := range invalidLengthNames {
		if _, err := domain.NewBoilerplateName(name); err == nil {
			t.Errorf("Expected invalid name, but got invalid: %s", name)
		}
	}

}
