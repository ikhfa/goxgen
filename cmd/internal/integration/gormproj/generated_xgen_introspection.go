// Code generated by github.com/goxgen/goxgen, DO NOT EDIT.

package gormproj

import (
    _ "embed"
    "encoding/json"
    "github.com/goxgen/goxgen/cmd/internal/integration/gormproj/generated"
)

// XgenIntrospection is the resolver for the XgenIntrospection field.
func (r *Resolver) XgenIntrospection() (*generated.XgenIntrospection, error) {
    var intr generated.XgenIntrospection
    err := json.Unmarshal(
		[]byte(`{"annotation":{"Action":[{"name":"UserInput","value":{"Action":"CREATE_MUTATION","Resource":"user","Route":"new","SchemaFieldName":"user_create"}},{"name":"UserInput","value":{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}},{"name":"CarInput","value":{"Action":"CREATE_MUTATION","Resource":"car","Route":"new","SchemaFieldName":"car_create"}},{"name":"CarInput","value":{"Action":"UPDATE_MUTATION","Resource":"car","Route":"update","SchemaFieldName":"car_update"}},{"name":"PhoneNumberInput","value":{"Action":"CREATE_MUTATION","Resource":"phone_number","Route":"new","SchemaFieldName":"phone_number_create"}},{"name":"PhoneNumberInput","value":{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"}}],"ListAction":[{"name":"CarBrowseInput","value":{"Action":"BROWSE_QUERY","Resource":"car","Route":"list","SchemaFieldName":"car_browse"}},{"name":"PhoneNumberBrowseInput","value":{"Action":"BROWSE_QUERY","Resource":"phone_number","Route":"list","SchemaFieldName":"phone_number_browse"}},{"name":"DeleteUsers","value":{"Action":"BATCH_DELETE_MUTATION","Resource":"user","Route":"delete","SchemaFieldName":"user_batch_delete"}},{"name":"ListUser","value":{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse"}}],"Resource":[{"name":"Car","value":{"DB":{"Table":"car"},"Name":"car","Primary":true,"Route":"car"}},{"name":"User","value":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}},{"name":"Phone","value":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}}]},"object":{"Car":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the todo","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"make"},"Description":"Car make","Label":"Make"}},"name":"make"},{"definition":{"Field":{"DB":{"Column":"done"},"Description":"Done of the todo","Label":"Done"}},"name":"done"},{"definition":{"Field":{"DB":{},"Description":"User of the todo","Label":"User"}},"name":"user"}],"object":{"Resource":{"DB":{"Table":"car"},"Name":"car","Primary":true,"Route":"car"}}},"CarBrowseInput":{"field":[{"definition":{"ActionField":{"Description":"ID","Label":"ID"}},"name":"id"},{"definition":{"ActionField":{"Description":"User ID","Label":"User ID"}},"name":"userId"},{"definition":{"ActionField":{"Description":"Make","Label":"Make"}},"name":"make"}],"object":{"ListAction":{"Action":"BROWSE_QUERY","Resource":"car","Route":"list","SchemaFieldName":"car_browse"}}},"CarInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the car","Label":"ID","MapTo":["Car.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Text of the todo","Label":"Make","MapTo":["Car.Make"]}},"name":"make"},{"definition":{"ActionField":{"Description":"Done of the todo","Label":"Done","MapTo":["Car.Done"]}},"name":"done"},{"definition":{"ActionField":{"Description":"User of the todo","Label":"User","MapTo":["Car.User"]}},"name":"user"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"car","Route":"update","SchemaFieldName":"car_update"}}},"DeleteUsers":{"field":[{"definition":{"ActionField":{"Description":"IDs of users","Label":"IDs"}},"name":"ids"}],"object":{"ListAction":{"Action":"BATCH_DELETE_MUTATION","Resource":"user","Route":"delete","SchemaFieldName":"user_batch_delete"}}},"ListUser":{"field":[{"definition":{"ActionField":{"Description":"ID","Label":"ID","MapTo":["User.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Name","Label":"Name","MapTo":["User.Name"]}},"name":"name"}],"object":{"ListAction":{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse"}}},"Phone":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the phone number","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"number"},"Description":"Number of phone","Label":"Number"}},"name":"number"},{"definition":{"Field":{"DB":{},"Description":"User of the todo","Label":"User"}},"name":"user"}],"object":{"Resource":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}}},"PhoneNumberBrowseInput":{"field":[{"definition":{"ActionField":{"Description":"ID","Label":"ID","MapTo":["Phone.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Number of phone","Label":"Number","MapTo":["Phone.Number"]}},"name":"number"}],"object":{"ListAction":{"Action":"BROWSE_QUERY","Resource":"phone_number","Route":"list","SchemaFieldName":"phone_number_browse"}}},"PhoneNumberInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the phone number","Label":"ID","MapTo":["Phone.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Number of phone","Label":"Name","MapTo":["Phone.Number"]}},"name":"number"},{"definition":{"ActionField":{"Description":"User of the phone","Label":"User","MapTo":["Phone.User"]}},"name":"user"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"}}},"User":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the todo","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"name","Unique":true},"Description":"Text of the todo","Label":"Text"}},"name":"name"},{"definition":{"Field":{"DB":{},"Description":"Cars of the todo","Label":"Cars"}},"name":"cars"},{"definition":{"Field":{"DB":{},"Description":"Phone numbers of the user","Label":"Phone Numbers"}},"name":"phoneNumbers"}],"object":{"Resource":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}}},"UserInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the user","Label":"ID","MapTo":["User.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Name","Label":"Name","MapTo":["User.Name"]}},"name":"name"},{"definition":{"ActionField":{"Description":"Cars of the user","Label":"Cars","MapTo":["User.Cars"]}},"name":"cars"},{"definition":{"ActionField":{"Description":"Phone numbers of the user","Label":"Phone Numbers","MapTo":["User.PhoneNumbers"]}},"name":"phones"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}}},"XgenCursorPaginationInput":{"field":[{"definition":{},"name":"first"},{"definition":{},"name":"after"},{"definition":{},"name":"last"},{"definition":{},"name":"before"}],"object":{}},"XgenPaginationInput":{"field":[{"definition":{},"name":"page"},{"definition":{},"name":"size"}],"object":{}},"XgenResourceActionType":{"field":[],"object":{}},"XgenResourceDbConfigInput":{"field":[{"definition":{},"name":"Table"}],"object":{}},"XgenResourceFieldDbConfigInput":{"field":[{"definition":{},"name":"Column"},{"definition":{},"name":"PrimaryKey"},{"definition":{},"name":"AutoIncrement"},{"definition":{},"name":"Unique"},{"definition":{},"name":"NotNull"},{"definition":{},"name":"Index"},{"definition":{},"name":"UniqueIndex"},{"definition":{},"name":"Size"},{"definition":{},"name":"Precision"},{"definition":{},"name":"Type"},{"definition":{},"name":"Scale"},{"definition":{},"name":"AutoIncrementIncrement"}],"object":{}},"XgenResourceListActionType":{"field":[],"object":{}}},"resource":{"car":{"actions":[{"Action":"BROWSE_QUERY","Resource":"car","Route":"list","SchemaFieldName":"car_browse"},{"Action":"CREATE_MUTATION","Resource":"car","Route":"new","SchemaFieldName":"car_create"},{"Action":"UPDATE_MUTATION","Resource":"car","Route":"update","SchemaFieldName":"car_update"}],"objectName":"Car","properties":{"DB":{"Table":"car"},"Name":"car","Primary":true,"Route":"car"}},"phone_number":{"actions":[{"Action":"CREATE_MUTATION","Resource":"phone_number","Route":"new","SchemaFieldName":"phone_number_create"},{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"},{"Action":"BROWSE_QUERY","Resource":"phone_number","Route":"list","SchemaFieldName":"phone_number_browse"}],"objectName":"Phone","properties":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}},"user":{"actions":[{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse"},{"Action":"BATCH_DELETE_MUTATION","Resource":"user","Route":"delete","SchemaFieldName":"user_batch_delete"},{"Action":"CREATE_MUTATION","Resource":"user","Route":"new","SchemaFieldName":"user_create"},{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}],"objectName":"User","properties":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}}}}`),
		&intr,
	)
    if err != nil {
        panic(err)
        return nil, err
    }
    return &intr, nil
}