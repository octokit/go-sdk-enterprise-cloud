package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecurityAndAnalysis_secret_scanning_validity_checks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The status property
    status *SecurityAndAnalysis_secret_scanning_validity_checks_status
}
// NewSecurityAndAnalysis_secret_scanning_validity_checks instantiates a new SecurityAndAnalysis_secret_scanning_validity_checks and sets the default values.
func NewSecurityAndAnalysis_secret_scanning_validity_checks()(*SecurityAndAnalysis_secret_scanning_validity_checks) {
    m := &SecurityAndAnalysis_secret_scanning_validity_checks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityAndAnalysis_secret_scanning_validity_checksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityAndAnalysis_secret_scanning_validity_checksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityAndAnalysis_secret_scanning_validity_checks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecurityAndAnalysis_secret_scanning_validity_checks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityAndAnalysis_secret_scanning_validity_checks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAndAnalysis_secret_scanning_validity_checks_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SecurityAndAnalysis_secret_scanning_validity_checks_status))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
// returns a *SecurityAndAnalysis_secret_scanning_validity_checks_status when successful
func (m *SecurityAndAnalysis_secret_scanning_validity_checks) GetStatus()(*SecurityAndAnalysis_secret_scanning_validity_checks_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *SecurityAndAnalysis_secret_scanning_validity_checks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *SecurityAndAnalysis_secret_scanning_validity_checks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatus sets the status property value. The status property
func (m *SecurityAndAnalysis_secret_scanning_validity_checks) SetStatus(value *SecurityAndAnalysis_secret_scanning_validity_checks_status)() {
    m.status = value
}
type SecurityAndAnalysis_secret_scanning_validity_checksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*SecurityAndAnalysis_secret_scanning_validity_checks_status)
    SetStatus(value *SecurityAndAnalysis_secret_scanning_validity_checks_status)()
}
