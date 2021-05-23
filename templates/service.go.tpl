package {{ .Package }}

import (
    "net/http"

    "github.com/emicklei/go-restful/v3"
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
    ws := s.ws

    ws.
        Path("/").
        Consumes(restful.MIME_JSON).
        Produces(restful.MIME_JSON)

    {{ range .Methods }}
    ws.Route(ws.{{ .HTTPMethod }}("{{ .Path }}").
        {{/* TODO Documentation, parameters and headers */}}
        Reads({{ .RequestType }}{}).
        Writes({{ .ResponseType }}{}).
        To(s.handle{{ .Name }}))
    {{ end }}
}

// GetWebService returns the underlying WebService
func (s *{{ .Name }}Service) GetWebService() *restful.WebService {
    return s.ws
}

{{ range .Methods }}
func (s *{{ $.Name }}Service) handle{{ .Name }}(req *restful.Request, rsp *restful.Response) {
    body := {{ .RequestType }}{}
    
    {{ if .ReadBody }}
    if err := req.ReadEntity(&body); err != nil {
        rsp.WriteError(http.StatusBadRequest, err)
        return
    }
    {{ end }}

    result, err := s.handler.{{ .Name }}(&body)
    if err != nil {
        // TODO do this better
        rsp.WriteError(http.StatusInternalServerError, err)
        return
    }

    rsp.WriteAsJson(result)
} 
{{ end }}