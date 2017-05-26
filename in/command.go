package in

import (
	"github.com/idahobean/npm-resource"
	. "github.com/idahobean/npm-resource/npm"
)

type Command struct {
	packageManager PackageManager
}

func NewCommand(packageManager PackageManager) *Command {
	return &Command {
		packageManager: packageManager,
	}
}

func (command *Command) Run(request Request) (Response, error) {
	err := command.packageManager.Install(
		request.Source.PackageName,
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}

	out, err := command.packageManager.View(
		request.Source.PackageName,
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}

	return Response {
		Version: resource.Version {
			Version: out.Version,
		},
		Metadata: []resource.MetadataPair {
			{
				Name: "name",
				Value: out.Name,
			},
			{
				Name: "homepage",
				Value: out.Homepage,
			},
		},
	}, nil
}
