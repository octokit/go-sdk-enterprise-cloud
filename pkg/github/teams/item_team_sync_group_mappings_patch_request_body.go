package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemTeamSyncGroupMappingsPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove.
    groups []ItemTeamSyncGroupMappingsPatchRequestBody_groupsable
    // The synced_at property
    synced_at *string
}
// NewItemTeamSyncGroupMappingsPatchRequestBody instantiates a new ItemTeamSyncGroupMappingsPatchRequestBody and sets the default values.
func NewItemTeamSyncGroupMappingsPatchRequestBody()(*ItemTeamSyncGroupMappingsPatchRequestBody) {
    m := &ItemTeamSyncGroupMappingsPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamSyncGroupMappingsPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamSyncGroupMappingsPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamSyncGroupMappingsPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemTeamSyncGroupMappingsPatchRequestBody_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemTeamSyncGroupMappingsPatchRequestBody_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemTeamSyncGroupMappingsPatchRequestBody_groupsable)
                }
            }
            m.SetGroups(res)
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
// GetGroups gets the groups property value. The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove.
// returns a []ItemTeamSyncGroupMappingsPatchRequestBody_groupsable when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) GetGroups()([]ItemTeamSyncGroupMappingsPatchRequestBody_groupsable) {
    return m.groups
}
// GetSyncedAt gets the synced_at property value. The synced_at property
// returns a *string when successful
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) GetSyncedAt()(*string) {
    return m.synced_at
}
// Serialize serializes information the current object
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroups sets the groups property value. The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove.
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) SetGroups(value []ItemTeamSyncGroupMappingsPatchRequestBody_groupsable)() {
    m.groups = value
}
// SetSyncedAt sets the synced_at property value. The synced_at property
func (m *ItemTeamSyncGroupMappingsPatchRequestBody) SetSyncedAt(value *string)() {
    m.synced_at = value
}
type ItemTeamSyncGroupMappingsPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroups()([]ItemTeamSyncGroupMappingsPatchRequestBody_groupsable)
    GetSyncedAt()(*string)
    SetGroups(value []ItemTeamSyncGroupMappingsPatchRequestBody_groupsable)()
    SetSyncedAt(value *string)()
}
