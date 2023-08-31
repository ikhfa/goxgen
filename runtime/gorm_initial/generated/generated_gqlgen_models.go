// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"
)

// This directive is used to mark the object as a resource action
type Action struct {
	Resource        string                 `json:"Resource" gorm:""`
	Action          XgenResourceActionType `json:"Action" gorm:""`
	Route           *string                `json:"Route,omitempty" gorm:""`
	SchemaFieldName *string                `json:"SchemaFieldName,omitempty" gorm:""`
}

type ActionAnnotationSingle struct {
	Name  *string `json:"name,omitempty" gorm:""`
	Value *Action `json:"value,omitempty" gorm:""`
}

// This directive is used to mark the object as a resource field
type ActionField struct {
	Label       *string `json:"Label,omitempty" gorm:""`
	Description *string `json:"Description,omitempty" gorm:""`
}

// This directive is used to mark the object as a resource field
type Field struct {
	Label       *string `json:"Label,omitempty" gorm:""`
	Description *string `json:"Description,omitempty" gorm:""`
}

// This directive is used to mark the object as a resource list action
type ListAction struct {
	Resource        string                     `json:"Resource" gorm:""`
	Action          XgenResourceListActionType `json:"Action" gorm:""`
	Route           *string                    `json:"Route,omitempty" gorm:""`
	Pagination      *bool                      `json:"Pagination,omitempty" gorm:""`
	SchemaFieldName *string                    `json:"SchemaFieldName,omitempty" gorm:""`
}

type ListActionAnnotationSingle struct {
	Name  *string     `json:"name,omitempty" gorm:""`
	Value *ListAction `json:"value,omitempty" gorm:""`
}

// This directive is used to mark the object as a resource
type Resource struct {
	Name    string  `json:"Name" gorm:""`
	Route   *string `json:"Route,omitempty" gorm:""`
	Primary *bool   `json:"Primary,omitempty" gorm:""`
}

type ResourceAnnotationSingle struct {
	Name  *string   `json:"name,omitempty" gorm:""`
	Value *Resource `json:"value,omitempty" gorm:""`
}

type XgenAnnotationMap struct {
	Action     []*ActionAnnotationSingle     `json:"Action" gorm:""`
	ListAction []*ListActionAnnotationSingle `json:"ListAction" gorm:""`
	Resource   []*ResourceAnnotationSingle   `json:"Resource" gorm:""`
}

type XgenCursorPaginationInput struct {
	First  int     `json:"first" gorm:""`
	After  *string `json:"after,omitempty" gorm:""`
	Last   int     `json:"last" gorm:""`
	Before *string `json:"before,omitempty" gorm:""`
}

type XgenCursorPaginationInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty" gorm:""`
	Field  []*XgenObjectField    `json:"field" gorm:""`
}

type XgenFieldDef struct {
	Field       *Field       `json:"Field,omitempty" gorm:""`
	ActionField *ActionField `json:"ActionField,omitempty" gorm:""`
}

type XgenIntrospection struct {
	Annotation *XgenAnnotationMap `json:"annotation,omitempty" gorm:""`
	Object     *XgenObjectMap     `json:"object,omitempty" gorm:""`
}

type XgenObjectDefinition struct {
	ListAction *ListAction `json:"ListAction,omitempty" gorm:""`
	Action     *Action     `json:"Action,omitempty" gorm:""`
	Resource   *Resource   `json:"Resource,omitempty" gorm:""`
}

type XgenObjectField struct {
	Name       *string       `json:"name,omitempty" gorm:""`
	Definition *XgenFieldDef `json:"definition,omitempty" gorm:""`
}

type XgenObjectMap struct {
	XgenCursorPaginationInput      *XgenCursorPaginationInputXgenDef      `json:"XgenCursorPaginationInput,omitempty" gorm:""`
	XgenResourceActionType         *XgenResourceActionTypeXgenDef         `json:"XgenResourceActionType,omitempty" gorm:""`
	XgenPaginationInput            *XgenPaginationInputXgenDef            `json:"XgenPaginationInput,omitempty" gorm:""`
	XgenResourceListActionType     *XgenResourceListActionTypeXgenDef     `json:"XgenResourceListActionType,omitempty" gorm:""`
	XgenResourceFieldDbConfigInput *XgenResourceFieldDbConfigInputXgenDef `json:"XgenResourceFieldDbConfigInput,omitempty" gorm:""`
	XgenResourceDbConfigInput      *XgenResourceDbConfigInputXgenDef      `json:"XgenResourceDbConfigInput,omitempty" gorm:""`
}

type XgenPaginationInput struct {
	Page  int `json:"page" gorm:""`
	Limit int `json:"limit" gorm:""`
}

type XgenPaginationInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty" gorm:""`
	Field  []*XgenObjectField    `json:"field" gorm:""`
}

type XgenResourceActionTypeXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty" gorm:""`
	Field  []*XgenObjectField    `json:"field" gorm:""`
}

type XgenResourceDbConfigInput struct {
	Table *string `json:"Table,omitempty" gorm:""`
}

type XgenResourceDbConfigInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty" gorm:""`
	Field  []*XgenObjectField    `json:"field" gorm:""`
}

type XgenResourceFieldDbConfigInput struct {
	Column                 *string `json:"Column,omitempty" gorm:""`
	PrimaryKey             *bool   `json:"PrimaryKey,omitempty" gorm:""`
	AutoIncrement          *bool   `json:"AutoIncrement,omitempty" gorm:""`
	Unique                 *bool   `json:"Unique,omitempty" gorm:""`
	NotNull                *bool   `json:"NotNull,omitempty" gorm:""`
	Index                  *bool   `json:"Index,omitempty" gorm:""`
	UniqueIndex            *bool   `json:"UniqueIndex,omitempty" gorm:""`
	Size                   *int    `json:"Size,omitempty" gorm:""`
	Precision              *int    `json:"Precision,omitempty" gorm:""`
	Type                   *string `json:"Type,omitempty" gorm:""`
	Scale                  *int    `json:"Scale,omitempty" gorm:""`
	AutoIncrementIncrement *int    `json:"AutoIncrementIncrement,omitempty" gorm:""`
}

type XgenResourceFieldDbConfigInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty" gorm:""`
	Field  []*XgenObjectField    `json:"field" gorm:""`
}

type XgenResourceListActionTypeXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty" gorm:""`
	Field  []*XgenObjectField    `json:"field" gorm:""`
}

type XgenResourceActionType string

const (
	XgenResourceActionTypeCreateMutation XgenResourceActionType = "CREATE_MUTATION"
	XgenResourceActionTypeReadQuery      XgenResourceActionType = "READ_QUERY"
	XgenResourceActionTypeUpdateMutation XgenResourceActionType = "UPDATE_MUTATION"
	XgenResourceActionTypeDeleteMutation XgenResourceActionType = "DELETE_MUTATION"
)

var AllXgenResourceActionType = []XgenResourceActionType{
	XgenResourceActionTypeCreateMutation,
	XgenResourceActionTypeReadQuery,
	XgenResourceActionTypeUpdateMutation,
	XgenResourceActionTypeDeleteMutation,
}

func (e XgenResourceActionType) IsValid() bool {
	switch e {
	case XgenResourceActionTypeCreateMutation, XgenResourceActionTypeReadQuery, XgenResourceActionTypeUpdateMutation, XgenResourceActionTypeDeleteMutation:
		return true
	}
	return false
}

func (e XgenResourceActionType) String() string {
	return string(e)
}

func (e *XgenResourceActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = XgenResourceActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid XgenResourceActionType", str)
	}
	return nil
}

func (e XgenResourceActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type XgenResourceListActionType string

const (
	XgenResourceListActionTypeBrowseQuery         XgenResourceListActionType = "BROWSE_QUERY"
	XgenResourceListActionTypeBatchDeleteMutation XgenResourceListActionType = "BATCH_DELETE_MUTATION"
)

var AllXgenResourceListActionType = []XgenResourceListActionType{
	XgenResourceListActionTypeBrowseQuery,
	XgenResourceListActionTypeBatchDeleteMutation,
}

func (e XgenResourceListActionType) IsValid() bool {
	switch e {
	case XgenResourceListActionTypeBrowseQuery, XgenResourceListActionTypeBatchDeleteMutation:
		return true
	}
	return false
}

func (e XgenResourceListActionType) String() string {
	return string(e)
}

func (e *XgenResourceListActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = XgenResourceListActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid XgenResourceListActionType", str)
	}
	return nil
}

func (e XgenResourceListActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
