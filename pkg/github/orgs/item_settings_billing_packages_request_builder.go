package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemSettingsBillingPackagesRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\billing\packages
type ItemSettingsBillingPackagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSettingsBillingPackagesRequestBuilderInternal instantiates a new ItemSettingsBillingPackagesRequestBuilder and sets the default values.
func NewItemSettingsBillingPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingPackagesRequestBuilder) {
    m := &ItemSettingsBillingPackagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/billing/packages", pathParameters),
    }
    return m
}
// NewItemSettingsBillingPackagesRequestBuilder instantiates a new ItemSettingsBillingPackagesRequestBuilder and sets the default values.
func NewItemSettingsBillingPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the free and paid storage used for GitHub Packages in gigabytes.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."OAuth app tokens and personal access tokens (classic) need the `repo` or `admin:org` scope to use this endpoint.
// returns a PackagesBillingUsageable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/billing/billing#get-github-packages-billing-for-an-organization
func (m *ItemSettingsBillingPackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackagesBillingUsageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreatePackagesBillingUsageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackagesBillingUsageable), nil
}
// ToGetRequestInformation gets the free and paid storage used for GitHub Packages in gigabytes.Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."OAuth app tokens and personal access tokens (classic) need the `repo` or `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsBillingPackagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsBillingPackagesRequestBuilder when successful
func (m *ItemSettingsBillingPackagesRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingPackagesRequestBuilder) {
    return NewItemSettingsBillingPackagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
