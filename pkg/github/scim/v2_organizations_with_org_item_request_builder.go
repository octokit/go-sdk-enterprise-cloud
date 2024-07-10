package scim

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V2OrganizationsWithOrgItemRequestBuilder builds and executes requests for operations under \scim\v2\organizations\{org}
type V2OrganizationsWithOrgItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV2OrganizationsWithOrgItemRequestBuilderInternal instantiates a new V2OrganizationsWithOrgItemRequestBuilder and sets the default values.
func NewV2OrganizationsWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2OrganizationsWithOrgItemRequestBuilder) {
    m := &V2OrganizationsWithOrgItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/organizations/{org}", pathParameters),
    }
    return m
}
// NewV2OrganizationsWithOrgItemRequestBuilder instantiates a new V2OrganizationsWithOrgItemRequestBuilder and sets the default values.
func NewV2OrganizationsWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2OrganizationsWithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2OrganizationsWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Users the Users property
// returns a *V2OrganizationsItemUsersRequestBuilder when successful
func (m *V2OrganizationsWithOrgItemRequestBuilder) Users()(*V2OrganizationsItemUsersRequestBuilder) {
    return NewV2OrganizationsItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
