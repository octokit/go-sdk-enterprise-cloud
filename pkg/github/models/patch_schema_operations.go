package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PatchSchema_Operations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The op property
    op *PatchSchema_Operations_op
    // The path property
    path *string
    // Corresponding 'value' of that field specified by 'path'
    value *string
}
// NewPatchSchema_Operations instantiates a new PatchSchema_Operations and sets the default values.
func NewPatchSchema_Operations()(*PatchSchema_Operations) {
    m := &PatchSchema_Operations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatchSchema_OperationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePatchSchema_OperationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatchSchema_Operations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PatchSchema_Operations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PatchSchema_Operations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["op"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePatchSchema_Operations_op)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOp(val.(*PatchSchema_Operations_op))
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetOp gets the op property value. The op property
// returns a *PatchSchema_Operations_op when successful
func (m *PatchSchema_Operations) GetOp()(*PatchSchema_Operations_op) {
    return m.op
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *PatchSchema_Operations) GetPath()(*string) {
    return m.path
}
// GetValue gets the value property value. Corresponding 'value' of that field specified by 'path'
// returns a *string when successful
func (m *PatchSchema_Operations) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *PatchSchema_Operations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOp() != nil {
        cast := (*m.GetOp()).String()
        err := writer.WriteStringValue("op", &cast)
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
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *PatchSchema_Operations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOp sets the op property value. The op property
func (m *PatchSchema_Operations) SetOp(value *PatchSchema_Operations_op)() {
    m.op = value
}
// SetPath sets the path property value. The path property
func (m *PatchSchema_Operations) SetPath(value *string)() {
    m.path = value
}
// SetValue sets the value property value. Corresponding 'value' of that field specified by 'path'
func (m *PatchSchema_Operations) SetValue(value *string)() {
    m.value = value
}
type PatchSchema_Operationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOp()(*PatchSchema_Operations_op)
    GetPath()(*string)
    GetValue()(*string)
    SetOp(value *PatchSchema_Operations_op)()
    SetPath(value *string)()
    SetValue(value *string)()
}
