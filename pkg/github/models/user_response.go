package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserResponse struct {
    // Whether the user active in the IdP.
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A human-readable name for the user.
    displayName *string
    // The emails for the user.
    emails []Usersable
    // A unique identifier for the resource as defined by the provisioning client.
    externalId *string
    // The name property
    name UserNameResponseable
    // The roles assigned to the user.
    roles []Usersable
    // The URIs that are used to indicate the namespaces of the SCIM schemas.
    schemas []UserResponse_schemas
    // The username for the user.
    userName *string
}
// NewUserResponse instantiates a new UserResponse and sets the default values.
func NewUserResponse()(*UserResponse) {
    m := &UserResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserResponse(), nil
}
// GetActive gets the active property value. Whether the user active in the IdP.
// returns a *bool when successful
func (m *UserResponse) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A human-readable name for the user.
// returns a *string when successful
func (m *UserResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmails gets the emails property value. The emails for the user.
// returns a []Usersable when successful
func (m *UserResponse) GetEmails()([]Usersable) {
    return m.emails
}
// GetExternalId gets the externalId property value. A unique identifier for the resource as defined by the provisioning client.
// returns a *string when successful
func (m *UserResponse) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateUsersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Usersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Usersable)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserNameResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val.(UserNameResponseable))
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUsersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Usersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Usersable)
                }
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseUserResponse_schemas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserResponse_schemas, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*UserResponse_schemas))
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
// GetName gets the name property value. The name property
// returns a UserNameResponseable when successful
func (m *UserResponse) GetName()(UserNameResponseable) {
    return m.name
}
// GetRoles gets the roles property value. The roles assigned to the user.
// returns a []Usersable when successful
func (m *UserResponse) GetRoles()([]Usersable) {
    return m.roles
}
// GetSchemas gets the schemas property value. The URIs that are used to indicate the namespaces of the SCIM schemas.
// returns a []UserResponse_schemas when successful
func (m *UserResponse) GetSchemas()([]UserResponse_schemas) {
    return m.schemas
}
// GetUserName gets the userName property value. The username for the user.
// returns a *string when successful
func (m *UserResponse) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *UserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteObjectValue("name", m.GetName())
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
        err := writer.WriteCollectionOfStringValues("schemas", SerializeUserResponse_schemas(m.GetSchemas()))
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
// SetActive sets the active property value. Whether the user active in the IdP.
func (m *UserResponse) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A human-readable name for the user.
func (m *UserResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmails sets the emails property value. The emails for the user.
func (m *UserResponse) SetEmails(value []Usersable)() {
    m.emails = value
}
// SetExternalId sets the externalId property value. A unique identifier for the resource as defined by the provisioning client.
func (m *UserResponse) SetExternalId(value *string)() {
    m.externalId = value
}
// SetName sets the name property value. The name property
func (m *UserResponse) SetName(value UserNameResponseable)() {
    m.name = value
}
// SetRoles sets the roles property value. The roles assigned to the user.
func (m *UserResponse) SetRoles(value []Usersable)() {
    m.roles = value
}
// SetSchemas sets the schemas property value. The URIs that are used to indicate the namespaces of the SCIM schemas.
func (m *UserResponse) SetSchemas(value []UserResponse_schemas)() {
    m.schemas = value
}
// SetUserName sets the userName property value. The username for the user.
func (m *UserResponse) SetUserName(value *string)() {
    m.userName = value
}
type UserResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetDisplayName()(*string)
    GetEmails()([]Usersable)
    GetExternalId()(*string)
    GetName()(UserNameResponseable)
    GetRoles()([]Usersable)
    GetSchemas()([]UserResponse_schemas)
    GetUserName()(*string)
    SetActive(value *bool)()
    SetDisplayName(value *string)()
    SetEmails(value []Usersable)()
    SetExternalId(value *string)()
    SetName(value UserNameResponseable)()
    SetRoles(value []Usersable)()
    SetSchemas(value []UserResponse_schemas)()
    SetUserName(value *string)()
}
