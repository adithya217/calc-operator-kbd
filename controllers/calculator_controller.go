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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	webappv1 "calc-operator/api/v1"
	"calc-operator/models"
	"calc-operator/services"
)

// CalculatorReconciler reconciles a Calculator object
type CalculatorReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=webapp.demo.calc-operator,resources=calculators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.demo.calc-operator,resources=calculators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.demo.calc-operator,resources=calculators/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Calculator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *CalculatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("reconciling calculator CR...")

	calculatorModel, err := models.NewCalculatorModel(ctx, r.Client, req.NamespacedName)
	if err != nil {
		logger.Error(err, "calculator data not found, can't reconcile!", "key", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	logger.Info("parsed data from CR succesfully!")

	resourceService := services.NewResourceService(logger, r.Client, ctx, r.Recorder)
	operation := calculatorModel.GetModel().Spec.Operation

	// resourceService.RecordEvent(calculatorModel.GetModel(), corev1.EventTypeNormal, "in progress",
	// 	fmt.Sprintf("in progress for reconciling CR with name %s", calculatorModel.GetModel().Name))

	service := services.NewCalculatorService(logger)
	computationResult, err := service.Compute(operation, calculatorModel.GetModel().Spec.Operands)
	if err != nil {
		logger.Error(err, "computation failed, can't reconcile!", "key", req.NamespacedName)
		calculatorModel.MarkFail(err.Error())
		resourceService.UpdateStatus(calculatorModel.GetModel())
		return ctrl.Result{}, err
	}

	calculatorModel.MarkSuccess(computationResult)
	resourceService.UpdateStatus(calculatorModel.GetModel())
	logger.Info("computation of CR succesful!", "result", computationResult)

	// resourceService.RecordEvent(calculatorModel.GetModel(), corev1.EventTypeNormal, "success",
	// 	fmt.Sprintf("successfully reconciled CR with name %s", calculatorModel.GetModel().Name))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CalculatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// TODO: Setup webhook for validation
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Calculator{}).
		Complete(r)
}
