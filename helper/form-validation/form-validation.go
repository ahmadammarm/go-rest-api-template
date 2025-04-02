package formvalidation

import "github.com/go-playground/validator/v10"

func FormValidationError(err error) map[string]string {
    errors := make(map[string]string)

    for _, err := range err.(validator.ValidationErrors) {
        errorField := err.StructField()
        errors[errorField] = err.Tag()
    }

    return errors
}