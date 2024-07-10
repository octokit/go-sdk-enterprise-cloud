package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalGroup_members struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An email attached to a user
    member_email *string
    // The internal user ID of the identity
    member_id *int32
    // The handle/login for the user
    member_login *string
    // The user display name/profile name
    member_name *string
}
// NewExternalGroup_members instantiates a new ExternalGroup_members and sets the default values.
func NewExternalGroup_members()(*ExternalGroup_members) {
    m := &ExternalGroup_members{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalGroup_membersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalGroup_membersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroup_members(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalGroup_members) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalGroup_members) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["member_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEmail(val)
        }
        return nil
    }
    res["member_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberId(val)
        }
        return nil
    }
    res["member_login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberLogin(val)
        }
        return nil
    }
    res["member_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberName(val)
        }
        return nil
    }
    return res
}
// GetMemberEmail gets the member_email property value. An email attached to a user
// returns a *string when successful
func (m *ExternalGroup_members) GetMemberEmail()(*string) {
    return m.member_email
}
// GetMemberId gets the member_id property value. The internal user ID of the identity
// returns a *int32 when successful
func (m *ExternalGroup_members) GetMemberId()(*int32) {
    return m.member_id
}
// GetMemberLogin gets the member_login property value. The handle/login for the user
// returns a *string when successful
func (m *ExternalGroup_members) GetMemberLogin()(*string) {
    return m.member_login
}
// GetMemberName gets the member_name property value. The user display name/profile name
// returns a *string when successful
func (m *ExternalGroup_members) GetMemberName()(*string) {
    return m.member_name
}
// Serialize serializes information the current object
func (m *ExternalGroup_members) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("member_email", m.GetMemberEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("member_id", m.GetMemberId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("member_login", m.GetMemberLogin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("member_name", m.GetMemberName())
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
func (m *ExternalGroup_members) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMemberEmail sets the member_email property value. An email attached to a user
func (m *ExternalGroup_members) SetMemberEmail(value *string)() {
    m.member_email = value
}
// SetMemberId sets the member_id property value. The internal user ID of the identity
func (m *ExternalGroup_members) SetMemberId(value *int32)() {
    m.member_id = value
}
// SetMemberLogin sets the member_login property value. The handle/login for the user
func (m *ExternalGroup_members) SetMemberLogin(value *string)() {
    m.member_login = value
}
// SetMemberName sets the member_name property value. The user display name/profile name
func (m *ExternalGroup_members) SetMemberName(value *string)() {
    m.member_name = value
}
type ExternalGroup_membersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMemberEmail()(*string)
    GetMemberId()(*int32)
    GetMemberLogin()(*string)
    GetMemberName()(*string)
    SetMemberEmail(value *string)()
    SetMemberId(value *int32)()
    SetMemberLogin(value *string)()
    SetMemberName(value *string)()
}
