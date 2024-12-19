package scim

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// V2EnterprisesItemUsersRequestBuilder builds and executes requests for operations under \scim\v2\enterprises\{enterprise}\Users
type V2EnterprisesItemUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V2EnterprisesItemUsersRequestBuilderGetQueryParameters lists provisioned SCIM enterprise members.When you remove a user with a SCIM-provisioned external identity from an enterprise using a `patch` with `active` flag to `false`, the user's metadata remains intact. This means they can potentially re-join the enterprise later. Although, while suspended, the user can't sign in. If you want to ensure the user can't re-join in the future, use the delete request. Only users who weren't permanently deleted will appear in the result list.
type V2EnterprisesItemUsersRequestBuilderGetQueryParameters struct {
    // Used for pagination: the number of results to return per page.
    Count *int32 `uriparametername:"count"`
    // If specified, only results that match the specified filter will be returned. Multiple filters are not supported. Possible filters are `userName`, `externalId`, `id`, and `displayName`. For example, `?filter="externalId eq '9138790-10932-109120392-12321'"`.
    Filter *string `uriparametername:"filter"`
    // Used for pagination: the starting index of the first result to return when paginating through values.
    StartIndex *int32 `uriparametername:"startIndex"`
}
// ByScim_user_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.scim.v2.enterprises.item.Users.item collection
// returns a *V2EnterprisesItemUsersWithScim_user_ItemRequestBuilder when successful
func (m *V2EnterprisesItemUsersRequestBuilder) ByScim_user_id(scim_user_id string)(*V2EnterprisesItemUsersWithScim_user_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if scim_user_id != "" {
        urlTplParams["scim_user_id"] = scim_user_id
    }
    return NewV2EnterprisesItemUsersWithScim_user_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV2EnterprisesItemUsersRequestBuilderInternal instantiates a new V2EnterprisesItemUsersRequestBuilder and sets the default values.
func NewV2EnterprisesItemUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemUsersRequestBuilder) {
    m := &V2EnterprisesItemUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/enterprises/{enterprise}/Users{?count*,filter*,startIndex*}", pathParameters),
    }
    return m
}
// NewV2EnterprisesItemUsersRequestBuilder instantiates a new V2EnterprisesItemUsersRequestBuilder and sets the default values.
func NewV2EnterprisesItemUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2EnterprisesItemUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists provisioned SCIM enterprise members.When you remove a user with a SCIM-provisioned external identity from an enterprise using a `patch` with `active` flag to `false`, the user's metadata remains intact. This means they can potentially re-join the enterprise later. Although, while suspended, the user can't sign in. If you want to ensure the user can't re-join in the future, use the delete request. Only users who weren't permanently deleted will appear in the result list.
// returns a ScimEnterpriseUserListable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/scim#list-scim-provisioned-identities-for-an-enterprise
func (m *V2EnterprisesItemUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemUsersRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimEnterpriseUserListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "429": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimEnterpriseUserListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimEnterpriseUserListable), nil
}
// Post creates an external identity for a new SCIM enterprise user.SCIM is responsible for user provisioning, not authentication. The actual user authentication is handled by SAML. However, with SCIM enabled, users must first be provisioned via SCIM before they can sign in through SAML.
// returns a ScimEnterpriseUserResponseable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/scim#provision-a-scim-enterprise-user
func (m *V2EnterprisesItemUsersRequestBuilder) Post(ctx context.Context, body i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Userable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimEnterpriseUserResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "429": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimEnterpriseUserResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimEnterpriseUserResponseable), nil
}
// ToGetRequestInformation lists provisioned SCIM enterprise members.When you remove a user with a SCIM-provisioned external identity from an enterprise using a `patch` with `active` flag to `false`, the user's metadata remains intact. This means they can potentially re-join the enterprise later. Although, while suspended, the user can't sign in. If you want to ensure the user can't re-join in the future, use the delete request. Only users who weren't permanently deleted will appear in the result list.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemUsersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    return requestInfo, nil
}
// ToPostRequestInformation creates an external identity for a new SCIM enterprise user.SCIM is responsible for user provisioning, not authentication. The actual user authentication is handled by SAML. However, with SCIM enabled, users must first be provisioned via SCIM before they can sign in through SAML.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemUsersRequestBuilder) ToPostRequestInformation(ctx context.Context, body i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Userable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V2EnterprisesItemUsersRequestBuilder when successful
func (m *V2EnterprisesItemUsersRequestBuilder) WithUrl(rawUrl string)(*V2EnterprisesItemUsersRequestBuilder) {
    return NewV2EnterprisesItemUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
