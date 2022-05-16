package validators

import (
	"calc-operator/constants"
	"fmt"

	"github.com/go-logr/logr"
)

type iBaseOperandValidator interface {
	validate(values []float64) ([]string, bool)
}

type operationValidator struct {
	logger logr.Logger
}

type IOperationValidator interface {
	Validate(operation constants.Operation, values []float64) ([]string, bool)
}

func NewOperationValidator(logger logr.Logger) IOperationValidator {
	return &operationValidator{
		logger: logger,
	}
}

func (ov operationValidator) Validate(operation constants.Operation, values []float64) ([]string, bool) {
	var error_messages []string
	var result bool

	// For any operation, minimum 1 operand is necessary
	error_messages, result = ov.validate_operand_count(values, min, 1)
	if !result {
		return error_messages, result
	}

	var validation_fn func([]float64) ([]string, bool)

	switch operation {
	case constants.ADD, constants.SUBTRACT, constants.MULTIPLY:
		{
			validation_fn = ov.validate_for_addition_subtraction_multiplication
		}
	case constants.DIVIDE:
		{
			validation_fn = ov.validate_for_division
		}
	case constants.FACTORIAL:
		{
			validation_fn = ov.validate_for_factorial
		}
	case constants.LOGARITHM:
		{
			validation_fn = ov.validate_for_logarithm
		}
	case constants.EXPONENT:
		{
			validation_fn = ov.validate_for_exponent
		}
	case constants.SIN, constants.COS, constants.SINH, constants.COSH:
		{
			validation_fn = ov.validate_for_trigonometric_functions
		}
	default:
		{
			return []string{fmt.Sprintf("unknown operation %s", operation)}, false
		}
	}

	error_messages, result = validation_fn(values)
	return error_messages, result
}

func (ov operationValidator) validate_operand_count(values []float64, countType operandCountType, count int) ([]string, bool) {
	validator := newOperandCountValidator(countType, count)
	return validator.validate(values)
}

func (ov operationValidator) validate_for_addition_subtraction_multiplication(values []float64) ([]string, bool) {
	// For addition, subtraction, multiplication - minimum 2 operands are required
	return ov.validate_operand_count(values, min, 2)
}

func (ov operationValidator) validate_for_division(values []float64) ([]string, bool) {
	// For division, minimum 2 operands are required.
	error_messages, result := ov.validate_operand_count(values, min, 2)
	if !result {
		return error_messages, result
	}

	// For division, any of the denominators can't have 0
	for index, value := range values {
		if index == 0 {
			continue
		}

		if value == 0 {
			result = false
			error_messages = append(error_messages, fmt.Sprintf("denominator at Index %d cannot be 0!", index))
		}
	}

	return error_messages, result
}

func (ov operationValidator) validate_for_factorial(values []float64) ([]string, bool) {
	// Factorial - only 1 operand
	error_messages, result := ov.validate_operand_count(values, exact, 1)
	if !result {
		return error_messages, result
	}

	operand := values[0]
	if operand < 0 {
		result = false
		error_messages = append(error_messages, fmt.Sprintf("factorial can't be computed for negative number %f!", operand))
	}

	// Factorial is possible only for whole numbers
	if operand != float64(int(operand)) {
		result = false
		error_messages = append(error_messages, fmt.Sprintf("factorial can't be computed for non-whole number %f!", operand))
	}

	return error_messages, result
}

func (ov operationValidator) validate_for_logarithm(values []float64) ([]string, bool) {
	// Logarithm - only 2 operands
	error_messages, result := ov.validate_operand_count(values, exact, 2)
	if !result {
		return error_messages, result
	}

	number, base := values[0], values[1]
	if number <= 0 {
		result = false
		error_messages = append(error_messages, fmt.Sprintf("number %f must be positive!", number))
	}
	if base <= 1 {
		result = false
		error_messages = append(error_messages, fmt.Sprintf("base %f must be > 1!", base))
	}

	return error_messages, result
}

func (ov operationValidator) validate_for_exponent(values []float64) ([]string, bool) {
	// Exponent - only 2 operands
	error_messages, result := ov.validate_operand_count(values, exact, 2)
	if !result {
		return error_messages, result
	}

	number, power := values[0], values[1]
	if number == 0 && power <= 0 {
		result = false
		error_messages = append(error_messages, fmt.Sprintln("0 ^ x where x <= 0 is not valid!"))
	}

	// Exponent - root of negative numbers
	if number < 0 && power > 0 && power < 1 {
		result = false
		error_messages = append(error_messages, fmt.Sprintln("x ^ y where x < 0 and 0 < y < 1 is not valid!"))
	}

	return error_messages, result
}

func (ov operationValidator) validate_for_trigonometric_functions(values []float64) ([]string, bool) {
	// sin, cos, sinh, cosh, - only 1 operand
	error_messages, result := ov.validate_operand_count(values, exact, 1)
	if !result {
		return error_messages, result
	}

	return error_messages, result
}
