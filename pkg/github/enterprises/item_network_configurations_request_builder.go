package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemNetworkConfigurationsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\network-configurations
type ItemNetworkConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemNetworkConfigurationsRequestBuilderGetQueryParameters lists all hosted compute network configurations configured in an enterprise.
type ItemNetworkConfigurationsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByNetwork_configuration_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterprises.item.networkConfigurations.item collection
// returns a *ItemNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder when successful
func (m *ItemNetworkConfigurationsRequestBuilder) ByNetwork_configuration_id(network_configuration_id string)(*ItemNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if network_configuration_id != "" {
        urlTplParams["network_configuration_id"] = network_configuration_id
    }
    return NewItemNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemNetworkConfigurationsRequestBuilderInternal instantiates a new ItemNetworkConfigurationsRequestBuilder and sets the default values.
func NewItemNetworkConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemNetworkConfigurationsRequestBuilder) {
    m := &ItemNetworkConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/network-configurations{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemNetworkConfigurationsRequestBuilder instantiates a new ItemNetworkConfigurationsRequestBuilder and sets the default values.
func NewItemNetworkConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemNetworkConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemNetworkConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all hosted compute network configurations configured in an enterprise.
// returns a ItemNetworkConfigurationsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/network-configurations#list-hosted-compute-network-configurations-for-an-enterprise
func (m *ItemNetworkConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemNetworkConfigurationsRequestBuilderGetQueryParameters])(ItemNetworkConfigurationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemNetworkConfigurationsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemNetworkConfigurationsGetResponseable), nil
}
// Post creates a hosted compute network configuration for an enterprise.
// returns a NetworkConfigurationable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/network-configurations#create-a-hosted-compute-network-configuration-for-an-enterprise
func (m *ItemNetworkConfigurationsRequestBuilder) Post(ctx context.Context, body ItemNetworkConfigurationsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.NetworkConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateNetworkConfigurationFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.NetworkConfigurationable), nil
}
// ToGetRequestInformation lists all hosted compute network configurations configured in an enterprise.
// returns a *RequestInformation when successful
func (m *ItemNetworkConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemNetworkConfigurationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates a hosted compute network configuration for an enterprise.
// returns a *RequestInformation when successful
func (m *ItemNetworkConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemNetworkConfigurationsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemNetworkConfigurationsRequestBuilder when successful
func (m *ItemNetworkConfigurationsRequestBuilder) WithUrl(rawUrl string)(*ItemNetworkConfigurationsRequestBuilder) {
    return NewItemNetworkConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
