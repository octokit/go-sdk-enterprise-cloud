package scim

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V2OrganizationsRequestBuilder builds and executes requests for operations under \scim\v2\organizations
type V2OrganizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByOrg gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.scim.v2.organizations.item collection
// returns a *V2OrganizationsWithOrgItemRequestBuilder when successful
func (m *V2OrganizationsRequestBuilder) ByOrg(org string)(*V2OrganizationsWithOrgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if org != "" {
        urlTplParams["org"] = org
    }
    return NewV2OrganizationsWithOrgItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV2OrganizationsRequestBuilderInternal instantiates a new V2OrganizationsRequestBuilder and sets the default values.
func NewV2OrganizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2OrganizationsRequestBuilder) {
    m := &V2OrganizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/organizations", pathParameters),
    }
    return m
}
// NewV2OrganizationsRequestBuilder instantiates a new V2OrganizationsRequestBuilder and sets the default values.
func NewV2OrganizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2OrganizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2OrganizationsRequestBuilderInternal(urlParams, requestAdapter)
}
