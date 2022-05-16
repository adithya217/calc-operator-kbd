package models

import (
	v1 "calc-operator/api/v1"
	"calc-operator/constants"
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type calculatorModel struct {
	data *v1.Calculator
}

func (cm calculatorModel) GetModel() v1.Calculator {
	return *cm.data
}

func (cm calculatorModel) MarkFail(reason string) {
	cm.data.Status.Status = constants.FAILED.String()
	cm.data.Status.Reason = reason
}

func (cm calculatorModel) MarkSuccess(result float64) {
	cm.data.Status.Status = constants.SUCCESS.String()
	cm.data.Status.Result = result
}

type ICalculatorModel interface {
	GetModel() v1.Calculator
	MarkFail(string)
	MarkSuccess(float64)
}

func NewCalculatorModel(ctx context.Context, client client.Client, name types.NamespacedName) (ICalculatorModel, error) {
	model := &calculatorModel{
		data: &v1.Calculator{},
	}

	if err := client.Get(ctx, name, model.data); err != nil {
		return model, fmt.Errorf("failed to get calculator data: %w", err)
	}

	return model, nil
}
