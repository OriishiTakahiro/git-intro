package main

import (
	"testing"
)

func TestValidateLength(t *testing.T) {
	if e := validate([]string{"a", "b", "0", "1"}); e != nil {
		t.Fatalf("長さ4つの配列でエラーが出ました%v", e)
	}
	if e := validate([]string{"a", "b"}); e == nil {
		t.Fatalf("長さ2つの配列でエラーが出ませんでした%v", e)
	}
}

func TestConvertibleToInt(t *testing.T) {
	if e := validate([]string{"a", "b", "0", "0"}); e != nil {
		t.Fatalf("args[2]=0 && args[3]=0 でエラーが出なかった%v", e)
	}
	if e := validate([]string{"a", "b", "c", "0"}); e == nil {
		t.Fatalf("args[2]=c でエラーが出なかった%v", e)
	}
	if e := validate([]string{"a", "b", "0", "d"}); e == nil {
		t.Fatalf("args[3]=d でエラーが出なかった%v", e)
	}
}
