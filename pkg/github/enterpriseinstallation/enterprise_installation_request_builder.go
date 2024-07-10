package enterpriseinstallation

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnterpriseInstallationRequestBuilder builds and executes requests for operations under \enterprise-installation
type EnterpriseInstallationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByEnterprise_or_org gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterpriseInstallation.item collection
// returns a *WithEnterprise_or_orgItemRequestBuilder when successful
func (m *EnterpriseInstallationRequestBuilder) ByEnterprise_or_org(enterprise_or_org string)(*WithEnterprise_or_orgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if enterprise_or_org != "" {
        urlTplParams["enterprise_or_org"] = enterprise_or_org
    }
    return NewWithEnterprise_or_orgItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterpriseInstallationRequestBuilderInternal instantiates a new EnterpriseInstallationRequestBuilder and sets the default values.
func NewEnterpriseInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseInstallationRequestBuilder) {
    m := &EnterpriseInstallationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprise-installation", pathParameters),
    }
    return m
}
// NewEnterpriseInstallationRequestBuilder instantiates a new EnterpriseInstallationRequestBuilder and sets the default values.
func NewEnterpriseInstallationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseInstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
