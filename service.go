package protogengosvc

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// Service contains all the info the service needs to be generated
type Service struct {
	Package  string
	Name     string
	RootPath string
	Methods  []*Method
}

// NewService creates a new service from a protobuf definition
func NewService(pkg string, svc *protogen.Service) *Service {
	return &Service{
		Package:  pkg,
		Name:     string(svc.Desc.Name()),
		RootPath: "/root", // TODO add support for this
		Methods:  NewMethods(svc.Methods),
	}
}

// Method is a single service method
type Method struct {
	Name         string
	Path         string
	HTTPMethod   string
	RequestType  string
	ResponseType string
	Parameters   []*Parameter
}

// NewMethod creates a method from a protobuf definition
func NewMethod(method *protogen.Method) *Method {
	name := string(method.Desc.Name())

	return &Method{
		Name:         name,
		Path:         "/todo", // TODO add support for this
		HTTPMethod:   httpMethodFromName(name),
		RequestType:  method.Input.GoIdent.GoName,
		ResponseType: method.Output.GoIdent.GoName,
	}
}

// NewMethods creates multiple methods from protobuf definitions
func NewMethods(methods []*protogen.Method) []*Method {
	var results []*Method
	for _, m := range methods {
		results = append(results, NewMethod(m))
	}
	return results
}

// Parameter is a paramenter on a Method
type Parameter struct {
	Name        string
	Description string
}

func httpMethodFromName(name string) string {
	switch {
	case strings.HasPrefix(name, "GET"):
		return "GET"
	case strings.HasPrefix(name, "POST"):
		return "POST"
	case strings.HasPrefix(name, "PUT"):
		return "PUT"
	case strings.HasPrefix(name, "DELETE"):
		return "DELETE"
	}

	return "POST"
}
