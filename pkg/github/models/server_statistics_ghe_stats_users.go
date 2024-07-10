package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_users struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The admin_users property
    admin_users *int32
    // The suspended_users property
    suspended_users *int32
    // The total_users property
    total_users *int32
}
// NewServerStatistics_ghe_stats_users instantiates a new ServerStatistics_ghe_stats_users and sets the default values.
func NewServerStatistics_ghe_stats_users()(*ServerStatistics_ghe_stats_users) {
    m := &ServerStatistics_ghe_stats_users{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_usersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_usersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_users(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_users) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdminUsers gets the admin_users property value. The admin_users property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_users) GetAdminUsers()(*int32) {
    return m.admin_users
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_users) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["admin_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminUsers(val)
        }
        return nil
    }
    res["suspended_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspendedUsers(val)
        }
        return nil
    }
    res["total_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUsers(val)
        }
        return nil
    }
    return res
}
// GetSuspendedUsers gets the suspended_users property value. The suspended_users property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_users) GetSuspendedUsers()(*int32) {
    return m.suspended_users
}
// GetTotalUsers gets the total_users property value. The total_users property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_users) GetTotalUsers()(*int32) {
    return m.total_users
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_users) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("admin_users", m.GetAdminUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("suspended_users", m.GetSuspendedUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_users", m.GetTotalUsers())
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
func (m *ServerStatistics_ghe_stats_users) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdminUsers sets the admin_users property value. The admin_users property
func (m *ServerStatistics_ghe_stats_users) SetAdminUsers(value *int32)() {
    m.admin_users = value
}
// SetSuspendedUsers sets the suspended_users property value. The suspended_users property
func (m *ServerStatistics_ghe_stats_users) SetSuspendedUsers(value *int32)() {
    m.suspended_users = value
}
// SetTotalUsers sets the total_users property value. The total_users property
func (m *ServerStatistics_ghe_stats_users) SetTotalUsers(value *int32)() {
    m.total_users = value
}
type ServerStatistics_ghe_stats_usersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminUsers()(*int32)
    GetSuspendedUsers()(*int32)
    GetTotalUsers()(*int32)
    SetAdminUsers(value *int32)()
    SetSuspendedUsers(value *int32)()
    SetTotalUsers(value *int32)()
}
