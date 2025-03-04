package parser

import "testing"

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
		t.Error("actual != expected")
	}
}

func TestGetFuncName2(
	t *testing.T,
) {
	input := `func (s *mystruct) myfunc() {`

	expected := "myfunc"

	actual := GetFunctionParser().GetFuncName(input)

	if actual != expected {
		t.Error("actual != expected")
	}
}
