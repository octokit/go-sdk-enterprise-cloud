package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemSettingsBillingSharedStorageRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\billing\shared-storage
type ItemSettingsBillingSharedStorageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSettingsBillingSharedStorageRequestBuilderInternal instantiates a new ItemSettingsBillingSharedStorageRequestBuilder and sets the default values.
func NewItemSettingsBillingSharedStorageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingSharedStorageRequestBuilder) {
    m := &ItemSettingsBillingSharedStorageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/billing/shared-storage", pathParameters),
    }
    return m
}
// NewItemSettingsBillingSharedStorageRequestBuilder instantiates a new ItemSettingsBillingSharedStorageRequestBuilder and sets the default values.
func NewItemSettingsBillingSharedStorageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingSharedStorageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingSharedStorageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the estimated paid and estimated total storage used for GitHub Actions and GitHub Packages.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."OAuth app tokens and personal access tokens (classic) need the `repo` or `admin:org` scope to use this endpoint.
// returns a CombinedBillingUsageable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/billing/billing#get-shared-storage-billing-for-an-organization
func (m *ItemSettingsBillingSharedStorageRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CombinedBillingUsageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateCombinedBillingUsageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CombinedBillingUsageable), nil
}
// ToGetRequestInformation gets the estimated paid and estimated total storage used for GitHub Actions and GitHub Packages.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."OAuth app tokens and personal access tokens (classic) need the `repo` or `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsBillingSharedStorageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsBillingSharedStorageRequestBuilder when successful
func (m *ItemSettingsBillingSharedStorageRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingSharedStorageRequestBuilder) {
    return NewItemSettingsBillingSharedStorageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
