package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemActionsRunnersRemoveTokenRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runners\remove-token
type ItemItemActionsRunnersRemoveTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal instantiates a new ItemItemActionsRunnersRemoveTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    m := &ItemItemActionsRunnersRemoveTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runners/remove-token", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersRemoveTokenRequestBuilder instantiates a new ItemItemActionsRunnersRemoveTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRemoveTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script to remove a self-hosted runner from an repository. The token expires after one hour.For example, you can replace `TOKEN` in the following example with the registration token provided by this endpoint to remove your self-hosted runner from an organization:```./config.sh remove --token TOKEN```Authenticated users must have admin access to the repository to use this endpoint.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a AuthenticationTokenable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/self-hosted-runners#create-a-remove-token-for-a-repository
func (m *ItemItemActionsRunnersRemoveTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AuthenticationTokenable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateAuthenticationTokenFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AuthenticationTokenable), nil
}
// ToPostRequestInformation returns a token that you can pass to the `config` script to remove a self-hosted runner from an repository. The token expires after one hour.For example, you can replace `TOKEN` in the following example with the registration token provided by this endpoint to remove your self-hosted runner from an organization:```./config.sh remove --token TOKEN```Authenticated users must have admin access to the repository to use this endpoint.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunnersRemoveTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunnersRemoveTokenRequestBuilder when successful
func (m *ItemItemActionsRunnersRemoveTokenRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemItemActionsRunnersRemoveTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
