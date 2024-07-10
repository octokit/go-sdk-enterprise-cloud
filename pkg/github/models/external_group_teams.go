package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalGroup_teams struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id for a team
    team_id *int32
    // The name of the team
    team_name *string
}
// NewExternalGroup_teams instantiates a new ExternalGroup_teams and sets the default values.
func NewExternalGroup_teams()(*ExternalGroup_teams) {
    m := &ExternalGroup_teams{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalGroup_teamsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalGroup_teamsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroup_teams(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalGroup_teams) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalGroup_teams) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["team_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    res["team_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamName(val)
        }
        return nil
    }
    return res
}
// GetTeamId gets the team_id property value. The id for a team
// returns a *int32 when successful
func (m *ExternalGroup_teams) GetTeamId()(*int32) {
    return m.team_id
}
// GetTeamName gets the team_name property value. The name of the team
// returns a *string when successful
func (m *ExternalGroup_teams) GetTeamName()(*string) {
    return m.team_name
}
// Serialize serializes information the current object
func (m *ExternalGroup_teams) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("team_id", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("team_name", m.GetTeamName())
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
func (m *ExternalGroup_teams) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTeamId sets the team_id property value. The id for a team
func (m *ExternalGroup_teams) SetTeamId(value *int32)() {
    m.team_id = value
}
// SetTeamName sets the team_name property value. The name of the team
func (m *ExternalGroup_teams) SetTeamName(value *string)() {
    m.team_name = value
}
type ExternalGroup_teamsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTeamId()(*int32)
    GetTeamName()(*string)
    SetTeamId(value *int32)()
    SetTeamName(value *string)()
}
