package services

import (
	"calc-operator/constants"
	"fmt"
	"math"

	"github.com/go-logr/logr"
)

type calculatorService struct {
	logger logr.Logger
}

type ICalculatorService interface {
	Compute(operation constants.Operation, values []float64) (float64, error)
}

func NewCalculatorService(logger logr.Logger) ICalculatorService {
	return &calculatorService{
		logger: logger,
	}
}

func (cs calculatorService) Compute(operation constants.Operation, values []float64) (float64, error) {
	var computation_fn func([]float64) (float64, error)

	switch operation {
	case constants.ADD:
		{
			computation_fn = compute_addition
		}
	case constants.SUBTRACT:
		{
			computation_fn = compute_subtraction
		}
	case constants.MULTIPLY:
		{
			computation_fn = compute_multiplication
		}
	case constants.DIVIDE:
		{
			computation_fn = compute_division
		}
	case constants.FACTORIAL:
		{
			computation_fn = compute_factorial
		}
	case constants.LOGARITHM:
		{
			computation_fn = compute_logarithm
		}
	case constants.EXPONENT:
		{
			computation_fn = compute_exponent
		}
	case constants.SIN:
		{
			computation_fn = compute_sine
		}
	case constants.COS:
		{
			computation_fn = compute_cosine
		}
	case constants.SINH:
		{
			computation_fn = compute_sine_hyperbolic
		}
	case constants.COSH:
		{
			computation_fn = compute_cosine_hyperbolic
		}
	default:
		{
			return 0, fmt.Errorf(fmt.Sprintf("Unknown operation %s!", operation))
		}
	}

	return computation_fn(values)
}

func compute_addition(values []float64) (float64, error) {
	var result float64 = 0
	for _, value := range values {
		result += value
	}

	return result, nil
}

func compute_subtraction(values []float64) (float64, error) {
	var result float64 = 0
	for _, value := range values {
		result -= value
	}

	return result, nil
}

func compute_multiplication(values []float64) (float64, error) {
	var result float64 = 0
	for _, value := range values {
		result *= value
	}

	return result, nil
}

func compute_division(values []float64) (float64, error) {
	var result float64 = 0
	for _, value := range values {
		result /= value
	}

	return result, nil
}

func compute_factorial(values []float64) (float64, error) {
	var result uint64
	operand := uint64(values[0])
	for i := uint64(1); i <= operand; i++ {
		result *= uint64(i) // mismatched types int64 and int
	}

	return float64(result), nil
}

func compute_logarithm(values []float64) (float64, error) {
	number, base := values[0], values[1]
	return math.Log2(number) / math.Log2(base), nil
}

func compute_exponent(values []float64) (float64, error) {
	number, power := values[0], values[1]
	return math.Pow(number, power), nil
}

func compute_sine(values []float64) (float64, error) {
	operand := values[0]
	return math.Sin(operand), nil
}

func compute_cosine(values []float64) (float64, error) {
	operand := values[0]
	return math.Cos(operand), nil
}

func compute_sine_hyperbolic(values []float64) (float64, error) {
	operand := values[0]
	return math.Sinh(operand), nil
}

func compute_cosine_hyperbolic(values []float64) (float64, error) {
	operand := values[0]
	return math.Cosh(operand), nil
}
