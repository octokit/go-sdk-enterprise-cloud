package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemPagesBuildsWithBuild_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\pages\builds\{build_id}
type ItemItemPagesBuildsWithBuild_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal instantiates a new ItemItemPagesBuildsWithBuild_ItemRequestBuilder and sets the default values.
func NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    m := &ItemItemPagesBuildsWithBuild_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/pages/builds/{build_id}", pathParameters),
    }
    return m
}
// NewItemItemPagesBuildsWithBuild_ItemRequestBuilder instantiates a new ItemItemPagesBuildsWithBuild_ItemRequestBuilder and sets the default values.
func NewItemItemPagesBuildsWithBuild_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about a GitHub Enterprise Cloud Pages build.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a PageBuildable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/pages/pages#get-apiname-pages-build
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PageBuildable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreatePageBuildFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PageBuildable), nil
}
// ToGetRequestInformation gets information about a GitHub Enterprise Cloud Pages build.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemPagesBuildsWithBuild_ItemRequestBuilder when successful
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    return NewItemItemPagesBuildsWithBuild_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
