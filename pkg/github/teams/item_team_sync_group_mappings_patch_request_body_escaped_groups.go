package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemTeamSyncGroupMappingsPatchRequestBody_groups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // Description of the IdP group.
    group_description *string
    // ID of the IdP group.
    group_id *string
    // Name of the IdP group.
    group_name *string
    // The id property
    id *string
    // The name property
    name *string
}
// NewItemTeamSyncGroupMappingsPatchRequestBody_groups instantiates a new ItemTeamSyncGroupMappingsPatchRequestBody_groups and sets the default values.
func NewItemTeamSyncGroupMappingsPatchRequestBody_groups()(*ItemTeamSyncGroupMappingsPatchRequestBody_groups) {
    m := &ItemTeamSyncGroupMappingsPatchRequestBody_groups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamSyncGroupMappingsPatchRequestBody_groupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamSyncGroupMappingsPatchRequestBody_groupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamSyncGroupMappingsPatchRequestBody_groups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    return res
}
// GetGroupDescription gets the group_description property value. Description of the IdP group.
// returns a *string when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetGroupDescription()(*string) {
    return m.group_description
}
// GetGroupId gets the group_id property value. ID of the IdP group.
// returns a *string when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetGroupId()(*string) {
    return m.group_id
}
// GetGroupName gets the group_name property value. Name of the IdP group.
// returns a *string when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetGroupName()(*string) {
    return m.group_name
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) SetDescription(value *string)() {
    m.description = value
}
// SetGroupDescription sets the group_description property value. Description of the IdP group.
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) SetGroupDescription(value *string)() {
    m.group_description = value
}
// SetGroupId sets the group_id property value. ID of the IdP group.
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) SetGroupId(value *string)() {
    m.group_id = value
}
// SetGroupName sets the group_name property value. Name of the IdP group.
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) SetGroupName(value *string)() {
    m.group_name = value
}
// SetId sets the id property value. The id property
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *ItemTeamSyncGroupMappingsPatchRequestBody_groups) SetName(value *string)() {
    m.name = value
}
type ItemTeamSyncGroupMappingsPatchRequestBody_groupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetGroupDescription()(*string)
    GetGroupId()(*string)
    GetGroupName()(*string)
    GetId()(*string)
    GetName()(*string)
    SetDescription(value *string)()
    SetGroupDescription(value *string)()
    SetGroupId(value *string)()
    SetGroupName(value *string)()
    SetId(value *string)()
    SetName(value *string)()
}
