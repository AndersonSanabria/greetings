package greetings

import (
	"regexp"
	"testing"
)

func TestHolaNombre(t *testing.T) {
	nombre := "Juan"
	want := regexp.MustCompile(`\b` + nombre + `\b`) // busca la palabra exacta Juan
	msg, err := Holam("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hola("Juan")= %q,%v,quiere coincidencia para  %#q,nil`, msg, err, want)

	}

}

func TestHolaEmpty(t *testing.T) {
	msg, err := Holam("")
	if msg != "" || err != nil {
		t.Fatalf(`Hola("")= %q,%v,quiere  "",error`, msg, err)
	}

}
