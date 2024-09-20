package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCode_security_and_analysisPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether GitHub Advanced Security is automatically enabled for new repositories. For more information, see "[About GitHub Advanced Security](https://docs.github.com/enterprise-cloud@latest//get-started/learning-about-github/about-github-advanced-security)."
    advanced_security_enabled_for_new_repositories *bool
    // Whether GitHub Advanced Security is automatically enabled for new user namespace repositories. For more information, see "[About GitHub Advanced Security](https://docs.github.com/enterprise-cloud@latest//get-started/learning-about-github/about-github-advanced-security)."
    advanced_security_enabled_new_user_namespace_repos *bool
    // Whether Dependabot alerts are automatically enabled for new repositories. For more information, see "[About Dependabot alerts](https://docs.github.com/enterprise-cloud@latest//code-security/dependabot/dependabot-alerts/about-dependabot-alerts)."
    dependabot_alerts_enabled_for_new_repositories *bool
    // Whether secret scanning is automatically enabled for new repositories. For more information, see "[About secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/about-secret-scanning)."
    secret_scanning_enabled_for_new_repositories *bool
    // Whether secret scanning of non-provider patterns is enabled for new repositories under this enterprise.
    secret_scanning_non_provider_patterns_enabled_for_new_repositories *bool
    // The URL that will be displayed to contributors who are blocked from pushing a secret. For more information, see "[Protecting pushes with secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/protecting-pushes-with-secret-scanning)."To disable this functionality, set this field to `null`.
    secret_scanning_push_protection_custom_link *string
    // Whether secret scanning push protection is automatically enabled for new repositories. For more information, see "[Protecting pushes with secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
    secret_scanning_push_protection_enabled_for_new_repositories *bool
}
// NewItemCode_security_and_analysisPatchRequestBody instantiates a new ItemCode_security_and_analysisPatchRequestBody and sets the default values.
func NewItemCode_security_and_analysisPatchRequestBody()(*ItemCode_security_and_analysisPatchRequestBody) {
    m := &ItemCode_security_and_analysisPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCode_security_and_analysisPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCode_security_and_analysisPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCode_security_and_analysisPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurityEnabledForNewRepositories gets the advanced_security_enabled_for_new_repositories property value. Whether GitHub Advanced Security is automatically enabled for new repositories. For more information, see "[About GitHub Advanced Security](https://docs.github.com/enterprise-cloud@latest//get-started/learning-about-github/about-github-advanced-security)."
// returns a *bool when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetAdvancedSecurityEnabledForNewRepositories()(*bool) {
    return m.advanced_security_enabled_for_new_repositories
}
// GetAdvancedSecurityEnabledNewUserNamespaceRepos gets the advanced_security_enabled_new_user_namespace_repos property value. Whether GitHub Advanced Security is automatically enabled for new user namespace repositories. For more information, see "[About GitHub Advanced Security](https://docs.github.com/enterprise-cloud@latest//get-started/learning-about-github/about-github-advanced-security)."
// returns a *bool when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetAdvancedSecurityEnabledNewUserNamespaceRepos()(*bool) {
    return m.advanced_security_enabled_new_user_namespace_repos
}
// GetDependabotAlertsEnabledForNewRepositories gets the dependabot_alerts_enabled_for_new_repositories property value. Whether Dependabot alerts are automatically enabled for new repositories. For more information, see "[About Dependabot alerts](https://docs.github.com/enterprise-cloud@latest//code-security/dependabot/dependabot-alerts/about-dependabot-alerts)."
// returns a *bool when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetDependabotAlertsEnabledForNewRepositories()(*bool) {
    return m.dependabot_alerts_enabled_for_new_repositories
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["advanced_security_enabled_new_user_namespace_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurityEnabledNewUserNamespaceRepos(val)
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
    return res
}
// GetSecretScanningEnabledForNewRepositories gets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories. For more information, see "[About secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/about-secret-scanning)."
// returns a *bool when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetSecretScanningEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_enabled_for_new_repositories
}
// GetSecretScanningNonProviderPatternsEnabledForNewRepositories gets the secret_scanning_non_provider_patterns_enabled_for_new_repositories property value. Whether secret scanning of non-provider patterns is enabled for new repositories under this enterprise.
// returns a *bool when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetSecretScanningNonProviderPatternsEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_non_provider_patterns_enabled_for_new_repositories
}
// GetSecretScanningPushProtectionCustomLink gets the secret_scanning_push_protection_custom_link property value. The URL that will be displayed to contributors who are blocked from pushing a secret. For more information, see "[Protecting pushes with secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/protecting-pushes-with-secret-scanning)."To disable this functionality, set this field to `null`.
// returns a *string when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetSecretScanningPushProtectionCustomLink()(*string) {
    return m.secret_scanning_push_protection_custom_link
}
// GetSecretScanningPushProtectionEnabledForNewRepositories gets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories. For more information, see "[Protecting pushes with secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
// returns a *bool when successful
func (m *ItemCode_security_and_analysisPatchRequestBody) GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_push_protection_enabled_for_new_repositories
}
// Serialize serializes information the current object
func (m *ItemCode_security_and_analysisPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("advanced_security_enabled_for_new_repositories", m.GetAdvancedSecurityEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("advanced_security_enabled_new_user_namespace_repos", m.GetAdvancedSecurityEnabledNewUserNamespaceRepos())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCode_security_and_analysisPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurityEnabledForNewRepositories sets the advanced_security_enabled_for_new_repositories property value. Whether GitHub Advanced Security is automatically enabled for new repositories. For more information, see "[About GitHub Advanced Security](https://docs.github.com/enterprise-cloud@latest//get-started/learning-about-github/about-github-advanced-security)."
func (m *ItemCode_security_and_analysisPatchRequestBody) SetAdvancedSecurityEnabledForNewRepositories(value *bool)() {
    m.advanced_security_enabled_for_new_repositories = value
}
// SetAdvancedSecurityEnabledNewUserNamespaceRepos sets the advanced_security_enabled_new_user_namespace_repos property value. Whether GitHub Advanced Security is automatically enabled for new user namespace repositories. For more information, see "[About GitHub Advanced Security](https://docs.github.com/enterprise-cloud@latest//get-started/learning-about-github/about-github-advanced-security)."
func (m *ItemCode_security_and_analysisPatchRequestBody) SetAdvancedSecurityEnabledNewUserNamespaceRepos(value *bool)() {
    m.advanced_security_enabled_new_user_namespace_repos = value
}
// SetDependabotAlertsEnabledForNewRepositories sets the dependabot_alerts_enabled_for_new_repositories property value. Whether Dependabot alerts are automatically enabled for new repositories. For more information, see "[About Dependabot alerts](https://docs.github.com/enterprise-cloud@latest//code-security/dependabot/dependabot-alerts/about-dependabot-alerts)."
func (m *ItemCode_security_and_analysisPatchRequestBody) SetDependabotAlertsEnabledForNewRepositories(value *bool)() {
    m.dependabot_alerts_enabled_for_new_repositories = value
}
// SetSecretScanningEnabledForNewRepositories sets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories. For more information, see "[About secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/about-secret-scanning)."
func (m *ItemCode_security_and_analysisPatchRequestBody) SetSecretScanningEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_enabled_for_new_repositories = value
}
// SetSecretScanningNonProviderPatternsEnabledForNewRepositories sets the secret_scanning_non_provider_patterns_enabled_for_new_repositories property value. Whether secret scanning of non-provider patterns is enabled for new repositories under this enterprise.
func (m *ItemCode_security_and_analysisPatchRequestBody) SetSecretScanningNonProviderPatternsEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_non_provider_patterns_enabled_for_new_repositories = value
}
// SetSecretScanningPushProtectionCustomLink sets the secret_scanning_push_protection_custom_link property value. The URL that will be displayed to contributors who are blocked from pushing a secret. For more information, see "[Protecting pushes with secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/protecting-pushes-with-secret-scanning)."To disable this functionality, set this field to `null`.
func (m *ItemCode_security_and_analysisPatchRequestBody) SetSecretScanningPushProtectionCustomLink(value *string)() {
    m.secret_scanning_push_protection_custom_link = value
}
// SetSecretScanningPushProtectionEnabledForNewRepositories sets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories. For more information, see "[Protecting pushes with secret scanning](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
func (m *ItemCode_security_and_analysisPatchRequestBody) SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_push_protection_enabled_for_new_repositories = value
}
type ItemCode_security_and_analysisPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurityEnabledForNewRepositories()(*bool)
    GetAdvancedSecurityEnabledNewUserNamespaceRepos()(*bool)
    GetDependabotAlertsEnabledForNewRepositories()(*bool)
    GetSecretScanningEnabledForNewRepositories()(*bool)
    GetSecretScanningNonProviderPatternsEnabledForNewRepositories()(*bool)
    GetSecretScanningPushProtectionCustomLink()(*string)
    GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool)
    SetAdvancedSecurityEnabledForNewRepositories(value *bool)()
    SetAdvancedSecurityEnabledNewUserNamespaceRepos(value *bool)()
    SetDependabotAlertsEnabledForNewRepositories(value *bool)()
    SetSecretScanningEnabledForNewRepositories(value *bool)()
    SetSecretScanningNonProviderPatternsEnabledForNewRepositories(value *bool)()
    SetSecretScanningPushProtectionCustomLink(value *string)()
    SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)()
}
