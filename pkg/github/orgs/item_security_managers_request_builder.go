package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemSecurityManagersRequestBuilder builds and executes requests for operations under \orgs\{org}\security-managers
type ItemSecurityManagersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSecurityManagersRequestBuilderInternal instantiates a new ItemSecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    m := &ItemSecurityManagersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/security-managers", pathParameters),
    }
    return m
}
// NewItemSecurityManagersRequestBuilder instantiates a new ItemSecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityManagersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Closing down notice:** This operation is closing down and will be removed starting January 1, 2026. Please use the "[Organization Roles](https://docs.github.com/enterprise-cloud@latest//rest/orgs/organization-roles)" endpoints instead.
// Deprecated: 
// returns a []TeamSimpleable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/security-managers#list-security-manager-teams
func (m *ItemSecurityManagersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.TeamSimpleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateTeamSimpleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.TeamSimpleable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.TeamSimpleable)
        }
    }
    return val, nil
}
// Teams the teams property
// returns a *ItemSecurityManagersTeamsRequestBuilder when successful
func (m *ItemSecurityManagersRequestBuilder) Teams()(*ItemSecurityManagersTeamsRequestBuilder) {
    return NewItemSecurityManagersTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation > [!WARNING]> **Closing down notice:** This operation is closing down and will be removed starting January 1, 2026. Please use the "[Organization Roles](https://docs.github.com/enterprise-cloud@latest//rest/orgs/organization-roles)" endpoints instead.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemSecurityManagersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemSecurityManagersRequestBuilder when successful
func (m *ItemSecurityManagersRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityManagersRequestBuilder) {
    return NewItemSecurityManagersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
