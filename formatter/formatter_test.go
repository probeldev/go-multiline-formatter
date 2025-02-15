package formatter

import "testing"

func TestOneLine2MultiLineDefineFunc1(
	t *testing.T,
) {
	formatter := GetFormatter()

	input := "func myfunc(a int, b int, c int) (int, error) {"

	actual := formatter.OneLine2MultiLine(input)

	expected := `func myfunc(
	a int,
	b int,
	c int,
) (
	int,
	error,
) {`

	if actual != expected {
		t.Error("actual != expected")
		t.Error("\n", actual)
		t.Error("\n", expected)
	}

}

func TestOneLine2MultiLineDefineFunc2(
	t *testing.T,
) {
	formatter := GetFormatter()

	input := "func myfunc(a int, b int, c int) error {"

	actual := formatter.OneLine2MultiLine(input)

	expected := `func myfunc(
	a int,
	b int,
	c int,
) error {`

	if actual != expected {
		t.Error("actual != expected")
		t.Error("\n", actual)
		t.Error("\n", expected)
	}

}

func TestOneLine2MultiLineDefineFunc3(
	t *testing.T,
) {
	formatter := GetFormatter()

	input := "func (s *mystruct) myfunc(a int, b int, c int) error {"

	actual := formatter.OneLine2MultiLine(input)

	expected := `func (s *mystruct) myfunc(
	a int,
	b int,
	c int,
) error {`

	if actual != expected {
		t.Error("actual != expected")
		t.Error("\n", actual)
		t.Error("\n", expected)
	}

}

func TestIsSturctFunc(
	t *testing.T,
) {
	formatter := GetFormatter()

	funcString := "func myfunc(a int, b int, c int) error"
	expected := false
	actual := formatter.isStuctFunc(funcString)

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
	actual = formatter.isStuctFunc(funcString)

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
	formatter := GetFormatter()

	funcString := "func (s *mystruct) myfunc(a int, b int, c int) error"
	expected := "(s *mystruct)"
	actual := formatter.getStructFromFunc(funcString)

	if expected != actual {
		t.Error(
			"actual != expected",
			"actual:", actual,
			"expected:", expected,
		)
	}
}

func TestOneLine2MultiLineRunFunc(
	t *testing.T,
) {
	formatter := GetFormatter()

	input := "myfunc(a, b, c)"

	actual := formatter.OneLine2MultiLine(input)

	expected := `myfunc(
	a,
	b,
	c,
)`

	if actual != expected {
		t.Error("actual != expected")
		t.Error("\n", actual)
		t.Error("\n", expected)
	}
}
