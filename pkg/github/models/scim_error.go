package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScimError scim Error
type ScimError struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The detail property
    detail *string
    // The documentation_url property
    documentation_url *string
    // The message property
    message *string
    // The schemas property
    schemas []string
    // The scimType property
    scimType *string
    // The status property
    status *int32
}
// NewScimError instantiates a new ScimError and sets the default values.
func NewScimError()(*ScimError) {
    m := &ScimError{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimError(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *ScimError) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimError) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetail gets the detail property value. The detail property
// returns a *string when successful
func (m *ScimError) GetDetail()(*string) {
    return m.detail
}
// GetDocumentationUrl gets the documentation_url property value. The documentation_url property
// returns a *string when successful
func (m *ScimError) GetDocumentationUrl()(*string) {
    return m.documentation_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val)
        }
        return nil
    }
    res["documentation_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
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
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    res["scimType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScimType(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *ScimError) GetMessage()(*string) {
    return m.message
}
// GetSchemas gets the schemas property value. The schemas property
// returns a []string when successful
func (m *ScimError) GetSchemas()([]string) {
    return m.schemas
}
// GetScimType gets the scimType property value. The scimType property
// returns a *string when successful
func (m *ScimError) GetScimType()(*string) {
    return m.scimType
}
// GetStatus gets the status property value. The status property
// returns a *int32 when successful
func (m *ScimError) GetStatus()(*int32) {
    return m.status
}
// Serialize serializes information the current object
func (m *ScimError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentation_url", m.GetDocumentationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        err := writer.WriteCollectionOfStringValues("schemas", m.GetSchemas())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scimType", m.GetScimType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("status", m.GetStatus())
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
func (m *ScimError) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetail sets the detail property value. The detail property
func (m *ScimError) SetDetail(value *string)() {
    m.detail = value
}
// SetDocumentationUrl sets the documentation_url property value. The documentation_url property
func (m *ScimError) SetDocumentationUrl(value *string)() {
    m.documentation_url = value
}
// SetMessage sets the message property value. The message property
func (m *ScimError) SetMessage(value *string)() {
    m.message = value
}
// SetSchemas sets the schemas property value. The schemas property
func (m *ScimError) SetSchemas(value []string)() {
    m.schemas = value
}
// SetScimType sets the scimType property value. The scimType property
func (m *ScimError) SetScimType(value *string)() {
    m.scimType = value
}
// SetStatus sets the status property value. The status property
func (m *ScimError) SetStatus(value *int32)() {
    m.status = value
}
type ScimErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetail()(*string)
    GetDocumentationUrl()(*string)
    GetMessage()(*string)
    GetSchemas()([]string)
    GetScimType()(*string)
    GetStatus()(*int32)
    SetDetail(value *string)()
    SetDocumentationUrl(value *string)()
    SetMessage(value *string)()
    SetSchemas(value []string)()
    SetScimType(value *string)()
    SetStatus(value *int32)()
}
