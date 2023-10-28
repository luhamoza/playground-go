package test

import (
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {
	expected := Player{"John", 90}
	have := Player{"John", 90}

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("Expected %+v, Have %+v", expected, have)
	}
}

func TestAddition(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 5
	)
	have := Addition(a, b)
	if have != expected {
		t.Errorf("Expected value:%d, Have value: %d", expected, have)
	}
}

// to test without using cached value
// $ go test -v ./... -count=1
