package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemTeamsItemExternalGroupsRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\external-groups
type ItemTeamsItemExternalGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemTeamsItemExternalGroupsRequestBuilderInternal instantiates a new ItemTeamsItemExternalGroupsRequestBuilder and sets the default values.
func NewItemTeamsItemExternalGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemExternalGroupsRequestBuilder) {
    m := &ItemTeamsItemExternalGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/external-groups", pathParameters),
    }
    return m
}
// NewItemTeamsItemExternalGroupsRequestBuilder instantiates a new ItemTeamsItemExternalGroupsRequestBuilder and sets the default values.
func NewItemTeamsItemExternalGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemExternalGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemExternalGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a connection between a team and an external group.You can manage team membership with your IdP using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see [GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/teams/external-groups#remove-the-connection-between-an-external-group-and-a-team
func (m *ItemTeamsItemExternalGroupsRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get lists a connection between a team and an external group.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a ExternalGroupsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/teams/external-groups#list-a-connection-between-an-external-group-and-a-team
func (m *ItemTeamsItemExternalGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ExternalGroupsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateExternalGroupsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ExternalGroupsable), nil
}
// Patch creates a connection between a team and an external group.  Only one external group can be linked to a team.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a ExternalGroupable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/teams/external-groups#update-the-connection-between-an-external-group-and-a-team
func (m *ItemTeamsItemExternalGroupsRequestBuilder) Patch(ctx context.Context, body ItemTeamsItemExternalGroupsPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ExternalGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateExternalGroupFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ExternalGroupable), nil
}
// ToDeleteRequestInformation deletes a connection between a team and an external group.You can manage team membership with your IdP using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see [GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemTeamsItemExternalGroupsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation lists a connection between a team and an external group.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemTeamsItemExternalGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation creates a connection between a team and an external group.  Only one external group can be linked to a team.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemTeamsItemExternalGroupsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemTeamsItemExternalGroupsPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamsItemExternalGroupsRequestBuilder when successful
func (m *ItemTeamsItemExternalGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamsItemExternalGroupsRequestBuilder) {
    return NewItemTeamsItemExternalGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
