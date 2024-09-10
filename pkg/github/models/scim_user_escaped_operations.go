package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScimUser_operations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The op property
    op *ScimUser_operations_op
    // The path property
    path *string
    // The value property
    value ScimUser_operations_ScimUser_operations_valueable
}
// ScimUser_operations_ScimUser_operations_value composed type wrapper for classes ScimUser_operations_valueMember1able, ScimUser_operations_valueMember2able, string
type ScimUser_operations_ScimUser_operations_value struct {
    // Composed type representation for type ScimUser_operations_valueMember1able
    scimUser_operations_valueMember1 ScimUser_operations_valueMember1able
    // Composed type representation for type ScimUser_operations_valueMember2able
    scimUser_operations_valueMember2 ScimUser_operations_valueMember2able
    // Composed type representation for type string
    string *string
}
// NewScimUser_operations_ScimUser_operations_value instantiates a new ScimUser_operations_ScimUser_operations_value and sets the default values.
func NewScimUser_operations_ScimUser_operations_value()(*ScimUser_operations_ScimUser_operations_value) {
    m := &ScimUser_operations_ScimUser_operations_value{
    }
    return m
}
// CreateScimUser_operations_ScimUser_operations_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimUser_operations_ScimUser_operations_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewScimUser_operations_ScimUser_operations_value()
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
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimUser_operations_ScimUser_operations_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ScimUser_operations_ScimUser_operations_value) GetIsComposedType()(bool) {
    return true
}
// GetScimUserOperationsValueMember1 gets the scimUser_operations_valueMember1 property value. Composed type representation for type ScimUser_operations_valueMember1able
// returns a ScimUser_operations_valueMember1able when successful
func (m *ScimUser_operations_ScimUser_operations_value) GetScimUserOperationsValueMember1()(ScimUser_operations_valueMember1able) {
    return m.scimUser_operations_valueMember1
}
// GetScimUserOperationsValueMember2 gets the scimUser_operations_valueMember2 property value. Composed type representation for type ScimUser_operations_valueMember2able
// returns a ScimUser_operations_valueMember2able when successful
func (m *ScimUser_operations_ScimUser_operations_value) GetScimUserOperationsValueMember2()(ScimUser_operations_valueMember2able) {
    return m.scimUser_operations_valueMember2
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ScimUser_operations_ScimUser_operations_value) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ScimUser_operations_ScimUser_operations_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetScimUserOperationsValueMember1() != nil {
        err := writer.WriteObjectValue("", m.GetScimUserOperationsValueMember1())
        if err != nil {
            return err
        }
    } else if m.GetScimUserOperationsValueMember2() != nil {
        err := writer.WriteObjectValue("", m.GetScimUserOperationsValueMember2())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetScimUserOperationsValueMember1 sets the scimUser_operations_valueMember1 property value. Composed type representation for type ScimUser_operations_valueMember1able
func (m *ScimUser_operations_ScimUser_operations_value) SetScimUserOperationsValueMember1(value ScimUser_operations_valueMember1able)() {
    m.scimUser_operations_valueMember1 = value
}
// SetScimUserOperationsValueMember2 sets the scimUser_operations_valueMember2 property value. Composed type representation for type ScimUser_operations_valueMember2able
func (m *ScimUser_operations_ScimUser_operations_value) SetScimUserOperationsValueMember2(value ScimUser_operations_valueMember2able)() {
    m.scimUser_operations_valueMember2 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ScimUser_operations_ScimUser_operations_value) SetString(value *string)() {
    m.string = value
}
type ScimUser_operations_ScimUser_operations_valueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScimUserOperationsValueMember1()(ScimUser_operations_valueMember1able)
    GetScimUserOperationsValueMember2()(ScimUser_operations_valueMember2able)
    GetString()(*string)
    SetScimUserOperationsValueMember1(value ScimUser_operations_valueMember1able)()
    SetScimUserOperationsValueMember2(value ScimUser_operations_valueMember2able)()
    SetString(value *string)()
}
// NewScimUser_operations instantiates a new ScimUser_operations and sets the default values.
func NewScimUser_operations()(*ScimUser_operations) {
    m := &ScimUser_operations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimUser_operationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimUser_operationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimUser_operations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimUser_operations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimUser_operations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["op"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScimUser_operations_op)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOp(val.(*ScimUser_operations_op))
        }
        return nil
    }
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
        val, err := n.GetObjectValue(CreateScimUser_operations_ScimUser_operations_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(ScimUser_operations_ScimUser_operations_valueable))
        }
        return nil
    }
    return res
}
// GetOp gets the op property value. The op property
// returns a *ScimUser_operations_op when successful
func (m *ScimUser_operations) GetOp()(*ScimUser_operations_op) {
    return m.op
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ScimUser_operations) GetPath()(*string) {
    return m.path
}
// GetValue gets the value property value. The value property
// returns a ScimUser_operations_ScimUser_operations_valueable when successful
func (m *ScimUser_operations) GetValue()(ScimUser_operations_ScimUser_operations_valueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ScimUser_operations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOp() != nil {
        cast := (*m.GetOp()).String()
        err := writer.WriteStringValue("op", &cast)
        if err != nil {
            return err
        }
    }
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
func (m *ScimUser_operations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOp sets the op property value. The op property
func (m *ScimUser_operations) SetOp(value *ScimUser_operations_op)() {
    m.op = value
}
// SetPath sets the path property value. The path property
func (m *ScimUser_operations) SetPath(value *string)() {
    m.path = value
}
// SetValue sets the value property value. The value property
func (m *ScimUser_operations) SetValue(value ScimUser_operations_ScimUser_operations_valueable)() {
    m.value = value
}
type ScimUser_operationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOp()(*ScimUser_operations_op)
    GetPath()(*string)
    GetValue()(ScimUser_operations_ScimUser_operations_valueable)
    SetOp(value *ScimUser_operations_op)()
    SetPath(value *string)()
    SetValue(value ScimUser_operations_ScimUser_operations_valueable)()
}
