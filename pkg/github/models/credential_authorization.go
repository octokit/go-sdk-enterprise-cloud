package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialAuthorization credential Authorization
type CredentialAuthorization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The expiry for the token. This will only be present when the credential is a token.
    authorized_credential_expires_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ID of the underlying token that was authorized by the user. This will remain unchanged across authorizations of the token.
    authorized_credential_id *int32
    // The note given to the token. This will only be present when the credential is a token.
    authorized_credential_note *string
    // The title given to the ssh key. This will only be present when the credential is an ssh key.
    authorized_credential_title *string
    // Date when the credential was last accessed. May be null if it was never accessed
    credential_accessed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date when the credential was authorized for use.
    credential_authorized_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique identifier for the authorization of the credential. Use this to revoke authorization of the underlying token or key.
    credential_id *int32
    // Human-readable description of the credential type.
    credential_type *string
    // Unique string to distinguish the credential. Only included in responses with credential_type of SSH Key.
    fingerprint *string
    // User login that owns the underlying credential.
    login *string
    // List of oauth scopes the token has been granted.
    scopes []string
    // Last eight characters of the credential. Only included in responses with credential_type of personal access token.
    token_last_eight *string
}
// NewCredentialAuthorization instantiates a new CredentialAuthorization and sets the default values.
func NewCredentialAuthorization()(*CredentialAuthorization) {
    m := &CredentialAuthorization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCredentialAuthorizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCredentialAuthorizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialAuthorization(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CredentialAuthorization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorizedCredentialExpiresAt gets the authorized_credential_expires_at property value. The expiry for the token. This will only be present when the credential is a token.
// returns a *Time when successful
func (m *CredentialAuthorization) GetAuthorizedCredentialExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.authorized_credential_expires_at
}
// GetAuthorizedCredentialId gets the authorized_credential_id property value. The ID of the underlying token that was authorized by the user. This will remain unchanged across authorizations of the token.
// returns a *int32 when successful
func (m *CredentialAuthorization) GetAuthorizedCredentialId()(*int32) {
    return m.authorized_credential_id
}
// GetAuthorizedCredentialNote gets the authorized_credential_note property value. The note given to the token. This will only be present when the credential is a token.
// returns a *string when successful
func (m *CredentialAuthorization) GetAuthorizedCredentialNote()(*string) {
    return m.authorized_credential_note
}
// GetAuthorizedCredentialTitle gets the authorized_credential_title property value. The title given to the ssh key. This will only be present when the credential is an ssh key.
// returns a *string when successful
func (m *CredentialAuthorization) GetAuthorizedCredentialTitle()(*string) {
    return m.authorized_credential_title
}
// GetCredentialAccessedAt gets the credential_accessed_at property value. Date when the credential was last accessed. May be null if it was never accessed
// returns a *Time when successful
func (m *CredentialAuthorization) GetCredentialAccessedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.credential_accessed_at
}
// GetCredentialAuthorizedAt gets the credential_authorized_at property value. Date when the credential was authorized for use.
// returns a *Time when successful
func (m *CredentialAuthorization) GetCredentialAuthorizedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.credential_authorized_at
}
// GetCredentialId gets the credential_id property value. Unique identifier for the authorization of the credential. Use this to revoke authorization of the underlying token or key.
// returns a *int32 when successful
func (m *CredentialAuthorization) GetCredentialId()(*int32) {
    return m.credential_id
}
// GetCredentialType gets the credential_type property value. Human-readable description of the credential type.
// returns a *string when successful
func (m *CredentialAuthorization) GetCredentialType()(*string) {
    return m.credential_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CredentialAuthorization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authorized_credential_expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedCredentialExpiresAt(val)
        }
        return nil
    }
    res["authorized_credential_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedCredentialId(val)
        }
        return nil
    }
    res["authorized_credential_note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedCredentialNote(val)
        }
        return nil
    }
    res["authorized_credential_title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedCredentialTitle(val)
        }
        return nil
    }
    res["credential_accessed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialAccessedAt(val)
        }
        return nil
    }
    res["credential_authorized_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialAuthorizedAt(val)
        }
        return nil
    }
    res["credential_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialId(val)
        }
        return nil
    }
    res["credential_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialType(val)
        }
        return nil
    }
    res["fingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprint(val)
        }
        return nil
    }
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
        }
        return nil
    }
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetScopes(res)
        }
        return nil
    }
    res["token_last_eight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenLastEight(val)
        }
        return nil
    }
    return res
}
// GetFingerprint gets the fingerprint property value. Unique string to distinguish the credential. Only included in responses with credential_type of SSH Key.
// returns a *string when successful
func (m *CredentialAuthorization) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetLogin gets the login property value. User login that owns the underlying credential.
// returns a *string when successful
func (m *CredentialAuthorization) GetLogin()(*string) {
    return m.login
}
// GetScopes gets the scopes property value. List of oauth scopes the token has been granted.
// returns a []string when successful
func (m *CredentialAuthorization) GetScopes()([]string) {
    return m.scopes
}
// GetTokenLastEight gets the token_last_eight property value. Last eight characters of the credential. Only included in responses with credential_type of personal access token.
// returns a *string when successful
func (m *CredentialAuthorization) GetTokenLastEight()(*string) {
    return m.token_last_eight
}
// Serialize serializes information the current object
func (m *CredentialAuthorization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("authorized_credential_expires_at", m.GetAuthorizedCredentialExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("authorized_credential_id", m.GetAuthorizedCredentialId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authorized_credential_note", m.GetAuthorizedCredentialNote())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authorized_credential_title", m.GetAuthorizedCredentialTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("credential_accessed_at", m.GetCredentialAccessedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("credential_authorized_at", m.GetCredentialAuthorizedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("credential_id", m.GetCredentialId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("credential_type", m.GetCredentialType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fingerprint", m.GetFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    if m.GetScopes() != nil {
        err := writer.WriteCollectionOfStringValues("scopes", m.GetScopes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token_last_eight", m.GetTokenLastEight())
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
func (m *CredentialAuthorization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorizedCredentialExpiresAt sets the authorized_credential_expires_at property value. The expiry for the token. This will only be present when the credential is a token.
func (m *CredentialAuthorization) SetAuthorizedCredentialExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.authorized_credential_expires_at = value
}
// SetAuthorizedCredentialId sets the authorized_credential_id property value. The ID of the underlying token that was authorized by the user. This will remain unchanged across authorizations of the token.
func (m *CredentialAuthorization) SetAuthorizedCredentialId(value *int32)() {
    m.authorized_credential_id = value
}
// SetAuthorizedCredentialNote sets the authorized_credential_note property value. The note given to the token. This will only be present when the credential is a token.
func (m *CredentialAuthorization) SetAuthorizedCredentialNote(value *string)() {
    m.authorized_credential_note = value
}
// SetAuthorizedCredentialTitle sets the authorized_credential_title property value. The title given to the ssh key. This will only be present when the credential is an ssh key.
func (m *CredentialAuthorization) SetAuthorizedCredentialTitle(value *string)() {
    m.authorized_credential_title = value
}
// SetCredentialAccessedAt sets the credential_accessed_at property value. Date when the credential was last accessed. May be null if it was never accessed
func (m *CredentialAuthorization) SetCredentialAccessedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.credential_accessed_at = value
}
// SetCredentialAuthorizedAt sets the credential_authorized_at property value. Date when the credential was authorized for use.
func (m *CredentialAuthorization) SetCredentialAuthorizedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.credential_authorized_at = value
}
// SetCredentialId sets the credential_id property value. Unique identifier for the authorization of the credential. Use this to revoke authorization of the underlying token or key.
func (m *CredentialAuthorization) SetCredentialId(value *int32)() {
    m.credential_id = value
}
// SetCredentialType sets the credential_type property value. Human-readable description of the credential type.
func (m *CredentialAuthorization) SetCredentialType(value *string)() {
    m.credential_type = value
}
// SetFingerprint sets the fingerprint property value. Unique string to distinguish the credential. Only included in responses with credential_type of SSH Key.
func (m *CredentialAuthorization) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetLogin sets the login property value. User login that owns the underlying credential.
func (m *CredentialAuthorization) SetLogin(value *string)() {
    m.login = value
}
// SetScopes sets the scopes property value. List of oauth scopes the token has been granted.
func (m *CredentialAuthorization) SetScopes(value []string)() {
    m.scopes = value
}
// SetTokenLastEight sets the token_last_eight property value. Last eight characters of the credential. Only included in responses with credential_type of personal access token.
func (m *CredentialAuthorization) SetTokenLastEight(value *string)() {
    m.token_last_eight = value
}
type CredentialAuthorizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizedCredentialExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAuthorizedCredentialId()(*int32)
    GetAuthorizedCredentialNote()(*string)
    GetAuthorizedCredentialTitle()(*string)
    GetCredentialAccessedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCredentialAuthorizedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCredentialId()(*int32)
    GetCredentialType()(*string)
    GetFingerprint()(*string)
    GetLogin()(*string)
    GetScopes()([]string)
    GetTokenLastEight()(*string)
    SetAuthorizedCredentialExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAuthorizedCredentialId(value *int32)()
    SetAuthorizedCredentialNote(value *string)()
    SetAuthorizedCredentialTitle(value *string)()
    SetCredentialAccessedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCredentialAuthorizedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCredentialId(value *int32)()
    SetCredentialType(value *string)()
    SetFingerprint(value *string)()
    SetLogin(value *string)()
    SetScopes(value []string)()
    SetTokenLastEight(value *string)()
}
