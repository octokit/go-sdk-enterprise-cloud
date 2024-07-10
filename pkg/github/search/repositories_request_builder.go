package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i890c9716ce04aa7efa4609d56b5b916d368eb72e245831b067173d022cefbff5 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/search/repositories"
)

// RepositoriesRequestBuilder builds and executes requests for operations under \search\repositories
type RepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RepositoriesRequestBuilderGetQueryParameters find repositories via various criteria. This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for repositories, you can get text match metadata for the **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to search for popular Tetris repositories written in assembly code, your query might look like this:`q=tetris+language:assembly&sort=stars&order=desc`This query searches for repositories with the word `tetris` in the name, the description, or the README. The results are limited to repositories where the primary language is assembly. The results are sorted by stars in descending order, so that the most popular repositories appear first in the search results.
type RepositoriesRequestBuilderGetQueryParameters struct {
    // Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`.
    Order *i890c9716ce04aa7efa4609d56b5b916d368eb72e245831b067173d022cefbff5.GetOrderQueryParameterType `uriparametername:"order"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub Enterprise Cloud. The REST API supports the same qualifiers as the web interface for GitHub Enterprise Cloud. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/enterprise-cloud@latest//rest/search/search#constructing-a-search-query). See "[Searching for repositories](https://docs.github.com/enterprise-cloud@latest//articles/searching-for-repositories/)" for a detailed list of qualifiers.
    Q *string `uriparametername:"q"`
    // Sorts the results of your query by number of `stars`, `forks`, or `help-wanted-issues` or how recently the items were `updated`. Default: [best match](https://docs.github.com/enterprise-cloud@latest//rest/search/search#ranking-search-results)
    Sort *i890c9716ce04aa7efa4609d56b5b916d368eb72e245831b067173d022cefbff5.GetSortQueryParameterType `uriparametername:"sort"`
}
// NewRepositoriesRequestBuilderInternal instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RepositoriesRequestBuilder) {
    m := &RepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search/repositories?q={q}{&order*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewRepositoriesRequestBuilder instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get find repositories via various criteria. This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for repositories, you can get text match metadata for the **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to search for popular Tetris repositories written in assembly code, your query might look like this:`q=tetris+language:assembly&sort=stars&order=desc`This query searches for repositories with the word `tetris` in the name, the description, or the README. The results are limited to repositories where the primary language is assembly. The results are sorted by stars in descending order, so that the most popular repositories appear first in the search results.
// returns a RepositoriesGetResponseable when successful
// returns a ValidationError error when the service returns a 422 status code
// returns a Repositories503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/search/search#search-repositories
func (m *RepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[RepositoriesRequestBuilderGetQueryParameters])(RepositoriesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
        "503": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateRepositories503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRepositoriesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RepositoriesGetResponseable), nil
}
// ToGetRequestInformation find repositories via various criteria. This method returns up to 100 results [per page](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api).When searching for repositories, you can get text match metadata for the **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://docs.github.com/enterprise-cloud@latest//rest/search/search#text-match-metadata).For example, if you want to search for popular Tetris repositories written in assembly code, your query might look like this:`q=tetris+language:assembly&sort=stars&order=desc`This query searches for repositories with the word `tetris` in the name, the description, or the README. The results are limited to repositories where the primary language is assembly. The results are sorted by stars in descending order, so that the most popular repositories appear first in the search results.
// returns a *RequestInformation when successful
func (m *RepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[RepositoriesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RepositoriesRequestBuilder when successful
func (m *RepositoriesRequestBuilder) WithUrl(rawUrl string)(*RepositoriesRequestBuilder) {
    return NewRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
