// Code generated by github.com/ikhfa/goxgen, DO NOT EDIT.

package generated

import (
    _ "embed"
    "encoding/json"
)

// XgenIntrospection is the resolver for the XgenIntrospection field.
func XgenIntrospectionValues() (*XgenIntrospection, error) {
    var intr XgenIntrospection
    err := json.Unmarshal(
		[]byte(`{"annotation":{"Action":[{"name":"CarInput","value":{"Action":"CREATE_MUTATION","Resource":"car","Route":"new","SchemaFieldName":"car_create"}},{"name":"CarInput","value":{"Action":"UPDATE_MUTATION","Resource":"car","Route":"update","SchemaFieldName":"car_update"}},{"name":"PhoneNumberInput","value":{"Action":"CREATE_MUTATION","Resource":"phone_number","Route":"new","SchemaFieldName":"phone_number_create"}},{"name":"PhoneNumberInput","value":{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"}},{"name":"UserInput","value":{"Action":"CREATE_MUTATION","Resource":"user","Route":"new","SchemaFieldName":"user_create"}},{"name":"UserInput","value":{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}}],"ListAction":[{"name":"ListUser","value":{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse","Sort":{"Default":[{"by":"name","direction":"ASC"}]},"sort":{}}},{"name":"PhoneNumberBrowseInput","value":{"Action":"BROWSE_QUERY","Resource":"phone_number","Route":"list","SchemaFieldName":"phone_number_browse","sort":{}}},{"name":"DeleteUsers","value":{"Action":"BATCH_DELETE_MUTATION","Resource":"user","Route":"delete","SchemaFieldName":"user_batch_delete","sort":{}}},{"name":"CarBrowseInput","value":{"Action":"BROWSE_QUERY","Resource":"car","Route":"list","SchemaFieldName":"car_browse","sort":{}}}],"Resource":[{"name":"Phone","value":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}},{"name":"User","value":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}},{"name":"Car","value":{"DB":{"Table":"car"},"Name":"car","Primary":true,"Route":"car"}}]},"object":{"Car":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the todo","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"make"},"Description":"Car make","Label":"Make"}},"name":"make"},{"definition":{"Field":{"DB":{"Column":"done"},"Description":"Done of the todo","Label":"Done"}},"name":"done"},{"definition":{"Field":{"DB":{},"Description":"User of the todo","Label":"User"}},"name":"user"}],"object":{"Resource":{"DB":{"Table":"car"},"Name":"car","Primary":true,"Route":"car"}}},"CarBrowseInput":{"field":[{"definition":{"ActionField":{"Description":"ID","Label":"ID"}},"name":"id"},{"definition":{"ActionField":{"Description":"User ID","Label":"User ID"}},"name":"userId"},{"definition":{"ActionField":{"Description":"Make","Label":"Make"}},"name":"make"}],"object":{"ListAction":{"Action":"BROWSE_QUERY","Resource":"car","Route":"list","SchemaFieldName":"car_browse","sort":{}}}},"CarInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the car","Label":"ID","MapTo":["Car.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Text of the todo","Label":"Make","MapTo":["Car.Make"]}},"name":"make"},{"definition":{"ActionField":{"Description":"Done of the todo","Label":"Done","MapTo":["Car.Done"]}},"name":"done"},{"definition":{"ActionField":{"Description":"User of the todo","Label":"User","MapTo":["Car.User"]}},"name":"user"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"car","Route":"update","SchemaFieldName":"car_update"}}},"DeleteUsers":{"field":[{"definition":{"ActionField":{"Description":"IDs of users","Label":"IDs"}},"name":"ids"}],"object":{"ListAction":{"Action":"BATCH_DELETE_MUTATION","Resource":"user","Route":"delete","SchemaFieldName":"user_batch_delete","sort":{}}}},"ListUser":{"field":[{"definition":{"ActionField":{"Description":"ID","Label":"ID","MapTo":["User.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Name","Label":"Name","MapTo":["User.Name"]}},"name":"name"}],"object":{"ListAction":{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse","Sort":{"Default":[{"by":"name","direction":"ASC"}]},"sort":{}}}},"Phone":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the phone number","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"number"},"Description":"Number of phone","Label":"Number"}},"name":"number"},{"definition":{"Field":{"DB":{},"Description":"User of the todo","Label":"User"}},"name":"user"}],"object":{"Resource":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}}},"PhoneNumberBrowseInput":{"field":[{"definition":{"ActionField":{"Description":"ID","Label":"ID","MapTo":["Phone.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Number of phone","Label":"Number","MapTo":["Phone.Number"]}},"name":"number"}],"object":{"ListAction":{"Action":"BROWSE_QUERY","Resource":"phone_number","Route":"list","SchemaFieldName":"phone_number_browse","sort":{}}}},"PhoneNumberInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the phone number","Label":"ID","MapTo":["Phone.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Number of phone","Label":"Name","MapTo":["Phone.Number"]}},"name":"number"},{"definition":{"ActionField":{"Description":"User of the phone","Label":"User","MapTo":["Phone.User"]}},"name":"user"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"}}},"User":{"field":[{"definition":{"Field":{"DB":{"Column":"id","PrimaryKey":true},"Description":"ID of the user","Label":"ID"}},"name":"id"},{"definition":{"Field":{"DB":{"Column":"name","Unique":true},"Description":"Text of the user","Label":"Text"}},"name":"name"},{"definition":{"Field":{"DB":{},"Description":"Cars of the user","Label":"Cars"}},"name":"cars"},{"definition":{"Field":{"DB":{},"Description":"Phone numbers of the user","Label":"Phone Numbers"}},"name":"phoneNumbers"}],"object":{"Resource":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}}},"UserInput":{"field":[{"definition":{"ActionField":{"Description":"ID of the user","Label":"ID","MapTo":["User.ID"]}},"name":"id"},{"definition":{"ActionField":{"Description":"Name","Label":"Name","MapTo":["User.Name"]}},"name":"name"},{"definition":{"ActionField":{"Description":"Cars of the user","Label":"Cars","MapTo":["User.Cars"]}},"name":"cars"},{"definition":{"ActionField":{"Description":"Phone numbers of the user","Label":"Phone Numbers","MapTo":["User.PhoneNumbers"]}},"name":"phones"}],"object":{"Action":{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}}},"XgenCursorPaginationInput":{"field":[{"definition":{},"name":"first"},{"definition":{},"name":"after"},{"definition":{},"name":"last"},{"definition":{},"name":"before"}],"object":{}},"XgenPaginationInput":{"field":[{"definition":{},"name":"page"},{"definition":{},"name":"size"}],"object":{}},"XgenResourceActionType":{"field":[],"object":{}},"XgenResourceDbConfigInput":{"field":[{"definition":{},"name":"Table"}],"object":{}},"XgenResourceFieldDbConfigInput":{"field":[{"definition":{},"name":"Column"},{"definition":{},"name":"PrimaryKey"},{"definition":{},"name":"AutoIncrement"},{"definition":{},"name":"Unique"},{"definition":{},"name":"NotNull"},{"definition":{},"name":"Index"},{"definition":{},"name":"UniqueIndex"},{"definition":{},"name":"Size"},{"definition":{},"name":"Precision"},{"definition":{},"name":"Type"},{"definition":{},"name":"Scale"},{"definition":{},"name":"AutoIncrementIncrement"}],"object":{}},"XgenResourceListActionType":{"field":[],"object":{}},"XgenSort":{"field":[{"definition":{},"name":"by"},{"definition":{},"name":"direction"}],"object":{}},"XgenSortDirection":{"field":[],"object":{}},"XgenSortInput":{"field":[{"definition":{},"name":"by"},{"definition":{},"name":"direction"}],"object":{}},"XgenSortResourceConfig":{"field":[{"definition":{},"name":"Disabled"},{"definition":{},"name":"Default"}],"object":{}},"XgenSortResourceConfigInput":{"field":[{"definition":{},"name":"Disabled"},{"definition":{"ToObjectType":{"type":"[XgenSort!]"}},"name":"Default"}],"object":{}}},"resource":{"car":{"actions":[{"Action":"CREATE_MUTATION","Resource":"car","Route":"new","SchemaFieldName":"car_create"},{"Action":"UPDATE_MUTATION","Resource":"car","Route":"update","SchemaFieldName":"car_update"},{"Action":"BROWSE_QUERY","Resource":"car","Route":"list","SchemaFieldName":"car_browse","sort":{}}],"objectName":"Car","properties":{"DB":{"Table":"car"},"Name":"car","Primary":true,"Route":"car"}},"phone_number":{"actions":[{"Action":"CREATE_MUTATION","Resource":"phone_number","Route":"new","SchemaFieldName":"phone_number_create"},{"Action":"UPDATE_MUTATION","Resource":"phone_number","Route":"update","SchemaFieldName":"phone_number_update"},{"Action":"BROWSE_QUERY","Resource":"phone_number","Route":"list","SchemaFieldName":"phone_number_browse","sort":{}}],"objectName":"Phone","properties":{"DB":{"Table":"phone_number"},"Name":"phone_number","Primary":true,"Route":"phone-number"}},"user":{"actions":[{"Action":"BATCH_DELETE_MUTATION","Resource":"user","Route":"delete","SchemaFieldName":"user_batch_delete","sort":{}},{"Action":"BROWSE_QUERY","Pagination":true,"Resource":"user","Route":"list","SchemaFieldName":"user_browse","Sort":{"Default":[{"by":"name","direction":"ASC"}]},"sort":{}},{"Action":"CREATE_MUTATION","Resource":"user","Route":"new","SchemaFieldName":"user_create"},{"Action":"UPDATE_MUTATION","Resource":"user","Route":"update","SchemaFieldName":"user_update"}],"objectName":"User","properties":{"DB":{"Table":"user"},"Name":"user","Primary":true,"Route":"user"}}}}`),
		&intr,
	)
    if err != nil {
        panic(err)
        return nil, err
    }
    return &intr, nil
}