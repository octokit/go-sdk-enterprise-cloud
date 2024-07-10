package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScimUser_roles struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display property
    display *string
    // The primary property
    primary *bool
    // The type property
    typeEscaped *string
    // The value property
    value *string
}
// NewScimUser_roles instantiates a new ScimUser_roles and sets the default values.
func NewScimUser_roles()(*ScimUser_roles) {
    m := &ScimUser_roles{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimUser_rolesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimUser_rolesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimUser_roles(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimUser_roles) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplay gets the display property value. The display property
// returns a *string when successful
func (m *ScimUser_roles) GetDisplay()(*string) {
    return m.display
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimUser_roles) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["primary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimary(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
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
// GetPrimary gets the primary property value. The primary property
// returns a *bool when successful
func (m *ScimUser_roles) GetPrimary()(*bool) {
    return m.primary
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *ScimUser_roles) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *ScimUser_roles) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ScimUser_roles) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("display", m.GetDisplay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("primary", m.GetPrimary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *ScimUser_roles) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplay sets the display property value. The display property
func (m *ScimUser_roles) SetDisplay(value *string)() {
    m.display = value
}
// SetPrimary sets the primary property value. The primary property
func (m *ScimUser_roles) SetPrimary(value *bool)() {
    m.primary = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ScimUser_roles) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The value property
func (m *ScimUser_roles) SetValue(value *string)() {
    m.value = value
}
type ScimUser_rolesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplay()(*string)
    GetPrimary()(*bool)
    GetTypeEscaped()(*string)
    GetValue()(*string)
    SetDisplay(value *string)()
    SetPrimary(value *bool)()
    SetTypeEscaped(value *string)()
    SetValue(value *string)()
}
