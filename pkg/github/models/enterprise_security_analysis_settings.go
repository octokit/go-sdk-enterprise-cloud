package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSecurityAnalysisSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether GitHub advanced security is automatically enabled for new repositories and repositories transferred tothis enterprise.
    advanced_security_enabled_for_new_repositories *bool
    // Whether GitHub Advanced Security is automatically enabled for new user namespace repositories.
    advanced_security_enabled_for_new_user_namespace_repositories *bool
    // Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to thisenterprise.
    dependabot_alerts_enabled_for_new_repositories *bool
    // Whether secret scanning is automatically enabled for new repositories and repositories transferred to thisenterprise.
    secret_scanning_enabled_for_new_repositories *bool
    // Whether secret scanning of non-provider patterns is enabled for new repositories under this enterprise.
    secret_scanning_non_provider_patterns_enabled_for_new_repositories *bool
    // An optional URL string to display to contributors who are blocked from pushing a secret.
    secret_scanning_push_protection_custom_link *string
    // Whether secret scanning push protection is automatically enabled for new repositories and repositoriestransferred to this enterprise.
    secret_scanning_push_protection_enabled_for_new_repositories *bool
    // Whether secret scanning automatic validity checks on supported partner tokens is enabled for all repositories under this enterprise.
    secret_scanning_validity_checks_enabled *bool
}
// NewEnterpriseSecurityAnalysisSettings instantiates a new EnterpriseSecurityAnalysisSettings and sets the default values.
func NewEnterpriseSecurityAnalysisSettings()(*EnterpriseSecurityAnalysisSettings) {
    m := &EnterpriseSecurityAnalysisSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSecurityAnalysisSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSecurityAnalysisSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSecurityAnalysisSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSecurityAnalysisSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurityEnabledForNewRepositories gets the advanced_security_enabled_for_new_repositories property value. Whether GitHub advanced security is automatically enabled for new repositories and repositories transferred tothis enterprise.
// returns a *bool when successful
func (m *EnterpriseSecurityAnalysisSettings) GetAdvancedSecurityEnabledForNewRepositories()(*bool) {
    return m.advanced_security_enabled_for_new_repositories
}
// GetAdvancedSecurityEnabledForNewUserNamespaceRepositories gets the advanced_security_enabled_for_new_user_namespace_repositories property value. Whether GitHub Advanced Security is automatically enabled for new user namespace repositories.
// returns a *bool when successful
func (m *EnterpriseSecurityAnalysisSettings) GetAdvancedSecurityEnabledForNewUserNamespaceRepositories()(*bool) {
    return m.advanced_security_enabled_for_new_user_namespace_repositories
}
// GetDependabotAlertsEnabledForNewRepositories gets the dependabot_alerts_enabled_for_new_repositories property value. Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to thisenterprise.
// returns a *bool when successful
func (m *EnterpriseSecurityAnalysisSettings) GetDependabotAlertsEnabledForNewRepositories()(*bool) {
    return m.dependabot_alerts_enabled_for_new_repositories
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSecurityAnalysisSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurityEnabledForNewRepositories(val)
        }
        return nil
    }
    res["advanced_security_enabled_for_new_user_namespace_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurityEnabledForNewUserNamespaceRepositories(val)
        }
        return nil
    }
    res["dependabot_alerts_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotAlertsEnabledForNewRepositories(val)
        }
        return nil
    }
    res["secret_scanning_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningEnabledForNewRepositories(val)
        }
        return nil
    }
    res["secret_scanning_non_provider_patterns_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningNonProviderPatternsEnabledForNewRepositories(val)
        }
        return nil
    }
    res["secret_scanning_push_protection_custom_link"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtectionCustomLink(val)
        }
        return nil
    }
    res["secret_scanning_push_protection_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtectionEnabledForNewRepositories(val)
        }
        return nil
    }
    res["secret_scanning_validity_checks_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningValidityChecksEnabled(val)
        }
        return nil
    }
    return res
}
// GetSecretScanningEnabledForNewRepositories gets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories and repositories transferred to thisenterprise.
// returns a *bool when successful
func (m *EnterpriseSecurityAnalysisSettings) GetSecretScanningEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_enabled_for_new_repositories
}
// GetSecretScanningNonProviderPatternsEnabledForNewRepositories gets the secret_scanning_non_provider_patterns_enabled_for_new_repositories property value. Whether secret scanning of non-provider patterns is enabled for new repositories under this enterprise.
// returns a *bool when successful
func (m *EnterpriseSecurityAnalysisSettings) GetSecretScanningNonProviderPatternsEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_non_provider_patterns_enabled_for_new_repositories
}
// GetSecretScanningPushProtectionCustomLink gets the secret_scanning_push_protection_custom_link property value. An optional URL string to display to contributors who are blocked from pushing a secret.
// returns a *string when successful
func (m *EnterpriseSecurityAnalysisSettings) GetSecretScanningPushProtectionCustomLink()(*string) {
    return m.secret_scanning_push_protection_custom_link
}
// GetSecretScanningPushProtectionEnabledForNewRepositories gets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories and repositoriestransferred to this enterprise.
// returns a *bool when successful
func (m *EnterpriseSecurityAnalysisSettings) GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_push_protection_enabled_for_new_repositories
}
// GetSecretScanningValidityChecksEnabled gets the secret_scanning_validity_checks_enabled property value. Whether secret scanning automatic validity checks on supported partner tokens is enabled for all repositories under this enterprise.
// returns a *bool when successful
func (m *EnterpriseSecurityAnalysisSettings) GetSecretScanningValidityChecksEnabled()(*bool) {
    return m.secret_scanning_validity_checks_enabled
}
// Serialize serializes information the current object
func (m *EnterpriseSecurityAnalysisSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("advanced_security_enabled_for_new_repositories", m.GetAdvancedSecurityEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("advanced_security_enabled_for_new_user_namespace_repositories", m.GetAdvancedSecurityEnabledForNewUserNamespaceRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dependabot_alerts_enabled_for_new_repositories", m.GetDependabotAlertsEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secret_scanning_enabled_for_new_repositories", m.GetSecretScanningEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secret_scanning_non_provider_patterns_enabled_for_new_repositories", m.GetSecretScanningNonProviderPatternsEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret_scanning_push_protection_custom_link", m.GetSecretScanningPushProtectionCustomLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secret_scanning_push_protection_enabled_for_new_repositories", m.GetSecretScanningPushProtectionEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secret_scanning_validity_checks_enabled", m.GetSecretScanningValidityChecksEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnterpriseSecurityAnalysisSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurityEnabledForNewRepositories sets the advanced_security_enabled_for_new_repositories property value. Whether GitHub advanced security is automatically enabled for new repositories and repositories transferred tothis enterprise.
func (m *EnterpriseSecurityAnalysisSettings) SetAdvancedSecurityEnabledForNewRepositories(value *bool)() {
    m.advanced_security_enabled_for_new_repositories = value
}
// SetAdvancedSecurityEnabledForNewUserNamespaceRepositories sets the advanced_security_enabled_for_new_user_namespace_repositories property value. Whether GitHub Advanced Security is automatically enabled for new user namespace repositories.
func (m *EnterpriseSecurityAnalysisSettings) SetAdvancedSecurityEnabledForNewUserNamespaceRepositories(value *bool)() {
    m.advanced_security_enabled_for_new_user_namespace_repositories = value
}
// SetDependabotAlertsEnabledForNewRepositories sets the dependabot_alerts_enabled_for_new_repositories property value. Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to thisenterprise.
func (m *EnterpriseSecurityAnalysisSettings) SetDependabotAlertsEnabledForNewRepositories(value *bool)() {
    m.dependabot_alerts_enabled_for_new_repositories = value
}
// SetSecretScanningEnabledForNewRepositories sets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories and repositories transferred to thisenterprise.
func (m *EnterpriseSecurityAnalysisSettings) SetSecretScanningEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_enabled_for_new_repositories = value
}
// SetSecretScanningNonProviderPatternsEnabledForNewRepositories sets the secret_scanning_non_provider_patterns_enabled_for_new_repositories property value. Whether secret scanning of non-provider patterns is enabled for new repositories under this enterprise.
func (m *EnterpriseSecurityAnalysisSettings) SetSecretScanningNonProviderPatternsEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_non_provider_patterns_enabled_for_new_repositories = value
}
// SetSecretScanningPushProtectionCustomLink sets the secret_scanning_push_protection_custom_link property value. An optional URL string to display to contributors who are blocked from pushing a secret.
func (m *EnterpriseSecurityAnalysisSettings) SetSecretScanningPushProtectionCustomLink(value *string)() {
    m.secret_scanning_push_protection_custom_link = value
}
// SetSecretScanningPushProtectionEnabledForNewRepositories sets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories and repositoriestransferred to this enterprise.
func (m *EnterpriseSecurityAnalysisSettings) SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_push_protection_enabled_for_new_repositories = value
}
// SetSecretScanningValidityChecksEnabled sets the secret_scanning_validity_checks_enabled property value. Whether secret scanning automatic validity checks on supported partner tokens is enabled for all repositories under this enterprise.
func (m *EnterpriseSecurityAnalysisSettings) SetSecretScanningValidityChecksEnabled(value *bool)() {
    m.secret_scanning_validity_checks_enabled = value
}
type EnterpriseSecurityAnalysisSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurityEnabledForNewRepositories()(*bool)
    GetAdvancedSecurityEnabledForNewUserNamespaceRepositories()(*bool)
    GetDependabotAlertsEnabledForNewRepositories()(*bool)
    GetSecretScanningEnabledForNewRepositories()(*bool)
    GetSecretScanningNonProviderPatternsEnabledForNewRepositories()(*bool)
    GetSecretScanningPushProtectionCustomLink()(*string)
    GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool)
    GetSecretScanningValidityChecksEnabled()(*bool)
    SetAdvancedSecurityEnabledForNewRepositories(value *bool)()
    SetAdvancedSecurityEnabledForNewUserNamespaceRepositories(value *bool)()
    SetDependabotAlertsEnabledForNewRepositories(value *bool)()
    SetSecretScanningEnabledForNewRepositories(value *bool)()
    SetSecretScanningNonProviderPatternsEnabledForNewRepositories(value *bool)()
    SetSecretScanningPushProtectionCustomLink(value *string)()
    SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)()
    SetSecretScanningValidityChecksEnabled(value *bool)()
}
