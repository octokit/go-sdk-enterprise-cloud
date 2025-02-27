package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\properties\schema\organizations\{org}
type ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCustom_property_name gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterprises.item.properties.schema.organizations.item.item collection
// returns a *ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder when successful
func (m *ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder) ByCustom_property_name(custom_property_name string)(*ItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if custom_property_name != "" {
        urlTplParams["custom_property_name"] = custom_property_name
    }
    return NewItemPropertiesSchemaOrganizationsItemWithCustom_property_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilderInternal instantiates a new ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder) {
    m := &ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/properties/schema/organizations/{org}", pathParameters),
    }
    return m
}
// NewItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder instantiates a new ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
