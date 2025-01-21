package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\network-settings\{network_settings_id}
type ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal instantiates a new ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder and sets the default values.
func NewItemNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    m := &ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/network-settings/{network_settings_id}", pathParameters),
    }
    return m
}
// NewItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder instantiates a new ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder and sets the default values.
func NewItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a hosted compute network settings resource configured for an enterprise.
// returns a NetworkSettingsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/network-configurations#get-a-hosted-compute-network-settings-resource-for-an-enterprise
func (m *ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.NetworkSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateNetworkSettingsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.NetworkSettingsable), nil
}
// ToGetRequestInformation gets a hosted compute network settings resource configured for an enterprise.
// returns a *RequestInformation when successful
func (m *ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder when successful
func (m *ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    return NewItemNetworkSettingsWithNetwork_settings_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
