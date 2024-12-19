package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SplunkConfig splunk Config for Audit Log Stream Configuration
type SplunkConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Domain of Splunk instance.
    domain *string
    // Encrypted Token.
    encrypted_token *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
    // The port number for connecting to Splunk.
    port *int32
    // SSL verification helps ensure your events are sent to your Splunk endpoint securely.
    ssl_verify *bool
}
// NewSplunkConfig instantiates a new SplunkConfig and sets the default values.
func NewSplunkConfig()(*SplunkConfig) {
    m := &SplunkConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSplunkConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSplunkConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSplunkConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SplunkConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDomain gets the domain property value. Domain of Splunk instance.
// returns a *string when successful
func (m *SplunkConfig) GetDomain()(*string) {
    return m.domain
}
// GetEncryptedToken gets the encrypted_token property value. Encrypted Token.
// returns a *string when successful
func (m *SplunkConfig) GetEncryptedToken()(*string) {
    return m.encrypted_token
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SplunkConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
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
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["ssl_verify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSslVerify(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
// returns a *string when successful
func (m *SplunkConfig) GetKeyId()(*string) {
    return m.key_id
}
// GetPort gets the port property value. The port number for connecting to Splunk.
// returns a *int32 when successful
func (m *SplunkConfig) GetPort()(*int32) {
    return m.port
}
// GetSslVerify gets the ssl_verify property value. SSL verification helps ensure your events are sent to your Splunk endpoint securely.
// returns a *bool when successful
func (m *SplunkConfig) GetSslVerify()(*bool) {
    return m.ssl_verify
}
// Serialize serializes information the current object
func (m *SplunkConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
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
    {
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ssl_verify", m.GetSslVerify())
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
func (m *SplunkConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDomain sets the domain property value. Domain of Splunk instance.
func (m *SplunkConfig) SetDomain(value *string)() {
    m.domain = value
}
// SetEncryptedToken sets the encrypted_token property value. Encrypted Token.
func (m *SplunkConfig) SetEncryptedToken(value *string)() {
    m.encrypted_token = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *SplunkConfig) SetKeyId(value *string)() {
    m.key_id = value
}
// SetPort sets the port property value. The port number for connecting to Splunk.
func (m *SplunkConfig) SetPort(value *int32)() {
    m.port = value
}
// SetSslVerify sets the ssl_verify property value. SSL verification helps ensure your events are sent to your Splunk endpoint securely.
func (m *SplunkConfig) SetSslVerify(value *bool)() {
    m.ssl_verify = value
}
type SplunkConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomain()(*string)
    GetEncryptedToken()(*string)
    GetKeyId()(*string)
    GetPort()(*int32)
    GetSslVerify()(*bool)
    SetDomain(value *string)()
    SetEncryptedToken(value *string)()
    SetKeyId(value *string)()
    SetPort(value *int32)()
    SetSslVerify(value *bool)()
}
