package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScimUser sCIM /Users provisioning endpoints
type ScimUser struct {
    // The active status of the User.
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the user, suitable for display to end-users
    displayName *string
    // user emails
    emails []ScimUser_emailsable
    // The ID of the User.
    externalId *string
    // associated groups
    groups []ScimUser_groupsable
    // Unique identifier of an external identity
    id *string
    // The meta property
    meta ScimUser_metaable
    // The name property
    name ScimUser_nameable
    // Set of operations to be performed
    operations []ScimUser_operationsable
    // The ID of the organization.
    organization_id *int32
    // The roles property
    roles []ScimUser_rolesable
    // SCIM schema used.
    schemas []string
    // Configured by the admin. Could be an email, login, or username
    userName *string
}
// NewScimUser instantiates a new ScimUser and sets the default values.
func NewScimUser()(*ScimUser) {
    m := &ScimUser{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimUser(), nil
}
// GetActive gets the active property value. The active status of the User.
// returns a *bool when successful
func (m *ScimUser) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimUser) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The name of the user, suitable for display to end-users
// returns a *string when successful
func (m *ScimUser) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmails gets the emails property value. user emails
// returns a []ScimUser_emailsable when successful
func (m *ScimUser) GetEmails()([]ScimUser_emailsable) {
    return m.emails
}
// GetExternalId gets the externalId property value. The ID of the User.
// returns a *string when successful
func (m *ScimUser) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateScimUser_emailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimUser_emailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScimUser_emailsable)
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
        val, err := n.GetCollectionOfObjectValues(CreateScimUser_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimUser_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScimUser_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["meta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScimUser_metaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeta(val.(ScimUser_metaable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScimUser_nameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val.(ScimUser_nameable))
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScimUser_operationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimUser_operationsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScimUser_operationsable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["organization_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationId(val)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScimUser_rolesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimUser_rolesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScimUser_rolesable)
                }
            }
            m.SetRoles(res)
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
// GetGroups gets the groups property value. associated groups
// returns a []ScimUser_groupsable when successful
func (m *ScimUser) GetGroups()([]ScimUser_groupsable) {
    return m.groups
}
// GetId gets the id property value. Unique identifier of an external identity
// returns a *string when successful
func (m *ScimUser) GetId()(*string) {
    return m.id
}
// GetMeta gets the meta property value. The meta property
// returns a ScimUser_metaable when successful
func (m *ScimUser) GetMeta()(ScimUser_metaable) {
    return m.meta
}
// GetName gets the name property value. The name property
// returns a ScimUser_nameable when successful
func (m *ScimUser) GetName()(ScimUser_nameable) {
    return m.name
}
// GetOperations gets the operations property value. Set of operations to be performed
// returns a []ScimUser_operationsable when successful
func (m *ScimUser) GetOperations()([]ScimUser_operationsable) {
    return m.operations
}
// GetOrganizationId gets the organization_id property value. The ID of the organization.
// returns a *int32 when successful
func (m *ScimUser) GetOrganizationId()(*int32) {
    return m.organization_id
}
// GetRoles gets the roles property value. The roles property
// returns a []ScimUser_rolesable when successful
func (m *ScimUser) GetRoles()([]ScimUser_rolesable) {
    return m.roles
}
// GetSchemas gets the schemas property value. SCIM schema used.
// returns a []string when successful
func (m *ScimUser) GetSchemas()([]string) {
    return m.schemas
}
// GetUserName gets the userName property value. Configured by the admin. Could be an email, login, or username
// returns a *string when successful
func (m *ScimUser) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *ScimUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("meta", m.GetMeta())
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
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("organization_id", m.GetOrganizationId())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoles()))
        for i, v := range m.GetRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("roles", cast)
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
// SetActive sets the active property value. The active status of the User.
func (m *ScimUser) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScimUser) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The name of the user, suitable for display to end-users
func (m *ScimUser) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmails sets the emails property value. user emails
func (m *ScimUser) SetEmails(value []ScimUser_emailsable)() {
    m.emails = value
}
// SetExternalId sets the externalId property value. The ID of the User.
func (m *ScimUser) SetExternalId(value *string)() {
    m.externalId = value
}
// SetGroups sets the groups property value. associated groups
func (m *ScimUser) SetGroups(value []ScimUser_groupsable)() {
    m.groups = value
}
// SetId sets the id property value. Unique identifier of an external identity
func (m *ScimUser) SetId(value *string)() {
    m.id = value
}
// SetMeta sets the meta property value. The meta property
func (m *ScimUser) SetMeta(value ScimUser_metaable)() {
    m.meta = value
}
// SetName sets the name property value. The name property
func (m *ScimUser) SetName(value ScimUser_nameable)() {
    m.name = value
}
// SetOperations sets the operations property value. Set of operations to be performed
func (m *ScimUser) SetOperations(value []ScimUser_operationsable)() {
    m.operations = value
}
// SetOrganizationId sets the organization_id property value. The ID of the organization.
func (m *ScimUser) SetOrganizationId(value *int32)() {
    m.organization_id = value
}
// SetRoles sets the roles property value. The roles property
func (m *ScimUser) SetRoles(value []ScimUser_rolesable)() {
    m.roles = value
}
// SetSchemas sets the schemas property value. SCIM schema used.
func (m *ScimUser) SetSchemas(value []string)() {
    m.schemas = value
}
// SetUserName sets the userName property value. Configured by the admin. Could be an email, login, or username
func (m *ScimUser) SetUserName(value *string)() {
    m.userName = value
}
type ScimUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetDisplayName()(*string)
    GetEmails()([]ScimUser_emailsable)
    GetExternalId()(*string)
    GetGroups()([]ScimUser_groupsable)
    GetId()(*string)
    GetMeta()(ScimUser_metaable)
    GetName()(ScimUser_nameable)
    GetOperations()([]ScimUser_operationsable)
    GetOrganizationId()(*int32)
    GetRoles()([]ScimUser_rolesable)
    GetSchemas()([]string)
    GetUserName()(*string)
    SetActive(value *bool)()
    SetDisplayName(value *string)()
    SetEmails(value []ScimUser_emailsable)()
    SetExternalId(value *string)()
    SetGroups(value []ScimUser_groupsable)()
    SetId(value *string)()
    SetMeta(value ScimUser_metaable)()
    SetName(value ScimUser_nameable)()
    SetOperations(value []ScimUser_operationsable)()
    SetOrganizationId(value *int32)()
    SetRoles(value []ScimUser_rolesable)()
    SetSchemas(value []string)()
    SetUserName(value *string)()
}
