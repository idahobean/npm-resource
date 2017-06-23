package check

import (
	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/npm"
)

type Command struct {
	packageManager npm.PackageManager
}

func NewCommand(packageManager npm.PackageManager) *Command {
	return &Command{
		packageManager: packageManager,
	}
}

func (command *Command) Run(request Request) ([]resource.Version, error) {
	out, err := command.packageManager.View(
		request.Source.PackageName,
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}

	return []resource.Version{
		{
			Version: out.Version,
		},
	}, nil
}
