package exec_test

import (
	"os/exec"
	"testing"
)

func TestLookPath(t *testing.T) {
	path, err := exec.LookPath("fortune")
	if err != nil {
		t.Fatal("installing fortune is in your future", err.Error())
	}

	t.Logf("fortune is available at %s\n", path)
}

func TestCombinedOutput(t *testing.T) {

	cmd := exec.Command("sh", "-c", "echo stdout;echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", stdoutStderr)
}
