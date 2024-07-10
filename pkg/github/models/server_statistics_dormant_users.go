package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_dormant_users struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dormancy_threshold property
    dormancy_threshold *string
    // The total_dormant_users property
    total_dormant_users *int32
}
// NewServerStatistics_dormant_users instantiates a new ServerStatistics_dormant_users and sets the default values.
func NewServerStatistics_dormant_users()(*ServerStatistics_dormant_users) {
    m := &ServerStatistics_dormant_users{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_dormant_usersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_dormant_usersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_dormant_users(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_dormant_users) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDormancyThreshold gets the dormancy_threshold property value. The dormancy_threshold property
// returns a *string when successful
func (m *ServerStatistics_dormant_users) GetDormancyThreshold()(*string) {
    return m.dormancy_threshold
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_dormant_users) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dormancy_threshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDormancyThreshold(val)
        }
        return nil
    }
    res["total_dormant_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDormantUsers(val)
        }
        return nil
    }
    return res
}
// GetTotalDormantUsers gets the total_dormant_users property value. The total_dormant_users property
// returns a *int32 when successful
func (m *ServerStatistics_dormant_users) GetTotalDormantUsers()(*int32) {
    return m.total_dormant_users
}
// Serialize serializes information the current object
func (m *ServerStatistics_dormant_users) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dormancy_threshold", m.GetDormancyThreshold())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_dormant_users", m.GetTotalDormantUsers())
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
func (m *ServerStatistics_dormant_users) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDormancyThreshold sets the dormancy_threshold property value. The dormancy_threshold property
func (m *ServerStatistics_dormant_users) SetDormancyThreshold(value *string)() {
    m.dormancy_threshold = value
}
// SetTotalDormantUsers sets the total_dormant_users property value. The total_dormant_users property
func (m *ServerStatistics_dormant_users) SetTotalDormantUsers(value *int32)() {
    m.total_dormant_users = value
}
type ServerStatistics_dormant_usersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDormancyThreshold()(*string)
    GetTotalDormantUsers()(*int32)
    SetDormancyThreshold(value *string)()
    SetTotalDormantUsers(value *int32)()
}
