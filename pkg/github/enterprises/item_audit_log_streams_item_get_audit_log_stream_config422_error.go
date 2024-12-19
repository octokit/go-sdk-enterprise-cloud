package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The errors property
    errors []string
}
// NewItemAuditLogStreamsItemGetAuditLogStreamConfig422Error instantiates a new ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error and sets the default values.
func NewItemAuditLogStreamsItemGetAuditLogStreamConfig422Error()(*ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) {
    m := &ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAuditLogStreamsItemGetAuditLogStreamConfig422ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAuditLogStreamsItemGetAuditLogStreamConfig422ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAuditLogStreamsItemGetAuditLogStreamConfig422Error(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetErrors gets the errors property value. The errors property
// returns a []string when successful
func (m *ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) GetErrors()([]string) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetErrors(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetErrors() != nil {
        err := writer.WriteCollectionOfStringValues("errors", m.GetErrors())
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
func (m *ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetErrors sets the errors property value. The errors property
func (m *ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error) SetErrors(value []string)() {
    m.errors = value
}
type ItemAuditLogStreamsItemGetAuditLogStreamConfig422Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrors()([]string)
    SetErrors(value []string)()
}
