package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetLicenseSyncStatus_server_instances struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hostname property
    hostname *string
    // The last_sync property
    last_sync GetLicenseSyncStatus_server_instances_last_syncable
    // The server_id property
    server_id *string
}
// NewGetLicenseSyncStatus_server_instances instantiates a new GetLicenseSyncStatus_server_instances and sets the default values.
func NewGetLicenseSyncStatus_server_instances()(*GetLicenseSyncStatus_server_instances) {
    m := &GetLicenseSyncStatus_server_instances{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetLicenseSyncStatus_server_instancesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetLicenseSyncStatus_server_instancesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetLicenseSyncStatus_server_instances(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetLicenseSyncStatus_server_instances) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetLicenseSyncStatus_server_instances) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
        }
        return nil
    }
    res["last_sync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGetLicenseSyncStatus_server_instances_last_syncFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSync(val.(GetLicenseSyncStatus_server_instances_last_syncable))
        }
        return nil
    }
    res["server_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerId(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. The hostname property
// returns a *string when successful
func (m *GetLicenseSyncStatus_server_instances) GetHostname()(*string) {
    return m.hostname
}
// GetLastSync gets the last_sync property value. The last_sync property
// returns a GetLicenseSyncStatus_server_instances_last_syncable when successful
func (m *GetLicenseSyncStatus_server_instances) GetLastSync()(GetLicenseSyncStatus_server_instances_last_syncable) {
    return m.last_sync
}
// GetServerId gets the server_id property value. The server_id property
// returns a *string when successful
func (m *GetLicenseSyncStatus_server_instances) GetServerId()(*string) {
    return m.server_id
}
// Serialize serializes information the current object
func (m *GetLicenseSyncStatus_server_instances) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("last_sync", m.GetLastSync())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("server_id", m.GetServerId())
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
func (m *GetLicenseSyncStatus_server_instances) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHostname sets the hostname property value. The hostname property
func (m *GetLicenseSyncStatus_server_instances) SetHostname(value *string)() {
    m.hostname = value
}
// SetLastSync sets the last_sync property value. The last_sync property
func (m *GetLicenseSyncStatus_server_instances) SetLastSync(value GetLicenseSyncStatus_server_instances_last_syncable)() {
    m.last_sync = value
}
// SetServerId sets the server_id property value. The server_id property
func (m *GetLicenseSyncStatus_server_instances) SetServerId(value *string)() {
    m.server_id = value
}
type GetLicenseSyncStatus_server_instancesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHostname()(*string)
    GetLastSync()(GetLicenseSyncStatus_server_instances_last_syncable)
    GetServerId()(*string)
    SetHostname(value *string)()
    SetLastSync(value GetLicenseSyncStatus_server_instances_last_syncable)()
    SetServerId(value *string)()
}
