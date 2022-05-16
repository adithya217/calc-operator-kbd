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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var calculatorlog = logf.Log.WithName("calculator-resource")

func (r *Calculator) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-webapp-demo-calc-operator-v1-calculator,mutating=true,failurePolicy=fail,sideEffects=None,groups=webapp.demo.calc-operator,resources=calculators,verbs=create;update,versions=v1,name=mcalculator.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Calculator{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Calculator) Default() {
	calculatorlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	if r.Spec.Operation == "" {
		r.Spec.Operation = constants.ADD
	}
	if len(r.Spec.Operands) == 0 {
		r.Spec.Operands = []float64{0, 1}
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-webapp-demo-calc-operator-v1-calculator,mutating=false,failurePolicy=fail,sideEffects=None,groups=webapp.demo.calc-operator,resources=calculators,verbs=create;update,versions=v1,name=vcalculator.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Calculator{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Calculator) ValidateCreate() error {
	calculatorlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Calculator) ValidateUpdate(old runtime.Object) error {
	calculatorlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Calculator) ValidateDelete() error {
	calculatorlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
