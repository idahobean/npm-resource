package out

import (
	"encoding/json"
	"net/url"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/idahobean/npm-resource"
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
	parsedUrl, err := url.Parse(request.Source.Registry)
	if err != nil {
		return Response{}, err
	}
	authToken := "//" + parsedUrl.Host + "/:_authToken=" + request.Source.Token
	ioutil.WriteFile(request.Params.Path + "/.npmrc", []byte(authToken), os.ModePerm)

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return Response{}, err
	}

	err := command.packageManager.Publish(
		filepath.Join(path, request.Params.Path),
		request.Params.Tag,
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
	var packageInfo interface{}
	err := json.Unmarshal(out, &packageInfo)

	return Response {
		Version: resource.Version {
			Version: out.version
		},
		Metadata: []resource.MetadataPair {
			{
				Name: "name",
				Value: out.name,
			}
			{
				Name: "homepage",
				Value: out.homepage,
			},
		},
	}, nil
}
