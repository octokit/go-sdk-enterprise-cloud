package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActionsRunnerGroupsItemOrganizationsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of organization IDs that can access the runner group.
    selected_organization_ids []int32
}
// NewItemActionsRunnerGroupsItemOrganizationsPutRequestBody instantiates a new ItemActionsRunnerGroupsItemOrganizationsPutRequestBody and sets the default values.
func NewItemActionsRunnerGroupsItemOrganizationsPutRequestBody()(*ItemActionsRunnerGroupsItemOrganizationsPutRequestBody) {
    m := &ItemActionsRunnerGroupsItemOrganizationsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRunnerGroupsItemOrganizationsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsRunnerGroupsItemOrganizationsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnerGroupsItemOrganizationsPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsRunnerGroupsItemOrganizationsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsRunnerGroupsItemOrganizationsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetSelectedOrganizationIds gets the selected_organization_ids property value. List of organization IDs that can access the runner group.
// returns a []int32 when successful
func (m *ItemActionsRunnerGroupsItemOrganizationsPutRequestBody) GetSelectedOrganizationIds()([]int32) {
    return m.selected_organization_ids
}
// Serialize serializes information the current object
func (m *ItemActionsRunnerGroupsItemOrganizationsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemActionsRunnerGroupsItemOrganizationsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSelectedOrganizationIds sets the selected_organization_ids property value. List of organization IDs that can access the runner group.
func (m *ItemActionsRunnerGroupsItemOrganizationsPutRequestBody) SetSelectedOrganizationIds(value []int32)() {
    m.selected_organization_ids = value
}
type ItemActionsRunnerGroupsItemOrganizationsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedOrganizationIds()([]int32)
    SetSelectedOrganizationIds(value []int32)()
}
