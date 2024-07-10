package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemTeamSyncGroupsRequestBuilder builds and executes requests for operations under \orgs\{org}\team-sync\groups
type ItemTeamSyncGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamSyncGroupsRequestBuilderGetQueryParameters lists IdP groups available in an organization.
type ItemTeamSyncGroupsRequestBuilderGetQueryParameters struct {
    // Page token
    Page *string `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Filters the results to return only those that begin with the value specified by this parameter. For example, a value of `ab` will return results that begin with "ab".
    Q *string `uriparametername:"q"`
}
// NewItemTeamSyncGroupsRequestBuilderInternal instantiates a new ItemTeamSyncGroupsRequestBuilder and sets the default values.
func NewItemTeamSyncGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamSyncGroupsRequestBuilder) {
    m := &ItemTeamSyncGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/team-sync/groups{?page*,per_page*,q*}", pathParameters),
    }
    return m
}
// NewItemTeamSyncGroupsRequestBuilder instantiates a new ItemTeamSyncGroupsRequestBuilder and sets the default values.
func NewItemTeamSyncGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamSyncGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamSyncGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists IdP groups available in an organization.
// returns a GroupMappingable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/teams/team-sync#list-idp-groups-for-an-organization
func (m *ItemTeamSyncGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamSyncGroupsRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GroupMappingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateGroupMappingFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GroupMappingable), nil
}
// ToGetRequestInformation lists IdP groups available in an organization.
// returns a *RequestInformation when successful
func (m *ItemTeamSyncGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamSyncGroupsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamSyncGroupsRequestBuilder when successful
func (m *ItemTeamSyncGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamSyncGroupsRequestBuilder) {
    return NewItemTeamSyncGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
