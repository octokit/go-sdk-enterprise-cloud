package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetLicenseSyncStatus information about the status of a license sync job for an enterprise.
type GetLicenseSyncStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The server_instances property
    server_instances []GetLicenseSyncStatus_server_instancesable
}
// NewGetLicenseSyncStatus instantiates a new GetLicenseSyncStatus and sets the default values.
func NewGetLicenseSyncStatus()(*GetLicenseSyncStatus) {
    m := &GetLicenseSyncStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetLicenseSyncStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetLicenseSyncStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetLicenseSyncStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetLicenseSyncStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetLicenseSyncStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["server_instances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGetLicenseSyncStatus_server_instancesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GetLicenseSyncStatus_server_instancesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GetLicenseSyncStatus_server_instancesable)
                }
            }
            m.SetServerInstances(res)
        }
        return nil
    }
    return res
}
// GetServerInstances gets the server_instances property value. The server_instances property
// returns a []GetLicenseSyncStatus_server_instancesable when successful
func (m *GetLicenseSyncStatus) GetServerInstances()([]GetLicenseSyncStatus_server_instancesable) {
    return m.server_instances
}
// Serialize serializes information the current object
func (m *GetLicenseSyncStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetServerInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServerInstances()))
        for i, v := range m.GetServerInstances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("server_instances", cast)
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
func (m *GetLicenseSyncStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetServerInstances sets the server_instances property value. The server_instances property
func (m *GetLicenseSyncStatus) SetServerInstances(value []GetLicenseSyncStatus_server_instancesable)() {
    m.server_instances = value
}
type GetLicenseSyncStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetServerInstances()([]GetLicenseSyncStatus_server_instancesable)
    SetServerInstances(value []GetLicenseSyncStatus_server_instancesable)()
}
