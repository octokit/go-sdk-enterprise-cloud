package teams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamSyncRequestBuilder builds and executes requests for operations under \teams\{team_id}\team-sync
type ItemTeamSyncRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemTeamSyncRequestBuilderInternal instantiates a new ItemTeamSyncRequestBuilder and sets the default values.
func NewItemTeamSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamSyncRequestBuilder) {
    m := &ItemTeamSyncRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/team-sync", pathParameters),
    }
    return m
}
// NewItemTeamSyncRequestBuilder instantiates a new ItemTeamSyncRequestBuilder and sets the default values.
func NewItemTeamSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamSyncRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamSyncRequestBuilderInternal(urlParams, requestAdapter)
}
// GroupMappings the groupMappings property
// returns a *ItemTeamSyncGroupMappingsRequestBuilder when successful
func (m *ItemTeamSyncRequestBuilder) GroupMappings()(*ItemTeamSyncGroupMappingsRequestBuilder) {
    return NewItemTeamSyncGroupMappingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
