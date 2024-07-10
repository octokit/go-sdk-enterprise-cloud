package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemCredentialAuthorizationsRequestBuilder builds and executes requests for operations under \orgs\{org}\credential-authorizations
type ItemCredentialAuthorizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCredentialAuthorizationsRequestBuilderGetQueryParameters lists all credential authorizations for an organization that uses SAML single sign-on (SSO). The credentials are either personal access tokens or SSH keys that organization members have authorized for the organization. For more information, see [About authentication with SAML single sign-on](https://docs.github.com/enterprise-cloud@latest//articles/about-authentication-with-saml-single-sign-on).The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
type ItemCredentialAuthorizationsRequestBuilderGetQueryParameters struct {
    // Limits the list of credentials authorizations for an organization to a specific login
    Login *string `uriparametername:"login"`
    // Page token
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByCredential_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.orgs.item.credentialAuthorizations.item collection
// returns a *ItemCredentialAuthorizationsWithCredential_ItemRequestBuilder when successful
func (m *ItemCredentialAuthorizationsRequestBuilder) ByCredential_id(credential_id int32)(*ItemCredentialAuthorizationsWithCredential_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["credential_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(credential_id), 10)
    return NewItemCredentialAuthorizationsWithCredential_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCredentialAuthorizationsRequestBuilderInternal instantiates a new ItemCredentialAuthorizationsRequestBuilder and sets the default values.
func NewItemCredentialAuthorizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCredentialAuthorizationsRequestBuilder) {
    m := &ItemCredentialAuthorizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/credential-authorizations{?login*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemCredentialAuthorizationsRequestBuilder instantiates a new ItemCredentialAuthorizationsRequestBuilder and sets the default values.
func NewItemCredentialAuthorizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCredentialAuthorizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCredentialAuthorizationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all credential authorizations for an organization that uses SAML single sign-on (SSO). The credentials are either personal access tokens or SSH keys that organization members have authorized for the organization. For more information, see [About authentication with SAML single sign-on](https://docs.github.com/enterprise-cloud@latest//articles/about-authentication-with-saml-single-sign-on).The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
// returns a []CredentialAuthorizationable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/orgs#list-saml-sso-authorizations-for-an-organization
func (m *ItemCredentialAuthorizationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCredentialAuthorizationsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CredentialAuthorizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateCredentialAuthorizationFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CredentialAuthorizationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CredentialAuthorizationable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all credential authorizations for an organization that uses SAML single sign-on (SSO). The credentials are either personal access tokens or SSH keys that organization members have authorized for the organization. For more information, see [About authentication with SAML single sign-on](https://docs.github.com/enterprise-cloud@latest//articles/about-authentication-with-saml-single-sign-on).The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemCredentialAuthorizationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCredentialAuthorizationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCredentialAuthorizationsRequestBuilder when successful
func (m *ItemCredentialAuthorizationsRequestBuilder) WithUrl(rawUrl string)(*ItemCredentialAuthorizationsRequestBuilder) {
    return NewItemCredentialAuthorizationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
