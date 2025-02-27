package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPropertiesSchemaOrganizationsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\properties\schema\organizations
type ItemPropertiesSchemaOrganizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByOrg gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterprises.item.properties.schema.organizations.item collection
// returns a *ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder when successful
func (m *ItemPropertiesSchemaOrganizationsRequestBuilder) ByOrg(org string)(*ItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if org != "" {
        urlTplParams["org"] = org
    }
    return NewItemPropertiesSchemaOrganizationsWithOrgItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPropertiesSchemaOrganizationsRequestBuilderInternal instantiates a new ItemPropertiesSchemaOrganizationsRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsRequestBuilder) {
    m := &ItemPropertiesSchemaOrganizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/properties/schema/organizations", pathParameters),
    }
    return m
}
// NewItemPropertiesSchemaOrganizationsRequestBuilder instantiates a new ItemPropertiesSchemaOrganizationsRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPropertiesSchemaOrganizationsRequestBuilderInternal(urlParams, requestAdapter)
}
