package controller

import (
	"github.com/angelokurtis/kube-operator/pkg/controller/kurtisservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kurtisservice.Add)
}
