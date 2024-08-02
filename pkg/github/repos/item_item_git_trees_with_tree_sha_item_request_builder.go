package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemGitTreesWithTree_shaItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\git\trees\{tree_sha}
type ItemItemGitTreesWithTree_shaItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemGitTreesWithTree_shaItemRequestBuilderGetQueryParameters returns a single tree using the SHA1 value or ref name for that tree.If `truncated` is `true` in the response then the number of items in the `tree` array exceeded our maximum limit. If you need to fetch more items, use the non-recursive method of fetching trees, and fetch one sub-tree at a time.> [!NOTE]> The limit for the `tree` array is 100,000 entries with a maximum size of 7 MB when using the `recursive` parameter.
type ItemItemGitTreesWithTree_shaItemRequestBuilderGetQueryParameters struct {
    // Setting this parameter to any value returns the objects or subtrees referenced by the tree specified in `:tree_sha`. For example, setting `recursive` to any of the following will enable returning objects or subtrees: `0`, `1`, `"true"`, and `"false"`. Omit this parameter to prevent recursively returning objects or subtrees.
    Recursive *string `uriparametername:"recursive"`
}
// NewItemItemGitTreesWithTree_shaItemRequestBuilderInternal instantiates a new ItemItemGitTreesWithTree_shaItemRequestBuilder and sets the default values.
func NewItemItemGitTreesWithTree_shaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitTreesWithTree_shaItemRequestBuilder) {
    m := &ItemItemGitTreesWithTree_shaItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/git/trees/{tree_sha}{?recursive*}", pathParameters),
    }
    return m
}
// NewItemItemGitTreesWithTree_shaItemRequestBuilder instantiates a new ItemItemGitTreesWithTree_shaItemRequestBuilder and sets the default values.
func NewItemItemGitTreesWithTree_shaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitTreesWithTree_shaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitTreesWithTree_shaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a single tree using the SHA1 value or ref name for that tree.If `truncated` is `true` in the response then the number of items in the `tree` array exceeded our maximum limit. If you need to fetch more items, use the non-recursive method of fetching trees, and fetch one sub-tree at a time.> [!NOTE]> The limit for the `tree` array is 100,000 entries with a maximum size of 7 MB when using the `recursive` parameter.
// returns a GitTreeable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 409 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/git/trees#get-a-tree
func (m *ItemItemGitTreesWithTree_shaItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemGitTreesWithTree_shaItemRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GitTreeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "409": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateGitTreeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GitTreeable), nil
}
// ToGetRequestInformation returns a single tree using the SHA1 value or ref name for that tree.If `truncated` is `true` in the response then the number of items in the `tree` array exceeded our maximum limit. If you need to fetch more items, use the non-recursive method of fetching trees, and fetch one sub-tree at a time.> [!NOTE]> The limit for the `tree` array is 100,000 entries with a maximum size of 7 MB when using the `recursive` parameter.
// returns a *RequestInformation when successful
func (m *ItemItemGitTreesWithTree_shaItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemGitTreesWithTree_shaItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemGitTreesWithTree_shaItemRequestBuilder when successful
func (m *ItemItemGitTreesWithTree_shaItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemGitTreesWithTree_shaItemRequestBuilder) {
    return NewItemItemGitTreesWithTree_shaItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
