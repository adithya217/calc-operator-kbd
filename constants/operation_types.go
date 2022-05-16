package constants

import (
	"fmt"
	"strings"
)

// Operation describes the type of operation to be performed by the Calculator.
// It must be mandatorily specified.
// +kubebuilder:validation:Enum=add;sub;mul;div;fact;log;exp;sin;cos;sinh;cosh
type Operation string

const (
	ADD       Operation = "add"
	SUBTRACT  Operation = "sub"
	MULTIPLY  Operation = "mul"
	DIVIDE    Operation = "div"
	FACTORIAL Operation = "fact"
	LOGARITHM Operation = "log"
	EXPONENT  Operation = "exp"
	SIN       Operation = "sin"
	COS       Operation = "cos"
	SINH      Operation = "sinh"
	COSH      Operation = "cosh"
)

// Golang doesn't have constant for map. This is a private map to be maintained.
// Whenever new operations are added above, add them here as well
// TODO: Maybe this functionality can be de-coupled into some "make" like process, for auto code gen.
var invertedOperationMap = map[string]Operation{
	"add":  ADD,
	"sub":  SUBTRACT,
	"mul":  MULTIPLY,
	"div":  DIVIDE,
	"fact": FACTORIAL,
	"log":  LOGARITHM,
	"exp":  EXPONENT,
	"sin":  SIN,
	"cos":  COS,
	"sinh": SINH,
	"cosh": COSH,
}

func (o Operation) String() string {
	return string(o)
}

func ParseOperation(operationString string) (Operation, error) {
	operation, found := invertedOperationMap[strings.ToLower(operationString)]
	if !found {
		return "", fmt.Errorf("cannot parse [%s] as Operation", operation)
	}
	return operation, nil
}
