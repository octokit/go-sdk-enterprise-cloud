package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamsItemTeamSyncRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\team-sync
type ItemTeamsItemTeamSyncRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemTeamsItemTeamSyncRequestBuilderInternal instantiates a new ItemTeamsItemTeamSyncRequestBuilder and sets the default values.
func NewItemTeamsItemTeamSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemTeamSyncRequestBuilder) {
    m := &ItemTeamsItemTeamSyncRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/team-sync", pathParameters),
    }
    return m
}
// NewItemTeamsItemTeamSyncRequestBuilder instantiates a new ItemTeamsItemTeamSyncRequestBuilder and sets the default values.
func NewItemTeamsItemTeamSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemTeamSyncRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemTeamSyncRequestBuilderInternal(urlParams, requestAdapter)
}
// GroupMappings the groupMappings property
// returns a *ItemTeamsItemTeamSyncGroupMappingsRequestBuilder when successful
func (m *ItemTeamsItemTeamSyncRequestBuilder) GroupMappings()(*ItemTeamsItemTeamSyncGroupMappingsRequestBuilder) {
    return NewItemTeamsItemTeamSyncGroupMappingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
