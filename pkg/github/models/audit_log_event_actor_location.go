package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuditLogEvent_actor_location struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The country_name property
    country_name *string
}
// NewAuditLogEvent_actor_location instantiates a new AuditLogEvent_actor_location and sets the default values.
func NewAuditLogEvent_actor_location()(*AuditLogEvent_actor_location) {
    m := &AuditLogEvent_actor_location{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuditLogEvent_actor_locationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditLogEvent_actor_locationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogEvent_actor_location(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuditLogEvent_actor_location) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCountryName gets the country_name property value. The country_name property
// returns a *string when successful
func (m *AuditLogEvent_actor_location) GetCountryName()(*string) {
    return m.country_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuditLogEvent_actor_location) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["country_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuditLogEvent_actor_location) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("country_name", m.GetCountryName())
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
func (m *AuditLogEvent_actor_location) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCountryName sets the country_name property value. The country_name property
func (m *AuditLogEvent_actor_location) SetCountryName(value *string)() {
    m.country_name = value
}
type AuditLogEvent_actor_locationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCountryName()(*string)
    SetCountryName(value *string)()
}
