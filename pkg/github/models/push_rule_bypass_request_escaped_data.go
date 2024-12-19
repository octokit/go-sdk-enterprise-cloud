package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PushRuleBypassRequest_data struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of rule that was violated.
    rule_type *string
    // The ID of the ruleset for the rules that were violated.
    ruleset_id *int32
    // The name of the ruleset for the rules that were violated.
    ruleset_name *string
    // The number of rule violations generated from the push associated with this request.
    total_violations *int32
}
// NewPushRuleBypassRequest_data instantiates a new PushRuleBypassRequest_data and sets the default values.
func NewPushRuleBypassRequest_data()(*PushRuleBypassRequest_data) {
    m := &PushRuleBypassRequest_data{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePushRuleBypassRequest_dataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePushRuleBypassRequest_dataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPushRuleBypassRequest_data(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PushRuleBypassRequest_data) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PushRuleBypassRequest_data) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["rule_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleType(val)
        }
        return nil
    }
    res["ruleset_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulesetId(val)
        }
        return nil
    }
    res["ruleset_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulesetName(val)
        }
        return nil
    }
    res["total_violations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalViolations(val)
        }
        return nil
    }
    return res
}
// GetRulesetId gets the ruleset_id property value. The ID of the ruleset for the rules that were violated.
// returns a *int32 when successful
func (m *PushRuleBypassRequest_data) GetRulesetId()(*int32) {
    return m.ruleset_id
}
// GetRulesetName gets the ruleset_name property value. The name of the ruleset for the rules that were violated.
// returns a *string when successful
func (m *PushRuleBypassRequest_data) GetRulesetName()(*string) {
    return m.ruleset_name
}
// GetRuleType gets the rule_type property value. The type of rule that was violated.
// returns a *string when successful
func (m *PushRuleBypassRequest_data) GetRuleType()(*string) {
    return m.rule_type
}
// GetTotalViolations gets the total_violations property value. The number of rule violations generated from the push associated with this request.
// returns a *int32 when successful
func (m *PushRuleBypassRequest_data) GetTotalViolations()(*int32) {
    return m.total_violations
}
// Serialize serializes information the current object
func (m *PushRuleBypassRequest_data) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("ruleset_id", m.GetRulesetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleset_name", m.GetRulesetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rule_type", m.GetRuleType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_violations", m.GetTotalViolations())
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
func (m *PushRuleBypassRequest_data) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRulesetId sets the ruleset_id property value. The ID of the ruleset for the rules that were violated.
func (m *PushRuleBypassRequest_data) SetRulesetId(value *int32)() {
    m.ruleset_id = value
}
// SetRulesetName sets the ruleset_name property value. The name of the ruleset for the rules that were violated.
func (m *PushRuleBypassRequest_data) SetRulesetName(value *string)() {
    m.ruleset_name = value
}
// SetRuleType sets the rule_type property value. The type of rule that was violated.
func (m *PushRuleBypassRequest_data) SetRuleType(value *string)() {
    m.rule_type = value
}
// SetTotalViolations sets the total_violations property value. The number of rule violations generated from the push associated with this request.
func (m *PushRuleBypassRequest_data) SetTotalViolations(value *int32)() {
    m.total_violations = value
}
type PushRuleBypassRequest_dataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRulesetId()(*int32)
    GetRulesetName()(*string)
    GetRuleType()(*string)
    GetTotalViolations()(*int32)
    SetRulesetId(value *int32)()
    SetRulesetName(value *string)()
    SetRuleType(value *string)()
    SetTotalViolations(value *int32)()
}
