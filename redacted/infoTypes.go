package redacted

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

type InfoType struct {
	Expr         string
	Name         string
	WordBoundary bool
	Generate     func() string
}

func AusDriversLicence() InfoType {
	name := "AusDriversLicence"
	expr := "[A-Z0-9][0-9]{5,7}"
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     func() string { return gofakeit.Regex(expr) },
	}
}

func AusPassport() InfoType {
	name := "AusPassport"
	expr := "[A-Z][0-9]{7}"
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     func() string { return gofakeit.Regex(expr) },
	}
}

func AusLicensePlate() InfoType {
	name := "AusLicencePlate"
	expr := `(([a-zA-Z0-9]{3})([\s,-.]?)([a-zA-Z0-9]{3}))`
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     func() string { return gofakeit.Regex(expr) },
	}
}

func AusTaxFileNumber() InfoType {
	name := "AusTaxFileNumber"
	expr := `[0-9]{3}( ?)[0-9]{3}[0-9]{2,3}`
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     func() string { return gofakeit.Regex(expr) },
	}
}

func Email() InfoType {
	name := "Email"
	expr := `([a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     func() string { return gofakeit.Email() },
	}
}

func AusPostCode() InfoType {
	name := "AusPostCode"
	expr := `(0[289][0-9]{2}\)|([1345689][0-9]{3})|(2[0-9][0-9]{2})|(290[0-9])|(291[0-9])|(7[0-4][0-9]{2})|(7[8-9][0-9]{2}))`
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     func() string { return gofakeit.Regex(expr) },
	}
}

func LongDigit(minLength int) InfoType {
	name := "LongDigit"
	expr := fmt.Sprintf("\\d{{%v}", minLength)
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     nil,
	}
}
