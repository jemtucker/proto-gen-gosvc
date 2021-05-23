package protogengosvc

import _ "embed"

var (
	//go:embed templates/handler.go.tpl
	templateHandler string

	//go:embed templates/service.go.tpl
	templateService string
)
