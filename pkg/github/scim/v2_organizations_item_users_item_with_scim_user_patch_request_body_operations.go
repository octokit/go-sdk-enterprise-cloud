package scim

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The path property
    path *string
    // The value property
    value V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable
}
// V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value composed type wrapper for classes string, V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able, []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able
type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value struct {
    // Composed type representation for type string
    string *string
    // Composed type representation for type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able
    v2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1 V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able
    // Composed type representation for type []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able
    v2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2 []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able
}
// NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value instantiates a new V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value and sets the default values.
func NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value()(*V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) {
    m := &V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value{
    }
    return m
}
// CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    } else if val, err := parseNode.GetCollectionOfObjectValues(CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able)
            }
        }
        result.SetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2(cast)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) GetString()(*string) {
    return m.string
}
// GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1 gets the V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1 property value. Composed type representation for type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able
// returns a V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1()(V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able) {
    return m.v2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1
}
// GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2 gets the V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2 property value. Composed type representation for type []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able
// returns a []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2()([]V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able) {
    return m.v2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2
}
// Serialize serializes information the current object
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2()))
        for i, v := range m.GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetString sets the string property value. Composed type representation for type string
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) SetString(value *string)() {
    m.string = value
}
// SetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1 sets the V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1 property value. Composed type representation for type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) SetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1(value V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able)() {
    m.v2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1 = value
}
// SetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2 sets the V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2 property value. Composed type representation for type []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_value) SetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2(value []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able)() {
    m.v2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2 = value
}
type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1()(V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able)
    GetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2()([]V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able)
    SetString(value *string)()
    SetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember1(value V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember1able)()
    SetV2OrganizationsItemUsersItemWithScimUserPatchRequestBodyOperationsValueMember2(value []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_valueMember2able)()
}
// NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations instantiates a new V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations and sets the default values.
func NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations()(*V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) {
    m := &V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_OperationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_OperationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable))
        }
        return nil
    }
    return res
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) GetPath()(*string) {
    return m.path
}
// GetValue gets the value property value. The value property
// returns a V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) GetValue()(V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPath sets the path property value. The path property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) SetPath(value *string)() {
    m.path = value
}
// SetValue sets the value property value. The value property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations) SetValue(value V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable)() {
    m.value = value
}
type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPath()(*string)
    GetValue()(V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable)
    SetPath(value *string)()
    SetValue(value V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operations_WithScim_user_PatchRequestBody_Operations_valueable)()
}
