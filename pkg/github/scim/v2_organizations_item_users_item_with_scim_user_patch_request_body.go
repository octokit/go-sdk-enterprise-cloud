package scim

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Set of operations to be performed
    operations []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable
    // The schemas property
    schemas []string
}
// NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody instantiates a new V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody and sets the default values.
func NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody()(*V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) {
    m := &V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_OperationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable)
                }
            }
            m.SetOperations(res)
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
    return res
}
// GetOperations gets the Operations property value. Set of operations to be performed
// returns a []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) GetOperations()([]V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable) {
    return m.operations
}
// GetSchemas gets the schemas property value. The schemas property
// returns a []string when successful
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) GetSchemas()([]string) {
    return m.schemas
}
// Serialize serializes information the current object
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteCollectionOfStringValues("schemas", m.GetSchemas())
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
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOperations sets the Operations property value. Set of operations to be performed
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) SetOperations(value []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable)() {
    m.operations = value
}
// SetSchemas sets the schemas property value. The schemas property
func (m *V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody) SetSchemas(value []string)() {
    m.schemas = value
}
type V2OrganizationsItemUsersItemWithScim_user_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOperations()([]V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable)
    GetSchemas()([]string)
    SetOperations(value []V2OrganizationsItemUsersItemWithScim_user_PatchRequestBody_Operationsable)()
    SetSchemas(value []string)()
}
