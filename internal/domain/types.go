package domain

import (
	"regexp"
	"strconv"
)

type BoilerplateId struct {
	value string
}

func NewBoilerplateId(value string) (*BoilerplateId, error) {
	if IsUUIDv4(value) {
		return &BoilerplateId{value}, nil
	}

	return nil, ErrBoilerplateBadRequest
}

func (b *BoilerplateId) Value() string {
	return b.value
}

// ---------------------------------------------------------------------
type BoilerplateName struct {
	value string
}

const BOILERPLATE_NAME_MAX_LENGTH = 100
const BOILERPLATE_NAME_MIN_LENGTH = 3

func NewBoilerplateName(value string) (*BoilerplateName, error) {
	if isValidBoilerplateName(value) {
		return &BoilerplateName{value}, nil
	}

	return nil, ErrBoilerplateBadRequest
}

func isValidBoilerplateName(value string) bool {
	if len(value) > BOILERPLATE_NAME_MAX_LENGTH || len(value) < BOILERPLATE_NAME_MIN_LENGTH {
		return false
	}

	pattern := `^[a-zA-Z]{1}[a-zA-Z ]+[a-zA-Z]$|^[a-zA-Z]{1}$`
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return false
	}

	return matched
}

func (b *BoilerplateName) Value() string {
	return b.value
}

//---------------------------------------------------------------------

type BoilerplateSearchTerm struct {
	value string
}

const BOILERPLATE_SEARCH_TERM_MAX_LENGTH = 25
const BOILERPLATE_SEARCH_TERM_MIN_LENGTH = 3

func NewBoilerplateSearchTerm(value string) (*BoilerplateSearchTerm, error) {
	if isValidBoilerplateSearchTerm(value) {
		return &BoilerplateSearchTerm{value}, nil
	}

	return nil, ErrBoilerplateBadRequest
}

func isValidBoilerplateSearchTerm(value string) bool {
	if len(value) > BOILERPLATE_SEARCH_TERM_MAX_LENGTH || len(value) < BOILERPLATE_SEARCH_TERM_MIN_LENGTH {
		return false
	}

	pattern := `^[a-zA-Z]{1}[a-zA-Z ]+[a-zA-Z]$|^[a-zA-Z]{1}$`
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return false
	}

	return matched
}

func (b *BoilerplateSearchTerm) Value() string {
	return b.value
}

// ---------------------------------------------------------------------
type BoilerplatePositiveNumber struct {
	value string
}

func NewBoilerplatePositiveNumber(value string) (*BoilerplatePositiveNumber, error) {
	if isValidBoilerplatePositiveNumber(value) {
		return &BoilerplatePositiveNumber{value}, nil
	}

	return nil, ErrBoilerplateBadRequest
}

func isValidBoilerplatePositiveNumber(value string) bool {
	number, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return number > 0
}

func (b *BoilerplatePositiveNumber) Value() string {
	return b.value
}

func (b *BoilerplatePositiveNumber) ValueToInt() int {
	number, _ := strconv.Atoi(b.value)

	return number
}
