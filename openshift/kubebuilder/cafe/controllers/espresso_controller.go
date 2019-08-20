/*

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

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	coffeev1 "github.com/sthaha/k8s-cafe/api/v1"
)

// EspressoReconciler reconciles a Espresso object
type EspressoReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=coffee.sthaha.cafe.xyz,resources=espressoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coffee.sthaha.cafe.xyz,resources=espressoes/status,verbs=get;update;patch

func (r *EspressoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("espresso", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *EspressoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&coffeev1.Espresso{}).
		Complete(r)
}
