package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemFine_grained_permissionsRequestBuilder builds and executes requests for operations under \orgs\{org}\fine_grained_permissions
type ItemFine_grained_permissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemFine_grained_permissionsRequestBuilderInternal instantiates a new ItemFine_grained_permissionsRequestBuilder and sets the default values.
func NewItemFine_grained_permissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFine_grained_permissionsRequestBuilder) {
    m := &ItemFine_grained_permissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/fine_grained_permissions", pathParameters),
    }
    return m
}
// NewItemFine_grained_permissionsRequestBuilder instantiates a new ItemFine_grained_permissionsRequestBuilder and sets the default values.
func NewItemFine_grained_permissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFine_grained_permissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFine_grained_permissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Closing down notice:** This operation is closing down and will be removed after September 6, 2023. Use the "[List fine-grained repository permissions](https://docs.github.com/enterprise-cloud@latest//rest/orgs/custom-roles#list-repository-fine-grained-permissions-for-an-organization)" endpoint instead.Lists the fine-grained permissions that can be used in custom repository roles for an organization. For more information, see "[About custom repository roles](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-repository-roles)."To use this endpoint the authenticated user must be an administrator of the organization or of a repository of the organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// Deprecated: 
// returns a []RepositoryFineGrainedPermissionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/custom-roles#closing-down---list-fine-grained-permissions-for-an-organization
func (m *ItemFine_grained_permissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RepositoryFineGrainedPermissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateRepositoryFineGrainedPermissionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RepositoryFineGrainedPermissionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RepositoryFineGrainedPermissionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation > [!WARNING]> **Closing down notice:** This operation is closing down and will be removed after September 6, 2023. Use the "[List fine-grained repository permissions](https://docs.github.com/enterprise-cloud@latest//rest/orgs/custom-roles#list-repository-fine-grained-permissions-for-an-organization)" endpoint instead.Lists the fine-grained permissions that can be used in custom repository roles for an organization. For more information, see "[About custom repository roles](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-repository-roles)."To use this endpoint the authenticated user must be an administrator of the organization or of a repository of the organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemFine_grained_permissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemFine_grained_permissionsRequestBuilder when successful
func (m *ItemFine_grained_permissionsRequestBuilder) WithUrl(rawUrl string)(*ItemFine_grained_permissionsRequestBuilder) {
    return NewItemFine_grained_permissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
