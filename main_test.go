package scraper

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	actual, err := Get()
	if err != nil {
		t.Error(err)
	}

	if actual == nil {
		t.Error("actual is nil")
	}

	if len(actual) == 0 {
		t.Error("actual is empty")
	}

	fmt.Printf("%+v", actual)
}
