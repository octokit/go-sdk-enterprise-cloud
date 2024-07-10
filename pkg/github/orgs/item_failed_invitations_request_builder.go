package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemFailed_invitationsRequestBuilder builds and executes requests for operations under \orgs\{org}\failed_invitations
type ItemFailed_invitationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFailed_invitationsRequestBuilderGetQueryParameters the return hash contains `failed_at` and `failed_reason` fields which represent the time at which the invitation failed and the reason for the failure.This endpoint is not available for [Enterprise Managed User (EMU) organizations](https://docs.github.com/enterprise-cloud@latest//admin/identity-and-access-management/using-enterprise-managed-users-for-iam/about-enterprise-managed-users).
type ItemFailed_invitationsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemFailed_invitationsRequestBuilderInternal instantiates a new ItemFailed_invitationsRequestBuilder and sets the default values.
func NewItemFailed_invitationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFailed_invitationsRequestBuilder) {
    m := &ItemFailed_invitationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/failed_invitations{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemFailed_invitationsRequestBuilder instantiates a new ItemFailed_invitationsRequestBuilder and sets the default values.
func NewItemFailed_invitationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFailed_invitationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFailed_invitationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the return hash contains `failed_at` and `failed_reason` fields which represent the time at which the invitation failed and the reason for the failure.This endpoint is not available for [Enterprise Managed User (EMU) organizations](https://docs.github.com/enterprise-cloud@latest//admin/identity-and-access-management/using-enterprise-managed-users-for-iam/about-enterprise-managed-users).
// returns a []OrganizationInvitationable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/members#list-failed-organization-invitations
func (m *ItemFailed_invitationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemFailed_invitationsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.OrganizationInvitationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateOrganizationInvitationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.OrganizationInvitationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.OrganizationInvitationable)
        }
    }
    return val, nil
}
// ToGetRequestInformation the return hash contains `failed_at` and `failed_reason` fields which represent the time at which the invitation failed and the reason for the failure.This endpoint is not available for [Enterprise Managed User (EMU) organizations](https://docs.github.com/enterprise-cloud@latest//admin/identity-and-access-management/using-enterprise-managed-users-for-iam/about-enterprise-managed-users).
// returns a *RequestInformation when successful
func (m *ItemFailed_invitationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemFailed_invitationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemFailed_invitationsRequestBuilder when successful
func (m *ItemFailed_invitationsRequestBuilder) WithUrl(rawUrl string)(*ItemFailed_invitationsRequestBuilder) {
    return NewItemFailed_invitationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
