package tmdgolangbase

import (
	"crypto/sha256"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

// DecimalStringToInt64 ...
func DecimalStringToInt64(text string) (number int64, err error) {
	number, err = strconv.ParseInt(text, 10, 64)
	return
}

// GetEnvVar ...
func GetEnvVar(key string) (value string, err error) {
	value = os.Getenv(key)
	if value == "" {
		err = newErrEnvironmentBlankVar(key)
	}
	return
}

// GetEnvVarInt64 ...
func GetEnvVarInt64(key string) (value int64, err error) {
	valueString, err := GetEnvVar(key)
	if err != nil {
		return
	}
	value, err = DecimalStringToInt64(valueString)
	return
}

func fieldValueOf(dataStruct interface{}, fieldName string) (fieldType reflect.Type, fieldValue interface{}, err error) {
	e := reflect.ValueOf(dataStruct).Elem()
	for i := 0; i < e.NumField(); i++ {
		if fieldName == e.Type().Field(i).Name {
			fieldType = e.Type().Field(i).Type
			fieldValue = e.Field(i).Interface()
			return
		}
	}
	err = newErrFieldNotFound(fieldName)
	return
}

// GetUintFieldValue ...
func GetUintFieldValue(dataStruct interface{}, fieldName string) (value uint, err error) {
	fieldType, fieldValue, err := fieldValueOf(dataStruct, fieldName)
	if err != nil {
		return
	}
	if reflect.TypeOf(uint(0)) != fieldType {
		err = newErrFieldInvalidType(fieldName, "uint")
		return
	}
	value = fieldValue.(uint)
	return
}

// GetInt64FieldValue ...
func GetInt64FieldValue(dataStruct interface{}, fieldName string) (value int64, err error) {
	fieldType, fieldValue, err := fieldValueOf(dataStruct, fieldName)
	if err != nil {
		return
	}
	if reflect.TypeOf(int64(0)) != fieldType {
		err = newErrFieldInvalidType(fieldName, "int64")
		return
	}
	value = fieldValue.(int64)
	return
}

// GetStringFieldValue ...
func GetStringFieldValue(dataStruct interface{}, fieldName string) (value string, err error) {
	fieldType, fieldValue, err := fieldValueOf(dataStruct, fieldName)
	if err != nil {
		return
	}
	if reflect.TypeOf(string("")) != fieldType {
		err = newErrFieldInvalidType(fieldName, "string")
		return
	}
	value = fieldValue.(string)
	return
}

// SumSHA256 ...
func SumSHA256(text string) (sum string) {
	textBytes := []byte(text)
	sumBytes := sha256.Sum256(textBytes)
	sum = fmt.Sprintf("%x", sumBytes)
	return
}

// TimeNow ...
func TimeNow() *time.Time {
	now := time.Now().UTC()
	return &now
}
