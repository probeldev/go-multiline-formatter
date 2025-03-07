package parser

import (
	"testing"
)

func TestIsStuctFunc(
	t *testing.T,
) {

	input := `func (s *structname) myfunc(
	a int,
	b int,
	c int,
) (
	int,
	error,
) {`

	expected := true

	actual := GetFunctionParser().IsStuctFunc(input)

	if actual != expected {
		t.Error("actual != expected")
	}

}

func TestIsStuctFunc2(
	t *testing.T,
) {

	input := `func myfunc() {`

	expected := false

	actual := GetFunctionParser().IsStuctFunc(input)

	if actual != expected {
		t.Error("actual != expected")
	}

}

func TestGetFuncName(
	t *testing.T,
) {
	input := `func myfunc() {`

	expected := "myfunc"

	actual := GetFunctionParser().GetFuncName(input)

	if actual != expected {
		t.Error(
			"actual != expected",
			"\n",
			"actual:",
			"\n",
			actual,
			"\n",
			"expected:",
			"\n",
			expected,
			"\n",
		)
	}
}

func TestGetFuncName2(
	t *testing.T,
) {
	input := `func (s *mystruct) myfunc() {`

	expected := "myfunc"

	actual := GetFunctionParser().GetFuncName(input)

	if actual != expected {
		t.Error(
			"actual != expected",
			"\n",
			"actual:",
			"\n",
			actual,
			"\n",
			"expected:",
			"\n",
			expected,
			"\n",
		)
	}
}

func TestIsSturctFunc(
	t *testing.T,
) {
	formatter := GetFunctionParser()

	funcString := "func myfunc(a int, b int, c int) error"
	expected := false
	actual := formatter.IsStuctFunc(funcString)

	if actual != expected {
		t.Error(
			"test case 1:",
			"actual != expected",
			"actual:", actual,
			"expected:", expected,
		)
	}

	funcString = "func (s *mystruct) myfunc(a int, b int, c int) error"
	expected = true
	actual = formatter.IsStuctFunc(funcString)

	if actual != expected {
		t.Error(
			"test case 2:",
			"actual != expected",
			"actual:", actual,
			"expected:", expected,
		)
	}
}

func TestGetStructFromFunc(
	t *testing.T,
) {
	formatter := GetFunctionParser()

	funcString := "func (s *mystruct) myfunc(a int, b int, c int) error"
	expected := "(s *mystruct)"
	actual := formatter.GetStructFromFunc(funcString)

	if expected != actual {
		t.Error(
			"actual != expected",
			"actual:", actual,
			"expected:", expected,
		)
	}
}

func TestGetOnlyStructNameFromFunc(
	t *testing.T,
) {
	formatter := GetFunctionParser()

	funcString := "func (s *mystruct) myfunc(a int, b int, c int) error"
	expected := "mystruct"
	actual, err := formatter.GetOnlyStructureNameFromFunc(funcString)

	if err != nil {
		t.Error(
			"error:",
			err,
		)
	}

	if expected != actual {
		t.Error(
			"actual != expected",
			"actual:", actual,
			"expected:", expected,
		)
	}
}
