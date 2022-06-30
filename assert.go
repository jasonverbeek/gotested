package gtest

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertTrue(t *testing.T, value bool, msg *string) {
	if !value {
		if msg != nil {
			t.Error(msg)
		}
		t.Error("Assertion failed; value is not true")
	}
}

func AssertFalse(t *testing.T, value bool, msg *string) {
	if value {
		if msg != nil {
			t.Error(msg)
		}
		t.Error("Assertion failed; Value is not false")
	}
}

func AssertEqual(t *testing.T, value any, expected any) {
	if reflect.TypeOf(value) != reflect.TypeOf(expected) {
		t.Errorf("Assertion failed; Type mismatch %T and %T", value, expected)
	}

	if value != expected {
		t.Errorf("Assertion failed; '%v' != '%v' for type %T", value, expected, expected)
	}
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func AssertInRange[num Number](t *testing.T, value num, min num, max num) {
	error := fmt.Sprintf("Assertion failed; value '%v' is not in range %v-%v (inclusive)", value, min, max)
	AssertTrue(t, value >= min, &error)
	AssertTrue(t, value <= max, &error)
}
