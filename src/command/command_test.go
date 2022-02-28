package command

import (
	"testing"
)

func Test_CommandExec(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Echo test", func(t *testing.T) {
		got := BashExecutor("printf \"%s\n\" HelloWorld")
		want := "HelloWorld"
		assertCorrectMessage(t, got, want)
	})
}
