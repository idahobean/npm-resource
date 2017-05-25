package check

import (
	"encoding/json"
	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/npm"
)

type Command struct {
	packageManager PackageManager
}

func NewCommand(packageManager PackageManager) *Command {
	return &Command {
		packageManager: PackageManager,
	}
}

func (command *Command) Run(request Request) (Response, error) {
	out, err := command.packageManager.View(
		request.Source.PackageName,
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}
	var packageInfo interface{}
	err = json.Unmarshal(out, &packageInfo)

	return Response {
		Version: resource.Version {
			Version: out.version,
		},
		Metadata: []resource.MetadataPair {
			{
				Name: "name",
				Value: out.name,
			},
			{
				Name: "homepage",
				Value: out.homepage,
			},
		},
	}, nil
}
