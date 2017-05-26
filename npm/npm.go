package npm

import (
	"os"
	"os/exec"

	simpleJson "github.com/bitly/go-simplejson"
)

type PackageManager interface {
	View(packageName string, registry string) (*Info, error)
	Install(packageName string, registry string) error
	Version(version string) error
	Publish(path string, tag string, registry string) error
}

type NPM struct {}

func NewNPM() *NPM {
	return &NPM {}
}

func (npm *NPM) View(packageName string, registry string) (*Info, error) {
	args := []string{ "view", packageName }

	if registry != "" {
		args = append(args, "--registry", registry)
	}

	out, err := npm.npm(args...).Output()
	if err != nil {
		return &Info{}, err
	}

	js, err := simpleJson.NewJson([]byte(out))
	if err != nil {
		return &Info{}, err
	}

	var info Info
	info.Name, err = js.Get("name").String()
	info.Version, err = js.Get("version").String()
	info.Homepage, err = js.Get("homepage").String()

	return &info, err
}

func (npm *NPM) Install(packageName string, registry string) error {
	args := []string{ "install", packageName }

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
	args := []string{ "publish", path }

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
