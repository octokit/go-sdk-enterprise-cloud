package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemLicenseSyncStatusRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\license-sync-status
type ItemLicenseSyncStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemLicenseSyncStatusRequestBuilderInternal instantiates a new ItemLicenseSyncStatusRequestBuilder and sets the default values.
func NewItemLicenseSyncStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLicenseSyncStatusRequestBuilder) {
    m := &ItemLicenseSyncStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/license-sync-status", pathParameters),
    }
    return m
}
// NewItemLicenseSyncStatusRequestBuilder instantiates a new ItemLicenseSyncStatusRequestBuilder and sets the default values.
func NewItemLicenseSyncStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLicenseSyncStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemLicenseSyncStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about the status of a license sync job for an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// returns a GetLicenseSyncStatusable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/license#get-a-license-sync-status
func (m *ItemLicenseSyncStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetLicenseSyncStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateGetLicenseSyncStatusFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetLicenseSyncStatusable), nil
}
// ToGetRequestInformation gets information about the status of a license sync job for an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemLicenseSyncStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemLicenseSyncStatusRequestBuilder when successful
func (m *ItemLicenseSyncStatusRequestBuilder) WithUrl(rawUrl string)(*ItemLicenseSyncStatusRequestBuilder) {
    return NewItemLicenseSyncStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
