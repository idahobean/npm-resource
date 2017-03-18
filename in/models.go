package in

import "github.com/idahobean/npm-resource"

type Request struct {
	Source  resource.Source  `json:"source"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
