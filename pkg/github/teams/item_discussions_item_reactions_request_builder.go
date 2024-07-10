package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    ib31ac189171b0fe4b57cce81e7455929005bc8f2b687cc9e203825feb6b2644b "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/teams/item/discussions/item/reactions"
)

// ItemDiscussionsItemReactionsRequestBuilder builds and executes requests for operations under \teams\{team_id}\discussions\{discussion_number}\reactions
type ItemDiscussionsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDiscussionsItemReactionsRequestBuilderGetQueryParameters **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion`](https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#list-reactions-for-a-team-discussion) endpoint.List the reactions to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
type ItemDiscussionsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion.
    Content *ib31ac189171b0fe4b57cce81e7455929005bc8f2b687cc9e203825feb6b2644b.GetContentQueryParameterType `uriparametername:"content"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemDiscussionsItemReactionsRequestBuilderInternal instantiates a new ItemDiscussionsItemReactionsRequestBuilder and sets the default values.
func NewItemDiscussionsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemReactionsRequestBuilder) {
    m := &ItemDiscussionsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/discussions/{discussion_number}/reactions{?content*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemDiscussionsItemReactionsRequestBuilder instantiates a new ItemDiscussionsItemReactionsRequestBuilder and sets the default values.
func NewItemDiscussionsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDiscussionsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion`](https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#list-reactions-for-a-team-discussion) endpoint.List the reactions to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a []Reactionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#list-reactions-for-a-team-discussion-legacy
func (m *ItemDiscussionsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsItemReactionsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Reactionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateReactionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Reactionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Reactionable)
        }
    }
    return val, nil
}
// Post **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create reaction for a team discussion`](https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#create-reaction-for-a-team-discussion) endpoint.Create a reaction to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).A response with an HTTP `200` status means that you already added the reaction type to this team discussion.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a Reactionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#create-reaction-for-a-team-discussion-legacy
func (m *ItemDiscussionsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemDiscussionsItemReactionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Reactionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateReactionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Reactionable), nil
}
// ToGetRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion`](https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#list-reactions-for-a-team-discussion) endpoint.List the reactions to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsItemReactionsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create reaction for a team discussion`](https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#create-reaction-for-a-team-discussion) endpoint.Create a reaction to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).A response with an HTTP `200` status means that you already added the reaction type to this team discussion.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDiscussionsItemReactionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemDiscussionsItemReactionsRequestBuilder when successful
func (m *ItemDiscussionsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemDiscussionsItemReactionsRequestBuilder) {
    return NewItemDiscussionsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
