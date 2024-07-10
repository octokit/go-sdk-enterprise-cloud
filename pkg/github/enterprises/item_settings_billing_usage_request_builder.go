package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemSettingsBillingUsageRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\settings\billing\usage
type ItemSettingsBillingUsageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsBillingUsageRequestBuilderGetQueryParameters gets a report of the total usage for an enterprise. To use this endpoint, you must be an administrator or billing manager of the enterprise.
type ItemSettingsBillingUsageRequestBuilderGetQueryParameters struct {
    // The ID corresponding to a cost center.
    Cost_center_id *string `uriparametername:"cost_center_id"`
    // If specified, only return results for a single day. The value of `day` is an integer between `1` and `31`.
    Day *int32 `uriparametername:"day"`
    // If specified, only return results for a single hour. The value of `hour` is an integer between `0` and `23`.
    Hour *int32 `uriparametername:"hour"`
    // If specified, only return results for a single month. The value of `month` is an integer between `1` and `12`.
    Month *int32 `uriparametername:"month"`
    // If specified, only return results for a single year. The value of `year` is an integer with four digits representing a year. For example, `2023`.
    Year *int32 `uriparametername:"year"`
}
// NewItemSettingsBillingUsageRequestBuilderInternal instantiates a new ItemSettingsBillingUsageRequestBuilder and sets the default values.
func NewItemSettingsBillingUsageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingUsageRequestBuilder) {
    m := &ItemSettingsBillingUsageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/settings/billing/usage{?cost_center_id*,day*,hour*,month*,year*}", pathParameters),
    }
    return m
}
// NewItemSettingsBillingUsageRequestBuilder instantiates a new ItemSettingsBillingUsageRequestBuilder and sets the default values.
func NewItemSettingsBillingUsageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingUsageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingUsageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a report of the total usage for an enterprise. To use this endpoint, you must be an administrator or billing manager of the enterprise.
// returns a BillingUsageReportable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 500 status code
// returns a BillingUsageReport503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/billing#get-billing-usage-report-for-an-enterprise
func (m *ItemSettingsBillingUsageRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsBillingUsageRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.BillingUsageReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "503": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBillingUsageReport503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBillingUsageReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.BillingUsageReportable), nil
}
// ToGetRequestInformation gets a report of the total usage for an enterprise. To use this endpoint, you must be an administrator or billing manager of the enterprise.
// returns a *RequestInformation when successful
func (m *ItemSettingsBillingUsageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsBillingUsageRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsBillingUsageRequestBuilder when successful
func (m *ItemSettingsBillingUsageRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingUsageRequestBuilder) {
    return NewItemSettingsBillingUsageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
