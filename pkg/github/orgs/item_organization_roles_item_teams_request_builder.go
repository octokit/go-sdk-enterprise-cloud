package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemOrganizationRolesItemTeamsRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\{role_id}\teams
type ItemOrganizationRolesItemTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters lists the teams that are assigned to an organization role. For more information on organization roles, see "[Using organization roles](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/using-organization-roles)."To use this endpoint, you must be an administrator for the organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemOrganizationRolesItemTeamsRequestBuilderInternal instantiates a new ItemOrganizationRolesItemTeamsRequestBuilder and sets the default values.
func NewItemOrganizationRolesItemTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesItemTeamsRequestBuilder) {
    m := &ItemOrganizationRolesItemTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/{role_id}/teams{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesItemTeamsRequestBuilder instantiates a new ItemOrganizationRolesItemTeamsRequestBuilder and sets the default values.
func NewItemOrganizationRolesItemTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesItemTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesItemTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the teams that are assigned to an organization role. For more information on organization roles, see "[Using organization roles](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/using-organization-roles)."To use this endpoint, you must be an administrator for the organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a []TeamRoleAssignmentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/organization-roles#list-teams-that-are-assigned-to-an-organization-role
func (m *ItemOrganizationRolesItemTeamsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.TeamRoleAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateTeamRoleAssignmentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.TeamRoleAssignmentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.TeamRoleAssignmentable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the teams that are assigned to an organization role. For more information on organization roles, see "[Using organization roles](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/using-organization-roles)."To use this endpoint, you must be an administrator for the organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemOrganizationRolesItemTeamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOrganizationRolesItemTeamsRequestBuilder when successful
func (m *ItemOrganizationRolesItemTeamsRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationRolesItemTeamsRequestBuilder) {
    return NewItemOrganizationRolesItemTeamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
