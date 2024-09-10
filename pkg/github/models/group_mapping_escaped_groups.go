package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupMapping_groups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // a description of the group
    group_description *string
    // The ID of the group
    group_id *string
    // The name of the group
    group_name *string
    // synchronization status for this group mapping
    status *string
    // the time of the last sync for this group-mapping
    synced_at *string
}
// NewGroupMapping_groups instantiates a new GroupMapping_groups and sets the default values.
func NewGroupMapping_groups()(*GroupMapping_groups) {
    m := &GroupMapping_groups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupMapping_groupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupMapping_groupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupMapping_groups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GroupMapping_groups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupMapping_groups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["group_description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDescription(val)
        }
        return nil
    }
    res["group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["synced_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncedAt(val)
        }
        return nil
    }
    return res
}
// GetGroupDescription gets the group_description property value. a description of the group
// returns a *string when successful
func (m *GroupMapping_groups) GetGroupDescription()(*string) {
    return m.group_description
}
// GetGroupId gets the group_id property value. The ID of the group
// returns a *string when successful
func (m *GroupMapping_groups) GetGroupId()(*string) {
    return m.group_id
}
// GetGroupName gets the group_name property value. The name of the group
// returns a *string when successful
func (m *GroupMapping_groups) GetGroupName()(*string) {
    return m.group_name
}
// GetStatus gets the status property value. synchronization status for this group mapping
// returns a *string when successful
func (m *GroupMapping_groups) GetStatus()(*string) {
    return m.status
}
// GetSyncedAt gets the synced_at property value. the time of the last sync for this group-mapping
// returns a *string when successful
func (m *GroupMapping_groups) GetSyncedAt()(*string) {
    return m.synced_at
}
// Serialize serializes information the current object
func (m *GroupMapping_groups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("group_description", m.GetGroupDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("group_id", m.GetGroupId())
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
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("synced_at", m.GetSyncedAt())
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
func (m *GroupMapping_groups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupDescription sets the group_description property value. a description of the group
func (m *GroupMapping_groups) SetGroupDescription(value *string)() {
    m.group_description = value
}
// SetGroupId sets the group_id property value. The ID of the group
func (m *GroupMapping_groups) SetGroupId(value *string)() {
    m.group_id = value
}
// SetGroupName sets the group_name property value. The name of the group
func (m *GroupMapping_groups) SetGroupName(value *string)() {
    m.group_name = value
}
// SetStatus sets the status property value. synchronization status for this group mapping
func (m *GroupMapping_groups) SetStatus(value *string)() {
    m.status = value
}
// SetSyncedAt sets the synced_at property value. the time of the last sync for this group-mapping
func (m *GroupMapping_groups) SetSyncedAt(value *string)() {
    m.synced_at = value
}
type GroupMapping_groupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupDescription()(*string)
    GetGroupId()(*string)
    GetGroupName()(*string)
    GetStatus()(*string)
    GetSyncedAt()(*string)
    SetGroupDescription(value *string)()
    SetGroupId(value *string)()
    SetGroupName(value *string)()
    SetStatus(value *string)()
    SetSyncedAt(value *string)()
}
