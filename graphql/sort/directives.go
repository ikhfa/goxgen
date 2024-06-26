package sort

import (
	"github.com/ikhfa/goxgen/consts"
	"github.com/ikhfa/goxgen/graphql/common"
	"github.com/vektah/gqlparser/v2/ast"
)

var (
	DirectionEnum = &ast.Definition{
		Kind: ast.Enum,
		Name: "XgenSortDirection",
		EnumValues: []*ast.EnumValueDefinition{
			{
				Name: consts.SortDirectionAsc,
			},
			{
				Name: consts.SortDirectionDesc,
			},
		},
	}
	InputObject = &ast.Definition{
		Kind: ast.InputObject,
		Name: "XgenSortInput",
		Fields: ast.FieldList{
			{
				Name: "by",
				Type: ast.NonNullNamedType("String", nil),
			},
			{
				Name: "direction",
				Type: ast.NamedType(DirectionEnum.Name, nil),
			},
		},
		Interfaces: []string{},
	}
	Object                    = common.ToObjectDefinition(*InputObject, "XgenSort")
	ResourceConfigInputObject = &ast.Definition{
		Kind: ast.InputObject,
		Name: "XgenSortResourceConfigInput",
		Fields: ast.FieldList{
			{
				Name:        "Disabled",
				Type:        ast.NamedType("Boolean", nil),
				Description: "If set to true, the sort will be disabled.",
			},
			{
				Name: "Default",
				Type: ast.ListType(ast.NonNullNamedType(InputObject.Name, nil), nil),
				Directives: ast.DirectiveList{
					{
						Name: consts.ToObjectType,
						Arguments: ast.ArgumentList{
							{
								Name: "type",
								Value: &ast.Value{
									Kind: ast.StringValue,
									Raw:  "[" + Object.Name + "!]",
								},
							},
						},
					},
				},
			},
		},
	}
	ResourceConfigObject = common.ToObjectDefinition(*ResourceConfigInputObject, "XgenSortResourceConfig")

	AllDefinitions = []*ast.Definition{
		DirectionEnum,
		InputObject,
		Object,
		ResourceConfigInputObject,
		ResourceConfigObject,
	}
)
