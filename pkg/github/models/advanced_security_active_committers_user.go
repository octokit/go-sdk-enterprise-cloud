package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AdvancedSecurityActiveCommittersUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The last_pushed_date property
    last_pushed_date *string
    // The last_pushed_email property
    last_pushed_email *string
    // The user_login property
    user_login *string
}
// NewAdvancedSecurityActiveCommittersUser instantiates a new AdvancedSecurityActiveCommittersUser and sets the default values.
func NewAdvancedSecurityActiveCommittersUser()(*AdvancedSecurityActiveCommittersUser) {
    m := &AdvancedSecurityActiveCommittersUser{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAdvancedSecurityActiveCommittersUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAdvancedSecurityActiveCommittersUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdvancedSecurityActiveCommittersUser(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AdvancedSecurityActiveCommittersUser) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AdvancedSecurityActiveCommittersUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["last_pushed_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastPushedDate(val)
        }
        return nil
    }
    res["last_pushed_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastPushedEmail(val)
        }
        return nil
    }
    res["user_login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserLogin(val)
        }
        return nil
    }
    return res
}
// GetLastPushedDate gets the last_pushed_date property value. The last_pushed_date property
// returns a *string when successful
func (m *AdvancedSecurityActiveCommittersUser) GetLastPushedDate()(*string) {
    return m.last_pushed_date
}
// GetLastPushedEmail gets the last_pushed_email property value. The last_pushed_email property
// returns a *string when successful
func (m *AdvancedSecurityActiveCommittersUser) GetLastPushedEmail()(*string) {
    return m.last_pushed_email
}
// GetUserLogin gets the user_login property value. The user_login property
// returns a *string when successful
func (m *AdvancedSecurityActiveCommittersUser) GetUserLogin()(*string) {
    return m.user_login
}
// Serialize serializes information the current object
func (m *AdvancedSecurityActiveCommittersUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("last_pushed_date", m.GetLastPushedDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_pushed_email", m.GetLastPushedEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_login", m.GetUserLogin())
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
func (m *AdvancedSecurityActiveCommittersUser) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLastPushedDate sets the last_pushed_date property value. The last_pushed_date property
func (m *AdvancedSecurityActiveCommittersUser) SetLastPushedDate(value *string)() {
    m.last_pushed_date = value
}
// SetLastPushedEmail sets the last_pushed_email property value. The last_pushed_email property
func (m *AdvancedSecurityActiveCommittersUser) SetLastPushedEmail(value *string)() {
    m.last_pushed_email = value
}
// SetUserLogin sets the user_login property value. The user_login property
func (m *AdvancedSecurityActiveCommittersUser) SetUserLogin(value *string)() {
    m.user_login = value
}
type AdvancedSecurityActiveCommittersUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastPushedDate()(*string)
    GetLastPushedEmail()(*string)
    GetUserLogin()(*string)
    SetLastPushedDate(value *string)()
    SetLastPushedEmail(value *string)()
    SetUserLogin(value *string)()
}
