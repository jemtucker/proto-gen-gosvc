package {{ .Package }}

// {{ .Name }}Handler is an interface for a function handler
type {{ .Name }}Handler interface {
    {{ range .Methods }}
    {{ .Name }}(*{{ .RequestType }}) (*{{ .ResponseType }}, error)
    {{ end }}
}