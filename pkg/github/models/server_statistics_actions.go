package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServerStatisticsActions actions metrics that are included in the Server Statistics payload/export from GHES
type ServerStatisticsActions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total number of repositories in a GHES installation that have Actions enabled
    number_of_repos_using_actions *int32
    // The percentage of repositories in a GHES installation that have Actions enabled
    percentage_of_repos_using_actions *string
}
// NewServerStatisticsActions instantiates a new ServerStatisticsActions and sets the default values.
func NewServerStatisticsActions()(*ServerStatisticsActions) {
    m := &ServerStatisticsActions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatisticsActionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatisticsActionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatisticsActions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatisticsActions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatisticsActions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["number_of_repos_using_actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfReposUsingActions(val)
        }
        return nil
    }
    res["percentage_of_repos_using_actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageOfReposUsingActions(val)
        }
        return nil
    }
    return res
}
// GetNumberOfReposUsingActions gets the number_of_repos_using_actions property value. The total number of repositories in a GHES installation that have Actions enabled
// returns a *int32 when successful
func (m *ServerStatisticsActions) GetNumberOfReposUsingActions()(*int32) {
    return m.number_of_repos_using_actions
}
// GetPercentageOfReposUsingActions gets the percentage_of_repos_using_actions property value. The percentage of repositories in a GHES installation that have Actions enabled
// returns a *string when successful
func (m *ServerStatisticsActions) GetPercentageOfReposUsingActions()(*string) {
    return m.percentage_of_repos_using_actions
}
// Serialize serializes information the current object
func (m *ServerStatisticsActions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("number_of_repos_using_actions", m.GetNumberOfReposUsingActions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("percentage_of_repos_using_actions", m.GetPercentageOfReposUsingActions())
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
func (m *ServerStatisticsActions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumberOfReposUsingActions sets the number_of_repos_using_actions property value. The total number of repositories in a GHES installation that have Actions enabled
func (m *ServerStatisticsActions) SetNumberOfReposUsingActions(value *int32)() {
    m.number_of_repos_using_actions = value
}
// SetPercentageOfReposUsingActions sets the percentage_of_repos_using_actions property value. The percentage of repositories in a GHES installation that have Actions enabled
func (m *ServerStatisticsActions) SetPercentageOfReposUsingActions(value *string)() {
    m.percentage_of_repos_using_actions = value
}
type ServerStatisticsActionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNumberOfReposUsingActions()(*int32)
    GetPercentageOfReposUsingActions()(*string)
    SetNumberOfReposUsingActions(value *int32)()
    SetPercentageOfReposUsingActions(value *string)()
}
