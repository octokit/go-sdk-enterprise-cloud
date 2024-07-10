package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetLicenseSyncStatus_server_instances_last_sync struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date property
    date *string
    // The error property
    error *string
    // The status property
    status *string
}
// NewGetLicenseSyncStatus_server_instances_last_sync instantiates a new GetLicenseSyncStatus_server_instances_last_sync and sets the default values.
func NewGetLicenseSyncStatus_server_instances_last_sync()(*GetLicenseSyncStatus_server_instances_last_sync) {
    m := &GetLicenseSyncStatus_server_instances_last_sync{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetLicenseSyncStatus_server_instances_last_syncFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetLicenseSyncStatus_server_instances_last_syncFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetLicenseSyncStatus_server_instances_last_sync(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetLicenseSyncStatus_server_instances_last_sync) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDate gets the date property value. The date property
// returns a *string when successful
func (m *GetLicenseSyncStatus_server_instances_last_sync) GetDate()(*string) {
    return m.date
}
// GetError gets the error property value. The error property
// returns a *string when successful
func (m *GetLicenseSyncStatus_server_instances_last_sync) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetLicenseSyncStatus_server_instances_last_sync) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *GetLicenseSyncStatus_server_instances_last_sync) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *GetLicenseSyncStatus_server_instances_last_sync) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *GetLicenseSyncStatus_server_instances_last_sync) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDate sets the date property value. The date property
func (m *GetLicenseSyncStatus_server_instances_last_sync) SetDate(value *string)() {
    m.date = value
}
// SetError sets the error property value. The error property
func (m *GetLicenseSyncStatus_server_instances_last_sync) SetError(value *string)() {
    m.error = value
}
// SetStatus sets the status property value. The status property
func (m *GetLicenseSyncStatus_server_instances_last_sync) SetStatus(value *string)() {
    m.status = value
}
type GetLicenseSyncStatus_server_instances_last_syncable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*string)
    GetError()(*string)
    GetStatus()(*string)
    SetDate(value *string)()
    SetError(value *string)()
    SetStatus(value *string)()
}
