package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comments property
    comments ServerStatistics_ghe_stats_commentsable
    // The gists property
    gists ServerStatistics_ghe_stats_gistsable
    // The hooks property
    hooks ServerStatistics_ghe_stats_hooksable
    // The issues property
    issues ServerStatistics_ghe_stats_issuesable
    // The milestones property
    milestones ServerStatistics_ghe_stats_milestonesable
    // The orgs property
    orgs ServerStatistics_ghe_stats_orgsable
    // The pages property
    pages ServerStatistics_ghe_stats_pagesable
    // The pulls property
    pulls ServerStatistics_ghe_stats_pullsable
    // The repos property
    repos ServerStatistics_ghe_stats_reposable
    // The users property
    users ServerStatistics_ghe_stats_usersable
}
// NewServerStatistics_ghe_stats instantiates a new ServerStatistics_ghe_stats and sets the default values.
func NewServerStatistics_ghe_stats()(*ServerStatistics_ghe_stats) {
    m := &ServerStatistics_ghe_stats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_statsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_statsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. The comments property
// returns a ServerStatistics_ghe_stats_commentsable when successful
func (m *ServerStatistics_ghe_stats) GetComments()(ServerStatistics_ghe_stats_commentsable) {
    return m.comments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_commentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val.(ServerStatistics_ghe_stats_commentsable))
        }
        return nil
    }
    res["gists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_gistsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGists(val.(ServerStatistics_ghe_stats_gistsable))
        }
        return nil
    }
    res["hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_hooksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooks(val.(ServerStatistics_ghe_stats_hooksable))
        }
        return nil
    }
    res["issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_issuesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssues(val.(ServerStatistics_ghe_stats_issuesable))
        }
        return nil
    }
    res["milestones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_milestonesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestones(val.(ServerStatistics_ghe_stats_milestonesable))
        }
        return nil
    }
    res["orgs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_orgsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrgs(val.(ServerStatistics_ghe_stats_orgsable))
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_pagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPages(val.(ServerStatistics_ghe_stats_pagesable))
        }
        return nil
    }
    res["pulls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_pullsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPulls(val.(ServerStatistics_ghe_stats_pullsable))
        }
        return nil
    }
    res["repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_reposFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepos(val.(ServerStatistics_ghe_stats_reposable))
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_stats_usersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsers(val.(ServerStatistics_ghe_stats_usersable))
        }
        return nil
    }
    return res
}
// GetGists gets the gists property value. The gists property
// returns a ServerStatistics_ghe_stats_gistsable when successful
func (m *ServerStatistics_ghe_stats) GetGists()(ServerStatistics_ghe_stats_gistsable) {
    return m.gists
}
// GetHooks gets the hooks property value. The hooks property
// returns a ServerStatistics_ghe_stats_hooksable when successful
func (m *ServerStatistics_ghe_stats) GetHooks()(ServerStatistics_ghe_stats_hooksable) {
    return m.hooks
}
// GetIssues gets the issues property value. The issues property
// returns a ServerStatistics_ghe_stats_issuesable when successful
func (m *ServerStatistics_ghe_stats) GetIssues()(ServerStatistics_ghe_stats_issuesable) {
    return m.issues
}
// GetMilestones gets the milestones property value. The milestones property
// returns a ServerStatistics_ghe_stats_milestonesable when successful
func (m *ServerStatistics_ghe_stats) GetMilestones()(ServerStatistics_ghe_stats_milestonesable) {
    return m.milestones
}
// GetOrgs gets the orgs property value. The orgs property
// returns a ServerStatistics_ghe_stats_orgsable when successful
func (m *ServerStatistics_ghe_stats) GetOrgs()(ServerStatistics_ghe_stats_orgsable) {
    return m.orgs
}
// GetPages gets the pages property value. The pages property
// returns a ServerStatistics_ghe_stats_pagesable when successful
func (m *ServerStatistics_ghe_stats) GetPages()(ServerStatistics_ghe_stats_pagesable) {
    return m.pages
}
// GetPulls gets the pulls property value. The pulls property
// returns a ServerStatistics_ghe_stats_pullsable when successful
func (m *ServerStatistics_ghe_stats) GetPulls()(ServerStatistics_ghe_stats_pullsable) {
    return m.pulls
}
// GetRepos gets the repos property value. The repos property
// returns a ServerStatistics_ghe_stats_reposable when successful
func (m *ServerStatistics_ghe_stats) GetRepos()(ServerStatistics_ghe_stats_reposable) {
    return m.repos
}
// GetUsers gets the users property value. The users property
// returns a ServerStatistics_ghe_stats_usersable when successful
func (m *ServerStatistics_ghe_stats) GetUsers()(ServerStatistics_ghe_stats_usersable) {
    return m.users
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("gists", m.GetGists())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hooks", m.GetHooks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("issues", m.GetIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("milestones", m.GetMilestones())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("orgs", m.GetOrgs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pages", m.GetPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pulls", m.GetPulls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repos", m.GetRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("users", m.GetUsers())
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
func (m *ServerStatistics_ghe_stats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. The comments property
func (m *ServerStatistics_ghe_stats) SetComments(value ServerStatistics_ghe_stats_commentsable)() {
    m.comments = value
}
// SetGists sets the gists property value. The gists property
func (m *ServerStatistics_ghe_stats) SetGists(value ServerStatistics_ghe_stats_gistsable)() {
    m.gists = value
}
// SetHooks sets the hooks property value. The hooks property
func (m *ServerStatistics_ghe_stats) SetHooks(value ServerStatistics_ghe_stats_hooksable)() {
    m.hooks = value
}
// SetIssues sets the issues property value. The issues property
func (m *ServerStatistics_ghe_stats) SetIssues(value ServerStatistics_ghe_stats_issuesable)() {
    m.issues = value
}
// SetMilestones sets the milestones property value. The milestones property
func (m *ServerStatistics_ghe_stats) SetMilestones(value ServerStatistics_ghe_stats_milestonesable)() {
    m.milestones = value
}
// SetOrgs sets the orgs property value. The orgs property
func (m *ServerStatistics_ghe_stats) SetOrgs(value ServerStatistics_ghe_stats_orgsable)() {
    m.orgs = value
}
// SetPages sets the pages property value. The pages property
func (m *ServerStatistics_ghe_stats) SetPages(value ServerStatistics_ghe_stats_pagesable)() {
    m.pages = value
}
// SetPulls sets the pulls property value. The pulls property
func (m *ServerStatistics_ghe_stats) SetPulls(value ServerStatistics_ghe_stats_pullsable)() {
    m.pulls = value
}
// SetRepos sets the repos property value. The repos property
func (m *ServerStatistics_ghe_stats) SetRepos(value ServerStatistics_ghe_stats_reposable)() {
    m.repos = value
}
// SetUsers sets the users property value. The users property
func (m *ServerStatistics_ghe_stats) SetUsers(value ServerStatistics_ghe_stats_usersable)() {
    m.users = value
}
type ServerStatistics_ghe_statsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(ServerStatistics_ghe_stats_commentsable)
    GetGists()(ServerStatistics_ghe_stats_gistsable)
    GetHooks()(ServerStatistics_ghe_stats_hooksable)
    GetIssues()(ServerStatistics_ghe_stats_issuesable)
    GetMilestones()(ServerStatistics_ghe_stats_milestonesable)
    GetOrgs()(ServerStatistics_ghe_stats_orgsable)
    GetPages()(ServerStatistics_ghe_stats_pagesable)
    GetPulls()(ServerStatistics_ghe_stats_pullsable)
    GetRepos()(ServerStatistics_ghe_stats_reposable)
    GetUsers()(ServerStatistics_ghe_stats_usersable)
    SetComments(value ServerStatistics_ghe_stats_commentsable)()
    SetGists(value ServerStatistics_ghe_stats_gistsable)()
    SetHooks(value ServerStatistics_ghe_stats_hooksable)()
    SetIssues(value ServerStatistics_ghe_stats_issuesable)()
    SetMilestones(value ServerStatistics_ghe_stats_milestonesable)()
    SetOrgs(value ServerStatistics_ghe_stats_orgsable)()
    SetPages(value ServerStatistics_ghe_stats_pagesable)()
    SetPulls(value ServerStatistics_ghe_stats_pullsable)()
    SetRepos(value ServerStatistics_ghe_stats_reposable)()
    SetUsers(value ServerStatistics_ghe_stats_usersable)()
}
