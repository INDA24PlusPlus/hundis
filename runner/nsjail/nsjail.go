package nsjail

import (
	"fmt"
	"os/exec"
)

/* type NsjailConfig struct {
	Mode       string
	Chroot     string
	Bindmounts map[string]string
	Cwd        string
	Env        map[string]string
	Command    []string
}

func NewNsjailConfig() *NsjailConfig {
	return &NsjailConfig{
		Mode:       "Mo",
		Bindmounts: make(map[string]string),
		Env:        make(map[string]string),
	}
}

func (n *NsjailConfig) SetChroot(path string) *NsjailConfig {
	n.Chroot = path
	return n
}

func (n *NsjailConfig) AddBindmount(src, dst string) *NsjailConfig {
	n.Bindmounts[src] = dst
	return n
}

func (n *NsjailConfig) SetCwd(cwd string) *NsjailConfig {
	n.Cwd = cwd
	return n
}

func (n *NsjailConfig) SetEnv(key, value string) *NsjailConfig {
	n.Env[key] = value
	return n
}

func (n *NsjailConfig) SetCommand(cmd ...string) *NsjailConfig {
	n.Command = cmd
	return n
}

func (n *NsjailConfig) BuildCommand() *exec.Cmd {
	args := []string{
		"-" + n.Mode,
	}
	if n.Chroot != "" {
		args = append(args, "--chroot", n.Chroot)
	}
	for src, dst := range n.Bindmounts {
		args = append(args, "--bindmount", fmt.Sprintf("%s:%s", src, dst))
	}
	if n.Cwd != "" {
		args = append(args, "--cwd", n.Cwd)
	}
	for k, v := range n.Env {
		args = append(args, "--env", fmt.Sprintf("%s=%s", k, v))
	}
	args = append(args, "--")
	args = append(args, n.Command...)
	return exec.Command("nsjail", args...)
}

func (n *NsjailConfig) Run() ([]byte, error) {
	cmd := n.BuildCommand()
	return cmd.CombinedOutput()
} */

func RunNsjail(configPath, mountPath string) ([]byte, error) {
	cmd := exec.Command("nsjail", "--config", configPath, "--bindmount", fmt.Sprintf("%s:/compile/", mountPath))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}
	return out, nil
}
