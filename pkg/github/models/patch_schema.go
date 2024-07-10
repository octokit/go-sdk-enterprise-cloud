package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PatchSchema struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // patch operations list
    operations []PatchSchema_Operationsable
    // The schemas property
    schemas []PatchSchema_schemas
}
// NewPatchSchema instantiates a new PatchSchema and sets the default values.
func NewPatchSchema()(*PatchSchema) {
    m := &PatchSchema{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatchSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePatchSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatchSchema(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PatchSchema) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PatchSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePatchSchema_OperationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PatchSchema_Operationsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PatchSchema_Operationsable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePatchSchema_schemas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PatchSchema_schemas, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*PatchSchema_schemas))
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    return res
}
// GetOperations gets the Operations property value. patch operations list
// returns a []PatchSchema_Operationsable when successful
func (m *PatchSchema) GetOperations()([]PatchSchema_Operationsable) {
    return m.operations
}
// GetSchemas gets the schemas property value. The schemas property
// returns a []PatchSchema_schemas when successful
func (m *PatchSchema) GetSchemas()([]PatchSchema_schemas) {
    return m.schemas
}
// Serialize serializes information the current object
func (m *PatchSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("Operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        err := writer.WriteCollectionOfStringValues("schemas", SerializePatchSchema_schemas(m.GetSchemas()))
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
func (m *PatchSchema) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOperations sets the Operations property value. patch operations list
func (m *PatchSchema) SetOperations(value []PatchSchema_Operationsable)() {
    m.operations = value
}
// SetSchemas sets the schemas property value. The schemas property
func (m *PatchSchema) SetSchemas(value []PatchSchema_schemas)() {
    m.schemas = value
}
type PatchSchemaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOperations()([]PatchSchema_Operationsable)
    GetSchemas()([]PatchSchema_schemas)
    SetOperations(value []PatchSchema_Operationsable)()
    SetSchemas(value []PatchSchema_schemas)()
}
