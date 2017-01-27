package rlib

import (
	"encoding/json"
	"testing"
)

// Data conversion tests

type foo1 struct {
	I1 int64
}

type bar1 struct {
	I1 XJSONYesNo
}

type foobar1 struct {
	I1 XJSONYesNo
	I2 XJSONYesNo
}

type foobar2 struct {
	I1 int64
	I2 int64
}

type testDataConv1 struct {
	I1     int64
	Expect XJSONYesNo
}

var td1 = []testDataConv1{
	{0, "no"},
	{1, "yes"},
	{99, "no"},
}
var td2 = []testDataConv1{
	{0, "no"},
	{1, "yes"},
	{0, "xyz"},
}

// TestConversion1 tests the conversion between int and XJSONYesNo
func TestConversion1(t *testing.T) {
	// int -> XJSONYesNo
	for i := 0; i < len(td1); i++ {
		var a foo1
		var b bar1
		a.I1 = td1[i].I1
		MigrateStructVals(&a, &b)
		if b.I1 != td1[i].Expect {
			t.Errorf("int2YesNo( %d )  Expect %s, got %s\n", td1[i].I1, td1[i].Expect, b.I1)
		}
	}

	// XJSONYesNo --> int
	for i := 0; i < len(td2); i++ {
		var a foo1
		var b bar1
		b.I1 = td2[i].Expect
		MigrateStructVals(&b, &a)
		if a.I1 != td2[i].I1 {
			t.Errorf("yesNo2Int( %s )  Expect %d, got %d\n", td2[i].Expect, td2[i].I1, a.I1)
		}
	}

	var b = foobar1{I1: "no", I2: "yes"}
	var a foobar2
	var c foobar1

	u, err := json.Marshal(&b)
	if err != nil {
		t.Errorf("Error marshaling json data: %s\n", err.Error())
		return
	}

	MigrateStructVals(&b, &a)

	err = json.Unmarshal(u, &c)
	if err != nil {
		t.Errorf("Error unmarshaling json data: %s\n", err.Error())
		return
	}

	if b.I1 != c.I1 || b.I2 != c.I2 {
		t.Errorf("Error: b != c after migration: b = %#v, c = %#v\n", b, c)
	}
}
