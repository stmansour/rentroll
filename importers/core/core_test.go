package core

import (
	"reflect"
	"testing"
)

// ======== TEST FOR METHODS DEFINED IN `struct_utils.go` ========

// Testing for `struct fields`
func TestGetStructFields(t *testing.T) {

	type Sample struct {
		a int
		b string
	}
	var ok bool

	// CASE: POSITIVE
	// expected output
	outputFields := []string{"a", "b"}

	// create var for struct
	var sample Sample
	fields, _ := GetStructFields(&sample)

	// check that both are equal
	if !reflect.DeepEqual(fields, outputFields) {
		t.Errorf("[TestGetStructFields] Expected field list is `%v`, but it returned `%v`", outputFields, fields)
	}

	// CASE: NEGATIVE
	var a int
	_, ok = GetStructFields(&a)
	if ok {
		t.Errorf("[TestGetStructFields] Expected `false` for ok, but it returned `%v` for `var a int`", ok)
	}
}

// ======== TEST FOR METHODS DEFINED IN `utils.go` ========

// Testing for `string in slice`
func TestStringInSlice(t *testing.T) {

	sampleSlice := []string{"a", "b", "c"}

	// CASE: POSITIVE
	var a = "a"
	ok := StringInSlice(a, sampleSlice)
	if !ok {
		t.Errorf("[TestStringInSlice] Expected `true`, but it returned `%v` for `%v` in slice `%v`", ok, a, sampleSlice)
	}

	// CASE: NEGATIVE
	var z = "z"
	ok = StringInSlice(z, sampleSlice)
	if ok {
		t.Errorf("[TestStringInSlice] Expected `false`, but it returned `%v` for `%v` in slice `%v`", ok, a, sampleSlice)
	}
}

// Testing for `string in slice`
func TestIntegerInSlice(t *testing.T) {

	sampleSlice := []int{1, 2, 3}

	// CASE: POSITIVE
	var a = 1
	ok := IntegerInSlice(a, sampleSlice)
	if !ok {
		t.Errorf("[TestIntegerInSlice] Expected `true`, but it returned `%v` for `%v` in slice `%v`", ok, a, sampleSlice)
	}

	// CASE: NEGATIVE
	var z = 8
	ok = IntegerInSlice(z, sampleSlice)
	if ok {
		t.Errorf("[TestIntegerInSlice] Expected `false`, but it returned `%v` for `%v` in slice `%v`", ok, a, sampleSlice)
	}
}

// Testing for `email validation`
func TestEmailValidation(t *testing.T) {

	validValues := []string{"test@test.c", "test@test."}
	invalidValues := []string{"test-a@", "testa.c"}

	// CASE: POSITIVE
	for _, email := range validValues {
		ok := IsValidEmail(email)
		if !ok {
			t.Errorf("[TestEmailValidation] Expected `true`, but it returned `%v` for `%s`", ok, email)
		}
	}

	// CASE: NEGATIVE
	for _, email := range invalidValues {
		ok := IsValidEmail(email)
		if ok {
			t.Errorf("[TestEmailValidation] Expected `false`, but it returned `%v` for `%s`", ok, email)
		}
	}
}
