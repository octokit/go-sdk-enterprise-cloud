package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithEnterpriseItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}
type WithEnterpriseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Actions the actions property
// returns a *ItemActionsRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Actions()(*ItemActionsRequestBuilder) {
    return NewItemActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Announcement the announcement property
// returns a *ItemAnnouncementRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Announcement()(*ItemAnnouncementRequestBuilder) {
    return NewItemAnnouncementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditLog the auditLog property
// returns a *ItemAuditLogRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) AuditLog()(*ItemAuditLogRequestBuilder) {
    return NewItemAuditLogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BySecurity_product gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterprises.item.item collection
// returns a *ItemWithSecurity_productItemRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) BySecurity_product(security_product string)(*ItemWithSecurity_productItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if security_product != "" {
        urlTplParams["security_product"] = security_product
    }
    return NewItemWithSecurity_productItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Code_security_and_analysis the code_security_and_analysis property
// returns a *ItemCode_security_and_analysisRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Code_security_and_analysis()(*ItemCode_security_and_analysisRequestBuilder) {
    return NewItemCode_security_and_analysisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CodeScanning the codeScanning property
// returns a *ItemCodeScanningRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) CodeScanning()(*ItemCodeScanningRequestBuilder) {
    return NewItemCodeScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithEnterpriseItemRequestBuilderInternal instantiates a new WithEnterpriseItemRequestBuilder and sets the default values.
func NewWithEnterpriseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterpriseItemRequestBuilder) {
    m := &WithEnterpriseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}", pathParameters),
    }
    return m
}
// NewWithEnterpriseItemRequestBuilder instantiates a new WithEnterpriseItemRequestBuilder and sets the default values.
func NewWithEnterpriseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterpriseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithEnterpriseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ConsumedLicenses the consumedLicenses property
// returns a *ItemConsumedLicensesRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) ConsumedLicenses()(*ItemConsumedLicensesRequestBuilder) {
    return NewItemConsumedLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copilot the copilot property
// returns a *ItemCopilotRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Copilot()(*ItemCopilotRequestBuilder) {
    return NewItemCopilotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dependabot the dependabot property
// returns a *ItemDependabotRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Dependabot()(*ItemDependabotRequestBuilder) {
    return NewItemDependabotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LicenseSyncStatus the licenseSyncStatus property
// returns a *ItemLicenseSyncStatusRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) LicenseSyncStatus()(*ItemLicenseSyncStatusRequestBuilder) {
    return NewItemLicenseSyncStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecretScanning the secretScanning property
// returns a *ItemSecretScanningRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) SecretScanning()(*ItemSecretScanningRequestBuilder) {
    return NewItemSecretScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
// returns a *ItemSettingsRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
