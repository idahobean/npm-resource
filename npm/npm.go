package npm

import (
	"os"
	"os/exec"
)

type PackageManager interface {
	View(packageName string, registry string) error
	Install(packageName string, registry string) error
	Version(version string) error
	Publish(path string, tag string, registry string) error
}

type NPM struct {}

func NewNPM() *NPM {
	return &NPM {}
}

func (npm *NPM) View(packageName string, registry string) error {
	args := []string{ args, "view", packageName }

	if registry != "" {
		args = append(args, "--registry", registry)
	}

	return npm.npm(args...).Output()
}

func (npm *NPM) Install(packageName string, registry string) error {
	args := []string{ args, "install", packageName }

	if registry != "" {
		args = append(args, "--registry", registry)
	}

	return npm.npm(args...).Run()
}

func (npm *NPM) Version(version string) error {
	args := []string{ "version", version }
	return npm.npm(args...).Run()
}

func (npm *NPM) Publish(path string, tag string, registry string) error {
	args := []string{ args, "publish", path }

	if tag != "" {
		args = append(args, "--tag", tag)
	}
	if registry != "" {
		args = append(args, "--registry", registry)
	}

	return npm.npm(args...).Run()
}

func (npm *NPM) npm(args ...string) *exec.Cmd {
	cmd := exec.Command("npm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
