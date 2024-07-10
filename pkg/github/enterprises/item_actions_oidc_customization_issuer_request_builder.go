package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemActionsOidcCustomizationIssuerRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\actions\oidc\customization\issuer
type ItemActionsOidcCustomizationIssuerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsOidcCustomizationIssuerRequestBuilderInternal instantiates a new ItemActionsOidcCustomizationIssuerRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationIssuerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationIssuerRequestBuilder) {
    m := &ItemActionsOidcCustomizationIssuerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/actions/oidc/customization/issuer", pathParameters),
    }
    return m
}
// NewItemActionsOidcCustomizationIssuerRequestBuilder instantiates a new ItemActionsOidcCustomizationIssuerRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationIssuerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationIssuerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsOidcCustomizationIssuerRequestBuilderInternal(urlParams, requestAdapter)
}
// Put sets the GitHub Actions OpenID Connect (OIDC) custom issuer policy for an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/oidc#set-the-github-actions-oidc-custom-issuer-policy-for-an-enterprise
func (m *ItemActionsOidcCustomizationIssuerRequestBuilder) Put(ctx context.Context, body i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsOidcCustomIssuerPolicyForEnterpriseable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPutRequestInformation sets the GitHub Actions OpenID Connect (OIDC) custom issuer policy for an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsOidcCustomizationIssuerRequestBuilder) ToPutRequestInformation(ctx context.Context, body i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ActionsOidcCustomIssuerPolicyForEnterpriseable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsOidcCustomizationIssuerRequestBuilder when successful
func (m *ItemActionsOidcCustomizationIssuerRequestBuilder) WithUrl(rawUrl string)(*ItemActionsOidcCustomizationIssuerRequestBuilder) {
    return NewItemActionsOidcCustomizationIssuerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
