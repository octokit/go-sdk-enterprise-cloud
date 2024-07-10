package enterprises

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnerGroupsItemOrganizationsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\actions\runner-groups\{runner_group_id}\organizations
type ItemActionsRunnerGroupsItemOrganizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnerGroupsItemOrganizationsRequestBuilderGetQueryParameters lists the organizations with access to a self-hosted runner group.OAuth app tokens and personal access tokens (classic) need the `manage_runners:enterprise` scope to use this endpoint.
type ItemActionsRunnerGroupsItemOrganizationsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByOrg_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterprises.item.actions.runnerGroups.item.organizations.item collection
// returns a *ItemActionsRunnerGroupsItemOrganizationsWithOrg_ItemRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) ByOrg_id(org_id int32)(*ItemActionsRunnerGroupsItemOrganizationsWithOrg_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["org_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(org_id), 10)
    return NewItemActionsRunnerGroupsItemOrganizationsWithOrg_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsRunnerGroupsItemOrganizationsRequestBuilderInternal instantiates a new ItemActionsRunnerGroupsItemOrganizationsRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemOrganizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) {
    m := &ItemActionsRunnerGroupsItemOrganizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsRunnerGroupsItemOrganizationsRequestBuilder instantiates a new ItemActionsRunnerGroupsItemOrganizationsRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemOrganizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnerGroupsItemOrganizationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the organizations with access to a self-hosted runner group.OAuth app tokens and personal access tokens (classic) need the `manage_runners:enterprise` scope to use this endpoint.
// returns a ItemActionsRunnerGroupsItemOrganizationsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/self-hosted-runner-groups#list-organization-access-to-a-self-hosted-runner-group-in-an-enterprise
func (m *ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemOrganizationsRequestBuilderGetQueryParameters])(ItemActionsRunnerGroupsItemOrganizationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnerGroupsItemOrganizationsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnerGroupsItemOrganizationsGetResponseable), nil
}
// Put replaces the list of organizations that have access to a self-hosted runner configured in an enterprise.OAuth app tokens and personal access tokens (classic) need the `manage_runners:enterprise` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/self-hosted-runner-groups#set-organization-access-for-a-self-hosted-runner-group-in-an-enterprise
func (m *ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) Put(ctx context.Context, body ItemActionsRunnerGroupsItemOrganizationsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToGetRequestInformation lists the organizations with access to a self-hosted runner group.OAuth app tokens and personal access tokens (classic) need the `manage_runners:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemOrganizationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation replaces the list of organizations that have access to a self-hosted runner configured in an enterprise.OAuth app tokens and personal access tokens (classic) need the `manage_runners:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemActionsRunnerGroupsItemOrganizationsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnerGroupsItemOrganizationsRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnerGroupsItemOrganizationsRequestBuilder) {
    return NewItemActionsRunnerGroupsItemOrganizationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
