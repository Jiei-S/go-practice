package sports

import (
	"testing"
)

func TestSoccer(t *testing.T) {
	got := 11

	if got != Soccer() {
		t.Errorf("%d != 11", got)
	}
}

func TestBaseball(t *testing.T) {
	got := 9

	if got != Baseball() {
		t.Errorf("%d != 9", got)
	}
}

func TestBasketball(t *testing.T) {
	got := 5

	if got != Basketball() {
		t.Errorf("%d != 5", got)
	}
}