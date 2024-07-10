package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemActionsPermissionsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\actions\permissions
type ItemActionsPermissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsPermissionsRequestBuilderInternal instantiates a new ItemActionsPermissionsRequestBuilder and sets the default values.
func NewItemActionsPermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsPermissionsRequestBuilder) {
    m := &ItemActionsPermissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/actions/permissions", pathParameters),
    }
    return m
}
// NewItemActionsPermissionsRequestBuilder instantiates a new ItemActionsPermissionsRequestBuilder and sets the default values.
func NewItemActionsPermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsPermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsPermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the GitHub Actions permissions policy for organizations and allowed actions and reusable workflows in an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a ActionsEnterprisePermissionsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/permissions#get-github-actions-permissions-for-an-enterprise
func (m *ItemActionsPermissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsEnterprisePermissionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateActionsEnterprisePermissionsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsEnterprisePermissionsable), nil
}
// Organizations the organizations property
// returns a *ItemActionsPermissionsOrganizationsRequestBuilder when successful
func (m *ItemActionsPermissionsRequestBuilder) Organizations()(*ItemActionsPermissionsOrganizationsRequestBuilder) {
    return NewItemActionsPermissionsOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put sets the GitHub Actions permissions policy for organizations and allowed actions and reusable workflows in an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/permissions#set-github-actions-permissions-for-an-enterprise
func (m *ItemActionsPermissionsRequestBuilder) Put(ctx context.Context, body ItemActionsPermissionsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// SelectedActions the selectedActions property
// returns a *ItemActionsPermissionsSelectedActionsRequestBuilder when successful
func (m *ItemActionsPermissionsRequestBuilder) SelectedActions()(*ItemActionsPermissionsSelectedActionsRequestBuilder) {
    return NewItemActionsPermissionsSelectedActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets the GitHub Actions permissions policy for organizations and allowed actions and reusable workflows in an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsPermissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation sets the GitHub Actions permissions policy for organizations and allowed actions and reusable workflows in an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsPermissionsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemActionsPermissionsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsPermissionsRequestBuilder when successful
func (m *ItemActionsPermissionsRequestBuilder) WithUrl(rawUrl string)(*ItemActionsPermissionsRequestBuilder) {
    return NewItemActionsPermissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workflow the workflow property
// returns a *ItemActionsPermissionsWorkflowRequestBuilder when successful
func (m *ItemActionsPermissionsRequestBuilder) Workflow()(*ItemActionsPermissionsWorkflowRequestBuilder) {
    return NewItemActionsPermissionsWorkflowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
