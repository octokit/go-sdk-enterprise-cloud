package scim

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V2EnterprisesWithEnterpriseItemRequestBuilder builds and executes requests for operations under \scim\v2\enterprises\{enterprise}
type V2EnterprisesWithEnterpriseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV2EnterprisesWithEnterpriseItemRequestBuilderInternal instantiates a new V2EnterprisesWithEnterpriseItemRequestBuilder and sets the default values.
func NewV2EnterprisesWithEnterpriseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesWithEnterpriseItemRequestBuilder) {
    m := &V2EnterprisesWithEnterpriseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/enterprises/{enterprise}", pathParameters),
    }
    return m
}
// NewV2EnterprisesWithEnterpriseItemRequestBuilder instantiates a new V2EnterprisesWithEnterpriseItemRequestBuilder and sets the default values.
func NewV2EnterprisesWithEnterpriseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesWithEnterpriseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2EnterprisesWithEnterpriseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Groups the Groups property
// returns a *V2EnterprisesItemGroupsRequestBuilder when successful
func (m *V2EnterprisesWithEnterpriseItemRequestBuilder) Groups()(*V2EnterprisesItemGroupsRequestBuilder) {
    return NewV2EnterprisesItemGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the Users property
// returns a *V2EnterprisesItemUsersRequestBuilder when successful
func (m *V2EnterprisesWithEnterpriseItemRequestBuilder) Users()(*V2EnterprisesItemUsersRequestBuilder) {
    return NewV2EnterprisesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
