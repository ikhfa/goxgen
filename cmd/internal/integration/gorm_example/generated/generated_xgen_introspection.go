// Code generated by github.com/goxgen/goxgen, DO NOT EDIT.

package generated

import (
    _ "embed"
    "encoding/json"
)

// XgenIntrospection is the resolver for the XgenIntrospection field.
func XgenIntrospectionValues() (*XgenIntrospection, error) {
    var intr XgenIntrospection
    err := json.Unmarshal(
		[]byte(`{"annotation":{"Action":[{"name":"PhoneNumberInput","value":{"Action":"CREATE_MUTATION","Resource":"phone_number","Route":"new","SchemaFieldName":"phone_number_create"}},{"name":"PhoneNumberInput","value":{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"}},{"name":"UserInput","value":{"Action":"CREATE_MUTATION","Resource":"user","Route":"new","SchemaFieldName":"user_create"}},{"name":"UserInput","value":{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}}],"ListAction":[{"name":"BrowseUserInput","value":{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse","Sort":{"Default":[{"by":"name","direction":"ASC"}]},"sort":{}}}],"Resource":[{"name":"Phone","value":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}},{"name":"User","value":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}}]},"object":{"BrowseUserInput":{"field":[{"definition":{"ActionField":{"Description":"ID","Label":"ID","MapTo":["User.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Name","Label":"Name","MapTo":["User.Name"]}},"name":"name"}],"object":{"ListAction":{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse","Sort":{"Default":[{"by":"name","direction":"ASC"}]},"sort":{}}}},"Phone":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the phone number","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"number"},"Description":"Number of phone","Label":"Number"}},"name":"number"},{"definition":{"Field":{"DB":{},"Description":"User of the todo","Label":"User"}},"name":"user"}],"object":{"Resource":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}}},"PhoneNumberInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the phone number","Label":"ID","MapTo":["Phone.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Number of phone","Label":"Name","MapTo":["Phone.Number"]}},"name":"number"},{"definition":{"ActionField":{"Description":"User of the phone","Label":"User","MapTo":["Phone.User"]}},"name":"user"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"}}},"User":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the user","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"name","Unique":true},"Description":"Text of the user","Label":"Text"}},"name":"name"},{"definition":{"Field":{"DB":{},"Description":"Phone numbers of the user","Label":"Phone Numbers"}},"name":"phoneNumbers"}],"object":{"Resource":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}}},"UserInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the user","Label":"ID","MapTo":["User.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Name","Label":"Name","MapTo":["User.Name"]}},"name":"name"},{"definition":{"ActionField":{"Description":"Phone numbers of the user","Label":"Phone Numbers","MapTo":["User.PhoneNumbers"]}},"name":"phones"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}}},"XgenCursorPaginationInput":{"field":[{"definition":{},"name":"first"},{"definition":{},"name":"after"},{"definition":{},"name":"last"},{"definition":{},"name":"before"}],"object":{}},"XgenPaginationInput":{"field":[{"definition":{},"name":"page"},{"definition":{},"name":"size"}],"object":{}},"XgenResourceActionType":{"field":[],"object":{}},"XgenResourceDbConfigInput":{"field":[{"definition":{},"name":"Table"}],"object":{}},"XgenResourceFieldDbConfigInput":{"field":[{"definition":{},"name":"Column"},{"definition":{},"name":"PrimaryKey"},{"definition":{},"name":"AutoIncrement"},{"definition":{},"name":"Unique"},{"definition":{},"name":"NotNull"},{"definition":{},"name":"Index"},{"definition":{},"name":"UniqueIndex"},{"definition":{},"name":"Size"},{"definition":{},"name":"Precision"},{"definition":{},"name":"Type"},{"definition":{},"name":"Scale"},{"definition":{},"name":"AutoIncrementIncrement"}],"object":{}},"XgenResourceListActionType":{"field":[],"object":{}},"XgenSort":{"field":[{"definition":{},"name":"by"},{"definition":{},"name":"direction"}],"object":{}},"XgenSortDirection":{"field":[],"object":{}},"XgenSortInput":{"field":[{"definition":{},"name":"by"},{"definition":{},"name":"direction"}],"object":{}},"XgenSortResourceConfig":{"field":[{"definition":{},"name":"Disabled"},{"definition":{},"name":"Default"}],"object":{}},"XgenSortResourceConfigInput":{"field":[{"definition":{},"name":"Disabled"},{"definition":{"ToObjectType":{"type":"[XgenSort!]"}},"name":"Default"}],"object":{}}},"resource":{"phone_number":{"actions":[{"Action":"CREATE_MUTATION","Resource":"phone_number","Route":"new","SchemaFieldName":"phone_number_create"},{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"}],"objectName":"Phone","properties":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}},"user":{"actions":[{"Action":"CREATE_MUTATION","Resource":"user","Route":"new","SchemaFieldName":"user_create"},{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"},{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse","Sort":{"Default":[{"by":"name","direction":"ASC"}]},"sort":{}}],"objectName":"User","properties":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}}}}`),
		&intr,
	)
    if err != nil {
        panic(err)
        return nil, err
    }
    return &intr, nil
}