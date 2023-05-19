package utils

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"go.mongodb.org/mongo-driver/bson"
)

func collectRequestedFields(ctx context.Context, selections ast.SelectionSet) []string {
	var fields []string

	operationCtx := graphql.GetOperationContext(ctx)
	if operationCtx == nil {
		return fields
	}
	queryDocument := operationCtx.Doc

	for _, selection := range selections {
		switch s := selection.(type) {
		case *ast.Field:
			fields = append(fields, s.Name)
		case *ast.InlineFragment:
			fields = append(fields, collectRequestedFields(ctx, s.SelectionSet)...)
		case *ast.FragmentSpread:
			fragment := findFragmentByName(queryDocument.Fragments, s.Name)
			if fragment != nil {
				fields = append(fields, collectRequestedFields(ctx, fragment.SelectionSet)...)
			}
		}
	}

	return fields
}
func findFragmentByName(fragments []*ast.FragmentDefinition, name string) *ast.FragmentDefinition {
	for _, fragment := range fragments {
		if fragment.Name == name {
			return fragment
		}
	}
	return nil
}
func createProjection(ctx context.Context, fieldMapping map[string]string) bson.M {
	fieldCtx := graphql.GetFieldContext(ctx)
	if fieldCtx == nil {
		return nil
	}

	requestedFields := collectRequestedFields(ctx, fieldCtx.Field.Selections)

	projection := bson.M{}
	for _, field := range requestedFields {
		if mongoField, ok := fieldMapping[field]; ok {
			projection[mongoField] = 1
		} else {
			projection[field] = 1
		}
	}

	return projection
}
