package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PushRuleBypassRequest_organization the organization associated with the repository the bypass request is for.
type PushRuleBypassRequest_organization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the organization.
    id *int32
    // The name of the organization.
    name *string
}
// NewPushRuleBypassRequest_organization instantiates a new PushRuleBypassRequest_organization and sets the default values.
func NewPushRuleBypassRequest_organization()(*PushRuleBypassRequest_organization) {
    m := &PushRuleBypassRequest_organization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePushRuleBypassRequest_organizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePushRuleBypassRequest_organizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPushRuleBypassRequest_organization(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PushRuleBypassRequest_organization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PushRuleBypassRequest_organization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetId gets the id property value. The ID of the organization.
// returns a *int32 when successful
func (m *PushRuleBypassRequest_organization) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the organization.
// returns a *string when successful
func (m *PushRuleBypassRequest_organization) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *PushRuleBypassRequest_organization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PushRuleBypassRequest_organization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The ID of the organization.
func (m *PushRuleBypassRequest_organization) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the organization.
func (m *PushRuleBypassRequest_organization) SetName(value *string)() {
    m.name = value
}
type PushRuleBypassRequest_organizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    SetId(value *int32)()
    SetName(value *string)()
}
