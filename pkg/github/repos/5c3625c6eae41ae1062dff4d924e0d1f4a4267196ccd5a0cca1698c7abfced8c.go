package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks use the `status` property to enable or disable secret scanning automatic validity checks on supported partner tokens for this repository.
type ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Can be `enabled` or `disabled`.
    status *string
}
// NewItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks instantiates a new ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks and sets the default values.
func NewItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks()(*ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks) {
    m := &ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Can be `enabled` or `disabled`.
// returns a *string when successful
func (m *ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatus sets the status property value. Can be `enabled` or `disabled`.
func (m *ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checks) SetStatus(value *string)() {
    m.status = value
}
type ItemItemRepoPatchRequestBody_security_and_analysis_secret_scanning_validity_checksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*string)
    SetStatus(value *string)()
}