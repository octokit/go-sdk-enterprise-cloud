package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalGroup information about an external group's usage and its members
type ExternalGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The internal ID of the group
    group_id *int32
    // The display name for the group
    group_name *string
    // An array of external members linked to this group
    members []ExternalGroup_membersable
    // An array of teams linked to this group
    teams []ExternalGroup_teamsable
    // The date when the group was last updated_at
    updated_at *string
}
// NewExternalGroup instantiates a new ExternalGroup and sets the default values.
func NewExternalGroup()(*ExternalGroup) {
    m := &ExternalGroup{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalGroup) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalGroup_membersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalGroup_membersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalGroup_membersable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalGroup_teamsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalGroup_teamsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalGroup_teamsable)
                }
            }
            m.SetTeams(res)
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
func (m *ExternalGroup) GetGroupId()(*int32) {
    return m.group_id
}
// GetGroupName gets the group_name property value. The display name for the group
// returns a *string when successful
func (m *ExternalGroup) GetGroupName()(*string) {
    return m.group_name
}
// GetMembers gets the members property value. An array of external members linked to this group
// returns a []ExternalGroup_membersable when successful
func (m *ExternalGroup) GetMembers()([]ExternalGroup_membersable) {
    return m.members
}
// GetTeams gets the teams property value. An array of teams linked to this group
// returns a []ExternalGroup_teamsable when successful
func (m *ExternalGroup) GetTeams()([]ExternalGroup_teamsable) {
    return m.teams
}
// GetUpdatedAt gets the updated_at property value. The date when the group was last updated_at
// returns a *string when successful
func (m *ExternalGroup) GetUpdatedAt()(*string) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *ExternalGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTeams()))
        for i, v := range m.GetTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("teams", cast)
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
func (m *ExternalGroup) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupId sets the group_id property value. The internal ID of the group
func (m *ExternalGroup) SetGroupId(value *int32)() {
    m.group_id = value
}
// SetGroupName sets the group_name property value. The display name for the group
func (m *ExternalGroup) SetGroupName(value *string)() {
    m.group_name = value
}
// SetMembers sets the members property value. An array of external members linked to this group
func (m *ExternalGroup) SetMembers(value []ExternalGroup_membersable)() {
    m.members = value
}
// SetTeams sets the teams property value. An array of teams linked to this group
func (m *ExternalGroup) SetTeams(value []ExternalGroup_teamsable)() {
    m.teams = value
}
// SetUpdatedAt sets the updated_at property value. The date when the group was last updated_at
func (m *ExternalGroup) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
type ExternalGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupId()(*int32)
    GetGroupName()(*string)
    GetMembers()([]ExternalGroup_membersable)
    GetTeams()([]ExternalGroup_teamsable)
    GetUpdatedAt()(*string)
    SetGroupId(value *int32)()
    SetGroupName(value *string)()
    SetMembers(value []ExternalGroup_membersable)()
    SetTeams(value []ExternalGroup_teamsable)()
    SetUpdatedAt(value *string)()
}
