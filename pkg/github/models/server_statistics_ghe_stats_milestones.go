package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_milestones struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The closed_milestones property
    closed_milestones *int32
    // The open_milestones property
    open_milestones *int32
    // The total_milestones property
    total_milestones *int32
}
// NewServerStatistics_ghe_stats_milestones instantiates a new ServerStatistics_ghe_stats_milestones and sets the default values.
func NewServerStatistics_ghe_stats_milestones()(*ServerStatistics_ghe_stats_milestones) {
    m := &ServerStatistics_ghe_stats_milestones{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_milestonesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_milestonesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_milestones(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_milestones) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClosedMilestones gets the closed_milestones property value. The closed_milestones property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_milestones) GetClosedMilestones()(*int32) {
    return m.closed_milestones
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_milestones) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["closed_milestones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedMilestones(val)
        }
        return nil
    }
    res["open_milestones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenMilestones(val)
        }
        return nil
    }
    res["total_milestones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalMilestones(val)
        }
        return nil
    }
    return res
}
// GetOpenMilestones gets the open_milestones property value. The open_milestones property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_milestones) GetOpenMilestones()(*int32) {
    return m.open_milestones
}
// GetTotalMilestones gets the total_milestones property value. The total_milestones property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_milestones) GetTotalMilestones()(*int32) {
    return m.total_milestones
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_milestones) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("closed_milestones", m.GetClosedMilestones())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("open_milestones", m.GetOpenMilestones())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_milestones", m.GetTotalMilestones())
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
func (m *ServerStatistics_ghe_stats_milestones) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClosedMilestones sets the closed_milestones property value. The closed_milestones property
func (m *ServerStatistics_ghe_stats_milestones) SetClosedMilestones(value *int32)() {
    m.closed_milestones = value
}
// SetOpenMilestones sets the open_milestones property value. The open_milestones property
func (m *ServerStatistics_ghe_stats_milestones) SetOpenMilestones(value *int32)() {
    m.open_milestones = value
}
// SetTotalMilestones sets the total_milestones property value. The total_milestones property
func (m *ServerStatistics_ghe_stats_milestones) SetTotalMilestones(value *int32)() {
    m.total_milestones = value
}
type ServerStatistics_ghe_stats_milestonesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedMilestones()(*int32)
    GetOpenMilestones()(*int32)
    GetTotalMilestones()(*int32)
    SetClosedMilestones(value *int32)()
    SetOpenMilestones(value *int32)()
    SetTotalMilestones(value *int32)()
}
