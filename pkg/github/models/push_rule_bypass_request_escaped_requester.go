package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PushRuleBypassRequest_requester the user who requested the bypass.
type PushRuleBypassRequest_requester struct {
    // The ID of the GitHub user who requested the bypass.
    actor_id *int32
    // The name of the GitHub user who requested the bypass.
    actor_name *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewPushRuleBypassRequest_requester instantiates a new PushRuleBypassRequest_requester and sets the default values.
func NewPushRuleBypassRequest_requester()(*PushRuleBypassRequest_requester) {
    m := &PushRuleBypassRequest_requester{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePushRuleBypassRequest_requesterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePushRuleBypassRequest_requesterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPushRuleBypassRequest_requester(), nil
}
// GetActorId gets the actor_id property value. The ID of the GitHub user who requested the bypass.
// returns a *int32 when successful
func (m *PushRuleBypassRequest_requester) GetActorId()(*int32) {
    return m.actor_id
}
// GetActorName gets the actor_name property value. The name of the GitHub user who requested the bypass.
// returns a *string when successful
func (m *PushRuleBypassRequest_requester) GetActorName()(*string) {
    return m.actor_name
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PushRuleBypassRequest_requester) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PushRuleBypassRequest_requester) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorId(val)
        }
        return nil
    }
    res["actor_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PushRuleBypassRequest_requester) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("actor_id", m.GetActorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("actor_name", m.GetActorName())
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
// SetActorId sets the actor_id property value. The ID of the GitHub user who requested the bypass.
func (m *PushRuleBypassRequest_requester) SetActorId(value *int32)() {
    m.actor_id = value
}
// SetActorName sets the actor_name property value. The name of the GitHub user who requested the bypass.
func (m *PushRuleBypassRequest_requester) SetActorName(value *string)() {
    m.actor_name = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PushRuleBypassRequest_requester) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type PushRuleBypassRequest_requesterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActorId()(*int32)
    GetActorName()(*string)
    SetActorId(value *int32)()
    SetActorName(value *string)()
}
