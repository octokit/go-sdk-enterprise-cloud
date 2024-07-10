package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetConsumedLicenses_users struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enterprise_server_emails property
    enterprise_server_emails []string
    // The enterprise_server_user property
    enterprise_server_user *bool
    // The enterprise_server_user_ids property
    enterprise_server_user_ids []string
    // All enterprise roles for a user.
    github_com_enterprise_roles []string
    // The github_com_login property
    github_com_login *string
    // The github_com_member_roles property
    github_com_member_roles []string
    // The github_com_name property
    github_com_name *string
    // The github_com_orgs_with_pending_invites property
    github_com_orgs_with_pending_invites []string
    // The github_com_profile property
    github_com_profile *string
    // The github_com_saml_name_id property
    github_com_saml_name_id *string
    // The github_com_two_factor_auth property
    github_com_two_factor_auth *bool
    // The github_com_user property
    github_com_user *bool
    // The github_com_verified_domain_emails property
    github_com_verified_domain_emails []string
    // The license_type property
    license_type *string
    // The total_user_accounts property
    total_user_accounts *int32
    // The visual_studio_license_status property
    visual_studio_license_status *string
    // The visual_studio_subscription_email property
    visual_studio_subscription_email *string
    // The visual_studio_subscription_user property
    visual_studio_subscription_user *bool
}
// NewGetConsumedLicenses_users instantiates a new GetConsumedLicenses_users and sets the default values.
func NewGetConsumedLicenses_users()(*GetConsumedLicenses_users) {
    m := &GetConsumedLicenses_users{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetConsumedLicenses_usersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetConsumedLicenses_usersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetConsumedLicenses_users(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetConsumedLicenses_users) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnterpriseServerEmails gets the enterprise_server_emails property value. The enterprise_server_emails property
// returns a []string when successful
func (m *GetConsumedLicenses_users) GetEnterpriseServerEmails()([]string) {
    return m.enterprise_server_emails
}
// GetEnterpriseServerUser gets the enterprise_server_user property value. The enterprise_server_user property
// returns a *bool when successful
func (m *GetConsumedLicenses_users) GetEnterpriseServerUser()(*bool) {
    return m.enterprise_server_user
}
// GetEnterpriseServerUserIds gets the enterprise_server_user_ids property value. The enterprise_server_user_ids property
// returns a []string when successful
func (m *GetConsumedLicenses_users) GetEnterpriseServerUserIds()([]string) {
    return m.enterprise_server_user_ids
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetConsumedLicenses_users) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enterprise_server_emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetEnterpriseServerEmails(res)
        }
        return nil
    }
    res["enterprise_server_user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseServerUser(val)
        }
        return nil
    }
    res["enterprise_server_user_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetEnterpriseServerUserIds(res)
        }
        return nil
    }
    res["github_com_enterprise_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGithubComEnterpriseRoles(res)
        }
        return nil
    }
    res["github_com_login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubComLogin(val)
        }
        return nil
    }
    res["github_com_member_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGithubComMemberRoles(res)
        }
        return nil
    }
    res["github_com_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubComName(val)
        }
        return nil
    }
    res["github_com_orgs_with_pending_invites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGithubComOrgsWithPendingInvites(res)
        }
        return nil
    }
    res["github_com_profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubComProfile(val)
        }
        return nil
    }
    res["github_com_saml_name_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubComSamlNameId(val)
        }
        return nil
    }
    res["github_com_two_factor_auth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubComTwoFactorAuth(val)
        }
        return nil
    }
    res["github_com_user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubComUser(val)
        }
        return nil
    }
    res["github_com_verified_domain_emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGithubComVerifiedDomainEmails(res)
        }
        return nil
    }
    res["license_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val)
        }
        return nil
    }
    res["total_user_accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUserAccounts(val)
        }
        return nil
    }
    res["visual_studio_license_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisualStudioLicenseStatus(val)
        }
        return nil
    }
    res["visual_studio_subscription_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisualStudioSubscriptionEmail(val)
        }
        return nil
    }
    res["visual_studio_subscription_user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisualStudioSubscriptionUser(val)
        }
        return nil
    }
    return res
}
// GetGithubComEnterpriseRoles gets the github_com_enterprise_roles property value. All enterprise roles for a user.
// returns a []string when successful
func (m *GetConsumedLicenses_users) GetGithubComEnterpriseRoles()([]string) {
    return m.github_com_enterprise_roles
}
// GetGithubComLogin gets the github_com_login property value. The github_com_login property
// returns a *string when successful
func (m *GetConsumedLicenses_users) GetGithubComLogin()(*string) {
    return m.github_com_login
}
// GetGithubComMemberRoles gets the github_com_member_roles property value. The github_com_member_roles property
// returns a []string when successful
func (m *GetConsumedLicenses_users) GetGithubComMemberRoles()([]string) {
    return m.github_com_member_roles
}
// GetGithubComName gets the github_com_name property value. The github_com_name property
// returns a *string when successful
func (m *GetConsumedLicenses_users) GetGithubComName()(*string) {
    return m.github_com_name
}
// GetGithubComOrgsWithPendingInvites gets the github_com_orgs_with_pending_invites property value. The github_com_orgs_with_pending_invites property
// returns a []string when successful
func (m *GetConsumedLicenses_users) GetGithubComOrgsWithPendingInvites()([]string) {
    return m.github_com_orgs_with_pending_invites
}
// GetGithubComProfile gets the github_com_profile property value. The github_com_profile property
// returns a *string when successful
func (m *GetConsumedLicenses_users) GetGithubComProfile()(*string) {
    return m.github_com_profile
}
// GetGithubComSamlNameId gets the github_com_saml_name_id property value. The github_com_saml_name_id property
// returns a *string when successful
func (m *GetConsumedLicenses_users) GetGithubComSamlNameId()(*string) {
    return m.github_com_saml_name_id
}
// GetGithubComTwoFactorAuth gets the github_com_two_factor_auth property value. The github_com_two_factor_auth property
// returns a *bool when successful
func (m *GetConsumedLicenses_users) GetGithubComTwoFactorAuth()(*bool) {
    return m.github_com_two_factor_auth
}
// GetGithubComUser gets the github_com_user property value. The github_com_user property
// returns a *bool when successful
func (m *GetConsumedLicenses_users) GetGithubComUser()(*bool) {
    return m.github_com_user
}
// GetGithubComVerifiedDomainEmails gets the github_com_verified_domain_emails property value. The github_com_verified_domain_emails property
// returns a []string when successful
func (m *GetConsumedLicenses_users) GetGithubComVerifiedDomainEmails()([]string) {
    return m.github_com_verified_domain_emails
}
// GetLicenseType gets the license_type property value. The license_type property
// returns a *string when successful
func (m *GetConsumedLicenses_users) GetLicenseType()(*string) {
    return m.license_type
}
// GetTotalUserAccounts gets the total_user_accounts property value. The total_user_accounts property
// returns a *int32 when successful
func (m *GetConsumedLicenses_users) GetTotalUserAccounts()(*int32) {
    return m.total_user_accounts
}
// GetVisualStudioLicenseStatus gets the visual_studio_license_status property value. The visual_studio_license_status property
// returns a *string when successful
func (m *GetConsumedLicenses_users) GetVisualStudioLicenseStatus()(*string) {
    return m.visual_studio_license_status
}
// GetVisualStudioSubscriptionEmail gets the visual_studio_subscription_email property value. The visual_studio_subscription_email property
// returns a *string when successful
func (m *GetConsumedLicenses_users) GetVisualStudioSubscriptionEmail()(*string) {
    return m.visual_studio_subscription_email
}
// GetVisualStudioSubscriptionUser gets the visual_studio_subscription_user property value. The visual_studio_subscription_user property
// returns a *bool when successful
func (m *GetConsumedLicenses_users) GetVisualStudioSubscriptionUser()(*bool) {
    return m.visual_studio_subscription_user
}
// Serialize serializes information the current object
func (m *GetConsumedLicenses_users) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnterpriseServerEmails() != nil {
        err := writer.WriteCollectionOfStringValues("enterprise_server_emails", m.GetEnterpriseServerEmails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enterprise_server_user", m.GetEnterpriseServerUser())
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseServerUserIds() != nil {
        err := writer.WriteCollectionOfStringValues("enterprise_server_user_ids", m.GetEnterpriseServerUserIds())
        if err != nil {
            return err
        }
    }
    if m.GetGithubComEnterpriseRoles() != nil {
        err := writer.WriteCollectionOfStringValues("github_com_enterprise_roles", m.GetGithubComEnterpriseRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("github_com_login", m.GetGithubComLogin())
        if err != nil {
            return err
        }
    }
    if m.GetGithubComMemberRoles() != nil {
        err := writer.WriteCollectionOfStringValues("github_com_member_roles", m.GetGithubComMemberRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("github_com_name", m.GetGithubComName())
        if err != nil {
            return err
        }
    }
    if m.GetGithubComOrgsWithPendingInvites() != nil {
        err := writer.WriteCollectionOfStringValues("github_com_orgs_with_pending_invites", m.GetGithubComOrgsWithPendingInvites())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("github_com_profile", m.GetGithubComProfile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("github_com_saml_name_id", m.GetGithubComSamlNameId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("github_com_two_factor_auth", m.GetGithubComTwoFactorAuth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("github_com_user", m.GetGithubComUser())
        if err != nil {
            return err
        }
    }
    if m.GetGithubComVerifiedDomainEmails() != nil {
        err := writer.WriteCollectionOfStringValues("github_com_verified_domain_emails", m.GetGithubComVerifiedDomainEmails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("license_type", m.GetLicenseType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_user_accounts", m.GetTotalUserAccounts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("visual_studio_license_status", m.GetVisualStudioLicenseStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("visual_studio_subscription_email", m.GetVisualStudioSubscriptionEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("visual_studio_subscription_user", m.GetVisualStudioSubscriptionUser())
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
func (m *GetConsumedLicenses_users) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnterpriseServerEmails sets the enterprise_server_emails property value. The enterprise_server_emails property
func (m *GetConsumedLicenses_users) SetEnterpriseServerEmails(value []string)() {
    m.enterprise_server_emails = value
}
// SetEnterpriseServerUser sets the enterprise_server_user property value. The enterprise_server_user property
func (m *GetConsumedLicenses_users) SetEnterpriseServerUser(value *bool)() {
    m.enterprise_server_user = value
}
// SetEnterpriseServerUserIds sets the enterprise_server_user_ids property value. The enterprise_server_user_ids property
func (m *GetConsumedLicenses_users) SetEnterpriseServerUserIds(value []string)() {
    m.enterprise_server_user_ids = value
}
// SetGithubComEnterpriseRoles sets the github_com_enterprise_roles property value. All enterprise roles for a user.
func (m *GetConsumedLicenses_users) SetGithubComEnterpriseRoles(value []string)() {
    m.github_com_enterprise_roles = value
}
// SetGithubComLogin sets the github_com_login property value. The github_com_login property
func (m *GetConsumedLicenses_users) SetGithubComLogin(value *string)() {
    m.github_com_login = value
}
// SetGithubComMemberRoles sets the github_com_member_roles property value. The github_com_member_roles property
func (m *GetConsumedLicenses_users) SetGithubComMemberRoles(value []string)() {
    m.github_com_member_roles = value
}
// SetGithubComName sets the github_com_name property value. The github_com_name property
func (m *GetConsumedLicenses_users) SetGithubComName(value *string)() {
    m.github_com_name = value
}
// SetGithubComOrgsWithPendingInvites sets the github_com_orgs_with_pending_invites property value. The github_com_orgs_with_pending_invites property
func (m *GetConsumedLicenses_users) SetGithubComOrgsWithPendingInvites(value []string)() {
    m.github_com_orgs_with_pending_invites = value
}
// SetGithubComProfile sets the github_com_profile property value. The github_com_profile property
func (m *GetConsumedLicenses_users) SetGithubComProfile(value *string)() {
    m.github_com_profile = value
}
// SetGithubComSamlNameId sets the github_com_saml_name_id property value. The github_com_saml_name_id property
func (m *GetConsumedLicenses_users) SetGithubComSamlNameId(value *string)() {
    m.github_com_saml_name_id = value
}
// SetGithubComTwoFactorAuth sets the github_com_two_factor_auth property value. The github_com_two_factor_auth property
func (m *GetConsumedLicenses_users) SetGithubComTwoFactorAuth(value *bool)() {
    m.github_com_two_factor_auth = value
}
// SetGithubComUser sets the github_com_user property value. The github_com_user property
func (m *GetConsumedLicenses_users) SetGithubComUser(value *bool)() {
    m.github_com_user = value
}
// SetGithubComVerifiedDomainEmails sets the github_com_verified_domain_emails property value. The github_com_verified_domain_emails property
func (m *GetConsumedLicenses_users) SetGithubComVerifiedDomainEmails(value []string)() {
    m.github_com_verified_domain_emails = value
}
// SetLicenseType sets the license_type property value. The license_type property
func (m *GetConsumedLicenses_users) SetLicenseType(value *string)() {
    m.license_type = value
}
// SetTotalUserAccounts sets the total_user_accounts property value. The total_user_accounts property
func (m *GetConsumedLicenses_users) SetTotalUserAccounts(value *int32)() {
    m.total_user_accounts = value
}
// SetVisualStudioLicenseStatus sets the visual_studio_license_status property value. The visual_studio_license_status property
func (m *GetConsumedLicenses_users) SetVisualStudioLicenseStatus(value *string)() {
    m.visual_studio_license_status = value
}
// SetVisualStudioSubscriptionEmail sets the visual_studio_subscription_email property value. The visual_studio_subscription_email property
func (m *GetConsumedLicenses_users) SetVisualStudioSubscriptionEmail(value *string)() {
    m.visual_studio_subscription_email = value
}
// SetVisualStudioSubscriptionUser sets the visual_studio_subscription_user property value. The visual_studio_subscription_user property
func (m *GetConsumedLicenses_users) SetVisualStudioSubscriptionUser(value *bool)() {
    m.visual_studio_subscription_user = value
}
type GetConsumedLicenses_usersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterpriseServerEmails()([]string)
    GetEnterpriseServerUser()(*bool)
    GetEnterpriseServerUserIds()([]string)
    GetGithubComEnterpriseRoles()([]string)
    GetGithubComLogin()(*string)
    GetGithubComMemberRoles()([]string)
    GetGithubComName()(*string)
    GetGithubComOrgsWithPendingInvites()([]string)
    GetGithubComProfile()(*string)
    GetGithubComSamlNameId()(*string)
    GetGithubComTwoFactorAuth()(*bool)
    GetGithubComUser()(*bool)
    GetGithubComVerifiedDomainEmails()([]string)
    GetLicenseType()(*string)
    GetTotalUserAccounts()(*int32)
    GetVisualStudioLicenseStatus()(*string)
    GetVisualStudioSubscriptionEmail()(*string)
    GetVisualStudioSubscriptionUser()(*bool)
    SetEnterpriseServerEmails(value []string)()
    SetEnterpriseServerUser(value *bool)()
    SetEnterpriseServerUserIds(value []string)()
    SetGithubComEnterpriseRoles(value []string)()
    SetGithubComLogin(value *string)()
    SetGithubComMemberRoles(value []string)()
    SetGithubComName(value *string)()
    SetGithubComOrgsWithPendingInvites(value []string)()
    SetGithubComProfile(value *string)()
    SetGithubComSamlNameId(value *string)()
    SetGithubComTwoFactorAuth(value *bool)()
    SetGithubComUser(value *bool)()
    SetGithubComVerifiedDomainEmails(value []string)()
    SetLicenseType(value *string)()
    SetTotalUserAccounts(value *int32)()
    SetVisualStudioLicenseStatus(value *string)()
    SetVisualStudioSubscriptionEmail(value *string)()
    SetVisualStudioSubscriptionUser(value *bool)()
}
