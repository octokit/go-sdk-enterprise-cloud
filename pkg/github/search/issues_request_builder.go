package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i3e2edbd634cff9e942e443075aff88f6e2831f06d459d0fdfd91aac38eef8bea "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/search/issues"
)

// IssuesRequestBuilder builds and executes requests for operations under \search\issues
type IssuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IssuesRequestBuilderGetQueryParameters find issues by state and keyword. This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for issues, you can get text match metadata for the issue **title**, issue **body**, and issue **comment body** fields when you pass the `text-match` media type. For more details about how to receive highlightedsearch results, see [Text match metadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to find the oldest unresolved Python bugs on Windows. Your query might look something like this.`q=windows+label:bug+language:python+state:open&sort=created&order=asc`This query searches for the keyword `windows`, within any open issue that is labeled as `bug`. The search runs across repositories whose primary language is Python. The results are sorted by creation date in ascending order, which means the oldest issues appear first in the search results.> [!NOTE]> For requests made by GitHub Apps with a user access token, you can't retrieve a combination of issues and pull requests in a single query. Requests that don't include the `is:issue` or `is:pull-request` qualifier will receive an HTTP `422 Unprocessable Entity` response. To get results for both issues and pull requests, you must send separate queries for issues and pull requests. For more information about the `is` qualifier, see "[Searching only issues or pull requests](https://docs.github.com/enterprise-cloud@latest//github/searching-for-information-on-github/searching-issues-and-pull-requests#search-only-issues-or-pull-requests)."
type IssuesRequestBuilderGetQueryParameters struct {
    // Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`.
    Order *i3e2edbd634cff9e942e443075aff88f6e2831f06d459d0fdfd91aac38eef8bea.GetOrderQueryParameterType `uriparametername:"order"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub Enterprise Cloud. The REST API supports the same qualifiers as the web interface for GitHub Enterprise Cloud. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/enterprise-cloud@latest//rest/search/search#constructing-a-search-query). See "[Searching issues and pull requests](https://docs.github.com/enterprise-cloud@latest//search-github/searching-on-github/searching-issues-and-pull-requests)" for a detailed list of qualifiers.
    Q *string `uriparametername:"q"`
    // Sorts the results of your query by the number of `comments`, `reactions`, `reactions-+1`, `reactions--1`, `reactions-smile`, `reactions-thinking_face`, `reactions-heart`, `reactions-tada`, or `interactions`. You can also sort results by how recently the items were `created` or `updated`, Default: [best match](https://docs.github.com/enterprise-cloud@latest//rest/search/search#ranking-search-results)
    Sort *i3e2edbd634cff9e942e443075aff88f6e2831f06d459d0fdfd91aac38eef8bea.GetSortQueryParameterType `uriparametername:"sort"`
}
// NewIssuesRequestBuilderInternal instantiates a new IssuesRequestBuilder and sets the default values.
func NewIssuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IssuesRequestBuilder) {
    m := &IssuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search/issues?q={q}{&order*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewIssuesRequestBuilder instantiates a new IssuesRequestBuilder and sets the default values.
func NewIssuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IssuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIssuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get find issues by state and keyword. This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for issues, you can get text match metadata for the issue **title**, issue **body**, and issue **comment body** fields when you pass the `text-match` media type. For more details about how to receive highlightedsearch results, see [Text match metadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to find the oldest unresolved Python bugs on Windows. Your query might look something like this.`q=windows+label:bug+language:python+state:open&sort=created&order=asc`This query searches for the keyword `windows`, within any open issue that is labeled as `bug`. The search runs across repositories whose primary language is Python. The results are sorted by creation date in ascending order, which means the oldest issues appear first in the search results.> [!NOTE]> For requests made by GitHub Apps with a user access token, you can't retrieve a combination of issues and pull requests in a single query. Requests that don't include the `is:issue` or `is:pull-request` qualifier will receive an HTTP `422 Unprocessable Entity` response. To get results for both issues and pull requests, you must send separate queries for issues and pull requests. For more information about the `is` qualifier, see "[Searching only issues or pull requests](https://docs.github.com/enterprise-cloud@latest//github/searching-for-information-on-github/searching-issues-and-pull-requests#search-only-issues-or-pull-requests)."
// returns a IssuesGetResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a Issues503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/search/search#search-issues-and-pull-requests
func (m *IssuesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[IssuesRequestBuilderGetQueryParameters])(IssuesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
        "503": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateIssues503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIssuesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IssuesGetResponseable), nil
}
// ToGetRequestInformation find issues by state and keyword. This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for issues, you can get text match metadata for the issue **title**, issue **body**, and issue **comment body** fields when you pass the `text-match` media type. For more details about how to receive highlightedsearch results, see [Text match metadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to find the oldest unresolved Python bugs on Windows. Your query might look something like this.`q=windows+label:bug+language:python+state:open&sort=created&order=asc`This query searches for the keyword `windows`, within any open issue that is labeled as `bug`. The search runs across repositories whose primary language is Python. The results are sorted by creation date in ascending order, which means the oldest issues appear first in the search results.> [!NOTE]> For requests made by GitHub Apps with a user access token, you can't retrieve a combination of issues and pull requests in a single query. Requests that don't include the `is:issue` or `is:pull-request` qualifier will receive an HTTP `422 Unprocessable Entity` response. To get results for both issues and pull requests, you must send separate queries for issues and pull requests. For more information about the `is` qualifier, see "[Searching only issues or pull requests](https://docs.github.com/enterprise-cloud@latest//github/searching-for-information-on-github/searching-issues-and-pull-requests#search-only-issues-or-pull-requests)."
// returns a *RequestInformation when successful
func (m *IssuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[IssuesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IssuesRequestBuilder when successful
func (m *IssuesRequestBuilder) WithUrl(rawUrl string)(*IssuesRequestBuilder) {
    return NewIssuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
