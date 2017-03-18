package in

import "github.com/idahobean/npm-resource"

type Request struct {
	Source  resource.Source  `json:"source"`
	Params  Params           `json:"params"`
}

type Params struct {
	Path    string           `json:"path"`
	Version string           `json:"version"`
	Tag     string           `json:"tag"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
