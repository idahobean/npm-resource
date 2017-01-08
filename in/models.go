package in

import "github.com/idahobean/cf-service-resource"

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
