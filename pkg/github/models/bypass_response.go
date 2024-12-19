package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BypassResponse a response made by a delegated bypasser to a bypass request.
type BypassResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date and time the response to the bypass request was created.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ID of the response to the bypass request.
    id *int32
    // The user who reviewed the bypass request.
    reviewer BypassResponse_reviewerable
    // The response status to the bypass request until dismissed.
    status *BypassResponse_status
}
// NewBypassResponse instantiates a new BypassResponse and sets the default values.
func NewBypassResponse()(*BypassResponse) {
    m := &BypassResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBypassResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBypassResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBypassResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BypassResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The date and time the response to the bypass request was created.
// returns a *Time when successful
func (m *BypassResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BypassResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["reviewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBypassResponse_reviewerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewer(val.(BypassResponse_reviewerable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBypassResponse_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*BypassResponse_status))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the response to the bypass request.
// returns a *int32 when successful
func (m *BypassResponse) GetId()(*int32) {
    return m.id
}
// GetReviewer gets the reviewer property value. The user who reviewed the bypass request.
// returns a BypassResponse_reviewerable when successful
func (m *BypassResponse) GetReviewer()(BypassResponse_reviewerable) {
    return m.reviewer
}
// GetStatus gets the status property value. The response status to the bypass request until dismissed.
// returns a *BypassResponse_status when successful
func (m *BypassResponse) GetStatus()(*BypassResponse_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *BypassResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reviewer", m.GetReviewer())
        if err != nil {
            return err
        }
    }
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
func (m *BypassResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The date and time the response to the bypass request was created.
func (m *BypassResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetId sets the id property value. The ID of the response to the bypass request.
func (m *BypassResponse) SetId(value *int32)() {
    m.id = value
}
// SetReviewer sets the reviewer property value. The user who reviewed the bypass request.
func (m *BypassResponse) SetReviewer(value BypassResponse_reviewerable)() {
    m.reviewer = value
}
// SetStatus sets the status property value. The response status to the bypass request until dismissed.
func (m *BypassResponse) SetStatus(value *BypassResponse_status)() {
    m.status = value
}
type BypassResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int32)
    GetReviewer()(BypassResponse_reviewerable)
    GetStatus()(*BypassResponse_status)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int32)()
    SetReviewer(value BypassResponse_reviewerable)()
    SetStatus(value *BypassResponse_status)()
}
