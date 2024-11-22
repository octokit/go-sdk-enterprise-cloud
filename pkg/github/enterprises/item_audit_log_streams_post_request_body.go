package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

type ItemAuditLogStreamsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This setting pauses or resumes a stream.
    enabled *bool
    // The vendor_specific property
    vendor_specific ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable
}
// ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific composed type wrapper for classes i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable
type ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific struct {
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable
    amazonS3AccessKeysConfig i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable
    amazonS3OidcConfig i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable
    azureBlobConfig i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable
    azureHubConfig i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable
    datadogConfig i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable
    googleCloudConfig i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable
    splunkConfig i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable
}
// NewItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific instantiates a new ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific and sets the default values.
func NewItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific()(*ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) {
    m := &ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific{
    }
    return m
}
// CreateItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetAmazonS3AccessKeysConfig gets the amazonS3AccessKeysConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable
// returns a AmazonS3AccessKeysConfigable when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetAmazonS3AccessKeysConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable) {
    return m.amazonS3AccessKeysConfig
}
// GetAmazonS3OidcConfig gets the amazonS3OidcConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable
// returns a AmazonS3OidcConfigable when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetAmazonS3OidcConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable) {
    return m.amazonS3OidcConfig
}
// GetAzureBlobConfig gets the azureBlobConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable
// returns a AzureBlobConfigable when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetAzureBlobConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable) {
    return m.azureBlobConfig
}
// GetAzureHubConfig gets the azureHubConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable
// returns a AzureHubConfigable when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetAzureHubConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable) {
    return m.azureHubConfig
}
// GetDatadogConfig gets the datadogConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable
// returns a DatadogConfigable when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetDatadogConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable) {
    return m.datadogConfig
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetGoogleCloudConfig gets the googleCloudConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable
// returns a GoogleCloudConfigable when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetGoogleCloudConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable) {
    return m.googleCloudConfig
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetIsComposedType()(bool) {
    return true
}
// GetSplunkConfig gets the splunkConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable
// returns a SplunkConfigable when successful
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) GetSplunkConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable) {
    return m.splunkConfig
}
// Serialize serializes information the current object
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAmazonS3AccessKeysConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAmazonS3AccessKeysConfig())
        if err != nil {
            return err
        }
    } else if m.GetAmazonS3OidcConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAmazonS3OidcConfig())
        if err != nil {
            return err
        }
    } else if m.GetAzureBlobConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAzureBlobConfig())
        if err != nil {
            return err
        }
    } else if m.GetAzureHubConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAzureHubConfig())
        if err != nil {
            return err
        }
    } else if m.GetDatadogConfig() != nil {
        err := writer.WriteObjectValue("", m.GetDatadogConfig())
        if err != nil {
            return err
        }
    } else if m.GetGoogleCloudConfig() != nil {
        err := writer.WriteObjectValue("", m.GetGoogleCloudConfig())
        if err != nil {
            return err
        }
    } else if m.GetSplunkConfig() != nil {
        err := writer.WriteObjectValue("", m.GetSplunkConfig())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAmazonS3AccessKeysConfig sets the amazonS3AccessKeysConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) SetAmazonS3AccessKeysConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable)() {
    m.amazonS3AccessKeysConfig = value
}
// SetAmazonS3OidcConfig sets the amazonS3OidcConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) SetAmazonS3OidcConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable)() {
    m.amazonS3OidcConfig = value
}
// SetAzureBlobConfig sets the azureBlobConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) SetAzureBlobConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable)() {
    m.azureBlobConfig = value
}
// SetAzureHubConfig sets the azureHubConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) SetAzureHubConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable)() {
    m.azureHubConfig = value
}
// SetDatadogConfig sets the datadogConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) SetDatadogConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable)() {
    m.datadogConfig = value
}
// SetGoogleCloudConfig sets the googleCloudConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) SetGoogleCloudConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable)() {
    m.googleCloudConfig = value
}
// SetSplunkConfig sets the splunkConfig property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable
func (m *ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specific) SetSplunkConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable)() {
    m.splunkConfig = value
}
type ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmazonS3AccessKeysConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable)
    GetAmazonS3OidcConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable)
    GetAzureBlobConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable)
    GetAzureHubConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable)
    GetDatadogConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable)
    GetGoogleCloudConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable)
    GetSplunkConfig()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable)
    SetAmazonS3AccessKeysConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3AccessKeysConfigable)()
    SetAmazonS3OidcConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AmazonS3OidcConfigable)()
    SetAzureBlobConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureBlobConfigable)()
    SetAzureHubConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.AzureHubConfigable)()
    SetDatadogConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DatadogConfigable)()
    SetGoogleCloudConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GoogleCloudConfigable)()
    SetSplunkConfig(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SplunkConfigable)()
}
// NewItemAuditLogStreamsPostRequestBody instantiates a new ItemAuditLogStreamsPostRequestBody and sets the default values.
func NewItemAuditLogStreamsPostRequestBody()(*ItemAuditLogStreamsPostRequestBody) {
    m := &ItemAuditLogStreamsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAuditLogStreamsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAuditLogStreamsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAuditLogStreamsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAuditLogStreamsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. This setting pauses or resumes a stream.
// returns a *bool when successful
func (m *ItemAuditLogStreamsPostRequestBody) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAuditLogStreamsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["vendor_specific"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorSpecific(val.(ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable))
        }
        return nil
    }
    return res
}
// GetVendorSpecific gets the vendor_specific property value. The vendor_specific property
// returns a ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable when successful
func (m *ItemAuditLogStreamsPostRequestBody) GetVendorSpecific()(ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable) {
    return m.vendor_specific
}
// Serialize serializes information the current object
func (m *ItemAuditLogStreamsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vendor_specific", m.GetVendorSpecific())
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
func (m *ItemAuditLogStreamsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. This setting pauses or resumes a stream.
func (m *ItemAuditLogStreamsPostRequestBody) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetVendorSpecific sets the vendor_specific property value. The vendor_specific property
func (m *ItemAuditLogStreamsPostRequestBody) SetVendorSpecific(value ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable)() {
    m.vendor_specific = value
}
type ItemAuditLogStreamsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetVendorSpecific()(ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable)
    SetEnabled(value *bool)()
    SetVendorSpecific(value ItemAuditLogStreamsPostRequestBody_StreamsPostRequestBody_vendor_specificable)()
}
