package scim

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ScimRequestBuilder builds and executes requests for operations under \scim
type ScimRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewScimRequestBuilderInternal instantiates a new ScimRequestBuilder and sets the default values.
func NewScimRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScimRequestBuilder) {
    m := &ScimRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim", pathParameters),
    }
    return m
}
// NewScimRequestBuilder instantiates a new ScimRequestBuilder and sets the default values.
func NewScimRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScimRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScimRequestBuilderInternal(urlParams, requestAdapter)
}
// V2 the v2 property
// returns a *V2RequestBuilder when successful
func (m *ScimRequestBuilder) V2()(*V2RequestBuilder) {
    return NewV2RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
