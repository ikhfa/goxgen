package resource

import (
	"github.com/ikhfa/goxgen/consts"
	"github.com/ikhfa/goxgen/graphql/actions/list"
	"github.com/ikhfa/goxgen/graphql/actions/singular"
	"github.com/ikhfa/goxgen/graphql/db"
	"github.com/ikhfa/goxgen/graphql/directives"
	"github.com/vektah/gqlparser/v2/ast"
)

var (
	AllDefinitions = []*ast.Definition{
		singular.ActionType,
		list.ActionType,
	}
)

var (
	pos       = &ast.Position{Src: &ast.Source{BuiltIn: false}}
	Directive = directives.ObjectDirectiveDefinition{
		Definition: &ast.DirectiveDefinition{
			Name:        consts.SchemaDefDirectiveResourceName,
			Description: `This directive is used to mark the object as a resource`,
			Position:    pos,
			Arguments: ast.ArgumentDefinitionList{
				{
					Name: consts.SchemaDefResourceDirectiveArgName,
					Type: ast.NonNullNamedType("String", nil),
				},
				{
					Name: "Route",
					Type: ast.NamedType("String", nil),
				},
				{
					Name: "Primary",
					Type: ast.NamedType("Boolean", nil),
				},
				{
					Name: consts.SchemaDefResourceDirectiveArgDb,
					Type: ast.NamedType(db.ResourceConfigInput.Name, nil),
					Directives: ast.DirectiveList{
						{Name: consts.ExcludeArgumentFromType},
					},
				},
			},
			Locations: []ast.DirectiveLocation{
				ast.LocationObject,
			},
		},
	}
)
