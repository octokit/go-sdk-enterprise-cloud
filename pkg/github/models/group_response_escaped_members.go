package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupResponse_members struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display name associated with the member
    display *string
    // The Ref property
    ref *string
    // The local unique identifier for the member
    value *string
}
// NewGroupResponse_members instantiates a new GroupResponse_members and sets the default values.
func NewGroupResponse_members()(*GroupResponse_members) {
    m := &GroupResponse_members{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupResponse_membersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupResponse_membersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupResponse_members(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GroupResponse_members) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplay gets the display property value. The display name associated with the member
// returns a *string when successful
func (m *GroupResponse_members) GetDisplay()(*string) {
    return m.display
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupResponse_members) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["display"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplay(val)
        }
        return nil
    }
    res["$ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetRef gets the $ref property value. The Ref property
// returns a *string when successful
func (m *GroupResponse_members) GetRef()(*string) {
    return m.ref
}
// GetValue gets the value property value. The local unique identifier for the member
// returns a *string when successful
func (m *GroupResponse_members) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupResponse_members) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("display", m.GetDisplay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("$ref", m.GetRef())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *GroupResponse_members) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplay sets the display property value. The display name associated with the member
func (m *GroupResponse_members) SetDisplay(value *string)() {
    m.display = value
}
// SetRef sets the $ref property value. The Ref property
func (m *GroupResponse_members) SetRef(value *string)() {
    m.ref = value
}
// SetValue sets the value property value. The local unique identifier for the member
func (m *GroupResponse_members) SetValue(value *string)() {
    m.value = value
}
type GroupResponse_membersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplay()(*string)
    GetRef()(*string)
    GetValue()(*string)
    SetDisplay(value *string)()
    SetRef(value *string)()
    SetValue(value *string)()
}
