package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HecConfig hec Config for Audit Log Stream Configuration
type HecConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Domain of Hec instance.
    domain *string
    // Encrypted Token.
    encrypted_token *string
    // Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
    key_id *string
    // Path to send events to.
    path *string
    // The port number for connecting to HEC.
    port *int32
    // SSL verification helps ensure your events are sent to your HEC endpoint securely.
    ssl_verify *bool
}
// NewHecConfig instantiates a new HecConfig and sets the default values.
func NewHecConfig()(*HecConfig) {
    m := &HecConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHecConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHecConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHecConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HecConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDomain gets the domain property value. Domain of Hec instance.
// returns a *string when successful
func (m *HecConfig) GetDomain()(*string) {
    return m.domain
}
// GetEncryptedToken gets the encrypted_token property value. Encrypted Token.
// returns a *string when successful
func (m *HecConfig) GetEncryptedToken()(*string) {
    return m.encrypted_token
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HecConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
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
func (m *HecConfig) GetKeyId()(*string) {
    return m.key_id
}
// GetPath gets the path property value. Path to send events to.
// returns a *string when successful
func (m *HecConfig) GetPath()(*string) {
    return m.path
}
// GetPort gets the port property value. The port number for connecting to HEC.
// returns a *int32 when successful
func (m *HecConfig) GetPort()(*int32) {
    return m.port
}
// GetSslVerify gets the ssl_verify property value. SSL verification helps ensure your events are sent to your HEC endpoint securely.
// returns a *bool when successful
func (m *HecConfig) GetSslVerify()(*bool) {
    return m.ssl_verify
}
// Serialize serializes information the current object
func (m *HecConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("path", m.GetPath())
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
func (m *HecConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDomain sets the domain property value. Domain of Hec instance.
func (m *HecConfig) SetDomain(value *string)() {
    m.domain = value
}
// SetEncryptedToken sets the encrypted_token property value. Encrypted Token.
func (m *HecConfig) SetEncryptedToken(value *string)() {
    m.encrypted_token = value
}
// SetKeyId sets the key_id property value. Key ID obtained from the audit log stream key endpoint used to encrypt secrets.
func (m *HecConfig) SetKeyId(value *string)() {
    m.key_id = value
}
// SetPath sets the path property value. Path to send events to.
func (m *HecConfig) SetPath(value *string)() {
    m.path = value
}
// SetPort sets the port property value. The port number for connecting to HEC.
func (m *HecConfig) SetPort(value *int32)() {
    m.port = value
}
// SetSslVerify sets the ssl_verify property value. SSL verification helps ensure your events are sent to your HEC endpoint securely.
func (m *HecConfig) SetSslVerify(value *bool)() {
    m.ssl_verify = value
}
type HecConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomain()(*string)
    GetEncryptedToken()(*string)
    GetKeyId()(*string)
    GetPath()(*string)
    GetPort()(*int32)
    GetSslVerify()(*bool)
    SetDomain(value *string)()
    SetEncryptedToken(value *string)()
    SetKeyId(value *string)()
    SetPath(value *string)()
    SetPort(value *int32)()
    SetSslVerify(value *bool)()
}
