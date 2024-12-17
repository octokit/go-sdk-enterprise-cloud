package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemBypassRequestsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-requests
type ItemItemBypassRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemBypassRequestsRequestBuilderInternal instantiates a new ItemItemBypassRequestsRequestBuilder and sets the default values.
func NewItemItemBypassRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsRequestBuilder) {
    m := &ItemItemBypassRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-requests", pathParameters),
    }
    return m
}
// NewItemItemBypassRequestsRequestBuilder instantiates a new ItemItemBypassRequestsRequestBuilder and sets the default values.
func NewItemItemBypassRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// PushRules the pushRules property
// returns a *ItemItemBypassRequestsPushRulesRequestBuilder when successful
func (m *ItemItemBypassRequestsRequestBuilder) PushRules()(*ItemItemBypassRequestsPushRulesRequestBuilder) {
    return NewItemItemBypassRequestsPushRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
