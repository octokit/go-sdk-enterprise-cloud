package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Users struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether this email address is the primary address.
    primary *bool
    // The type of email address.
    typeEscaped *string
    // The email address.
    value *string
}
// NewUsers instantiates a new Users and sets the default values.
func NewUsers()(*Users) {
    m := &Users{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUsersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Users) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Users) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetPrimary gets the primary property value. Whether this email address is the primary address.
// returns a *bool when successful
func (m *Users) GetPrimary()(*bool) {
    return m.primary
}
// GetTypeEscaped gets the type property value. The type of email address.
// returns a *string when successful
func (m *Users) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValue gets the value property value. The email address.
// returns a *string when successful
func (m *Users) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *Users) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Users) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrimary sets the primary property value. Whether this email address is the primary address.
func (m *Users) SetPrimary(value *bool)() {
    m.primary = value
}
// SetTypeEscaped sets the type property value. The type of email address.
func (m *Users) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The email address.
func (m *Users) SetValue(value *string)() {
    m.value = value
}
type Usersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrimary()(*bool)
    GetTypeEscaped()(*string)
    GetValue()(*string)
    SetPrimary(value *bool)()
    SetTypeEscaped(value *string)()
    SetValue(value *string)()
}
