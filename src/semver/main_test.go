package main

import (
	"testing"
)

func Test_SemVer(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}
	t.Run("Prod tagging", func(t *testing.T) {
		got := SemVer("1.0.0")
		want := true
		assertCorrectMessage(t, got, want)
	})
	t.Run("Dev tagging", func(t *testing.T) {
		got := SemVer("1.0.0-rc.1")
		want := true
		assertCorrectMessage(t, got, want)
	})
	t.Run("Error tagging w/ extra field", func(t *testing.T) {
		got := SemVer("1.1.1.1")
		want := false
		assertCorrectMessage(t, got, want)
	})
	t.Run("Error tagging w/ wrong delimiter", func(t *testing.T) {
		got := SemVer("1.1.1-rc-1")
		want := false
		assertCorrectMessage(t, got, want)
	})
	t.Run("Error tagging", func(t *testing.T) {
		got := SemVer("feature/bool")
		want := false
		assertCorrectMessage(t, got, want)
	})
}
