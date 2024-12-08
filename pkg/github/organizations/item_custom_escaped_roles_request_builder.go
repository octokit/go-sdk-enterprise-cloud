package organizations

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCustom_rolesRequestBuilder builds and executes requests for operations under \organizations\{organization_-id}\custom_roles
type ItemCustom_rolesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCustom_rolesRequestBuilderInternal instantiates a new ItemCustom_rolesRequestBuilder and sets the default values.
func NewItemCustom_rolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustom_rolesRequestBuilder) {
    m := &ItemCustom_rolesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organizations/{organization_%2Did}/custom_roles", pathParameters),
    }
    return m
}
// NewItemCustom_rolesRequestBuilder instantiates a new ItemCustom_rolesRequestBuilder and sets the default values.
func NewItemCustom_rolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustom_rolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCustom_rolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Closing down notice:** This operation is closing down and will be removed in the future. Use the "[List custom repository roles](https://docs.github.com/enterprise-cloud@latest//rest/orgs/custom-roles#list-custom-repository-roles-in-an-organization)" endpoint instead.List the custom repository roles available in this organization. For more information on custom repository roles, see "[About custom repository roles](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-repository-roles)."The authenticated user must be administrator of the organization or of a repository of the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// Deprecated: 
// returns a ItemCustom_rolesGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/custom-roles#closing-down---list-custom-repository-roles-in-an-organization
func (m *ItemCustom_rolesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemCustom_rolesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCustom_rolesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCustom_rolesGetResponseable), nil
}
// ToGetRequestInformation > [!WARNING]> **Closing down notice:** This operation is closing down and will be removed in the future. Use the "[List custom repository roles](https://docs.github.com/enterprise-cloud@latest//rest/orgs/custom-roles#list-custom-repository-roles-in-an-organization)" endpoint instead.List the custom repository roles available in this organization. For more information on custom repository roles, see "[About custom repository roles](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-repository-roles)."The authenticated user must be administrator of the organization or of a repository of the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemCustom_rolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemCustom_rolesRequestBuilder when successful
func (m *ItemCustom_rolesRequestBuilder) WithUrl(rawUrl string)(*ItemCustom_rolesRequestBuilder) {
    return NewItemCustom_rolesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
