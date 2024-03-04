package cmdutil

import (
	"context"
	"testing"
	"time"
)

func TestRunCmdWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	ec, err := RunCmdWithContext(ctx, false, "bash", "-c", "echo hello && sleep 5")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("cmd success")
	t.Log("stdout:", ec.StdOut())
	t.Log("stderr:", ec.StdErr())
}
