package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemSettingsBillingAdvancedSecurityRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\billing\advanced-security
type ItemSettingsBillingAdvancedSecurityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters gets the GitHub Advanced Security active committers for an organization per repository.Each distinct user login across all repositories is counted as a single Advanced Security seat, so the `total_advanced_security_committers` is not the sum of advanced_security_committers for each repository.If this organization defers to an enterprise for billing, the `total_advanced_security_committers` returned from the organization API may include some users that are in more than one organization, so they will only consume a single Advanced Security seat at the enterprise level.The total number of repositories with committer information is tracked by the `total_count` field.
type ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemSettingsBillingAdvancedSecurityRequestBuilderInternal instantiates a new ItemSettingsBillingAdvancedSecurityRequestBuilder and sets the default values.
func NewItemSettingsBillingAdvancedSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingAdvancedSecurityRequestBuilder) {
    m := &ItemSettingsBillingAdvancedSecurityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/billing/advanced-security{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemSettingsBillingAdvancedSecurityRequestBuilder instantiates a new ItemSettingsBillingAdvancedSecurityRequestBuilder and sets the default values.
func NewItemSettingsBillingAdvancedSecurityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingAdvancedSecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingAdvancedSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the GitHub Advanced Security active committers for an organization per repository.Each distinct user login across all repositories is counted as a single Advanced Security seat, so the `total_advanced_security_committers` is not the sum of advanced_security_committers for each repository.If this organization defers to an enterprise for billing, the `total_advanced_security_committers` returned from the organization API may include some users that are in more than one organization, so they will only consume a single Advanced Security seat at the enterprise level.The total number of repositories with committer information is tracked by the `total_count` field.
// returns a AdvancedSecurityActiveCommittersable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/billing/billing#get-github-advanced-security-active-committers-for-an-organization
func (m *ItemSettingsBillingAdvancedSecurityRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AdvancedSecurityActiveCommittersable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateAdvancedSecurityActiveCommittersFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AdvancedSecurityActiveCommittersable), nil
}
// ToGetRequestInformation gets the GitHub Advanced Security active committers for an organization per repository.Each distinct user login across all repositories is counted as a single Advanced Security seat, so the `total_advanced_security_committers` is not the sum of advanced_security_committers for each repository.If this organization defers to an enterprise for billing, the `total_advanced_security_committers` returned from the organization API may include some users that are in more than one organization, so they will only consume a single Advanced Security seat at the enterprise level.The total number of repositories with committer information is tracked by the `total_count` field.
// returns a *RequestInformation when successful
func (m *ItemSettingsBillingAdvancedSecurityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsBillingAdvancedSecurityRequestBuilder when successful
func (m *ItemSettingsBillingAdvancedSecurityRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingAdvancedSecurityRequestBuilder) {
    return NewItemSettingsBillingAdvancedSecurityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
