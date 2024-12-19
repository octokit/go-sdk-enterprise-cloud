package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

type ItemCodeSecurityConfigurationsItemDefaultsPutResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A code security configuration
    configuration i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CodeSecurityConfigurationable
}
// NewItemCodeSecurityConfigurationsItemDefaultsPutResponse instantiates a new ItemCodeSecurityConfigurationsItemDefaultsPutResponse and sets the default values.
func NewItemCodeSecurityConfigurationsItemDefaultsPutResponse()(*ItemCodeSecurityConfigurationsItemDefaultsPutResponse) {
    m := &ItemCodeSecurityConfigurationsItemDefaultsPutResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCodeSecurityConfigurationsItemDefaultsPutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsItemDefaultsPutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsItemDefaultsPutResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCodeSecurityConfigurationsItemDefaultsPutResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfiguration gets the configuration property value. A code security configuration
// returns a CodeSecurityConfigurationable when successful
func (m *ItemCodeSecurityConfigurationsItemDefaultsPutResponse) GetConfiguration()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CodeSecurityConfigurationable) {
    return m.configuration
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsItemDefaultsPutResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateCodeSecurityConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CodeSecurityConfigurationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsItemDefaultsPutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configuration", m.GetConfiguration())
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
func (m *ItemCodeSecurityConfigurationsItemDefaultsPutResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfiguration sets the configuration property value. A code security configuration
func (m *ItemCodeSecurityConfigurationsItemDefaultsPutResponse) SetConfiguration(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CodeSecurityConfigurationable)() {
    m.configuration = value
}
type ItemCodeSecurityConfigurationsItemDefaultsPutResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CodeSecurityConfigurationable)
    SetConfiguration(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CodeSecurityConfigurationable)()
}
