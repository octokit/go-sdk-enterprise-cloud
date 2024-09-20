package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description of the IdP group.
    group_description *string
    // ID of the IdP group.
    group_id *string
    // Name of the IdP group.
    group_name *string
}
// NewItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups instantiates a new ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups and sets the default values.
func NewItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups()(*ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) {
    m := &ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetGroupDescription gets the group_description property value. Description of the IdP group.
// returns a *string when successful
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) GetGroupDescription()(*string) {
    return m.group_description
}
// GetGroupId gets the group_id property value. ID of the IdP group.
// returns a *string when successful
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) GetGroupId()(*string) {
    return m.group_id
}
// GetGroupName gets the group_name property value. Name of the IdP group.
// returns a *string when successful
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) GetGroupName()(*string) {
    return m.group_name
}
// Serialize serializes information the current object
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupDescription sets the group_description property value. Description of the IdP group.
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) SetGroupDescription(value *string)() {
    m.group_description = value
}
// SetGroupId sets the group_id property value. ID of the IdP group.
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) SetGroupId(value *string)() {
    m.group_id = value
}
// SetGroupName sets the group_name property value. Name of the IdP group.
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groups) SetGroupName(value *string)() {
    m.group_name = value
}
type ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupDescription()(*string)
    GetGroupId()(*string)
    GetGroupName()(*string)
    SetGroupDescription(value *string)()
    SetGroupId(value *string)()
    SetGroupName(value *string)()
}
