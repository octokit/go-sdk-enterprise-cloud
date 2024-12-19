package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureBlobConfig azure Blob Config for audit log streaming configuration.
type AzureBlobConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The encrypted_sas_url property
    encrypted_sas_url *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
}
// NewAzureBlobConfig instantiates a new AzureBlobConfig and sets the default values.
func NewAzureBlobConfig()(*AzureBlobConfig) {
    m := &AzureBlobConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAzureBlobConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureBlobConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureBlobConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AzureBlobConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncryptedSasUrl gets the encrypted_sas_url property value. The encrypted_sas_url property
// returns a *string when successful
func (m *AzureBlobConfig) GetEncryptedSasUrl()(*string) {
    return m.encrypted_sas_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureBlobConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encrypted_sas_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedSasUrl(val)
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
    return res
}
// GetKeyId gets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
// returns a *string when successful
func (m *AzureBlobConfig) GetKeyId()(*string) {
    return m.key_id
}
// Serialize serializes information the current object
func (m *AzureBlobConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encrypted_sas_url", m.GetEncryptedSasUrl())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureBlobConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncryptedSasUrl sets the encrypted_sas_url property value. The encrypted_sas_url property
func (m *AzureBlobConfig) SetEncryptedSasUrl(value *string)() {
    m.encrypted_sas_url = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *AzureBlobConfig) SetKeyId(value *string)() {
    m.key_id = value
}
type AzureBlobConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedSasUrl()(*string)
    GetKeyId()(*string)
    SetEncryptedSasUrl(value *string)()
    SetKeyId(value *string)()
}
