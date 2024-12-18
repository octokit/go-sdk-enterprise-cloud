package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserName struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The family name of the user.
    familyName *string
    // The full name, including all middle names, titles, and suffixes as appropriate, formatted for display.
    formatted *string
    // The given name of the user.
    givenName *string
    // The middle name(s) of the user.
    middleName *string
}
// NewUserName instantiates a new UserName and sets the default values.
func NewUserName()(*UserName) {
    m := &UserName{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserName(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserName) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFamilyName gets the familyName property value. The family name of the user.
// returns a *string when successful
func (m *UserName) GetFamilyName()(*string) {
    return m.familyName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserName) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["middleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    return res
}
// GetFormatted gets the formatted property value. The full name, including all middle names, titles, and suffixes as appropriate, formatted for display.
// returns a *string when successful
func (m *UserName) GetFormatted()(*string) {
    return m.formatted
}
// GetGivenName gets the givenName property value. The given name of the user.
// returns a *string when successful
func (m *UserName) GetGivenName()(*string) {
    return m.givenName
}
// GetMiddleName gets the middleName property value. The middle name(s) of the user.
// returns a *string when successful
func (m *UserName) GetMiddleName()(*string) {
    return m.middleName
}
// Serialize serializes information the current object
func (m *UserName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("middleName", m.GetMiddleName())
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
func (m *UserName) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFamilyName sets the familyName property value. The family name of the user.
func (m *UserName) SetFamilyName(value *string)() {
    m.familyName = value
}
// SetFormatted sets the formatted property value. The full name, including all middle names, titles, and suffixes as appropriate, formatted for display.
func (m *UserName) SetFormatted(value *string)() {
    m.formatted = value
}
// SetGivenName sets the givenName property value. The given name of the user.
func (m *UserName) SetGivenName(value *string)() {
    m.givenName = value
}
// SetMiddleName sets the middleName property value. The middle name(s) of the user.
func (m *UserName) SetMiddleName(value *string)() {
    m.middleName = value
}
type UserNameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFamilyName()(*string)
    GetFormatted()(*string)
    GetGivenName()(*string)
    GetMiddleName()(*string)
    SetFamilyName(value *string)()
    SetFormatted(value *string)()
    SetGivenName(value *string)()
    SetMiddleName(value *string)()
}
