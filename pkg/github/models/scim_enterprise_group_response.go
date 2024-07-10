package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScimEnterpriseGroupResponse struct {
    GroupResponse
    // The internally generated id for the group object.
    id *string
    // The metadata associated with the creation/updates to the user.
    meta Metaable
}
// NewScimEnterpriseGroupResponse instantiates a new ScimEnterpriseGroupResponse and sets the default values.
func NewScimEnterpriseGroupResponse()(*ScimEnterpriseGroupResponse) {
    m := &ScimEnterpriseGroupResponse{
        GroupResponse: *NewGroupResponse(),
    }
    return m
}
// CreateScimEnterpriseGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimEnterpriseGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimEnterpriseGroupResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimEnterpriseGroupResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupResponse.GetFieldDeserializers()
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
// GetId gets the id property value. The internally generated id for the group object.
// returns a *string when successful
func (m *ScimEnterpriseGroupResponse) GetId()(*string) {
    return m.id
}
// GetMeta gets the meta property value. The metadata associated with the creation/updates to the user.
// returns a Metaable when successful
func (m *ScimEnterpriseGroupResponse) GetMeta()(Metaable) {
    return m.meta
}
// Serialize serializes information the current object
func (m *ScimEnterpriseGroupResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupResponse.Serialize(writer)
    if err != nil {
        return err
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
// SetId sets the id property value. The internally generated id for the group object.
func (m *ScimEnterpriseGroupResponse) SetId(value *string)() {
    m.id = value
}
// SetMeta sets the meta property value. The metadata associated with the creation/updates to the user.
func (m *ScimEnterpriseGroupResponse) SetMeta(value Metaable)() {
    m.meta = value
}
type ScimEnterpriseGroupResponseable interface {
    GroupResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetMeta()(Metaable)
    SetId(value *string)()
    SetMeta(value Metaable)()
}
