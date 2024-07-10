package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0e912ee9b5d0ddd25cb2c0551e5151ad2ba9e55b6635357ebf3961fc887df2bb "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/search/commits"
)

// CommitsRequestBuilder builds and executes requests for operations under \search\commits
type CommitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CommitsRequestBuilderGetQueryParameters find commits via various criteria on the default branch (usually `main`). This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text matchmetadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:`q=repo:octocat/Spoon-Knife+css`
type CommitsRequestBuilderGetQueryParameters struct {
    // Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`.
    Order *i0e912ee9b5d0ddd25cb2c0551e5151ad2ba9e55b6635357ebf3961fc887df2bb.GetOrderQueryParameterType `uriparametername:"order"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub Enterprise Cloud. The REST API supports the same qualifiers as the web interface for GitHub Enterprise Cloud. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/enterprise-cloud@latest//rest/search/search#constructing-a-search-query). See "[Searching commits](https://docs.github.com/enterprise-cloud@latest//search-github/searching-on-github/searching-commits)" for a detailed list of qualifiers.
    Q *string `uriparametername:"q"`
    // Sorts the results of your query by `author-date` or `committer-date`. Default: [best match](https://docs.github.com/enterprise-cloud@latest//rest/search/search#ranking-search-results)
    Sort *i0e912ee9b5d0ddd25cb2c0551e5151ad2ba9e55b6635357ebf3961fc887df2bb.GetSortQueryParameterType `uriparametername:"sort"`
}
// NewCommitsRequestBuilderInternal instantiates a new CommitsRequestBuilder and sets the default values.
func NewCommitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommitsRequestBuilder) {
    m := &CommitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search/commits?q={q}{&order*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewCommitsRequestBuilder instantiates a new CommitsRequestBuilder and sets the default values.
func NewCommitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get find commits via various criteria on the default branch (usually `main`). This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text matchmetadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:`q=repo:octocat/Spoon-Knife+css`
// returns a CommitsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/search/search#search-commits
func (m *CommitsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[CommitsRequestBuilderGetQueryParameters])(CommitsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCommitsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CommitsGetResponseable), nil
}
// ToGetRequestInformation find commits via various criteria on the default branch (usually `main`). This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text matchmetadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:`q=repo:octocat/Spoon-Knife+css`
// returns a *RequestInformation when successful
func (m *CommitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[CommitsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CommitsRequestBuilder when successful
func (m *CommitsRequestBuilder) WithUrl(rawUrl string)(*CommitsRequestBuilder) {
    return NewCommitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
