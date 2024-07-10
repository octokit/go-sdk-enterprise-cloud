package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetAllCostCenters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The costCenters property
    costCenters []GetAllCostCenters_costCentersable
}
// NewGetAllCostCenters instantiates a new GetAllCostCenters and sets the default values.
func NewGetAllCostCenters()(*GetAllCostCenters) {
    m := &GetAllCostCenters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetAllCostCentersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetAllCostCentersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAllCostCenters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetAllCostCenters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCostCenters gets the costCenters property value. The costCenters property
// returns a []GetAllCostCenters_costCentersable when successful
func (m *GetAllCostCenters) GetCostCenters()([]GetAllCostCenters_costCentersable) {
    return m.costCenters
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetAllCostCenters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["costCenters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGetAllCostCenters_costCentersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GetAllCostCenters_costCentersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GetAllCostCenters_costCentersable)
                }
            }
            m.SetCostCenters(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GetAllCostCenters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCostCenters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCostCenters()))
        for i, v := range m.GetCostCenters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("costCenters", cast)
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
func (m *GetAllCostCenters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCostCenters sets the costCenters property value. The costCenters property
func (m *GetAllCostCenters) SetCostCenters(value []GetAllCostCenters_costCentersable)() {
    m.costCenters = value
}
type GetAllCostCentersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCostCenters()([]GetAllCostCenters_costCentersable)
    SetCostCenters(value []GetAllCostCenters_costCentersable)()
}
