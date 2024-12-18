package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    ie961c9b74390e268ab48ff84f76179ac01ff30f851de45cdc9b8a186e7198301 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/teams/item/discussions/item/reactions"
)

// ItemTeamsItemDiscussionsItemReactionsRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\discussions\{discussion_number}\reactions
type ItemTeamsItemDiscussionsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetQueryParameters list the reactions to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).> [!NOTE]> You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
type ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion.
    Content *ie961c9b74390e268ab48ff84f76179ac01ff30f851de45cdc9b8a186e7198301.GetContentQueryParameterType `uriparametername:"content"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByReaction_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.orgs.item.teams.item.discussions.item.reactions.item collection
// returns a *ItemTeamsItemDiscussionsItemReactionsWithReaction_ItemRequestBuilder when successful
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) ByReaction_id(reaction_id int32)(*ItemTeamsItemDiscussionsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["reaction_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(reaction_id), 10)
    return NewItemTeamsItemDiscussionsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamsItemDiscussionsItemReactionsRequestBuilderInternal instantiates a new ItemTeamsItemDiscussionsItemReactionsRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemReactionsRequestBuilder) {
    m := &ItemTeamsItemDiscussionsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions{?content*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemTeamsItemDiscussionsItemReactionsRequestBuilder instantiates a new ItemTeamsItemDiscussionsItemReactionsRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemDiscussionsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the reactions to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).> [!NOTE]> You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// returns a []Reactionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#list-reactions-for-a-team-discussion
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Reactionable, error) {
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
// Post create a reaction to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).A response with an HTTP `200` status means that you already added the reaction type to this team discussion.> [!NOTE]> You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// returns a Reactionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/reactions/reactions#create-reaction-for-a-team-discussion
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemTeamsItemDiscussionsItemReactionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Reactionable, error) {
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
// ToGetRequestInformation list the reactions to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).> [!NOTE]> You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamsItemDiscussionsItemReactionsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a reaction to a [team discussion](https://docs.github.com/enterprise-cloud@latest//rest/teams/discussions#get-a-discussion).A response with an HTTP `200` status means that you already added the reaction type to this team discussion.> [!NOTE]> You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamsItemDiscussionsItemReactionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamsItemDiscussionsItemReactionsRequestBuilder when successful
func (m *ItemTeamsItemDiscussionsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamsItemDiscussionsItemReactionsRequestBuilder) {
    return NewItemTeamsItemDiscussionsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
