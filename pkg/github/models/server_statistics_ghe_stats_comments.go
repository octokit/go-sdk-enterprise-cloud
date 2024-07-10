package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics_ghe_stats_comments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total_commit_comments property
    total_commit_comments *int32
    // The total_gist_comments property
    total_gist_comments *int32
    // The total_issue_comments property
    total_issue_comments *int32
    // The total_pull_request_comments property
    total_pull_request_comments *int32
}
// NewServerStatistics_ghe_stats_comments instantiates a new ServerStatistics_ghe_stats_comments and sets the default values.
func NewServerStatistics_ghe_stats_comments()(*ServerStatistics_ghe_stats_comments) {
    m := &ServerStatistics_ghe_stats_comments{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatistics_ghe_stats_commentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatistics_ghe_stats_commentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics_ghe_stats_comments(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics_ghe_stats_comments) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics_ghe_stats_comments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["total_commit_comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCommitComments(val)
        }
        return nil
    }
    res["total_gist_comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalGistComments(val)
        }
        return nil
    }
    res["total_issue_comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalIssueComments(val)
        }
        return nil
    }
    res["total_pull_request_comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPullRequestComments(val)
        }
        return nil
    }
    return res
}
// GetTotalCommitComments gets the total_commit_comments property value. The total_commit_comments property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_comments) GetTotalCommitComments()(*int32) {
    return m.total_commit_comments
}
// GetTotalGistComments gets the total_gist_comments property value. The total_gist_comments property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_comments) GetTotalGistComments()(*int32) {
    return m.total_gist_comments
}
// GetTotalIssueComments gets the total_issue_comments property value. The total_issue_comments property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_comments) GetTotalIssueComments()(*int32) {
    return m.total_issue_comments
}
// GetTotalPullRequestComments gets the total_pull_request_comments property value. The total_pull_request_comments property
// returns a *int32 when successful
func (m *ServerStatistics_ghe_stats_comments) GetTotalPullRequestComments()(*int32) {
    return m.total_pull_request_comments
}
// Serialize serializes information the current object
func (m *ServerStatistics_ghe_stats_comments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total_commit_comments", m.GetTotalCommitComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_gist_comments", m.GetTotalGistComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_issue_comments", m.GetTotalIssueComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_pull_request_comments", m.GetTotalPullRequestComments())
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
func (m *ServerStatistics_ghe_stats_comments) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalCommitComments sets the total_commit_comments property value. The total_commit_comments property
func (m *ServerStatistics_ghe_stats_comments) SetTotalCommitComments(value *int32)() {
    m.total_commit_comments = value
}
// SetTotalGistComments sets the total_gist_comments property value. The total_gist_comments property
func (m *ServerStatistics_ghe_stats_comments) SetTotalGistComments(value *int32)() {
    m.total_gist_comments = value
}
// SetTotalIssueComments sets the total_issue_comments property value. The total_issue_comments property
func (m *ServerStatistics_ghe_stats_comments) SetTotalIssueComments(value *int32)() {
    m.total_issue_comments = value
}
// SetTotalPullRequestComments sets the total_pull_request_comments property value. The total_pull_request_comments property
func (m *ServerStatistics_ghe_stats_comments) SetTotalPullRequestComments(value *int32)() {
    m.total_pull_request_comments = value
}
type ServerStatistics_ghe_stats_commentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalCommitComments()(*int32)
    GetTotalGistComments()(*int32)
    GetTotalIssueComments()(*int32)
    GetTotalPullRequestComments()(*int32)
    SetTotalCommitComments(value *int32)()
    SetTotalGistComments(value *int32)()
    SetTotalIssueComments(value *int32)()
    SetTotalPullRequestComments(value *int32)()
}
