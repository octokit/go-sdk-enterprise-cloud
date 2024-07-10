package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServerStatisticsPackages packages metrics that are included in the Server Statistics payload/export from GHES
type ServerStatisticsPackages struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The details of the package ecosystems that are enabled in a GHES installation
    ecosystems []ServerStatisticsPackages_ecosystemsable
    // Whether GitHub Packages is enabled globally in a GHES installation
    registry_enabled *bool
    // Whether a beta registry is enabled in a GHES installation
    registry_v2_enabled *bool
}
// NewServerStatisticsPackages instantiates a new ServerStatisticsPackages and sets the default values.
func NewServerStatisticsPackages()(*ServerStatisticsPackages) {
    m := &ServerStatisticsPackages{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatisticsPackagesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatisticsPackagesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatisticsPackages(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatisticsPackages) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEcosystems gets the ecosystems property value. The details of the package ecosystems that are enabled in a GHES installation
// returns a []ServerStatisticsPackages_ecosystemsable when successful
func (m *ServerStatisticsPackages) GetEcosystems()([]ServerStatisticsPackages_ecosystemsable) {
    return m.ecosystems
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatisticsPackages) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ecosystems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServerStatisticsPackages_ecosystemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServerStatisticsPackages_ecosystemsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServerStatisticsPackages_ecosystemsable)
                }
            }
            m.SetEcosystems(res)
        }
        return nil
    }
    res["registry_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryEnabled(val)
        }
        return nil
    }
    res["registry_v2_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryV2Enabled(val)
        }
        return nil
    }
    return res
}
// GetRegistryEnabled gets the registry_enabled property value. Whether GitHub Packages is enabled globally in a GHES installation
// returns a *bool when successful
func (m *ServerStatisticsPackages) GetRegistryEnabled()(*bool) {
    return m.registry_enabled
}
// GetRegistryV2Enabled gets the registry_v2_enabled property value. Whether a beta registry is enabled in a GHES installation
// returns a *bool when successful
func (m *ServerStatisticsPackages) GetRegistryV2Enabled()(*bool) {
    return m.registry_v2_enabled
}
// Serialize serializes information the current object
func (m *ServerStatisticsPackages) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEcosystems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEcosystems()))
        for i, v := range m.GetEcosystems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ecosystems", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("registry_enabled", m.GetRegistryEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("registry_v2_enabled", m.GetRegistryV2Enabled())
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
func (m *ServerStatisticsPackages) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEcosystems sets the ecosystems property value. The details of the package ecosystems that are enabled in a GHES installation
func (m *ServerStatisticsPackages) SetEcosystems(value []ServerStatisticsPackages_ecosystemsable)() {
    m.ecosystems = value
}
// SetRegistryEnabled sets the registry_enabled property value. Whether GitHub Packages is enabled globally in a GHES installation
func (m *ServerStatisticsPackages) SetRegistryEnabled(value *bool)() {
    m.registry_enabled = value
}
// SetRegistryV2Enabled sets the registry_v2_enabled property value. Whether a beta registry is enabled in a GHES installation
func (m *ServerStatisticsPackages) SetRegistryV2Enabled(value *bool)() {
    m.registry_v2_enabled = value
}
type ServerStatisticsPackagesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEcosystems()([]ServerStatisticsPackages_ecosystemsable)
    GetRegistryEnabled()(*bool)
    GetRegistryV2Enabled()(*bool)
    SetEcosystems(value []ServerStatisticsPackages_ecosystemsable)()
    SetRegistryEnabled(value *bool)()
    SetRegistryV2Enabled(value *bool)()
}
