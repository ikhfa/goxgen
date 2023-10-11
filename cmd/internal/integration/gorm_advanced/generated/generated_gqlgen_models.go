// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"

	"gorm.io/gorm"
)

// This directive is used to mark the object as a resource action
type Action struct {
	Resource        string                 `json:"Resource"`
	Action          XgenResourceActionType `json:"Action"`
	Route           *string                `json:"Route,omitempty"`
	SchemaFieldName *string                `json:"SchemaFieldName,omitempty"`
}

type ActionAnnotationSingle struct {
	Name  *string `json:"name,omitempty"`
	Value *Action `json:"value,omitempty"`
}

// This directive is used to mark the object as a resource field
type ActionField struct {
	Label       *string `json:"Label,omitempty"`
	Description *string `json:"Description,omitempty"`
	// Map field to resource field, {resource}.{field}, eg. user.id
	MapTo []string `json:"MapTo,omitempty"`
}

type Car struct {
	ID          int    `json:"id" gorm:"column:id;primaryKey;"`
	Make        string `json:"make" gorm:"column:make;"`
	Done        bool   `json:"done" gorm:"column:done;"`
	User        *User  `json:"user" gorm:""`
	UserID      int    ``
	*gorm.Model ``
}

type CarBrowseInput struct {
	ID     *int    `json:"id,omitempty" mapto:""`
	UserID *int    `json:"userId,omitempty" mapto:""`
	Make   *string `json:"make,omitempty" mapto:""`
}

type CarBrowseInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type CarInput struct {
	ID   *int       `json:"id,omitempty" mapto:"Car.ID"`
	Make *string    `json:"make,omitempty" mapto:"Car.Make"`
	Done *bool      `json:"done,omitempty" mapto:"Car.Done"`
	User *UserInput `json:"user,omitempty" mapto:"Car.User"`
}

type CarInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type CarSingleSortInput struct {
	Field     CarSortableField   `json:"field"`
	Direction *XgenSortDirection `json:"direction,omitempty"`
}

type CarSortInput struct {
	By []*CarSingleSortInput `json:"by,omitempty"`
}

type CarXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type DeleteUsers struct {
	Ids []int `json:"ids,omitempty" mapto:""`
}

type DeleteUsersXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

// This directive is used to mark the object as a resource field
type Field struct {
	Label       *string `json:"Label,omitempty"`
	Description *string `json:"Description,omitempty"`
}

// This directive is used to mark the object as a resource list action
type ListAction struct {
	Resource        string                     `json:"Resource"`
	Action          XgenResourceListActionType `json:"Action"`
	Route           *string                    `json:"Route,omitempty"`
	Pagination      *bool                      `json:"Pagination,omitempty"`
	Sort            *XgenSortResourceConfig    `json:"Sort,omitempty"`
	SchemaFieldName *string                    `json:"SchemaFieldName,omitempty"`
}

type ListActionAnnotationSingle struct {
	Name  *string     `json:"name,omitempty"`
	Value *ListAction `json:"value,omitempty"`
}

type ListUser struct {
	ID   *int    `json:"id,omitempty" mapto:"User.ID"`
	Name *string `json:"name,omitempty" mapto:"User.Name"`
}

type ListUserXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type Phone struct {
	ID          int    `json:"id" gorm:"column:id;primaryKey;"`
	Number      string `json:"number" gorm:"column:number;"`
	User        *User  `json:"user" gorm:""`
	UserID      int    ``
	*gorm.Model ``
}

type PhoneNumberBrowseInput struct {
	ID     *int `json:"id,omitempty" mapto:"Phone.ID"`
	Number *int `json:"number,omitempty" mapto:"Phone.Number"`
}

type PhoneNumberBrowseInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type PhoneNumberInput struct {
	ID     *int       `json:"id,omitempty" mapto:"Phone.ID"`
	Number *string    `json:"number,omitempty" mapto:"Phone.Number"`
	User   *UserInput `json:"user,omitempty" mapto:"Phone.User"`
}

type PhoneNumberInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type PhoneNumberSingleSortInput struct {
	Field     PhoneNumberSortableField `json:"field"`
	Direction *XgenSortDirection       `json:"direction,omitempty"`
}

type PhoneNumberSortInput struct {
	By []*PhoneNumberSingleSortInput `json:"by,omitempty"`
}

type PhoneXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

// This directive is used to mark the object as a resource
type Resource struct {
	Name    string  `json:"Name"`
	Route   *string `json:"Route,omitempty"`
	Primary *bool   `json:"Primary,omitempty"`
}

type ResourceAnnotationSingle struct {
	Name  *string   `json:"name,omitempty"`
	Value *Resource `json:"value,omitempty"`
}

type User struct {
	ID           int      `json:"id" gorm:"column:id;primaryKey;"`
	Name         string   `json:"name" gorm:"column:name;unique;"`
	Cars         []*Car   `json:"cars" gorm:""`
	PhoneNumbers []*Phone `json:"phoneNumbers" gorm:""`
	*gorm.Model  ``
}

type UserInput struct {
	ID     *int                `json:"id,omitempty" mapto:"User.ID"`
	Name   *string             `json:"name,omitempty" mapto:"User.Name"`
	Cars   []*CarInput         `json:"cars,omitempty" mapto:"User.Cars"`
	Phones []*PhoneNumberInput `json:"phones,omitempty" mapto:"User.PhoneNumbers"`
}

type UserInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type UserSingleSortInput struct {
	Field     UserSortableField  `json:"field"`
	Direction *XgenSortDirection `json:"direction,omitempty"`
}

type UserSortInput struct {
	By []*UserSingleSortInput `json:"by,omitempty"`
}

type UserXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenAnnotationMap struct {
	ListAction []*ListActionAnnotationSingle `json:"ListAction"`
	Resource   []*ResourceAnnotationSingle   `json:"Resource"`
	Action     []*ActionAnnotationSingle     `json:"Action"`
}

type XgenCursorPaginationInput struct {
	First  int     `json:"first"`
	After  *string `json:"after,omitempty"`
	Last   int     `json:"last"`
	Before *string `json:"before,omitempty"`
}

type XgenCursorPaginationInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenFieldDef struct {
	Field       *Field       `json:"Field,omitempty"`
	ActionField *ActionField `json:"ActionField,omitempty"`
}

type XgenIntrospection struct {
	Annotation *XgenAnnotationMap `json:"annotation,omitempty"`
	Object     *XgenObjectMap     `json:"object,omitempty"`
	Resource   *XgenResourceMap   `json:"resource,omitempty"`
}

type XgenObjectDefinition struct {
	Action     *Action     `json:"Action,omitempty"`
	Resource   *Resource   `json:"Resource,omitempty"`
	ListAction *ListAction `json:"ListAction,omitempty"`
}

type XgenObjectField struct {
	Name       *string       `json:"name,omitempty"`
	Definition *XgenFieldDef `json:"definition,omitempty"`
}

type XgenObjectMap struct {
	XgenPaginationInput            *XgenPaginationInputXgenDef            `json:"XgenPaginationInput,omitempty"`
	XgenResourceFieldDbConfigInput *XgenResourceFieldDbConfigInputXgenDef `json:"XgenResourceFieldDbConfigInput,omitempty"`
	XgenResourceListActionType     *XgenResourceListActionTypeXgenDef     `json:"XgenResourceListActionType,omitempty"`
	XgenResourceActionType         *XgenResourceActionTypeXgenDef         `json:"XgenResourceActionType,omitempty"`
	CarInput                       *CarInputXgenDef                       `json:"CarInput,omitempty"`
	ListUser                       *ListUserXgenDef                       `json:"ListUser,omitempty"`
	Phone                          *PhoneXgenDef                          `json:"Phone,omitempty"`
	DeleteUsers                    *DeleteUsersXgenDef                    `json:"DeleteUsers,omitempty"`
	XgenSortInput                  *XgenSortInputXgenDef                  `json:"XgenSortInput,omitempty"`
	Car                            *CarXgenDef                            `json:"Car,omitempty"`
	XgenCursorPaginationInput      *XgenCursorPaginationInputXgenDef      `json:"XgenCursorPaginationInput,omitempty"`
	XgenSortResourceConfig         *XgenSortResourceConfigXgenDef         `json:"XgenSortResourceConfig,omitempty"`
	XgenResourceDbConfigInput      *XgenResourceDbConfigInputXgenDef      `json:"XgenResourceDbConfigInput,omitempty"`
	PhoneNumberBrowseInput         *PhoneNumberBrowseInputXgenDef         `json:"PhoneNumberBrowseInput,omitempty"`
	XgenSort                       *XgenSortXgenDef                       `json:"XgenSort,omitempty"`
	UserInput                      *UserInputXgenDef                      `json:"UserInput,omitempty"`
	User                           *UserXgenDef                           `json:"User,omitempty"`
	CarBrowseInput                 *CarBrowseInputXgenDef                 `json:"CarBrowseInput,omitempty"`
	XgenSortDirection              *XgenSortDirectionXgenDef              `json:"XgenSortDirection,omitempty"`
	XgenSortResourceConfigInput    *XgenSortResourceConfigInputXgenDef    `json:"XgenSortResourceConfigInput,omitempty"`
	PhoneNumberInput               *PhoneNumberInputXgenDef               `json:"PhoneNumberInput,omitempty"`
}

type XgenPaginationInput struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type XgenPaginationInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenResourceAction struct {
	Resource        string                 `json:"Resource"`
	Action          XgenResourceActionType `json:"Action"`
	Route           *string                `json:"Route,omitempty"`
	SchemaFieldName *string                `json:"SchemaFieldName,omitempty"`
}

type XgenResourceActionTypeXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenResourceDbConfigInput struct {
	Table *string `json:"Table,omitempty"`
}

type XgenResourceDbConfigInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenResourceDefinition struct {
	ObjectName *string               `json:"objectName,omitempty"`
	Properties *XgenResourceProperty `json:"properties,omitempty"`
	Actions    []*XgenResourceAction `json:"actions"`
}

type XgenResourceFieldDbConfigInput struct {
	Column                 *string `json:"Column,omitempty"`
	PrimaryKey             *bool   `json:"PrimaryKey,omitempty"`
	AutoIncrement          *bool   `json:"AutoIncrement,omitempty"`
	Unique                 *bool   `json:"Unique,omitempty"`
	NotNull                *bool   `json:"NotNull,omitempty"`
	Index                  *bool   `json:"Index,omitempty"`
	UniqueIndex            *bool   `json:"UniqueIndex,omitempty"`
	Size                   *int    `json:"Size,omitempty"`
	Precision              *int    `json:"Precision,omitempty"`
	Type                   *string `json:"Type,omitempty"`
	Scale                  *int    `json:"Scale,omitempty"`
	AutoIncrementIncrement *int    `json:"AutoIncrementIncrement,omitempty"`
}

type XgenResourceFieldDbConfigInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenResourceListActionTypeXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenResourceMap struct {
	User        *XgenResourceDefinition `json:"user,omitempty"`
	PhoneNumber *XgenResourceDefinition `json:"phone_number,omitempty"`
	Car         *XgenResourceDefinition `json:"car,omitempty"`
}

type XgenResourceProperty struct {
	Name    string  `json:"Name"`
	Route   *string `json:"Route,omitempty"`
	Primary *bool   `json:"Primary,omitempty"`
}

type XgenSort struct {
	By        string             `json:"by"`
	Direction *XgenSortDirection `json:"direction,omitempty"`
}

type XgenSortDirectionXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenSortInput struct {
	By        string             `json:"by"`
	Direction *XgenSortDirection `json:"direction,omitempty"`
}

type XgenSortInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenSortResourceConfig struct {
	// If set to true, the sort will be disabled.
	Disabled *bool       `json:"Disabled,omitempty"`
	Default  []*XgenSort `json:"Default,omitempty"`
}

type XgenSortResourceConfigInput struct {
	// If set to true, the sort will be disabled.
	Disabled *bool            `json:"Disabled,omitempty"`
	Default  []*XgenSortInput `json:"Default,omitempty"`
}

type XgenSortResourceConfigInputXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenSortResourceConfigXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type XgenSortXgenDef struct {
	Object *XgenObjectDefinition `json:"object,omitempty"`
	Field  []*XgenObjectField    `json:"field"`
}

type CarSortableField string

const (
	CarSortableFieldID   CarSortableField = "id"
	CarSortableFieldMake CarSortableField = "make"
	CarSortableFieldDone CarSortableField = "done"
	CarSortableFieldUser CarSortableField = "user"
)

var AllCarSortableField = []CarSortableField{
	CarSortableFieldID,
	CarSortableFieldMake,
	CarSortableFieldDone,
	CarSortableFieldUser,
}

func (e CarSortableField) IsValid() bool {
	switch e {
	case CarSortableFieldID, CarSortableFieldMake, CarSortableFieldDone, CarSortableFieldUser:
		return true
	}
	return false
}

func (e CarSortableField) String() string {
	return string(e)
}

func (e *CarSortableField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CarSortableField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CAR_SORTABLE_FIELD", str)
	}
	return nil
}

func (e CarSortableField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PhoneNumberSortableField string

const (
	PhoneNumberSortableFieldID     PhoneNumberSortableField = "id"
	PhoneNumberSortableFieldNumber PhoneNumberSortableField = "number"
	PhoneNumberSortableFieldUser   PhoneNumberSortableField = "user"
)

var AllPhoneNumberSortableField = []PhoneNumberSortableField{
	PhoneNumberSortableFieldID,
	PhoneNumberSortableFieldNumber,
	PhoneNumberSortableFieldUser,
}

func (e PhoneNumberSortableField) IsValid() bool {
	switch e {
	case PhoneNumberSortableFieldID, PhoneNumberSortableFieldNumber, PhoneNumberSortableFieldUser:
		return true
	}
	return false
}

func (e PhoneNumberSortableField) String() string {
	return string(e)
}

func (e *PhoneNumberSortableField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PhoneNumberSortableField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PHONE_NUMBER_SORTABLE_FIELD", str)
	}
	return nil
}

func (e PhoneNumberSortableField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserSortableField string

const (
	UserSortableFieldID           UserSortableField = "id"
	UserSortableFieldName         UserSortableField = "name"
	UserSortableFieldCars         UserSortableField = "cars"
	UserSortableFieldPhoneNumbers UserSortableField = "phoneNumbers"
)

var AllUserSortableField = []UserSortableField{
	UserSortableFieldID,
	UserSortableFieldName,
	UserSortableFieldCars,
	UserSortableFieldPhoneNumbers,
}

func (e UserSortableField) IsValid() bool {
	switch e {
	case UserSortableFieldID, UserSortableFieldName, UserSortableFieldCars, UserSortableFieldPhoneNumbers:
		return true
	}
	return false
}

func (e UserSortableField) String() string {
	return string(e)
}

func (e *UserSortableField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSortableField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid USER_SORTABLE_FIELD", str)
	}
	return nil
}

func (e UserSortableField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

type XgenSortDirection string

const (
	XgenSortDirectionAsc  XgenSortDirection = "ASC"
	XgenSortDirectionDesc XgenSortDirection = "DESC"
)

var AllXgenSortDirection = []XgenSortDirection{
	XgenSortDirectionAsc,
	XgenSortDirectionDesc,
}

func (e XgenSortDirection) IsValid() bool {
	switch e {
	case XgenSortDirectionAsc, XgenSortDirectionDesc:
		return true
	}
	return false
}

func (e XgenSortDirection) String() string {
	return string(e)
}

func (e *XgenSortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = XgenSortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid XgenSortDirection", str)
	}
	return nil
}

func (e XgenSortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
