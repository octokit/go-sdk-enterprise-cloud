package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScimEnterpriseUserResponse struct {
    UserResponse
    // Provisioned SCIM groups that the user is a member of.
    groups []ScimEnterpriseUserResponse_groupsable
    // The internally generated id for the user object.
    id *string
    // The metadata associated with the creation/updates to the user.
    meta Metaable
}
// NewScimEnterpriseUserResponse instantiates a new ScimEnterpriseUserResponse and sets the default values.
func NewScimEnterpriseUserResponse()(*ScimEnterpriseUserResponse) {
    m := &ScimEnterpriseUserResponse{
        UserResponse: *NewUserResponse(),
    }
    return m
}
// CreateScimEnterpriseUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimEnterpriseUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimEnterpriseUserResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimEnterpriseUserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserResponse.GetFieldDeserializers()
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScimEnterpriseUserResponse_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimEnterpriseUserResponse_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScimEnterpriseUserResponse_groupsable)
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
        val, err := n.GetObjectValue(CreateMetaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeta(val.(Metaable))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Provisioned SCIM groups that the user is a member of.
// returns a []ScimEnterpriseUserResponse_groupsable when successful
func (m *ScimEnterpriseUserResponse) GetGroups()([]ScimEnterpriseUserResponse_groupsable) {
    return m.groups
}
// GetId gets the id property value. The internally generated id for the user object.
// returns a *string when successful
func (m *ScimEnterpriseUserResponse) GetId()(*string) {
    return m.id
}
// GetMeta gets the meta property value. The metadata associated with the creation/updates to the user.
// returns a Metaable when successful
func (m *ScimEnterpriseUserResponse) GetMeta()(Metaable) {
    return m.meta
}
// Serialize serializes information the current object
func (m *ScimEnterpriseUserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meta", m.GetMeta())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroups sets the groups property value. Provisioned SCIM groups that the user is a member of.
func (m *ScimEnterpriseUserResponse) SetGroups(value []ScimEnterpriseUserResponse_groupsable)() {
    m.groups = value
}
// SetId sets the id property value. The internally generated id for the user object.
func (m *ScimEnterpriseUserResponse) SetId(value *string)() {
    m.id = value
}
// SetMeta sets the meta property value. The metadata associated with the creation/updates to the user.
func (m *ScimEnterpriseUserResponse) SetMeta(value Metaable)() {
    m.meta = value
}
type ScimEnterpriseUserResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserResponseable
    GetGroups()([]ScimEnterpriseUserResponse_groupsable)
    GetId()(*string)
    GetMeta()(Metaable)
    SetGroups(value []ScimEnterpriseUserResponse_groupsable)()
    SetId(value *string)()
    SetMeta(value Metaable)()
}
