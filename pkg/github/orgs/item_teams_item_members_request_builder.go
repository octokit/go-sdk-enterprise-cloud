package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i6e223be6fe8a5e809c1f58c35fde308f3d99429fe7b30b7670732fa95bae78a6 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/teams/item/members"
)

// ItemTeamsItemMembersRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\members
type ItemTeamsItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamsItemMembersRequestBuilderGetQueryParameters team members will include the members of child teams.To list members in a team, the team must be visible to the authenticated user.
type ItemTeamsItemMembersRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Filters members returned by their role in the team.
    Role *i6e223be6fe8a5e809c1f58c35fde308f3d99429fe7b30b7670732fa95bae78a6.GetRoleQueryParameterType `uriparametername:"role"`
}
// NewItemTeamsItemMembersRequestBuilderInternal instantiates a new ItemTeamsItemMembersRequestBuilder and sets the default values.
func NewItemTeamsItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemMembersRequestBuilder) {
    m := &ItemTeamsItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/members{?page*,per_page*,role*}", pathParameters),
    }
    return m
}
// NewItemTeamsItemMembersRequestBuilder instantiates a new ItemTeamsItemMembersRequestBuilder and sets the default values.
func NewItemTeamsItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get team members will include the members of child teams.To list members in a team, the team must be visible to the authenticated user.
// returns a []SimpleUserable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/teams/members#list-team-members
func (m *ItemTeamsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamsItemMembersRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateSimpleUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SimpleUserable)
        }
    }
    return val, nil
}
// ToGetRequestInformation team members will include the members of child teams.To list members in a team, the team must be visible to the authenticated user.
// returns a *RequestInformation when successful
func (m *ItemTeamsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamsItemMembersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamsItemMembersRequestBuilder when successful
func (m *ItemTeamsItemMembersRequestBuilder) WithUrl(rawUrl string)(*ItemTeamsItemMembersRequestBuilder) {
    return NewItemTeamsItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
