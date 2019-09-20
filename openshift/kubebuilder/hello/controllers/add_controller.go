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

	arithmeticv1 "github.com/sthaha/k8s-ctrl-simple/api/v1"
)

// AddReconciler reconciles a Add object
type AddReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=arithmetic.simple.math.xyz,resources=adds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=arithmetic.simple.math.xyz,resources=adds/status,verbs=get;update;patch

func (r *AddReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("add", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *AddReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&arithmeticv1.Add{}).
		Complete(r)
}
