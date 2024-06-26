package db

import (
	"github.com/ikhfa/goxgen/consts"
	"github.com/vektah/gqlparser/v2/ast"
)

var (
	ResourceConfigInput = &ast.Definition{
		Kind: ast.InputObject,
		Name: "XgenResourceDbConfigInput",
		Fields: ast.FieldList{
			{
				Name: "Table",
				Type: ast.NamedType("String", nil),
			},
		},
	}
	ResourceFieldConfigInput = &ast.Definition{
		Kind: ast.InputObject,
		Name: consts.SchemaDefFieldDbConfigInputType,
		Fields: ast.FieldList{
			{
				Name: "Column",
				Type: ast.NamedType("String", nil),
			},
			{
				Name: "PrimaryKey",
				Type: ast.NamedType("Boolean", nil),
			},
			{
				Name: "AutoIncrement",
				Type: ast.NamedType("Boolean", nil),
			},
			{
				Name: "Unique",
				Type: ast.NamedType("Boolean", nil),
			},
			{
				Name: "NotNull",
				Type: ast.NamedType("Boolean", nil),
			},
			{
				Name: "Index",
				Type: ast.NamedType("Boolean", nil),
			},
			{
				Name: "UniqueIndex",
				Type: ast.NamedType("Boolean", nil),
			},
			{
				Name: "Size",
				Type: ast.NamedType("Int", nil),
			},
			{
				Name: "Precision",
				Type: ast.NamedType("Int", nil),
			},
			{
				Name: "Type",
				Type: ast.NamedType("String", nil),
			},
			{
				Name: "Scale",
				Type: ast.NamedType("Int", nil),
			},
			{
				Name: "AutoIncrementIncrement",
				Type: ast.NamedType("Int", nil),
			},
		},
	}

	AllDefinitions = []*ast.Definition{
		ResourceConfigInput,
		ResourceFieldConfigInput,
	}
)
