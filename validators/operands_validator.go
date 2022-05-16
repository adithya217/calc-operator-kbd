package validators

import (
	"calc-operator/constants"
	"fmt"

	"github.com/go-logr/logr"
)

type iBaseOperandValidator interface {
	validate(values []float64) ([]string, bool)
}

type operandsValidator struct {
	logger logr.Logger
}

type IOperandsValidator interface {
	Validate(operation constants.Operation, values []float64) ([]string, bool)
}

func NewOperandsValidator(logger logr.Logger) IOperandsValidator {
	return &operandsValidator{
		logger: logger,
	}
}

func (ov operandsValidator) Validate(operation constants.Operation, values []float64) ([]string, bool) {
	var errorMessages []string
	var result bool

	// For any operation, minimum 1 operand is necessary
	errorMessages, result = ov.validateOperandCount(values, min, 1)
	if !result {
		return errorMessages, result
	}

	var validationFn func([]float64) ([]string, bool)

	switch operation {
	case constants.ADD, constants.SUBTRACT, constants.MULTIPLY:
		{
			validationFn = ov.validateForAdditionSubtractionMultiplication
		}
	case constants.DIVIDE:
		{
			validationFn = ov.validateForDivision
		}
	case constants.FACTORIAL:
		{
			validationFn = ov.validateForFactorial
		}
	case constants.LOGARITHM:
		{
			validationFn = ov.validateForLogarithm
		}
	case constants.EXPONENT:
		{
			validationFn = ov.validateForExponent
		}
	case constants.SIN, constants.COS, constants.SINH, constants.COSH:
		{
			validationFn = ov.validateForTrigonometricFunctions
		}
	default:
		{
			return []string{fmt.Sprintf("cannot validate operands due to unknown operation %s", operation)}, false
		}
	}

	errorMessages, result = validationFn(values)
	return errorMessages, result
}

func (ov operandsValidator) validateOperandCount(values []float64, countType operandCountType, count int) ([]string, bool) {
	validator := newOperandCountValidator(countType, count)
	return validator.validate(values)
}

func (ov operandsValidator) validateForAdditionSubtractionMultiplication(values []float64) ([]string, bool) {
	// For addition, subtraction, multiplication - minimum 2 operands are required
	return ov.validateOperandCount(values, min, 2)
}

func (ov operandsValidator) validateForDivision(values []float64) ([]string, bool) {
	// For division, minimum 2 operands are required.
	errorMessages, result := ov.validateOperandCount(values, min, 2)
	if !result {
		return errorMessages, result
	}

	// For division, any of the denominators can't have 0
	for index, value := range values {
		if index == 0 {
			continue
		}

		if value == 0 {
			result = false
			errorMessages = append(errorMessages, fmt.Sprintf("denominator at Index %d cannot be 0!", index))
		}
	}

	return errorMessages, result
}

func (ov operandsValidator) validateForFactorial(values []float64) ([]string, bool) {
	// Factorial - only 1 operand
	errorMessages, result := ov.validateOperandCount(values, exact, 1)
	if !result {
		return errorMessages, result
	}

	operand := values[0]
	if operand < 0 {
		result = false
		errorMessages = append(errorMessages, fmt.Sprintf("factorial can't be computed for negative number %f!", operand))
	}

	// Factorial is possible only for whole numbers
	if operand != float64(int(operand)) {
		result = false
		errorMessages = append(errorMessages, fmt.Sprintf("factorial can't be computed for non-whole number %f!", operand))
	}

	return errorMessages, result
}

func (ov operandsValidator) validateForLogarithm(values []float64) ([]string, bool) {
	// Logarithm - only 2 operands
	errorMessages, result := ov.validateOperandCount(values, exact, 2)
	if !result {
		return errorMessages, result
	}

	number, base := values[0], values[1]
	if number <= 0 {
		result = false
		errorMessages = append(errorMessages, fmt.Sprintf("number %f must be positive!", number))
	}
	if base <= 1 {
		result = false
		errorMessages = append(errorMessages, fmt.Sprintf("base %f must be > 1!", base))
	}

	return errorMessages, result
}

func (ov operandsValidator) validateForExponent(values []float64) ([]string, bool) {
	// Exponent - only 2 operands
	errorMessages, result := ov.validateOperandCount(values, exact, 2)
	if !result {
		return errorMessages, result
	}

	number, power := values[0], values[1]
	if number == 0 && power <= 0 {
		result = false
		errorMessages = append(errorMessages, fmt.Sprintln("0 ^ x where x <= 0 is not valid!"))
	}

	// Exponent - root of negative numbers
	if number < 0 && power > 0 && power < 1 {
		result = false
		errorMessages = append(errorMessages, fmt.Sprintln("x ^ y where x < 0 and 0 < y < 1 is not valid!"))
	}

	return errorMessages, result
}

func (ov operandsValidator) validateForTrigonometricFunctions(values []float64) ([]string, bool) {
	// sin, cos, sinh, cosh, - only 1 operand
	errorMessages, result := ov.validateOperandCount(values, exact, 1)
	if !result {
		return errorMessages, result
	}

	return errorMessages, result
}
