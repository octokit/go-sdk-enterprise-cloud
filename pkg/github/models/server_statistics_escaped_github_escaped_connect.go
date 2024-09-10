package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_github_connect struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The features_enabled property
    features_enabled []string
}
// NewServerStatistics_github_connect instantiates a new ServerStatistics_github_connect and sets the default values.
func NewServerStatistics_github_connect()(*ServerStatistics_github_connect) {
    m := &ServerStatistics_github_connect{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_github_connectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_github_connectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_github_connect(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_github_connect) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFeaturesEnabled gets the features_enabled property value. The features_enabled property
// returns a []string when successful
func (m *ServerStatistics_github_connect) GetFeaturesEnabled()([]string) {
    return m.features_enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_github_connect) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["features_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetFeaturesEnabled(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServerStatistics_github_connect) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFeaturesEnabled() != nil {
        err := writer.WriteCollectionOfStringValues("features_enabled", m.GetFeaturesEnabled())
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
func (m *ServerStatistics_github_connect) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFeaturesEnabled sets the features_enabled property value. The features_enabled property
func (m *ServerStatistics_github_connect) SetFeaturesEnabled(value []string)() {
    m.features_enabled = value
}
type ServerStatistics_github_connectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFeaturesEnabled()([]string)
    SetFeaturesEnabled(value []string)()
}
