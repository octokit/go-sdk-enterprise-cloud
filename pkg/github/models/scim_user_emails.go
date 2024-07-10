package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScimUser_emails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The primary property
    primary *bool
    // The type property
    typeEscaped *string
    // The value property
    value *string
}
// NewScimUser_emails instantiates a new ScimUser_emails and sets the default values.
func NewScimUser_emails()(*ScimUser_emails) {
    m := &ScimUser_emails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimUser_emailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimUser_emailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimUser_emails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimUser_emails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimUser_emails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ScimUser_emails) GetPrimary()(*bool) {
    return m.primary
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *ScimUser_emails) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *ScimUser_emails) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ScimUser_emails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ScimUser_emails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrimary sets the primary property value. The primary property
func (m *ScimUser_emails) SetPrimary(value *bool)() {
    m.primary = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ScimUser_emails) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The value property
func (m *ScimUser_emails) SetValue(value *string)() {
    m.value = value
}
type ScimUser_emailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrimary()(*bool)
    GetTypeEscaped()(*string)
    GetValue()(*string)
    SetPrimary(value *bool)()
    SetTypeEscaped(value *string)()
    SetValue(value *string)()
}
