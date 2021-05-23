package protogengosvc

import (
	"fmt"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

// Protogen is a protobuf generator
type Protogen struct {
	Generator *protogen.Plugin
}

// Generate does generation for a single protobuf file
func (p *Protogen) Generate(file *protogen.File) error {
	for _, svc := range file.Services {
		if err := p.genService(svc, file); err != nil {
			return err
		}
	}

	return nil
}

func (p *Protogen) genService(svc *protogen.Service, file *protogen.File) error {
	templates := []struct {
		name    string
		content string
	}{
		{name: "service", content: templateService},
		{name: "handler", content: templateHandler},
	}

	service, err := NewService(string(file.GoPackageName), svc)
	if err != nil {
		return err
	}

	for _, t := range templates {
		if err := p.genTemplate(t.name, t.content, service, file); err != nil {
			return fmt.Errorf("error writing template %q: %w", t.name, err)
		}
	}

	return nil
}

func (p *Protogen) genTemplate(prefix, tmpl string, data interface{}, file *protogen.File) error {
	g := p.Generator.NewGeneratedFile(
		fmt.Sprintf("%s.%s.go", file.GeneratedFilenamePrefix, prefix),
		file.GoImportPath,
	)

	g.P("// Code generated by protoc-gen-gosvc. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// 	protoc-gen-gosvc v0.0.0")
	g.P(
		fmt.Sprintf(
			"// 	protoc           v%d.%d.%d",
			*p.Generator.Request.CompilerVersion.Major,
			*p.Generator.Request.CompilerVersion.Minor,
			*p.Generator.Request.CompilerVersion.Patch,
		),
	)
	g.P("// source: ", file.Desc.Path())
	g.P()

	t, err := template.New(prefix).Parse(tmpl)
	if err != nil {
		return err
	}

	if err := t.Execute(g, data); err != nil {
		return err
	}

	return nil
}
