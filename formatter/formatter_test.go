package formatter

// Закомментировал пока, что бы устанавлвиалось через nix

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

func TestOneLine2MultiLineRunFunc2(
	t *testing.T,
) {
	formatter := GetFormatter()

	input := "a, b := myfunc(a, b, c)"

	actual := formatter.OneLine2MultiLine(input)

	expected := `a, b := myfunc(
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

// TODO
// func TestOneLine2MultiLineRunFunc3(
// 	t *testing.T,
// ) {
// 	formatter := GetFormatter()
//
// 	input := `a, b := myfunc(a, b, func(c int, b int,){
// 		log.Println(c, b)
// 	}, c)`
//
// 	actual := formatter.OneLine2MultiLine(input)
//
// 	expected := `a, b := myfunc(
// 	a,
// 	b,
// 	func(c int, b int) {
// 		log.Println(c, b)
// 	},
// 	c,
// )`
//
// 	if actual != expected {
// 		t.Error("actual != expected")
// 		t.Error("\n", actual)
// 		t.Error("\n", expected)
// 	}
// }

// func TestOneLine2MultiLineRunFunc4(
// 	t *testing.T,
// ) {
// 	formatter := GetFormatter()
//
// 	input := `a, b := myfunc(a, b, func(c int, b int,){
// 		log.Println(c, b)
// 	}, c)`
//
// 	actual := formatter.OneLine2MultiLine(input)
//
// 	expected := `a, b := myfunc(
// 	a,
// 	b,
// 	func(c int, b int) {
// 		d := c + b
// 		log.Println(d)
// 	},
// 	c,
// )`
//
// 	if actual != expected {
// 		t.Error("actual != expected")
// 		t.Error("\n", actual)
// 		t.Error("\n", expected)
// 	}
// }

func TestOneLine2MultiLineRunFunc4(
	t *testing.T,
) {
	formatter := GetFormatter()

	input := `err = stmt.QueryRow(deal.Name, deal.Number).Scan(&ldeal.Id, &ldeal.Name, &ldeal.Number)`

	actual := formatter.OneLine2MultiLine(input)

	expected := `err = stmt.QueryRow(
	deal.Name,
	deal.Number,
).Scan(
	&ldeal.Id,
	&ldeal.Name,
	&ldeal.Number,
)`

	if actual != expected {
		t.Error("actual != expected")
		t.Error("\n", actual)
		t.Error("\n", expected)
	}
}

// func TestOneLine2MultiLineIf(
// 	t *testing.T,
// ) {
//
// 	formatter := GetFormatter()
//
// 	input := `if category.Url == categoryUrl && category.SiteId == siteDb.Id && category.ParentId == partntId {`
//
// 	actual := formatter.OneLine2MultiLine(input)
//
// 	expected := `if category.Url == categoryUrl &&
// category.SiteId == siteDb.Id &&
// category.ParentId == partntId {`
//
// 	if actual != expected {
// 		t.Error("actual != expected")
// 		t.Error("\n", actual)
// 		t.Error("\n", expected)
// 	}
// }
