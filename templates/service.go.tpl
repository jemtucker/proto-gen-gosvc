package {{ .Package }}

import (
    restful "github.com/emicklei/go-restful/v3"
)

type {{ .Name }}Service struct {
    ws      *restful.WebService
    handler {{ .Name }}Handler
}

// New{{ .Name }}Service instantiates a new service
func New{{ .Name }}Service(handler {{ .Name }}Handler) *{{ .Name }}Service {
    return &{{ .Name }}Service{
        ws:      &restful.WebService{},
        handler: handler,
    }
}

// RegisterRoutes registers all routes for the {{ .Name }}Service
func (s *{{ .Name }}Service) RegisterRoutes() {
    ws.
        Path("{{ .RootPath }}").
        Consumes(restful.MIME_JSON).
        Produces(restful.MIME_JSON)

    {{ range .Methods }}
    ws.Route(ws.{{ .HTTPMethod }}("{{ .Path }}").To(u.findUser).
        {{/* TODO Doc("{{ .Documentation }}"). */}}
        {{ range .Parameters }}
        Param(
            ws.
                PathParameter("{{ .Name }}", "{{ .Description }}").
                DataType("string"),
        ).
        {{ end }}
        Reads({{ .RequestType }}{}).
        Writes({{ .ResponseType }}{}))	
    {{ end }}
}

// GetWebService returns the underlying WebService
func (s *{{ .Name }}Service) GetWebService() *restful.WebService {
    return s.ws
}
