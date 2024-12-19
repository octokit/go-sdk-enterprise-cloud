package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BypassResponse_reviewer the user who reviewed the bypass request.
type BypassResponse_reviewer struct {
    // The ID of the GitHub user who reviewed the bypass request.
    actor_id *int32
    // The name of the GitHub user who reviewed the bypass request.
    actor_name *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewBypassResponse_reviewer instantiates a new BypassResponse_reviewer and sets the default values.
func NewBypassResponse_reviewer()(*BypassResponse_reviewer) {
    m := &BypassResponse_reviewer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBypassResponse_reviewerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBypassResponse_reviewerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBypassResponse_reviewer(), nil
}
// GetActorId gets the actor_id property value. The ID of the GitHub user who reviewed the bypass request.
// returns a *int32 when successful
func (m *BypassResponse_reviewer) GetActorId()(*int32) {
    return m.actor_id
}
// GetActorName gets the actor_name property value. The name of the GitHub user who reviewed the bypass request.
// returns a *string when successful
func (m *BypassResponse_reviewer) GetActorName()(*string) {
    return m.actor_name
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BypassResponse_reviewer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BypassResponse_reviewer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *BypassResponse_reviewer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetActorId sets the actor_id property value. The ID of the GitHub user who reviewed the bypass request.
func (m *BypassResponse_reviewer) SetActorId(value *int32)() {
    m.actor_id = value
}
// SetActorName sets the actor_name property value. The name of the GitHub user who reviewed the bypass request.
func (m *BypassResponse_reviewer) SetActorName(value *string)() {
    m.actor_name = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BypassResponse_reviewer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type BypassResponse_reviewerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActorId()(*int32)
    GetActorName()(*string)
    SetActorId(value *int32)()
    SetActorName(value *string)()
}
