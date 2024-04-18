package common

import (
	"github.com/ikhfa/goxgen/consts"
	"github.com/ikhfa/goxgen/graphql/directives"
	"github.com/vektah/gqlparser/v2/ast"
)

var (
	ToObjectType = directives.InputObjectDirectiveDefinition{
		Definition: &ast.DirectiveDefinition{
			Name:        consts.ToObjectType,
			Description: `This directive is used to define the object type`,
			Position:    pos,
			Locations: []ast.DirectiveLocation{
				ast.LocationArgumentDefinition,
				ast.LocationInputFieldDefinition,
				ast.LocationFieldDefinition,
			},
			Arguments: ast.ArgumentDefinitionList{
				{
					Name: "type",
					Type: ast.NonNullNamedType("String", nil),
				},
			},
		},
	}
)
