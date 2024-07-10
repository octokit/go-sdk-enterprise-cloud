package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Meta the metadata associated with the creation/updates to the user.
type Meta struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A date and time when the user was created.
    created *string
    // A data and time when the user was last modified.
    lastModified *string
    // A URL location of an object
    location *string
    // A type of a resource
    resourceType *Meta_resourceType
}
// NewMeta instantiates a new Meta and sets the default values.
func NewMeta()(*Meta) {
    m := &Meta{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMetaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMetaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeta(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Meta) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreated gets the created property value. A date and time when the user was created.
// returns a *string when successful
func (m *Meta) GetCreated()(*string) {
    return m.created
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Meta) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreated(val)
        }
        return nil
    }
    res["lastModified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModified(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeta_resourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val.(*Meta_resourceType))
        }
        return nil
    }
    return res
}
// GetLastModified gets the lastModified property value. A data and time when the user was last modified.
// returns a *string when successful
func (m *Meta) GetLastModified()(*string) {
    return m.lastModified
}
// GetLocation gets the location property value. A URL location of an object
// returns a *string when successful
func (m *Meta) GetLocation()(*string) {
    return m.location
}
// GetResourceType gets the resourceType property value. A type of a resource
// returns a *Meta_resourceType when successful
func (m *Meta) GetResourceType()(*Meta_resourceType) {
    return m.resourceType
}
// Serialize serializes information the current object
func (m *Meta) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created", m.GetCreated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastModified", m.GetLastModified())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    if m.GetResourceType() != nil {
        cast := (*m.GetResourceType()).String()
        err := writer.WriteStringValue("resourceType", &cast)
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
func (m *Meta) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreated sets the created property value. A date and time when the user was created.
func (m *Meta) SetCreated(value *string)() {
    m.created = value
}
// SetLastModified sets the lastModified property value. A data and time when the user was last modified.
func (m *Meta) SetLastModified(value *string)() {
    m.lastModified = value
}
// SetLocation sets the location property value. A URL location of an object
func (m *Meta) SetLocation(value *string)() {
    m.location = value
}
// SetResourceType sets the resourceType property value. A type of a resource
func (m *Meta) SetResourceType(value *Meta_resourceType)() {
    m.resourceType = value
}
type Metaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreated()(*string)
    GetLastModified()(*string)
    GetLocation()(*string)
    GetResourceType()(*Meta_resourceType)
    SetCreated(value *string)()
    SetLastModified(value *string)()
    SetLocation(value *string)()
    SetResourceType(value *Meta_resourceType)()
}
