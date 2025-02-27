package repos

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemBypassResponsesSecretScanningRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-responses\secret-scanning
type ItemItemBypassResponsesSecretScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByBypass_response_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.repos.item.item.bypassResponses.secretScanning.item collection
// returns a *ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder when successful
func (m *ItemItemBypassResponsesSecretScanningRequestBuilder) ByBypass_response_id(bypass_response_id int32)(*ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["bypass_response_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(bypass_response_id), 10)
    return NewItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemBypassResponsesSecretScanningRequestBuilderInternal instantiates a new ItemItemBypassResponsesSecretScanningRequestBuilder and sets the default values.
func NewItemItemBypassResponsesSecretScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassResponsesSecretScanningRequestBuilder) {
    m := &ItemItemBypassResponsesSecretScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-responses/secret-scanning", pathParameters),
    }
    return m
}
// NewItemItemBypassResponsesSecretScanningRequestBuilder instantiates a new ItemItemBypassResponsesSecretScanningRequestBuilder and sets the default values.
func NewItemItemBypassResponsesSecretScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassResponsesSecretScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassResponsesSecretScanningRequestBuilderInternal(urlParams, requestAdapter)
}
