package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DatadogConfig datadog Config for audit log streaming configuration.
type DatadogConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Encrypted Splunk token.
    encrypted_token *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
    // Datadog Site to use.
    site *DatadogConfig_site
}
// NewDatadogConfig instantiates a new DatadogConfig and sets the default values.
func NewDatadogConfig()(*DatadogConfig) {
    m := &DatadogConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDatadogConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDatadogConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDatadogConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DatadogConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncryptedToken gets the encrypted_token property value. Encrypted Splunk token.
// returns a *string when successful
func (m *DatadogConfig) GetEncryptedToken()(*string) {
    return m.encrypted_token
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DatadogConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encrypted_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedToken(val)
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
    res["site"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDatadogConfig_site)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val.(*DatadogConfig_site))
        }
        return nil
    }
    return res
}
// GetKeyId gets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
// returns a *string when successful
func (m *DatadogConfig) GetKeyId()(*string) {
    return m.key_id
}
// GetSite gets the site property value. Datadog Site to use.
// returns a *DatadogConfig_site when successful
func (m *DatadogConfig) GetSite()(*DatadogConfig_site) {
    return m.site
}
// Serialize serializes information the current object
func (m *DatadogConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encrypted_token", m.GetEncryptedToken())
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
    if m.GetSite() != nil {
        cast := (*m.GetSite()).String()
        err := writer.WriteStringValue("site", &cast)
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
func (m *DatadogConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncryptedToken sets the encrypted_token property value. Encrypted Splunk token.
func (m *DatadogConfig) SetEncryptedToken(value *string)() {
    m.encrypted_token = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *DatadogConfig) SetKeyId(value *string)() {
    m.key_id = value
}
// SetSite sets the site property value. Datadog Site to use.
func (m *DatadogConfig) SetSite(value *DatadogConfig_site)() {
    m.site = value
}
type DatadogConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedToken()(*string)
    GetKeyId()(*string)
    GetSite()(*DatadogConfig_site)
    SetEncryptedToken(value *string)()
    SetKeyId(value *string)()
    SetSite(value *DatadogConfig_site)()
}
