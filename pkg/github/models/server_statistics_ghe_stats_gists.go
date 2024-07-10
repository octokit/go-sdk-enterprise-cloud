package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_gists struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The private_gists property
    private_gists *int32
    // The public_gists property
    public_gists *int32
    // The total_gists property
    total_gists *int32
}
// NewServerStatistics_ghe_stats_gists instantiates a new ServerStatistics_ghe_stats_gists and sets the default values.
func NewServerStatistics_ghe_stats_gists()(*ServerStatistics_ghe_stats_gists) {
    m := &ServerStatistics_ghe_stats_gists{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_gistsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_gistsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_gists(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_gists) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_gists) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["private_gists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateGists(val)
        }
        return nil
    }
    res["public_gists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicGists(val)
        }
        return nil
    }
    res["total_gists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalGists(val)
        }
        return nil
    }
    return res
}
// GetPrivateGists gets the private_gists property value. The private_gists property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_gists) GetPrivateGists()(*int32) {
    return m.private_gists
}
// GetPublicGists gets the public_gists property value. The public_gists property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_gists) GetPublicGists()(*int32) {
    return m.public_gists
}
// GetTotalGists gets the total_gists property value. The total_gists property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_gists) GetTotalGists()(*int32) {
    return m.total_gists
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_gists) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("private_gists", m.GetPrivateGists())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("public_gists", m.GetPublicGists())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_gists", m.GetTotalGists())
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
func (m *ServerStatistics_ghe_stats_gists) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateGists sets the private_gists property value. The private_gists property
func (m *ServerStatistics_ghe_stats_gists) SetPrivateGists(value *int32)() {
    m.private_gists = value
}
// SetPublicGists sets the public_gists property value. The public_gists property
func (m *ServerStatistics_ghe_stats_gists) SetPublicGists(value *int32)() {
    m.public_gists = value
}
// SetTotalGists sets the total_gists property value. The total_gists property
func (m *ServerStatistics_ghe_stats_gists) SetTotalGists(value *int32)() {
    m.total_gists = value
}
type ServerStatistics_ghe_stats_gistsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateGists()(*int32)
    GetPublicGists()(*int32)
    GetTotalGists()(*int32)
    SetPrivateGists(value *int32)()
    SetPublicGists(value *int32)()
    SetTotalGists(value *int32)()
}
