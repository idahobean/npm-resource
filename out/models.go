package out

import "github.com/idahobean/npm-resource"

type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Path     string `json:"path"`
	Version  string `json:"version"`
	Tag      string `json:"tag"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
