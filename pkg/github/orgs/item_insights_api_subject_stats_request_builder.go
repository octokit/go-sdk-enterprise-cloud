package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i8fa9bb5dfd2b8a7a2339db11c5b176f7eff27ccbeaff093675e5ca1dd9827b74 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/insights/api/subjectstats"
)

// ItemInsightsApiSubjectStatsRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\subject-stats
type ItemInsightsApiSubjectStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters get API request statistics for all subjects within an organization within a specified time frame. Subjects can be users or GitHub Apps.
type ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i8fa9bb5dfd2b8a7a2339db11c5b176f7eff27ccbeaff093675e5ca1dd9827b74.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The maximum timestamp to query for stats. Defaults to the time 30 days ago. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Min_timestamp *string `uriparametername:"min_timestamp"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    Sort []i8fa9bb5dfd2b8a7a2339db11c5b176f7eff27ccbeaff093675e5ca1dd9827b74.GetSortQueryParameterType `uriparametername:"sort"`
    // Providing a substring will filter results where the subject name contains the substring. This is a case-insensitive search.
    Subject_name_substring *string `uriparametername:"subject_name_substring"`
}
// NewItemInsightsApiSubjectStatsRequestBuilderInternal instantiates a new ItemInsightsApiSubjectStatsRequestBuilder and sets the default values.
func NewItemInsightsApiSubjectStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSubjectStatsRequestBuilder) {
    m := &ItemInsightsApiSubjectStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/subject-stats?min_timestamp={min_timestamp}{&direction*,max_timestamp*,page*,per_page*,sort*,subject_name_substring*}", pathParameters),
    }
    return m
}
// NewItemInsightsApiSubjectStatsRequestBuilder instantiates a new ItemInsightsApiSubjectStatsRequestBuilder and sets the default values.
func NewItemInsightsApiSubjectStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSubjectStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiSubjectStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get API request statistics for all subjects within an organization within a specified time frame. Subjects can be users or GitHub Apps.
// returns a []ApiInsightsSubjectStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/api-insights#get-subject-stats
func (m *ItemInsightsApiSubjectStatsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsSubjectStatsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateApiInsightsSubjectStatsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsSubjectStatsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsSubjectStatsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation get API request statistics for all subjects within an organization within a specified time frame. Subjects can be users or GitHub Apps.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiSubjectStatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiSubjectStatsRequestBuilder when successful
func (m *ItemInsightsApiSubjectStatsRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiSubjectStatsRequestBuilder) {
    return NewItemInsightsApiSubjectStatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
