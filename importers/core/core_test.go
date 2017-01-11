package core

import (
	"reflect"
	"testing"
)

// ======== TEST FOR METHODS DEFINED IN `struct_utils.go` ========

// Testing for `field map of struct`
func TestBuildFieldMap(t *testing.T) {

	type Sample struct {
		a int
		b string
	}

	// CASE: POSITIVE
	// expected output
	outputSampleMap := map[string]int{
		"a": 0,
		"b": 1,
	}

	// create var for struct
	var sample Sample
	fmap, ok := BuildFieldMap(&sample)

	// check that both are equal
	if !reflect.DeepEqual(fmap, outputSampleMap) {
		t.Errorf("[TestBuildFieldMap] Expected field map is `%v`, but it returned `%v`", outputSampleMap, fmap)
	}

	// CASE: NEGATIVE
	var a int
	fmap, ok = BuildFieldMap(&a)
	if ok {
		t.Errorf("[TestBuildFieldMap] Expected `false` for ok, but it returned `%v` for `var a int`", ok)
	}
}

// Testing for `struct fields`
func TestGetStructFields(t *testing.T) {

	type Sample struct {
		a int
		b string
	}

	// CASE: POSITIVE
	// expected output
	outputFields := []string{"a", "b"}

	// create var for struct
	var sample Sample
	fields, ok := GetStructFields(&sample)

	// check that both are equal
	if !reflect.DeepEqual(fields, outputFields) {
		t.Errorf("[TestGetStructFields] Expected field list is `%v`, but it returned `%v`", outputFields, fields)
	}

	// CASE: NEGATIVE
	var a int
	fields, ok = GetStructFields(&a)
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

// Testing for `string is integer value?`
func TestIsIntString(t *testing.T) {

	validValues := []string{"10", "+10", "-100"}
	invalidValues := []string{"+", "a", "1000.123", "-1.123"}

	// CASE: POSITIVE
	for _, value := range validValues {
		ok := IsIntString(value)
		if !ok {
			t.Errorf("[TestIsIntString] Expected `true`, but it returned `%v` for `%s`", ok, value)
		}
	}

	// CASE: NEGATIVE
	for _, value := range invalidValues {
		ok := IsIntString(value)
		if ok {
			t.Errorf("[TestIsIntString] Expected `false`, but it returned `%v` for `%s`", ok, value)
		}
	}
}

// Testing for `string is positive integer value?`
func TestIsUIntString(t *testing.T) {

	validValues := []string{"10", "0"}
	invalidValues := []string{"+", "a", "1000.123", "-1.123"}

	// CASE: POSITIVE
	for _, value := range validValues {
		ok := IsUIntString(value)
		if !ok {
			t.Errorf("[TestIsUIntString] Expected `true`, but it returned `%v` for `%s`", ok, value)
		}
	}

	// CASE: NEGATIVE
	for _, value := range invalidValues {
		ok := IsUIntString(value)
		if ok {
			t.Errorf("[TestIsUIntString] Expected `false`, but it returned `%v` for `%s`", ok, value)
		}
	}
}

// Testing for `string is float value?`
func TestIsFloatString(t *testing.T) {

	validValues := []string{"10.0", "0", "-100.123", "10"}
	invalidValues := []string{"+", "a"}

	// CASE: POSITIVE
	for _, value := range validValues {
		ok := IsFloatString(value)
		if !ok {
			t.Errorf("[TestIsFloatString] Expected `true`, but it returned `%v` for `%s`", ok, value)
		}
	}

	// CASE: NEGATIVE
	for _, value := range invalidValues {
		ok := IsFloatString(value)
		if ok {
			t.Errorf("[TestIsFloatString] Expected `false`, but it returned `%v` for `%s`", ok, value)
		}
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
