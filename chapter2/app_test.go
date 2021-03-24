package main

import (
	"testing"
)

func TestApp(t *testing.T) {
	got := "App"

	if got != App() {
		t.Errorf("%s != App", got)
	}
}