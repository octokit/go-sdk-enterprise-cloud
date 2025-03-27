package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // ID of the bypass review.
    bypass_review_id *int32
}
// NewItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse instantiates a new ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse and sets the default values.
func NewItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse()(*ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse) {
    m := &ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassReviewId gets the bypass_review_id property value. ID of the bypass review.
// returns a *int32 when successful
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse) GetBypassReviewId()(*int32) {
    return m.bypass_review_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bypass_review_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassReviewId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("bypass_review_id", m.GetBypassReviewId())
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
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassReviewId sets the bypass_review_id property value. ID of the bypass review.
func (m *ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponse) SetBypassReviewId(value *int32)() {
    m.bypass_review_id = value
}
type ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassReviewId()(*int32)
    SetBypassReviewId(value *int32)()
}
