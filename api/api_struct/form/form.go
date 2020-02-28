//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package form

// SetForm represent the user request model
type SetForm struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Validate represent the validation method from UserForm
func (v *SetForm) Validate() []string {
	errs := []string{}
	if len(v.Key) < 1 {
		errs = append(errs, "Key can't be empty")
	}

	if len(v.Value) < 1 {
		errs = append(errs, "Value can't be empty")
	}

	return errs
}
