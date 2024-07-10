package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody struct {
    // The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove.
    groups []ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable
}
// NewItemTeamsItemTeamSyncGroupMappingsPatchRequestBody instantiates a new ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody and sets the default values.
func NewItemTeamsItemTeamSyncGroupMappingsPatchRequestBody()(*ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody) {
    m := &ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody{
    }
    return m
}
// CreateItemTeamsItemTeamSyncGroupMappingsPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamsItemTeamSyncGroupMappingsPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamsItemTeamSyncGroupMappingsPatchRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove.
// returns a []ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable when successful
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody) GetGroups()([]ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable) {
    return m.groups
}
// Serialize serializes information the current object
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetGroups sets the groups property value. The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove.
func (m *ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody) SetGroups(value []ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable)() {
    m.groups = value
}
type ItemTeamsItemTeamSyncGroupMappingsPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroups()([]ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable)
    SetGroups(value []ItemTeamsItemTeamSyncGroupMappingsPatchRequestBody_groupsable)()
}
