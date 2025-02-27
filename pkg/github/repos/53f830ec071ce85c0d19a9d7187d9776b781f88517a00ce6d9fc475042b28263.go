package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A message to include with the review. Has a maximum character length of 2048.
    message *string
}
// NewItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody instantiates a new ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody and sets the default values.
func NewItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody()(*ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody) {
    m := &ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. A message to include with the review. Has a maximum character length of 2048.
// returns a *string when successful
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. A message to include with the review. Has a maximum character length of 2048.
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBody) SetMessage(value *string)() {
    m.message = value
}
type ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(*string)
    SetMessage(value *string)()
}
