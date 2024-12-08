package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AmazonS3OidcConfig amazon S3 OIDC Config for audit log streaming configuration.
type AmazonS3OidcConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The arn_role property
    arn_role *string
    // Authentication Type for Amazon S3.
    authentication_type *AmazonS3OidcConfig_authentication_type
    // Amazon S3 Bucket Name.
    bucket *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
    // AWS S3 Bucket Region.
    region *string
}
// NewAmazonS3OidcConfig instantiates a new AmazonS3OidcConfig and sets the default values.
func NewAmazonS3OidcConfig()(*AmazonS3OidcConfig) {
    m := &AmazonS3OidcConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAmazonS3OidcConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAmazonS3OidcConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAmazonS3OidcConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AmazonS3OidcConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArnRole gets the arn_role property value. The arn_role property
// returns a *string when successful
func (m *AmazonS3OidcConfig) GetArnRole()(*string) {
    return m.arn_role
}
// GetAuthenticationType gets the authentication_type property value. Authentication Type for Amazon S3.
// returns a *AmazonS3OidcConfig_authentication_type when successful
func (m *AmazonS3OidcConfig) GetAuthenticationType()(*AmazonS3OidcConfig_authentication_type) {
    return m.authentication_type
}
// GetBucket gets the bucket property value. Amazon S3 Bucket Name.
// returns a *string when successful
func (m *AmazonS3OidcConfig) GetBucket()(*string) {
    return m.bucket
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AmazonS3OidcConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["arn_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArnRole(val)
        }
        return nil
    }
    res["authentication_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAmazonS3OidcConfig_authentication_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val.(*AmazonS3OidcConfig_authentication_type))
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
func (m *AmazonS3OidcConfig) GetKeyId()(*string) {
    return m.key_id
}
// GetRegion gets the region property value. AWS S3 Bucket Region.
// returns a *string when successful
func (m *AmazonS3OidcConfig) GetRegion()(*string) {
    return m.region
}
// Serialize serializes information the current object
func (m *AmazonS3OidcConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("arn_role", m.GetArnRole())
        if err != nil {
            return err
        }
    }
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
func (m *AmazonS3OidcConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArnRole sets the arn_role property value. The arn_role property
func (m *AmazonS3OidcConfig) SetArnRole(value *string)() {
    m.arn_role = value
}
// SetAuthenticationType sets the authentication_type property value. Authentication Type for Amazon S3.
func (m *AmazonS3OidcConfig) SetAuthenticationType(value *AmazonS3OidcConfig_authentication_type)() {
    m.authentication_type = value
}
// SetBucket sets the bucket property value. Amazon S3 Bucket Name.
func (m *AmazonS3OidcConfig) SetBucket(value *string)() {
    m.bucket = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *AmazonS3OidcConfig) SetKeyId(value *string)() {
    m.key_id = value
}
// SetRegion sets the region property value. AWS S3 Bucket Region.
func (m *AmazonS3OidcConfig) SetRegion(value *string)() {
    m.region = value
}
type AmazonS3OidcConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArnRole()(*string)
    GetAuthenticationType()(*AmazonS3OidcConfig_authentication_type)
    GetBucket()(*string)
    GetKeyId()(*string)
    GetRegion()(*string)
    SetArnRole(value *string)()
    SetAuthenticationType(value *AmazonS3OidcConfig_authentication_type)()
    SetBucket(value *string)()
    SetKeyId(value *string)()
    SetRegion(value *string)()
}
