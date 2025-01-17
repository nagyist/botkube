package helm

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func runHelmCLIBinary(ctx context.Context, cfg Config, args []string) (string, error) {
	cmd := exec.CommandContext(ctx, helmBinaryName, args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env,
		fmt.Sprintf("HELM_DRIVER=%s", cfg.HelmDriver),
		fmt.Sprintf("HELM_CACHE_HOME=%s", cfg.HelmCacheDir),
		fmt.Sprintf("HELM_CONFIG_HOME=%s", cfg.HelmConfigDir),
	)

	out, err := cmd.CombinedOutput()
	return string(out), err
}
