package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_pulls struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The mergeable_pulls property
    mergeable_pulls *int32
    // The merged_pulls property
    merged_pulls *int32
    // The total_pulls property
    total_pulls *int32
    // The unmergeable_pulls property
    unmergeable_pulls *int32
}
// NewServerStatistics_ghe_stats_pulls instantiates a new ServerStatistics_ghe_stats_pulls and sets the default values.
func NewServerStatistics_ghe_stats_pulls()(*ServerStatistics_ghe_stats_pulls) {
    m := &ServerStatistics_ghe_stats_pulls{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_pullsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_pullsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_pulls(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_pulls) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_pulls) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mergeable_pulls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeablePulls(val)
        }
        return nil
    }
    res["merged_pulls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergedPulls(val)
        }
        return nil
    }
    res["total_pulls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPulls(val)
        }
        return nil
    }
    res["unmergeable_pulls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnmergeablePulls(val)
        }
        return nil
    }
    return res
}
// GetMergeablePulls gets the mergeable_pulls property value. The mergeable_pulls property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_pulls) GetMergeablePulls()(*int32) {
    return m.mergeable_pulls
}
// GetMergedPulls gets the merged_pulls property value. The merged_pulls property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_pulls) GetMergedPulls()(*int32) {
    return m.merged_pulls
}
// GetTotalPulls gets the total_pulls property value. The total_pulls property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_pulls) GetTotalPulls()(*int32) {
    return m.total_pulls
}
// GetUnmergeablePulls gets the unmergeable_pulls property value. The unmergeable_pulls property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_pulls) GetUnmergeablePulls()(*int32) {
    return m.unmergeable_pulls
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_pulls) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("mergeable_pulls", m.GetMergeablePulls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("merged_pulls", m.GetMergedPulls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_pulls", m.GetTotalPulls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unmergeable_pulls", m.GetUnmergeablePulls())
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
func (m *ServerStatistics_ghe_stats_pulls) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMergeablePulls sets the mergeable_pulls property value. The mergeable_pulls property
func (m *ServerStatistics_ghe_stats_pulls) SetMergeablePulls(value *int32)() {
    m.mergeable_pulls = value
}
// SetMergedPulls sets the merged_pulls property value. The merged_pulls property
func (m *ServerStatistics_ghe_stats_pulls) SetMergedPulls(value *int32)() {
    m.merged_pulls = value
}
// SetTotalPulls sets the total_pulls property value. The total_pulls property
func (m *ServerStatistics_ghe_stats_pulls) SetTotalPulls(value *int32)() {
    m.total_pulls = value
}
// SetUnmergeablePulls sets the unmergeable_pulls property value. The unmergeable_pulls property
func (m *ServerStatistics_ghe_stats_pulls) SetUnmergeablePulls(value *int32)() {
    m.unmergeable_pulls = value
}
type ServerStatistics_ghe_stats_pullsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMergeablePulls()(*int32)
    GetMergedPulls()(*int32)
    GetTotalPulls()(*int32)
    GetUnmergeablePulls()(*int32)
    SetMergeablePulls(value *int32)()
    SetMergedPulls(value *int32)()
    SetTotalPulls(value *int32)()
    SetUnmergeablePulls(value *int32)()
}
