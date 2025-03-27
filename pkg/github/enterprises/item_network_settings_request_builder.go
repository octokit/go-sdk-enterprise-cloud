package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemNetworkSettingsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\network-settings
type ItemNetworkSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByNetwork_settings_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterprises.item.networkSettings.item collection
// returns a *ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder when successful
func (m *ItemNetworkSettingsRequestBuilder) ByNetwork_settings_id(network_settings_id string)(*ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if network_settings_id != "" {
        urlTplParams["network_settings_id"] = network_settings_id
    }
    return NewItemNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemNetworkSettingsRequestBuilderInternal instantiates a new ItemNetworkSettingsRequestBuilder and sets the default values.
func NewItemNetworkSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemNetworkSettingsRequestBuilder) {
    m := &ItemNetworkSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/network-settings", pathParameters),
    }
    return m
}
// NewItemNetworkSettingsRequestBuilder instantiates a new ItemNetworkSettingsRequestBuilder and sets the default values.
func NewItemNetworkSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemNetworkSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemNetworkSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
