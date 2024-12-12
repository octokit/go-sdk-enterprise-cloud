package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GoogleCloudConfig google Cloud Config for audit log streaming configuration.
type GoogleCloudConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Google Cloud Bucket Name
    bucket *string
    // The encrypted_json_credentials property
    encrypted_json_credentials *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
}
// NewGoogleCloudConfig instantiates a new GoogleCloudConfig and sets the default values.
func NewGoogleCloudConfig()(*GoogleCloudConfig) {
    m := &GoogleCloudConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGoogleCloudConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGoogleCloudConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGoogleCloudConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GoogleCloudConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBucket gets the bucket property value. Google Cloud Bucket Name
// returns a *string when successful
func (m *GoogleCloudConfig) GetBucket()(*string) {
    return m.bucket
}
// GetEncryptedJsonCredentials gets the encrypted_json_credentials property value. The encrypted_json_credentials property
// returns a *string when successful
func (m *GoogleCloudConfig) GetEncryptedJsonCredentials()(*string) {
    return m.encrypted_json_credentials
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GoogleCloudConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bucket"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucket(val)
        }
        return nil
    }
    res["encrypted_json_credentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedJsonCredentials(val)
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
func (m *GoogleCloudConfig) GetKeyId()(*string) {
    return m.key_id
}
// Serialize serializes information the current object
func (m *GoogleCloudConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("bucket", m.GetBucket())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encrypted_json_credentials", m.GetEncryptedJsonCredentials())
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
func (m *GoogleCloudConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBucket sets the bucket property value. Google Cloud Bucket Name
func (m *GoogleCloudConfig) SetBucket(value *string)() {
    m.bucket = value
}
// SetEncryptedJsonCredentials sets the encrypted_json_credentials property value. The encrypted_json_credentials property
func (m *GoogleCloudConfig) SetEncryptedJsonCredentials(value *string)() {
    m.encrypted_json_credentials = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *GoogleCloudConfig) SetKeyId(value *string)() {
    m.key_id = value
}
type GoogleCloudConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBucket()(*string)
    GetEncryptedJsonCredentials()(*string)
    GetKeyId()(*string)
    SetBucket(value *string)()
    SetEncryptedJsonCredentials(value *string)()
    SetKeyId(value *string)()
}
