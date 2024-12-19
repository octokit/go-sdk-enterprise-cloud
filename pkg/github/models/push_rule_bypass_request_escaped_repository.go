package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PushRuleBypassRequest_repository the repository the bypass request is for.
type PushRuleBypassRequest_repository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full name of the repository the bypass request is for.
    full_name *string
    // The ID of the repository the bypass request is for.
    id *int32
    // The name of the repository the bypass request is for.
    name *string
}
// NewPushRuleBypassRequest_repository instantiates a new PushRuleBypassRequest_repository and sets the default values.
func NewPushRuleBypassRequest_repository()(*PushRuleBypassRequest_repository) {
    m := &PushRuleBypassRequest_repository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePushRuleBypassRequest_repositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePushRuleBypassRequest_repositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPushRuleBypassRequest_repository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PushRuleBypassRequest_repository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PushRuleBypassRequest_repository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetFullName gets the full_name property value. The full name of the repository the bypass request is for.
// returns a *string when successful
func (m *PushRuleBypassRequest_repository) GetFullName()(*string) {
    return m.full_name
}
// GetId gets the id property value. The ID of the repository the bypass request is for.
// returns a *int32 when successful
func (m *PushRuleBypassRequest_repository) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the repository the bypass request is for.
// returns a *string when successful
func (m *PushRuleBypassRequest_repository) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *PushRuleBypassRequest_repository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("full_name", m.GetFullName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *PushRuleBypassRequest_repository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFullName sets the full_name property value. The full name of the repository the bypass request is for.
func (m *PushRuleBypassRequest_repository) SetFullName(value *string)() {
    m.full_name = value
}
// SetId sets the id property value. The ID of the repository the bypass request is for.
func (m *PushRuleBypassRequest_repository) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the repository the bypass request is for.
func (m *PushRuleBypassRequest_repository) SetName(value *string)() {
    m.name = value
}
type PushRuleBypassRequest_repositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFullName()(*string)
    GetId()(*int32)
    GetName()(*string)
    SetFullName(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
}
