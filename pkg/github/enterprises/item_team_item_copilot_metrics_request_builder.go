package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemTeamItemCopilotMetricsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\team\{team_slug}\copilot\metrics
type ItemTeamItemCopilotMetricsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamItemCopilotMetricsRequestBuilderGetQueryParameters > [!NOTE]> This endpoint is only applicable to dedicated enterprise accounts for Copilot Business. See "[About enterprise accounts for Copilot Business](https://docs.github.com/enterprise-cloud@latest//admin/copilot-business-only/about-enterprise-accounts-for-copilot-business)."Use this endpoint to see a breakdown of aggregated metrics for various GitHub Copilot features. See the response schema tab for detailed metrics definitions.The response contains metrics for up to 28 days prior. Metrics are processed once per day for the previous day,and the response will only include data up until yesterday. In order for an end user to be counted towards these metrics,they must have telemetry enabled in their IDE.> [!NOTE]> This endpoint will only return results for a given day if the enterprise team had **five or more members with active Copilot licenses** on that day, as evaluated at the end of that day.To access this endpoint, the Copilot Metrics API access policy must be enabled or set to "no policy" for the enterprise within GitHub settings.Only owners and billing managers for the enterprise that contains the enterprise team can view Copilot metrics for the enterprise team.OAuth app tokens and personal access tokens (classic) need either the `manage_billing:copilot` or `read:enterprise` scopes to use this endpoint.
type ItemTeamItemCopilotMetricsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of days of metrics to display per page (max 28). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Show usage metrics since this date. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:MM:SSZ`). Maximum value is 28 days ago.
    Since *string `uriparametername:"since"`
    // Show usage metrics until this date. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:MM:SSZ`) and should not preceed the `since` date if it is passed.
    Until *string `uriparametername:"until"`
}
// NewItemTeamItemCopilotMetricsRequestBuilderInternal instantiates a new ItemTeamItemCopilotMetricsRequestBuilder and sets the default values.
func NewItemTeamItemCopilotMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamItemCopilotMetricsRequestBuilder) {
    m := &ItemTeamItemCopilotMetricsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/team/{team_slug}/copilot/metrics{?page*,per_page*,since*,until*}", pathParameters),
    }
    return m
}
// NewItemTeamItemCopilotMetricsRequestBuilder instantiates a new ItemTeamItemCopilotMetricsRequestBuilder and sets the default values.
func NewItemTeamItemCopilotMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamItemCopilotMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamItemCopilotMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!NOTE]> This endpoint is only applicable to dedicated enterprise accounts for Copilot Business. See "[About enterprise accounts for Copilot Business](https://docs.github.com/enterprise-cloud@latest//admin/copilot-business-only/about-enterprise-accounts-for-copilot-business)."Use this endpoint to see a breakdown of aggregated metrics for various GitHub Copilot features. See the response schema tab for detailed metrics definitions.The response contains metrics for up to 28 days prior. Metrics are processed once per day for the previous day,and the response will only include data up until yesterday. In order for an end user to be counted towards these metrics,they must have telemetry enabled in their IDE.> [!NOTE]> This endpoint will only return results for a given day if the enterprise team had **five or more members with active Copilot licenses** on that day, as evaluated at the end of that day.To access this endpoint, the Copilot Metrics API access policy must be enabled or set to "no policy" for the enterprise within GitHub settings.Only owners and billing managers for the enterprise that contains the enterprise team can view Copilot metrics for the enterprise team.OAuth app tokens and personal access tokens (classic) need either the `manage_billing:copilot` or `read:enterprise` scopes to use this endpoint.
// returns a []CopilotUsageMetricsDayable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 422 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/copilot/copilot-metrics#get-copilot-metrics-for-an-enterprise-team
func (m *ItemTeamItemCopilotMetricsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamItemCopilotMetricsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CopilotUsageMetricsDayable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateCopilotUsageMetricsDayFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CopilotUsageMetricsDayable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CopilotUsageMetricsDayable)
        }
    }
    return val, nil
}
// ToGetRequestInformation > [!NOTE]> This endpoint is only applicable to dedicated enterprise accounts for Copilot Business. See "[About enterprise accounts for Copilot Business](https://docs.github.com/enterprise-cloud@latest//admin/copilot-business-only/about-enterprise-accounts-for-copilot-business)."Use this endpoint to see a breakdown of aggregated metrics for various GitHub Copilot features. See the response schema tab for detailed metrics definitions.The response contains metrics for up to 28 days prior. Metrics are processed once per day for the previous day,and the response will only include data up until yesterday. In order for an end user to be counted towards these metrics,they must have telemetry enabled in their IDE.> [!NOTE]> This endpoint will only return results for a given day if the enterprise team had **five or more members with active Copilot licenses** on that day, as evaluated at the end of that day.To access this endpoint, the Copilot Metrics API access policy must be enabled or set to "no policy" for the enterprise within GitHub settings.Only owners and billing managers for the enterprise that contains the enterprise team can view Copilot metrics for the enterprise team.OAuth app tokens and personal access tokens (classic) need either the `manage_billing:copilot` or `read:enterprise` scopes to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemTeamItemCopilotMetricsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamItemCopilotMetricsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamItemCopilotMetricsRequestBuilder when successful
func (m *ItemTeamItemCopilotMetricsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamItemCopilotMetricsRequestBuilder) {
    return NewItemTeamItemCopilotMetricsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
