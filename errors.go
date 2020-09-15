package tmdgolangbase

import "fmt"

func newErrEnvironmentBlankVar(key string) (err error) {
	err = fmt.Errorf("%v environment variable is blank", key)
	return
}

func newErrFieldNotFound(fieldName string) (err error) {
	err = fmt.Errorf("field %s not found", fieldName)
	return
}

func newErrFieldInvalidType(fieldName, fieldType string) (err error) {
	err = fmt.Errorf("%s field is not a/an %v type", fieldName, fieldType)
	return
}
