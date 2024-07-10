package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalGroups_groups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The internal ID of the group
    group_id *int32
    // The display name of the group
    group_name *string
    // The time of the last update for this group
    updated_at *string
}
// NewExternalGroups_groups instantiates a new ExternalGroups_groups and sets the default values.
func NewExternalGroups_groups()(*ExternalGroups_groups) {
    m := &ExternalGroups_groups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalGroups_groupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalGroups_groupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroups_groups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalGroups_groups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalGroups_groups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["group_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the group_id property value. The internal ID of the group
// returns a *int32 when successful
func (m *ExternalGroups_groups) GetGroupId()(*int32) {
    return m.group_id
}
// GetGroupName gets the group_name property value. The display name of the group
// returns a *string when successful
func (m *ExternalGroups_groups) GetGroupName()(*string) {
    return m.group_name
}
// GetUpdatedAt gets the updated_at property value. The time of the last update for this group
// returns a *string when successful
func (m *ExternalGroups_groups) GetUpdatedAt()(*string) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *ExternalGroups_groups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("group_id", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("group_name", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
func (m *ExternalGroups_groups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupId sets the group_id property value. The internal ID of the group
func (m *ExternalGroups_groups) SetGroupId(value *int32)() {
    m.group_id = value
}
// SetGroupName sets the group_name property value. The display name of the group
func (m *ExternalGroups_groups) SetGroupName(value *string)() {
    m.group_name = value
}
// SetUpdatedAt sets the updated_at property value. The time of the last update for this group
func (m *ExternalGroups_groups) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
type ExternalGroups_groupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupId()(*int32)
    GetGroupName()(*string)
    GetUpdatedAt()(*string)
    SetGroupId(value *int32)()
    SetGroupName(value *string)()
    SetUpdatedAt(value *string)()
}
