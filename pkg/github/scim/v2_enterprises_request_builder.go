package scim

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V2EnterprisesRequestBuilder builds and executes requests for operations under \scim\v2\enterprises
type V2EnterprisesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByEnterprise gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.scim.v2.enterprises.item collection
// returns a *V2EnterprisesWithEnterpriseItemRequestBuilder when successful
func (m *V2EnterprisesRequestBuilder) ByEnterprise(enterprise string)(*V2EnterprisesWithEnterpriseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if enterprise != "" {
        urlTplParams["enterprise"] = enterprise
    }
    return NewV2EnterprisesWithEnterpriseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV2EnterprisesRequestBuilderInternal instantiates a new V2EnterprisesRequestBuilder and sets the default values.
func NewV2EnterprisesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesRequestBuilder) {
    m := &V2EnterprisesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/enterprises", pathParameters),
    }
    return m
}
// NewV2EnterprisesRequestBuilder instantiates a new V2EnterprisesRequestBuilder and sets the default values.
func NewV2EnterprisesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2EnterprisesRequestBuilderInternal(urlParams, requestAdapter)
}
