package protogengosvc

import (
	"errors"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Service contains all the info the service needs to be generated
type Service struct {
	Package string
	Name    string
	Methods []*Method
}

// NewService creates a new service from a protobuf definition
func NewService(pkg string, svc *protogen.Service) (*Service, error) {
	methods, err := NewMethods(svc.Methods)
	if err != nil {
		return nil, err
	}

	return &Service{
		Package: pkg,
		Name:    string(svc.Desc.Name()),
		Methods: methods,
	}, nil
}

// Method is a single service method
type Method struct {
	Name         string
	Path         string
	HTTPMethod   string
	RequestType  string
	ResponseType string
}

// NewMethod creates a method from a protobuf definition
func NewMethod(method *protogen.Method) (*Method, error) {
	options, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, errors.New("invalid option type")
	}

	rule, ok := proto.GetExtension(options, annotations.E_Http).(*annotations.HttpRule)
	if !ok {
		return nil, errors.New("invalid http rule type")
	}

	var m, p string
	switch rule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		m = "GET"
		p = rule.GetGet()
	case *annotations.HttpRule_Put:
		m = "PUT"
		p = rule.GetPut()
	case *annotations.HttpRule_Post:
		m = "POST"
		p = rule.GetPost()
	case *annotations.HttpRule_Delete:
		m = "DELETE"
		p = rule.GetDelete()
	case *annotations.HttpRule_Patch:
		m = "PATCH"
		p = rule.GetPatch()
	default:
		return nil, errors.New("invalid pattern type")
	}

	return &Method{
		Name:         string(method.Desc.Name()),
		HTTPMethod:   m,
		Path:         p,
		RequestType:  method.Input.GoIdent.GoName,
		ResponseType: method.Output.GoIdent.GoName,
	}, nil
}

// NewMethods creates multiple methods from protobuf definitions
func NewMethods(methods []*protogen.Method) ([]*Method, error) {
	var results []*Method
	for _, m := range methods {
		n, err := NewMethod(m)
		if err != nil {
			return nil, err
		}

		results = append(results, n)
	}
	return results, nil
}

func (m *Method) ReadBody() bool {
	switch m.HTTPMethod {
	case "GET", "DELETE":
		return false
	default:
		return true
	}
}

// Parameter is a paramenter on a Method
type Parameter struct {
	Name        string
	Description string
}
