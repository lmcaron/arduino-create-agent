// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// indexes views
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package views

import (
	goa "goa.design/goa"
)

// Operation is the viewed result type that is projected based on a view.
type Operation struct {
	// Type to project
	Projected *OperationView
	// View to render
	View string
}

// OperationView is a type that runs validations on a projected type.
type OperationView struct {
	// The status of the operation
	Status *string
}

// ValidateOperation runs the validations defined on the viewed result type
// Operation.
func ValidateOperation(result *Operation) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateOperationView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateOperationView runs the validations defined on OperationView using
// the "default" view.
func ValidateOperationView(result *OperationView) (err error) {
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	return
}
