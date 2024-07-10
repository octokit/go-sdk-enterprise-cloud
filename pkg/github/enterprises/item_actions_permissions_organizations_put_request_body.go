package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActionsPermissionsOrganizationsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of organization IDs to enable for GitHub Actions.
    selected_organization_ids []int32
}
// NewItemActionsPermissionsOrganizationsPutRequestBody instantiates a new ItemActionsPermissionsOrganizationsPutRequestBody and sets the default values.
func NewItemActionsPermissionsOrganizationsPutRequestBody()(*ItemActionsPermissionsOrganizationsPutRequestBody) {
    m := &ItemActionsPermissionsOrganizationsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsPermissionsOrganizationsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsPermissionsOrganizationsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsPermissionsOrganizationsPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsPermissionsOrganizationsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsPermissionsOrganizationsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["selected_organization_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetSelectedOrganizationIds(res)
        }
        return nil
    }
    return res
}
// GetSelectedOrganizationIds gets the selected_organization_ids property value. List of organization IDs to enable for GitHub Actions.
// returns a []int32 when successful
func (m *ItemActionsPermissionsOrganizationsPutRequestBody) GetSelectedOrganizationIds()([]int32) {
    return m.selected_organization_ids
}
// Serialize serializes information the current object
func (m *ItemActionsPermissionsOrganizationsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedOrganizationIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_organization_ids", m.GetSelectedOrganizationIds())
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
func (m *ItemActionsPermissionsOrganizationsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSelectedOrganizationIds sets the selected_organization_ids property value. List of organization IDs to enable for GitHub Actions.
func (m *ItemActionsPermissionsOrganizationsPutRequestBody) SetSelectedOrganizationIds(value []int32)() {
    m.selected_organization_ids = value
}
type ItemActionsPermissionsOrganizationsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedOrganizationIds()([]int32)
    SetSelectedOrganizationIds(value []int32)()
}
