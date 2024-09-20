package scim

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V2OrganizationsItemUsersItemWithScim_user_PutRequestBody struct {
    // The active property
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the user, suitable for display to end-users
    displayName *string
    // user emails
    emails []V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable
    // The externalId property
    externalId *string
    // The groups property
    groups []string
    // The name property
    name V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameable
    // The schemas property
    schemas []string
    // Configured by the admin. Could be an email, login, or username
    userName *string
}
// NewV2OrganizationsItemUsersItemWithScim_user_PutRequestBody instantiates a new V2OrganizationsItemUsersItemWithScim_user_PutRequestBody and sets the default values.
func NewV2OrganizationsItemUsersItemWithScim_user_PutRequestBody()(*V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) {
    m := &V2OrganizationsItemUsersItemWithScim_user_PutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV2OrganizationsItemUsersItemWithScim_user_PutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV2OrganizationsItemUsersItemWithScim_user_PutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV2OrganizationsItemUsersItemWithScim_user_PutRequestBody(), nil
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The name of the user, suitable for display to end-users
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmails gets the emails property value. user emails
// returns a []V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetEmails()([]V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable) {
    return m.emails
}
// GetExternalId gets the externalId property value. The externalId property
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable)
                }
            }
            m.SetEmails(res)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetGroups(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val.(V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameable))
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSchemas(res)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
// returns a []string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetGroups()([]string) {
    return m.groups
}
// GetName gets the name property value. The name property
// returns a V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameable when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetName()(V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameable) {
    return m.name
}
// GetSchemas gets the schemas property value. The schemas property
// returns a []string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetSchemas()([]string) {
    return m.schemas
}
// GetUserName gets the userName property value. Configured by the admin. Could be an email, login, or username
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEmails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmails()))
        for i, v := range m.GetEmails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("emails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        err := writer.WriteCollectionOfStringValues("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        err := writer.WriteCollectionOfStringValues("schemas", m.GetSchemas())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userName", m.GetUserName())
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
// SetActive sets the active property value. The active property
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The name of the user, suitable for display to end-users
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmails sets the emails property value. user emails
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetEmails(value []V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable)() {
    m.emails = value
}
// SetExternalId sets the externalId property value. The externalId property
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetExternalId(value *string)() {
    m.externalId = value
}
// SetGroups sets the groups property value. The groups property
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetGroups(value []string)() {
    m.groups = value
}
// SetName sets the name property value. The name property
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetName(value V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameable)() {
    m.name = value
}
// SetSchemas sets the schemas property value. The schemas property
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetSchemas(value []string)() {
    m.schemas = value
}
// SetUserName sets the userName property value. Configured by the admin. Could be an email, login, or username
func (m *V2OrganizationsItemUsersItemWithScim_user_PutRequestBody) SetUserName(value *string)() {
    m.userName = value
}
type V2OrganizationsItemUsersItemWithScim_user_PutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetDisplayName()(*string)
    GetEmails()([]V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable)
    GetExternalId()(*string)
    GetGroups()([]string)
    GetName()(V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameable)
    GetSchemas()([]string)
    GetUserName()(*string)
    SetActive(value *bool)()
    SetDisplayName(value *string)()
    SetEmails(value []V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_emailsable)()
    SetExternalId(value *string)()
    SetGroups(value []string)()
    SetName(value V2OrganizationsItemUsersItemWithScim_user_PutRequestBody_nameable)()
    SetSchemas(value []string)()
    SetUserName(value *string)()
}
