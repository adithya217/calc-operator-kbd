package validators

import "fmt"

type operandCountType string

const (
	min   operandCountType = "min"
	max   operandCountType = "max"
	exact operandCountType = "exact"
)

type operandCountValidator struct {
	ocType operandCountType
	count  int
}

func newOperandCountValidator(ocType operandCountType, count int) iBaseOperandValidator {
	return &operandCountValidator{
		ocType: ocType,
		count:  count,
	}
}

func (ocv operandCountValidator) validate(values []float64) ([]string, bool) {
	switch ocv.ocType {
	case min:
		return ocv.validate_min(values)
	case max:
		return ocv.validate_max(values)
	case exact:
		return ocv.validate_exact(values)
	default:
		return []string{fmt.Sprintf("unknown operandCountType %s!", ocv.ocType)}, false
	}
}

func (ocv operandCountValidator) validate_min(values []float64) ([]string, bool) {
	count := len(values)

	if count < ocv.count {
		return []string{fmt.Sprintf("expected at-least %d values but found %d values", ocv.count, count)}, false
	}

	return []string{}, true
}

func (ocv operandCountValidator) validate_max(values []float64) ([]string, bool) {
	count := len(values)

	if count > ocv.count {
		return []string{fmt.Sprintf("expected at-most %d values but found %d values", ocv.count, count)}, false
	}

	return []string{}, true
}

func (ocv operandCountValidator) validate_exact(values []float64) ([]string, bool) {
	count := len(values)

	if count != ocv.count {
		return []string{fmt.Sprintf("expected exactly %d values but found %d values", ocv.count, count)}, false
	}

	return []string{}, true
}
