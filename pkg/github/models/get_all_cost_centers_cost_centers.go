package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetAllCostCenters_costCenters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // ID of the cost center.
    id *string
    // Name of the cost center.
    name *string
    // The resources property
    resources []GetAllCostCenters_costCenters_resourcesable
}
// NewGetAllCostCenters_costCenters instantiates a new GetAllCostCenters_costCenters and sets the default values.
func NewGetAllCostCenters_costCenters()(*GetAllCostCenters_costCenters) {
    m := &GetAllCostCenters_costCenters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetAllCostCenters_costCentersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetAllCostCenters_costCentersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAllCostCenters_costCenters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetAllCostCenters_costCenters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetAllCostCenters_costCenters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGetAllCostCenters_costCenters_resourcesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GetAllCostCenters_costCenters_resourcesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GetAllCostCenters_costCenters_resourcesable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. ID of the cost center.
// returns a *string when successful
func (m *GetAllCostCenters_costCenters) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the cost center.
// returns a *string when successful
func (m *GetAllCostCenters_costCenters) GetName()(*string) {
    return m.name
}
// GetResources gets the resources property value. The resources property
// returns a []GetAllCostCenters_costCenters_resourcesable when successful
func (m *GetAllCostCenters_costCenters) GetResources()([]GetAllCostCenters_costCenters_resourcesable) {
    return m.resources
}
// Serialize serializes information the current object
func (m *GetAllCostCenters_costCenters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("resources", cast)
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
func (m *GetAllCostCenters_costCenters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. ID of the cost center.
func (m *GetAllCostCenters_costCenters) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the cost center.
func (m *GetAllCostCenters_costCenters) SetName(value *string)() {
    m.name = value
}
// SetResources sets the resources property value. The resources property
func (m *GetAllCostCenters_costCenters) SetResources(value []GetAllCostCenters_costCenters_resourcesable)() {
    m.resources = value
}
type GetAllCostCenters_costCentersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetResources()([]GetAllCostCenters_costCenters_resourcesable)
    SetId(value *string)()
    SetName(value *string)()
    SetResources(value []GetAllCostCenters_costCenters_resourcesable)()
}
