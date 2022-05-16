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
	var computationFn func([]float64) (float64, error)

	switch operation {
	case constants.ADD:
		{
			computationFn = computeAddition
		}
	case constants.SUBTRACT:
		{
			computationFn = computeSubtraction
		}
	case constants.MULTIPLY:
		{
			computationFn = computeMultiplication
		}
	case constants.DIVIDE:
		{
			computationFn = computeDivision
		}
	case constants.FACTORIAL:
		{
			computationFn = computeFactorial
		}
	case constants.LOGARITHM:
		{
			computationFn = computeLogarithm
		}
	case constants.EXPONENT:
		{
			computationFn = computeExponent
		}
	case constants.SIN:
		{
			computationFn = computeSine
		}
	case constants.COS:
		{
			computationFn = computeCosine
		}
	case constants.SINH:
		{
			computationFn = computeSineHyperbolic
		}
	case constants.COSH:
		{
			computationFn = computeCosineHyperbolic
		}
	default:
		{
			return 0, fmt.Errorf(fmt.Sprintf("unknown operation %s!", operation))
		}
	}

	return computationFn(values)
}

func computeAddition(values []float64) (float64, error) {
	var result float64 = 0
	for _, value := range values {
		result += value
	}

	return result, nil
}

func computeSubtraction(values []float64) (float64, error) {
	var result float64 = values[0]
	for index, value := range values {
		if index == 0 {
			continue
		}
		result -= value
	}

	return result, nil
}

func computeMultiplication(values []float64) (float64, error) {
	var result float64 = 1
	for _, value := range values {
		result *= value
	}

	return result, nil
}

func computeDivision(values []float64) (float64, error) {
	var result float64 = values[0]
	for index, value := range values {
		if index == 0 {
			continue
		}
		result /= value
	}

	return result, nil
}

func computeFactorial(values []float64) (float64, error) {
	var result uint64 = 1
	operand := uint64(values[0])
	for i := uint64(1); i <= operand; i++ {
		result *= uint64(i) // mismatched types int64 and int
	}

	return float64(result), nil
}

func computeLogarithm(values []float64) (float64, error) {
	number, base := values[0], values[1]
	return math.Log2(number) / math.Log2(base), nil
}

func computeExponent(values []float64) (float64, error) {
	number, power := values[0], values[1]
	return math.Pow(number, power), nil
}

func computeSine(values []float64) (float64, error) {
	operand := values[0]
	return math.Sin(operand), nil
}

func computeCosine(values []float64) (float64, error) {
	operand := values[0]
	return math.Cos(operand), nil
}

func computeSineHyperbolic(values []float64) (float64, error) {
	operand := values[0]
	return math.Sinh(operand), nil
}

func computeCosineHyperbolic(values []float64) (float64, error) {
	operand := values[0]
	return math.Cosh(operand), nil
}
