package scim

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1 struct {
    // The active property
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The externalId property
    externalId *string
    // The familyName property
    familyName *string
    // The givenName property
    givenName *string
    // The userName property
    userName *string
}
// NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1 instantiates a new V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1 and sets the default values.
func NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1()(*V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) {
    m := &V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1(), nil
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExternalId gets the externalId property value. The externalId property
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) GetExternalId()(*string) {
    return m.externalId
}
// GetFamilyName gets the familyName property value. The familyName property
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) GetFamilyName()(*string) {
    return m.familyName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
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
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetGivenName gets the givenName property value. The givenName property
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) GetGivenName()(*string) {
    return m.givenName
}
// GetUserName gets the userName property value. The userName property
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("familyName", m.GetFamilyName())
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
        err := writer.WriteStringValue("userName", m.GetUserName())
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
// SetActive sets the active property value. The active property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExternalId sets the externalId property value. The externalId property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) SetExternalId(value *string)() {
    m.externalId = value
}
// SetFamilyName sets the familyName property value. The familyName property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) SetFamilyName(value *string)() {
    m.familyName = value
}
// SetGivenName sets the givenName property value. The givenName property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) SetGivenName(value *string)() {
    m.givenName = value
}
// SetUserName sets the userName property value. The userName property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1) SetUserName(value *string)() {
    m.userName = value
}
type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetExternalId()(*string)
    GetFamilyName()(*string)
    GetGivenName()(*string)
    GetUserName()(*string)
    SetActive(value *bool)()
    SetExternalId(value *string)()
    SetFamilyName(value *string)()
    SetGivenName(value *string)()
    SetUserName(value *string)()
}
