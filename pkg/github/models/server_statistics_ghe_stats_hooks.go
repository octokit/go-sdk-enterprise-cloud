package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_hooks struct {
    // The active_hooks property
    active_hooks *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The inactive_hooks property
    inactive_hooks *int32
    // The total_hooks property
    total_hooks *int32
}
// NewServerStatistics_ghe_stats_hooks instantiates a new ServerStatistics_ghe_stats_hooks and sets the default values.
func NewServerStatistics_ghe_stats_hooks()(*ServerStatistics_ghe_stats_hooks) {
    m := &ServerStatistics_ghe_stats_hooks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_hooksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_hooksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_hooks(), nil
}
// GetActiveHooks gets the active_hooks property value. The active_hooks property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_hooks) GetActiveHooks()(*int32) {
    return m.active_hooks
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_hooks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_hooks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active_hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveHooks(val)
        }
        return nil
    }
    res["inactive_hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactiveHooks(val)
        }
        return nil
    }
    res["total_hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalHooks(val)
        }
        return nil
    }
    return res
}
// GetInactiveHooks gets the inactive_hooks property value. The inactive_hooks property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_hooks) GetInactiveHooks()(*int32) {
    return m.inactive_hooks
}
// GetTotalHooks gets the total_hooks property value. The total_hooks property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_hooks) GetTotalHooks()(*int32) {
    return m.total_hooks
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_hooks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("active_hooks", m.GetActiveHooks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inactive_hooks", m.GetInactiveHooks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_hooks", m.GetTotalHooks())
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
// SetActiveHooks sets the active_hooks property value. The active_hooks property
func (m *ServerStatistics_ghe_stats_hooks) SetActiveHooks(value *int32)() {
    m.active_hooks = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServerStatistics_ghe_stats_hooks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInactiveHooks sets the inactive_hooks property value. The inactive_hooks property
func (m *ServerStatistics_ghe_stats_hooks) SetInactiveHooks(value *int32)() {
    m.inactive_hooks = value
}
// SetTotalHooks sets the total_hooks property value. The total_hooks property
func (m *ServerStatistics_ghe_stats_hooks) SetTotalHooks(value *int32)() {
    m.total_hooks = value
}
type ServerStatistics_ghe_stats_hooksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveHooks()(*int32)
    GetInactiveHooks()(*int32)
    GetTotalHooks()(*int32)
    SetActiveHooks(value *int32)()
    SetInactiveHooks(value *int32)()
    SetTotalHooks(value *int32)()
}
