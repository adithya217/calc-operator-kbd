/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"calc-operator/constants"
	"calc-operator/validators"
	"fmt"
	"runtime"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CalculatorSpec defines the desired state of Calculator
type CalculatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The type of the operation
	Operation constants.Operation `json:"operation,omitempty"`

	Operands []float64 `json:"operands,omitempty"`

	// Foo is an example field of Calculator. Edit calculator_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
}

// CalculatorStatus defines the observed state of Calculator
type CalculatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Indicates status of operation - success/failure/in-progress etc.,
	Status string `json:"status,omitempty"`

	// Reason for failure if status is failure
	Reason string `json:"reason,omitempty"`

	// Result of the operation
	Result float64 `json:"result,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Calculator is the Schema for the calculators API
type Calculator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CalculatorSpec   `json:"spec,omitempty"`
	Status CalculatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CalculatorList contains a list of Calculator
type CalculatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Calculator `json:"items"`
}

func init() {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("Calculator types init called from %s#%d\n", file, no)
	}
	SchemeBuilder.Register(&Calculator{}, &CalculatorList{})
}

func (calc *Calculator) SetDefaults() {
	calc.setDefaultOperation()
	calc.setDefaultOperands()
}

func (calc *Calculator) setDefaultOperation() {
	if calc.Spec.Operation == "" {
		calc.Spec.Operation = constants.MULTIPLY
	}
}

func (calc *Calculator) setDefaultOperands() {
	if len(calc.Spec.Operands) == 0 {
		calc.Spec.Operands = []float64{1, 1}
	}
}

func (r *Calculator) validate() error {
	validator := validators.NewOperandsValidator(calculatorlog)
	errorMessages, validationResult := validator.Validate(r.Spec.Operation, r.Spec.Operands)
	if validationResult {
		calculatorlog.Info("validated operands from CR succesfully!", "name", r.Name)
		return nil
	}

	err := fmt.Errorf("validation failed for operands: %s", errorMessages)
	calculatorlog.Error(err, "operands validation failed!", "name", r.Name)

	operandsFieldPath := field.NewPath(constants.SPEC).Child(constants.OPERANDS)
	operandsErr := field.Invalid(operandsFieldPath, r.Spec.Operands, err.Error())

	var allErrors field.ErrorList
	allErrors = append(allErrors, operandsErr)
	return errors.NewInvalid(
		schema.GroupKind{Group: "webapp.demo.calc-operator", Kind: "Calculator"},
		r.Name, allErrors)
}
