package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScimUser_name struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The familyName property
    familyName *string
    // The formatted property
    formatted *string
    // The givenName property
    givenName *string
}
// NewScimUser_name instantiates a new ScimUser_name and sets the default values.
func NewScimUser_name()(*ScimUser_name) {
    m := &ScimUser_name{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimUser_nameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimUser_nameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimUser_name(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimUser_name) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFamilyName gets the familyName property value. The familyName property
// returns a *string when successful
func (m *ScimUser_name) GetFamilyName()(*string) {
    return m.familyName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimUser_name) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["familyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFamilyName(val)
        }
        return nil
    }
    res["formatted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormatted(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    return res
}
// GetFormatted gets the formatted property value. The formatted property
// returns a *string when successful
func (m *ScimUser_name) GetFormatted()(*string) {
    return m.formatted
}
// GetGivenName gets the givenName property value. The givenName property
// returns a *string when successful
func (m *ScimUser_name) GetGivenName()(*string) {
    return m.givenName
}
// Serialize serializes information the current object
func (m *ScimUser_name) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("familyName", m.GetFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formatted", m.GetFormatted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("givenName", m.GetGivenName())
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
func (m *ScimUser_name) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFamilyName sets the familyName property value. The familyName property
func (m *ScimUser_name) SetFamilyName(value *string)() {
    m.familyName = value
}
// SetFormatted sets the formatted property value. The formatted property
func (m *ScimUser_name) SetFormatted(value *string)() {
    m.formatted = value
}
// SetGivenName sets the givenName property value. The givenName property
func (m *ScimUser_name) SetGivenName(value *string)() {
    m.givenName = value
}
type ScimUser_nameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFamilyName()(*string)
    GetFormatted()(*string)
    GetGivenName()(*string)
    SetFamilyName(value *string)()
    SetFormatted(value *string)()
    SetGivenName(value *string)()
}
