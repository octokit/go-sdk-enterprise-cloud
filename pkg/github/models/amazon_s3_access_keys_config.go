package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AmazonS3AccessKeysConfig amazon S3 Access Keys Config for audit log streaming configuration.
type AmazonS3AccessKeysConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Authentication Type for Amazon S3.
    authentication_type *AmazonS3AccessKeysConfig_authentication_type
    // Amazon S3 Bucket Name.
    bucket *string
    // Encrypted AWS Access Key ID.
    encrypted_access_key_id *string
    // Encrypted AWS Secret Key.
    encrypted_secret_key *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
    // Amazon S3 Bucket Name.
    region *string
}
// NewAmazonS3AccessKeysConfig instantiates a new AmazonS3AccessKeysConfig and sets the default values.
func NewAmazonS3AccessKeysConfig()(*AmazonS3AccessKeysConfig) {
    m := &AmazonS3AccessKeysConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAmazonS3AccessKeysConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAmazonS3AccessKeysConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAmazonS3AccessKeysConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AmazonS3AccessKeysConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthenticationType gets the authentication_type property value. Authentication Type for Amazon S3.
// returns a *AmazonS3AccessKeysConfig_authentication_type when successful
func (m *AmazonS3AccessKeysConfig) GetAuthenticationType()(*AmazonS3AccessKeysConfig_authentication_type) {
    return m.authentication_type
}
// GetBucket gets the bucket property value. Amazon S3 Bucket Name.
// returns a *string when successful
func (m *AmazonS3AccessKeysConfig) GetBucket()(*string) {
    return m.bucket
}
// GetEncryptedAccessKeyId gets the encrypted_access_key_id property value. Encrypted AWS Access Key ID.
// returns a *string when successful
func (m *AmazonS3AccessKeysConfig) GetEncryptedAccessKeyId()(*string) {
    return m.encrypted_access_key_id
}
// GetEncryptedSecretKey gets the encrypted_secret_key property value. Encrypted AWS Secret Key.
// returns a *string when successful
func (m *AmazonS3AccessKeysConfig) GetEncryptedSecretKey()(*string) {
    return m.encrypted_secret_key
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AmazonS3AccessKeysConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authentication_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAmazonS3AccessKeysConfig_authentication_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val.(*AmazonS3AccessKeysConfig_authentication_type))
        }
        return nil
    }
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
    res["encrypted_access_key_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedAccessKeyId(val)
        }
        return nil
    }
    res["encrypted_secret_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedSecretKey(val)
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
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
// returns a *string when successful
func (m *AmazonS3AccessKeysConfig) GetKeyId()(*string) {
    return m.key_id
}
// GetRegion gets the region property value. Amazon S3 Bucket Name.
// returns a *string when successful
func (m *AmazonS3AccessKeysConfig) GetRegion()(*string) {
    return m.region
}
// Serialize serializes information the current object
func (m *AmazonS3AccessKeysConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAuthenticationType() != nil {
        cast := (*m.GetAuthenticationType()).String()
        err := writer.WriteStringValue("authentication_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bucket", m.GetBucket())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encrypted_access_key_id", m.GetEncryptedAccessKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encrypted_secret_key", m.GetEncryptedSecretKey())
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
        err := writer.WriteStringValue("region", m.GetRegion())
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
func (m *AmazonS3AccessKeysConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthenticationType sets the authentication_type property value. Authentication Type for Amazon S3.
func (m *AmazonS3AccessKeysConfig) SetAuthenticationType(value *AmazonS3AccessKeysConfig_authentication_type)() {
    m.authentication_type = value
}
// SetBucket sets the bucket property value. Amazon S3 Bucket Name.
func (m *AmazonS3AccessKeysConfig) SetBucket(value *string)() {
    m.bucket = value
}
// SetEncryptedAccessKeyId sets the encrypted_access_key_id property value. Encrypted AWS Access Key ID.
func (m *AmazonS3AccessKeysConfig) SetEncryptedAccessKeyId(value *string)() {
    m.encrypted_access_key_id = value
}
// SetEncryptedSecretKey sets the encrypted_secret_key property value. Encrypted AWS Secret Key.
func (m *AmazonS3AccessKeysConfig) SetEncryptedSecretKey(value *string)() {
    m.encrypted_secret_key = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *AmazonS3AccessKeysConfig) SetKeyId(value *string)() {
    m.key_id = value
}
// SetRegion sets the region property value. Amazon S3 Bucket Name.
func (m *AmazonS3AccessKeysConfig) SetRegion(value *string)() {
    m.region = value
}
type AmazonS3AccessKeysConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationType()(*AmazonS3AccessKeysConfig_authentication_type)
    GetBucket()(*string)
    GetEncryptedAccessKeyId()(*string)
    GetEncryptedSecretKey()(*string)
    GetKeyId()(*string)
    GetRegion()(*string)
    SetAuthenticationType(value *AmazonS3AccessKeysConfig_authentication_type)()
    SetBucket(value *string)()
    SetEncryptedAccessKeyId(value *string)()
    SetEncryptedSecretKey(value *string)()
    SetKeyId(value *string)()
    SetRegion(value *string)()
}
