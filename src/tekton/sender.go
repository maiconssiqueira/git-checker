package tekton

import (
	"os"
)

func ResultSender(key string, value string) {
	// cmd := fmt.Sprintf("echo -n %s > $(params.%s)", value, key)
	// command.BashExecutor(cmd)
	os.Setenv(key, value)
}
