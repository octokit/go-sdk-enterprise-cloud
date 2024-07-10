package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalGroups a list of external groups available to be connected to a team
type ExternalGroups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An array of external groups available to be mapped to a team
    groups []ExternalGroups_groupsable
}
// NewExternalGroups instantiates a new ExternalGroups and sets the default values.
func NewExternalGroups()(*ExternalGroups) {
    m := &ExternalGroups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalGroupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalGroupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalGroups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalGroups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalGroups_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalGroups_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalGroups_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. An array of external groups available to be mapped to a team
// returns a []ExternalGroups_groupsable when successful
func (m *ExternalGroups) GetGroups()([]ExternalGroups_groupsable) {
    return m.groups
}
// Serialize serializes information the current object
func (m *ExternalGroups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("groups", cast)
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
func (m *ExternalGroups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroups sets the groups property value. An array of external groups available to be mapped to a team
func (m *ExternalGroups) SetGroups(value []ExternalGroups_groupsable)() {
    m.groups = value
}
type ExternalGroupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroups()([]ExternalGroups_groupsable)
    SetGroups(value []ExternalGroups_groupsable)()
}
