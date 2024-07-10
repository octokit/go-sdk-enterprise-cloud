package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_orgs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The disabled_orgs property
    disabled_orgs *int32
    // The total_orgs property
    total_orgs *int32
    // The total_team_members property
    total_team_members *int32
    // The total_teams property
    total_teams *int32
}
// NewServerStatistics_ghe_stats_orgs instantiates a new ServerStatistics_ghe_stats_orgs and sets the default values.
func NewServerStatistics_ghe_stats_orgs()(*ServerStatistics_ghe_stats_orgs) {
    m := &ServerStatistics_ghe_stats_orgs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_orgsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_orgsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_orgs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_orgs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisabledOrgs gets the disabled_orgs property value. The disabled_orgs property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_orgs) GetDisabledOrgs()(*int32) {
    return m.disabled_orgs
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_orgs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["disabled_orgs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledOrgs(val)
        }
        return nil
    }
    res["total_orgs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalOrgs(val)
        }
        return nil
    }
    res["total_team_members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTeamMembers(val)
        }
        return nil
    }
    res["total_teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTeams(val)
        }
        return nil
    }
    return res
}
// GetTotalOrgs gets the total_orgs property value. The total_orgs property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_orgs) GetTotalOrgs()(*int32) {
    return m.total_orgs
}
// GetTotalTeamMembers gets the total_team_members property value. The total_team_members property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_orgs) GetTotalTeamMembers()(*int32) {
    return m.total_team_members
}
// GetTotalTeams gets the total_teams property value. The total_teams property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_orgs) GetTotalTeams()(*int32) {
    return m.total_teams
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_orgs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("disabled_orgs", m.GetDisabledOrgs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_orgs", m.GetTotalOrgs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_teams", m.GetTotalTeams())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_team_members", m.GetTotalTeamMembers())
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
func (m *ServerStatistics_ghe_stats_orgs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisabledOrgs sets the disabled_orgs property value. The disabled_orgs property
func (m *ServerStatistics_ghe_stats_orgs) SetDisabledOrgs(value *int32)() {
    m.disabled_orgs = value
}
// SetTotalOrgs sets the total_orgs property value. The total_orgs property
func (m *ServerStatistics_ghe_stats_orgs) SetTotalOrgs(value *int32)() {
    m.total_orgs = value
}
// SetTotalTeamMembers sets the total_team_members property value. The total_team_members property
func (m *ServerStatistics_ghe_stats_orgs) SetTotalTeamMembers(value *int32)() {
    m.total_team_members = value
}
// SetTotalTeams sets the total_teams property value. The total_teams property
func (m *ServerStatistics_ghe_stats_orgs) SetTotalTeams(value *int32)() {
    m.total_teams = value
}
type ServerStatistics_ghe_stats_orgsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisabledOrgs()(*int32)
    GetTotalOrgs()(*int32)
    GetTotalTeamMembers()(*int32)
    GetTotalTeams()(*int32)
    SetDisabledOrgs(value *int32)()
    SetTotalOrgs(value *int32)()
    SetTotalTeamMembers(value *int32)()
    SetTotalTeams(value *int32)()
}
