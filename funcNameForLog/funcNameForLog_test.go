package funcnameforlog

import (
	"testing"
)

func TestAddFuncName(
	t *testing.T,
) {

	input := "func myfunc() {"

	expected := `func myfunc() {
	fn:="myfunc"
`

	actual := GetFuncNameForLog().AddFuncName(input)

	if actual != expected {
		t.Error("actual != expected")
	}
}

func TestAddFuncName2(
	t *testing.T,
) {

	input := "func (s *structname) myfunc() {"

	expected := `func (s *structname) myfunc() {
	fn:="structname:myfunc"
`
	actual := GetFuncNameForLog().AddFuncName(input)

	if actual != expected {
		t.Error("actual != expected")
	}
}

func TestAddFuncName3(
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

	expected := `func (s *structname) myfunc(
	a int,
	b int,
	c int,
) (
	int,
	error,
) {
	fn:="structname:myfunc"
`
	actual := GetFuncNameForLog().AddFuncName(input)

	if actual != expected {
		t.Error("actual != expected")
	}
}
