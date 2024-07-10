package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_repos struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fork_repos property
    fork_repos *int32
    // The org_repos property
    org_repos *int32
    // The root_repos property
    root_repos *int32
    // The total_pushes property
    total_pushes *int32
    // The total_repos property
    total_repos *int32
    // The total_wikis property
    total_wikis *int32
}
// NewServerStatistics_ghe_stats_repos instantiates a new ServerStatistics_ghe_stats_repos and sets the default values.
func NewServerStatistics_ghe_stats_repos()(*ServerStatistics_ghe_stats_repos) {
    m := &ServerStatistics_ghe_stats_repos{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_reposFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_reposFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_repos(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_repos) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_repos) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fork_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForkRepos(val)
        }
        return nil
    }
    res["org_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrgRepos(val)
        }
        return nil
    }
    res["root_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootRepos(val)
        }
        return nil
    }
    res["total_pushes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPushes(val)
        }
        return nil
    }
    res["total_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRepos(val)
        }
        return nil
    }
    res["total_wikis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalWikis(val)
        }
        return nil
    }
    return res
}
// GetForkRepos gets the fork_repos property value. The fork_repos property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_repos) GetForkRepos()(*int32) {
    return m.fork_repos
}
// GetOrgRepos gets the org_repos property value. The org_repos property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_repos) GetOrgRepos()(*int32) {
    return m.org_repos
}
// GetRootRepos gets the root_repos property value. The root_repos property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_repos) GetRootRepos()(*int32) {
    return m.root_repos
}
// GetTotalPushes gets the total_pushes property value. The total_pushes property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_repos) GetTotalPushes()(*int32) {
    return m.total_pushes
}
// GetTotalRepos gets the total_repos property value. The total_repos property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_repos) GetTotalRepos()(*int32) {
    return m.total_repos
}
// GetTotalWikis gets the total_wikis property value. The total_wikis property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_repos) GetTotalWikis()(*int32) {
    return m.total_wikis
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_repos) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("fork_repos", m.GetForkRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("org_repos", m.GetOrgRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("root_repos", m.GetRootRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_pushes", m.GetTotalPushes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_repos", m.GetTotalRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_wikis", m.GetTotalWikis())
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
func (m *ServerStatistics_ghe_stats_repos) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetForkRepos sets the fork_repos property value. The fork_repos property
func (m *ServerStatistics_ghe_stats_repos) SetForkRepos(value *int32)() {
    m.fork_repos = value
}
// SetOrgRepos sets the org_repos property value. The org_repos property
func (m *ServerStatistics_ghe_stats_repos) SetOrgRepos(value *int32)() {
    m.org_repos = value
}
// SetRootRepos sets the root_repos property value. The root_repos property
func (m *ServerStatistics_ghe_stats_repos) SetRootRepos(value *int32)() {
    m.root_repos = value
}
// SetTotalPushes sets the total_pushes property value. The total_pushes property
func (m *ServerStatistics_ghe_stats_repos) SetTotalPushes(value *int32)() {
    m.total_pushes = value
}
// SetTotalRepos sets the total_repos property value. The total_repos property
func (m *ServerStatistics_ghe_stats_repos) SetTotalRepos(value *int32)() {
    m.total_repos = value
}
// SetTotalWikis sets the total_wikis property value. The total_wikis property
func (m *ServerStatistics_ghe_stats_repos) SetTotalWikis(value *int32)() {
    m.total_wikis = value
}
type ServerStatistics_ghe_stats_reposable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetForkRepos()(*int32)
    GetOrgRepos()(*int32)
    GetRootRepos()(*int32)
    GetTotalPushes()(*int32)
    GetTotalRepos()(*int32)
    GetTotalWikis()(*int32)
    SetForkRepos(value *int32)()
    SetOrgRepos(value *int32)()
    SetRootRepos(value *int32)()
    SetTotalPushes(value *int32)()
    SetTotalRepos(value *int32)()
    SetTotalWikis(value *int32)()
}
