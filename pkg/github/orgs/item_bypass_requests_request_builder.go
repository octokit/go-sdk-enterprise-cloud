package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemBypassRequestsRequestBuilder builds and executes requests for operations under \orgs\{org}\bypass-requests
type ItemBypassRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemBypassRequestsRequestBuilderInternal instantiates a new ItemBypassRequestsRequestBuilder and sets the default values.
func NewItemBypassRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBypassRequestsRequestBuilder) {
    m := &ItemBypassRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/bypass-requests", pathParameters),
    }
    return m
}
// NewItemBypassRequestsRequestBuilder instantiates a new ItemBypassRequestsRequestBuilder and sets the default values.
func NewItemBypassRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBypassRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBypassRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// PushRules the pushRules property
// returns a *ItemBypassRequestsPushRulesRequestBuilder when successful
func (m *ItemBypassRequestsRequestBuilder) PushRules()(*ItemBypassRequestsPushRulesRequestBuilder) {
    return NewItemBypassRequestsPushRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
