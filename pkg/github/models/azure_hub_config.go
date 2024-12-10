package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureHubConfig azure Event Hubs Config for audit log streaming configuration.
type AzureHubConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Encrypted Connection String for Azure Event Hubs
    encrypted_connstring *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
    // Instance name of Azure Event Hubs
    name *string
}
// NewAzureHubConfig instantiates a new AzureHubConfig and sets the default values.
func NewAzureHubConfig()(*AzureHubConfig) {
    m := &AzureHubConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAzureHubConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureHubConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureHubConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AzureHubConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncryptedConnstring gets the encrypted_connstring property value. Encrypted Connection String for Azure Event Hubs
// returns a *string when successful
func (m *AzureHubConfig) GetEncryptedConnstring()(*string) {
    return m.encrypted_connstring
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureHubConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encrypted_connstring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedConnstring(val)
        }
        return nil
    }
    res["key_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
// returns a *string when successful
func (m *AzureHubConfig) GetKeyId()(*string) {
    return m.key_id
}
// GetName gets the name property value. Instance name of Azure Event Hubs
// returns a *string when successful
func (m *AzureHubConfig) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *AzureHubConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encrypted_connstring", m.GetEncryptedConnstring())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key_id", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *AzureHubConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncryptedConnstring sets the encrypted_connstring property value. Encrypted Connection String for Azure Event Hubs
func (m *AzureHubConfig) SetEncryptedConnstring(value *string)() {
    m.encrypted_connstring = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *AzureHubConfig) SetKeyId(value *string)() {
    m.key_id = value
}
// SetName sets the name property value. Instance name of Azure Event Hubs
func (m *AzureHubConfig) SetName(value *string)() {
    m.name = value
}
type AzureHubConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedConnstring()(*string)
    GetKeyId()(*string)
    GetName()(*string)
    SetEncryptedConnstring(value *string)()
    SetKeyId(value *string)()
    SetName(value *string)()
}
