package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\properties\schema\organizations\{org}\{custom_property_name}
type ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilderInternal instantiates a new ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder) {
    m := &ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/properties/schema/organizations/{org}/{custom_property_name}", pathParameters),
    }
    return m
}
// NewItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder instantiates a new ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Promote the promote property
// returns a *ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder when successful
func (m *ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder) Promote()(*ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder) {
    return NewItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
