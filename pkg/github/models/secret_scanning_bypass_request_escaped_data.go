package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretScanningBypassRequest_data struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The branch in the repo where the secret was located during the request.
    branch *string
    // The reason the bypass was requested.
    bypass_reason *SecretScanningBypassRequest_data_bypass_reason
    // The path in the repo where the secret was located during the request.
    path *string
    // The type of secret that secret scanning detected.
    secret_type *string
}
// NewSecretScanningBypassRequest_data instantiates a new SecretScanningBypassRequest_data and sets the default values.
func NewSecretScanningBypassRequest_data()(*SecretScanningBypassRequest_data) {
    m := &SecretScanningBypassRequest_data{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningBypassRequest_dataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningBypassRequest_dataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningBypassRequest_data(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningBypassRequest_data) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranch gets the branch property value. The branch in the repo where the secret was located during the request.
// returns a *string when successful
func (m *SecretScanningBypassRequest_data) GetBranch()(*string) {
    return m.branch
}
// GetBypassReason gets the bypass_reason property value. The reason the bypass was requested.
// returns a *SecretScanningBypassRequest_data_bypass_reason when successful
func (m *SecretScanningBypassRequest_data) GetBypassReason()(*SecretScanningBypassRequest_data_bypass_reason) {
    return m.bypass_reason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningBypassRequest_data) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranch(val)
        }
        return nil
    }
    res["bypass_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningBypassRequest_data_bypass_reason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassReason(val.(*SecretScanningBypassRequest_data_bypass_reason))
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
    res["secret_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretType(val)
        }
        return nil
    }
    return res
}
// GetPath gets the path property value. The path in the repo where the secret was located during the request.
// returns a *string when successful
func (m *SecretScanningBypassRequest_data) GetPath()(*string) {
    return m.path
}
// GetSecretType gets the secret_type property value. The type of secret that secret scanning detected.
// returns a *string when successful
func (m *SecretScanningBypassRequest_data) GetSecretType()(*string) {
    return m.secret_type
}
// Serialize serializes information the current object
func (m *SecretScanningBypassRequest_data) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("branch", m.GetBranch())
        if err != nil {
            return err
        }
    }
    if m.GetBypassReason() != nil {
        cast := (*m.GetBypassReason()).String()
        err := writer.WriteStringValue("bypass_reason", &cast)
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
        err := writer.WriteStringValue("secret_type", m.GetSecretType())
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
func (m *SecretScanningBypassRequest_data) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranch sets the branch property value. The branch in the repo where the secret was located during the request.
func (m *SecretScanningBypassRequest_data) SetBranch(value *string)() {
    m.branch = value
}
// SetBypassReason sets the bypass_reason property value. The reason the bypass was requested.
func (m *SecretScanningBypassRequest_data) SetBypassReason(value *SecretScanningBypassRequest_data_bypass_reason)() {
    m.bypass_reason = value
}
// SetPath sets the path property value. The path in the repo where the secret was located during the request.
func (m *SecretScanningBypassRequest_data) SetPath(value *string)() {
    m.path = value
}
// SetSecretType sets the secret_type property value. The type of secret that secret scanning detected.
func (m *SecretScanningBypassRequest_data) SetSecretType(value *string)() {
    m.secret_type = value
}
type SecretScanningBypassRequest_dataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranch()(*string)
    GetBypassReason()(*SecretScanningBypassRequest_data_bypass_reason)
    GetPath()(*string)
    GetSecretType()(*string)
    SetBranch(value *string)()
    SetBypassReason(value *SecretScanningBypassRequest_data_bypass_reason)()
    SetPath(value *string)()
    SetSecretType(value *string)()
}
