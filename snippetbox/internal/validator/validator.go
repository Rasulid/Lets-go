package validator

type Validator struct {
	FieldError map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldError) == 0
}

func (v *Validator) AddFieldError(key, massage string) {
	if v.FieldError == nil {
		v.FieldError = make(map[string]string)
	}

	if _, exists := v.FieldError[key]; !exists {
		v.FieldError[key] = massage
	}

}
