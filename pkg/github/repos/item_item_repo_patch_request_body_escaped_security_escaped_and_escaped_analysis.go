package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemRepoPatchRequestBody_security_and_analysis specify which security and analysis features to enable or disable for the repository.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/enterprise-cloud@latest//organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."For example, to enable GitHub Advanced Security, use this data in the body of the `PATCH` request:`{ "security_and_analysis": {"advanced_security": { "status": "enabled" } } }`.You can check which security and analysis features are currently enabled by using a `GET /repos/{owner}/{repo}` request.
type ItemItemRepoPatchRequestBody_security_and_analysis struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Use the `status` property to enable or disable GitHub Advanced Security for this repository. For more information, see "[About GitHub Advanced Security](/github/getting-started-with-github/learning-about-github/about-github-advanced-security)."
    advanced_security ItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityable
    // Use the `status` property to enable or disable secret scanning for this repository. For more information, see "[About secret scanning](/code-security/secret-security/about-secret-scanning)."
    secret_scanning ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningable
    // Use the `status` property to enable or disable secret scanning AI detection for this repository. For more information, see "[Responsible detection of generic secrets with AI](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/using-advanced-secret-scanning-and-push-protection-features/generic-secret-detection/responsible-ai-generic-secrets)."
    secret_scanning_ai_detection ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionable
    // Use the `status` property to enable or disable secret scanning non-provider patterns for this repository. For more information, see "[Supported secret scanning patterns](/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
    secret_scanning_non_provider_patterns ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsable
    // Use the `status` property to enable or disable secret scanning push protection for this repository. For more information, see "[Protecting pushes with secret scanning](/code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
    secret_scanning_push_protection ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable
    // Use the `status` property to enable or disable secret scanning automatic validity checks on supported partner tokens for this repository.
    secret_scanning_validity_checks ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable
}
// NewItemItemRepoPatchRequestBody_security_and_analysis instantiates a new ItemItemRepoPatchRequestBody_security_and_analysis and sets the default values.
func NewItemItemRepoPatchRequestBody_security_and_analysis()(*ItemItemRepoPatchRequestBody_security_and_analysis) {
    m := &ItemItemRepoPatchRequestBody_security_and_analysis{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemRepoPatchRequestBody_security_and_analysisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemRepoPatchRequestBody_security_and_analysisFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemRepoPatchRequestBody_security_and_analysis(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurity gets the advanced_security property value. Use the `status` property to enable or disable GitHub Advanced Security for this repository. For more information, see "[About GitHub Advanced Security](/github/getting-started-with-github/learning-about-github/about-github-advanced-security)."
// returns a ItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityable when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetAdvancedSecurity()(ItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityable) {
    return m.advanced_security
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurity(val.(ItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityable))
        }
        return nil
    }
    res["secret_scanning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanning(val.(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningable))
        }
        return nil
    }
    res["secret_scanning_ai_detection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningAiDetection(val.(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionable))
        }
        return nil
    }
    res["secret_scanning_non_provider_patterns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningNonProviderPatterns(val.(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsable))
        }
        return nil
    }
    res["secret_scanning_push_protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtection(val.(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable))
        }
        return nil
    }
    res["secret_scanning_validity_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningValidityChecks(val.(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable))
        }
        return nil
    }
    return res
}
// GetSecretScanning gets the secret_scanning property value. Use the `status` property to enable or disable secret scanning for this repository. For more information, see "[About secret scanning](/code-security/secret-security/about-secret-scanning)."
// returns a ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningable when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetSecretScanning()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningable) {
    return m.secret_scanning
}
// GetSecretScanningAiDetection gets the secret_scanning_ai_detection property value. Use the `status` property to enable or disable secret scanning AI detection for this repository. For more information, see "[Responsible detection of generic secrets with AI](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/using-advanced-secret-scanning-and-push-protection-features/generic-secret-detection/responsible-ai-generic-secrets)."
// returns a ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionable when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetSecretScanningAiDetection()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionable) {
    return m.secret_scanning_ai_detection
}
// GetSecretScanningNonProviderPatterns gets the secret_scanning_non_provider_patterns property value. Use the `status` property to enable or disable secret scanning non-provider patterns for this repository. For more information, see "[Supported secret scanning patterns](/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
// returns a ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsable when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetSecretScanningNonProviderPatterns()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsable) {
    return m.secret_scanning_non_provider_patterns
}
// GetSecretScanningPushProtection gets the secret_scanning_push_protection property value. Use the `status` property to enable or disable secret scanning push protection for this repository. For more information, see "[Protecting pushes with secret scanning](/code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
// returns a ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetSecretScanningPushProtection()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable) {
    return m.secret_scanning_push_protection
}
// GetSecretScanningValidityChecks gets the secret_scanning_validity_checks property value. Use the `status` property to enable or disable secret scanning automatic validity checks on supported partner tokens for this repository.
// returns a ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) GetSecretScanningValidityChecks()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable) {
    return m.secret_scanning_validity_checks
}
// Serialize serializes information the current object
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("advanced_security", m.GetAdvancedSecurity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning", m.GetSecretScanning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning_ai_detection", m.GetSecretScanningAiDetection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning_non_provider_patterns", m.GetSecretScanningNonProviderPatterns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning_push_protection", m.GetSecretScanningPushProtection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning_validity_checks", m.GetSecretScanningValidityChecks())
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
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurity sets the advanced_security property value. Use the `status` property to enable or disable GitHub Advanced Security for this repository. For more information, see "[About GitHub Advanced Security](/github/getting-started-with-github/learning-about-github/about-github-advanced-security)."
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) SetAdvancedSecurity(value ItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityable)() {
    m.advanced_security = value
}
// SetSecretScanning sets the secret_scanning property value. Use the `status` property to enable or disable secret scanning for this repository. For more information, see "[About secret scanning](/code-security/secret-security/about-secret-scanning)."
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) SetSecretScanning(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningable)() {
    m.secret_scanning = value
}
// SetSecretScanningAiDetection sets the secret_scanning_ai_detection property value. Use the `status` property to enable or disable secret scanning AI detection for this repository. For more information, see "[Responsible detection of generic secrets with AI](https://docs.github.com/enterprise-cloud@latest//code-security/secret-scanning/using-advanced-secret-scanning-and-push-protection-features/generic-secret-detection/responsible-ai-generic-secrets)."
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) SetSecretScanningAiDetection(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionable)() {
    m.secret_scanning_ai_detection = value
}
// SetSecretScanningNonProviderPatterns sets the secret_scanning_non_provider_patterns property value. Use the `status` property to enable or disable secret scanning non-provider patterns for this repository. For more information, see "[Supported secret scanning patterns](/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) SetSecretScanningNonProviderPatterns(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsable)() {
    m.secret_scanning_non_provider_patterns = value
}
// SetSecretScanningPushProtection sets the secret_scanning_push_protection property value. Use the `status` property to enable or disable secret scanning push protection for this repository. For more information, see "[Protecting pushes with secret scanning](/code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) SetSecretScanningPushProtection(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable)() {
    m.secret_scanning_push_protection = value
}
// SetSecretScanningValidityChecks sets the secret_scanning_validity_checks property value. Use the `status` property to enable or disable secret scanning automatic validity checks on supported partner tokens for this repository.
func (m *ItemItemRepoPatchRequestBody_security_and_analysis) SetSecretScanningValidityChecks(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable)() {
    m.secret_scanning_validity_checks = value
}
type ItemItemRepoPatchRequestBody_security_and_analysisable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurity()(ItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityable)
    GetSecretScanning()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningable)
    GetSecretScanningAiDetection()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionable)
    GetSecretScanningNonProviderPatterns()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsable)
    GetSecretScanningPushProtection()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable)
    GetSecretScanningValidityChecks()(ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable)
    SetAdvancedSecurity(value ItemItemRepoPatchRequestBody_security_and_analysis_advanced_securityable)()
    SetSecretScanning(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanningable)()
    SetSecretScanningAiDetection(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_ai_detectionable)()
    SetSecretScanningNonProviderPatterns(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_non_provider_patternsable)()
    SetSecretScanningPushProtection(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable)()
    SetSecretScanningValidityChecks(value ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable)()
}
