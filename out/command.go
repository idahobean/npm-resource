package out

import (
	"net/url"
	"io/ioutil"
	"os"
	"path/filepath"
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

	err = command.packageManager.Publish(
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
