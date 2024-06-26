package validation

import (
	"fmt"
	"github.com/ikhfa/goxgen/consts"
	"github.com/ikhfa/goxgen/graphql/common"
	"github.com/ikhfa/goxgen/graphql/directives"
	"github.com/ikhfa/goxgen/graphql/generator"
	"github.com/vektah/gqlparser/v2/ast"
)

func SchemaGeneratorHook(schema *ast.Schema, bundle *directives.DirectiveDefinitionBundle) generator.SchemaHook {
	return func(_ *ast.SchemaDocument) error {
		objects := common.GetDefinedObjects(schema)
		for _, object := range objects {
			resActionDirs := append(
				object.Directives.ForNames(consts.SchemaDefDirectiveActionName),
				object.Directives.ForNames(consts.SchemaDefDirectiveListActionName)...,
			)
			for _, resActionDir := range resActionDirs {
				xgenDirDef := bundle.GetInputObjectDirectiveDefinition(resActionDir.Name)
				if xgenDirDef != nil && xgenDirDef.Validate != nil {
					err := xgenDirDef.Validate(resActionDir, object)
					if err != nil {
						return fmt.Errorf("failed to validate object %s: %w", object.Name, err)
					}
				}
			}

			for _, field := range object.Fields {
				resActionFieldDirs := field.Directives.ForNames(consts.SchemaDefDirectiveActionFieldName)
				for _, resActionFieldDir := range resActionFieldDirs {
					xgenDirDef := bundle.GetInputFieldDirectiveDefinition(resActionFieldDir.Name)
					if xgenDirDef != nil && xgenDirDef.Validate != nil {
						err := xgenDirDef.Validate(resActionFieldDir, field)
						if err != nil {
							return fmt.Errorf("failed to validate field %s: %w", field.Name, err)
						}
					}
				}
			}
		}
		return nil
	}
}
