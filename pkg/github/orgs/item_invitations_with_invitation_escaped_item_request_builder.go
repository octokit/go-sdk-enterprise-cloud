package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemInvitationsWithInvitation_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\invitations\{invitation_id}
type ItemInvitationsWithInvitation_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemInvitationsWithInvitation_ItemRequestBuilderInternal instantiates a new ItemInvitationsWithInvitation_ItemRequestBuilder and sets the default values.
func NewItemInvitationsWithInvitation_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    m := &ItemInvitationsWithInvitation_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/invitations/{invitation_id}", pathParameters),
    }
    return m
}
// NewItemInvitationsWithInvitation_ItemRequestBuilder instantiates a new ItemInvitationsWithInvitation_ItemRequestBuilder and sets the default values.
func NewItemInvitationsWithInvitation_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInvitationsWithInvitation_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete cancel an organization invitation. In order to cancel an organization invitation, the authenticated user must be an organization owner.This endpoint is not available for [Enterprise Managed User (EMU) organizations](https://docs.github.com/enterprise-cloud@latest//admin/identity-and-access-management/using-enterprise-managed-users-for-iam/about-enterprise-managed-users).This endpoint triggers [notifications](https://docs.github.com/enterprise-cloud@latest//github/managing-subscriptions-and-notifications-on-github/about-notifications).
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/members#cancel-an-organization-invitation
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Teams the teams property
// returns a *ItemInvitationsItemTeamsRequestBuilder when successful
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) Teams()(*ItemInvitationsItemTeamsRequestBuilder) {
    return NewItemInvitationsItemTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation cancel an organization invitation. In order to cancel an organization invitation, the authenticated user must be an organization owner.This endpoint is not available for [Enterprise Managed User (EMU) organizations](https://docs.github.com/enterprise-cloud@latest//admin/identity-and-access-management/using-enterprise-managed-users-for-iam/about-enterprise-managed-users).This endpoint triggers [notifications](https://docs.github.com/enterprise-cloud@latest//github/managing-subscriptions-and-notifications-on-github/about-notifications).
// returns a *RequestInformation when successful
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInvitationsWithInvitation_ItemRequestBuilder when successful
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    return NewItemInvitationsWithInvitation_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
