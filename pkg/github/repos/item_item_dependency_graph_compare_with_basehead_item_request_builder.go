package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\dependency-graph\compare\{basehead}
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters struct {
    // The full path, relative to the repository root, of the dependency manifest file.
    Name *string `uriparametername:"name"`
}
// NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal instantiates a new ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    m := &ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/dependency-graph/compare/{basehead}{?name*}", pathParameters),
    }
    return m
}
// NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder instantiates a new ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
// returns a []DependencyGraphDiffable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/dependency-graph/dependency-review#get-a-diff-of-the-dependencies-between-commits
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DependencyGraphDiffable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateDependencyGraphDiffFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DependencyGraphDiffable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DependencyGraphDiffable)
        }
    }
    return val, nil
}
// ToGetRequestInformation gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
// returns a *RequestInformation when successful
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder when successful
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    return NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
