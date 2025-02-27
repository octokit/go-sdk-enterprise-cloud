package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemBypassResponsesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-responses
type ItemItemBypassResponsesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemBypassResponsesRequestBuilderInternal instantiates a new ItemItemBypassResponsesRequestBuilder and sets the default values.
func NewItemItemBypassResponsesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassResponsesRequestBuilder) {
    m := &ItemItemBypassResponsesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-responses", pathParameters),
    }
    return m
}
// NewItemItemBypassResponsesRequestBuilder instantiates a new ItemItemBypassResponsesRequestBuilder and sets the default values.
func NewItemItemBypassResponsesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassResponsesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassResponsesRequestBuilderInternal(urlParams, requestAdapter)
}
// SecretScanning the secretScanning property
// returns a *ItemItemBypassResponsesSecretScanningRequestBuilder when successful
func (m *ItemItemBypassResponsesRequestBuilder) SecretScanning()(*ItemItemBypassResponsesSecretScanningRequestBuilder) {
    return NewItemItemBypassResponsesSecretScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
