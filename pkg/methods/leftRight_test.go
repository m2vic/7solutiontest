package pkg_test

import (
	pkg "gofortest/pkg/methods"
	"testing"
)

func TestLeftRight(t *testing.T) {
	result := pkg.LeftRight("LLRR=")
	result1 := pkg.LeftRight("==RLL")
	result2 := pkg.LeftRight("=LLRR")
	result3 := pkg.LeftRight("RRL=R")

	expected := "210122"
	expected1 := "000210"
	expected2 := "221012"
	expected3 := "012001"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	if result1 != expected1 {
		t.Errorf("Expected %s, got %s", expected1, result1)
	}

	if result2 != expected2 {
		t.Errorf("Expected %s, got %s", expected2, result2)
	}

	if result3 != expected3 {
		t.Errorf("Expected %s, got %s", expected3, result3)
	}
}
