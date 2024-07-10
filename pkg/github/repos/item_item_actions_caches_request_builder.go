package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i496bb68fcc18ee3e149c324e776371af2675eaa1506db2cbc8692b1a7d86e871 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos/item/item/actions/caches"
)

// ItemItemActionsCachesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\caches
type ItemItemActionsCachesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsCachesRequestBuilderDeleteQueryParameters deletes one or more GitHub Actions caches for a repository, using a complete cache key. By default, all caches that match the provided key are deleted, but you can optionally provide a Git ref to restrict deletions to caches that match both the provided key and the Git ref.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
type ItemItemActionsCachesRequestBuilderDeleteQueryParameters struct {
    // A key for identifying the cache.
    Key *string `uriparametername:"key"`
    // The full Git reference for narrowing down the cache. The `ref` for a branch should be formatted as `refs/heads/<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
}
// ItemItemActionsCachesRequestBuilderGetQueryParameters lists the GitHub Actions caches for a repository.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
type ItemItemActionsCachesRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i496bb68fcc18ee3e149c324e776371af2675eaa1506db2cbc8692b1a7d86e871.GetDirectionQueryParameterType `uriparametername:"direction"`
    // An explicit key or prefix for identifying the cache
    Key *string `uriparametername:"key"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The full Git reference for narrowing down the cache. The `ref` for a branch should be formatted as `refs/heads/<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
    // The property to sort the results by. `created_at` means when the cache was created. `last_accessed_at` means when the cache was last accessed. `size_in_bytes` is the size of the cache in bytes.
    Sort *i496bb68fcc18ee3e149c324e776371af2675eaa1506db2cbc8692b1a7d86e871.GetSortQueryParameterType `uriparametername:"sort"`
}
// ByCache_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.repos.item.item.actions.caches.item collection
// returns a *ItemItemActionsCachesWithCache_ItemRequestBuilder when successful
func (m *ItemItemActionsCachesRequestBuilder) ByCache_id(cache_id int32)(*ItemItemActionsCachesWithCache_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["cache_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(cache_id), 10)
    return NewItemItemActionsCachesWithCache_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsCachesRequestBuilderInternal instantiates a new ItemItemActionsCachesRequestBuilder and sets the default values.
func NewItemItemActionsCachesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCachesRequestBuilder) {
    m := &ItemItemActionsCachesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/caches{?direction*,key*,page*,per_page*,ref*,sort*}", pathParameters),
    }
    return m
}
// NewItemItemActionsCachesRequestBuilder instantiates a new ItemItemActionsCachesRequestBuilder and sets the default values.
func NewItemItemActionsCachesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCachesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsCachesRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes one or more GitHub Actions caches for a repository, using a complete cache key. By default, all caches that match the provided key are deleted, but you can optionally provide a Git ref to restrict deletions to caches that match both the provided key and the Git ref.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ActionsCacheListable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/cache#delete-github-actions-caches-for-a-repository-using-a-cache-key
func (m *ItemItemActionsCachesRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsCachesRequestBuilderDeleteQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsCacheListable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateActionsCacheListFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsCacheListable), nil
}
// Get lists the GitHub Actions caches for a repository.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ActionsCacheListable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/cache#list-github-actions-caches-for-a-repository
func (m *ItemItemActionsCachesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsCachesRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsCacheListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateActionsCacheListFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsCacheListable), nil
}
// ToDeleteRequestInformation deletes one or more GitHub Actions caches for a repository, using a complete cache key. By default, all caches that match the provided key are deleted, but you can optionally provide a Git ref to restrict deletions to caches that match both the provided key and the Git ref.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsCachesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsCachesRequestBuilderDeleteQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/caches?key={key}{&ref*}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation lists the GitHub Actions caches for a repository.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsCachesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsCachesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsCachesRequestBuilder when successful
func (m *ItemItemActionsCachesRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsCachesRequestBuilder) {
    return NewItemItemActionsCachesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
