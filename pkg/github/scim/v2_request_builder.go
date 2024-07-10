package scim

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V2RequestBuilder builds and executes requests for operations under \scim\v2
type V2RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV2RequestBuilderInternal instantiates a new V2RequestBuilder and sets the default values.
func NewV2RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2RequestBuilder) {
    m := &V2RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2", pathParameters),
    }
    return m
}
// NewV2RequestBuilder instantiates a new V2RequestBuilder and sets the default values.
func NewV2RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2RequestBuilderInternal(urlParams, requestAdapter)
}
// Enterprises the enterprises property
// returns a *V2EnterprisesRequestBuilder when successful
func (m *V2RequestBuilder) Enterprises()(*V2EnterprisesRequestBuilder) {
    return NewV2EnterprisesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
// returns a *V2OrganizationsRequestBuilder when successful
func (m *V2RequestBuilder) Organizations()(*V2OrganizationsRequestBuilder) {
    return NewV2OrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
