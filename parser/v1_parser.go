package parser

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/tkuchiki/tfmbm/resource"
)

type V1Parser struct {
}

func NewV1Parser() *V1Parser {
	return &V1Parser{}
}

func (v1 *V1Parser) ParseBinding(r io.Reader) ([]*resource.BindingResource, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	f, err := hcl.ParseBytes(b)
	if err != nil {
		return nil, err
	}

	resources := make([]*resource.BindingResource, 0)

	switch node := f.Node.(type) {
	case *ast.ObjectList:
		for _, item := range node.Items {
			if !(item.Keys[0].Token.Text == "resource" && item.Keys[1].Token.Text == `"google_project_iam_binding"`) {
				continue
			}

			switch node2 := item.Val.(type) {
			case *ast.ObjectType:
				resource := resource.NewBindingResource()
				for _, item2 := range node2.List.Items {
					key := item2.Keys[0].Token.Text

					if key == "project" {
						project := strings.Trim(item2.Val.(*ast.LiteralType).Token.Text, `"`)
						resource.Project = project
					} else if key == "role" {
						role := strings.Trim(item2.Val.(*ast.LiteralType).Token.Text, `"`)
						resource.Role = role
					} else if key == "members" {
						literals := item2.Val.(*ast.ListType).List
						resource.Members = make([]string, 0, len(literals))
						for _, literal := range literals {
							member := strings.Trim(literal.(*ast.LiteralType).Token.Text, `"`)
							resource.Members = append(resource.Members, member)
						}
					}
				}
				resources = append(resources, resource)
			}
		}
	}

	return resources, nil
}
