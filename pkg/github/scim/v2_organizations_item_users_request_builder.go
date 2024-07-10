package scim

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// V2OrganizationsItemUsersRequestBuilder builds and executes requests for operations under \scim\v2\organizations\{org}\Users
type V2OrganizationsItemUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V2OrganizationsItemUsersRequestBuilderGetQueryParameters retrieves a paginated list of all provisioned organization members, including pending invitations. If you provide the `filter` parameter, the resources for all matching provisions members are returned.When a user with a SAML-provisioned external identity leaves (or is removed from) an organization, the account's metadata is immediately removed. However, the returned list of user accounts might not always match the organization or enterprise member list you see on GitHub Enterprise Cloud. This can happen in certain cases where an external identity associated with an organization will not match an organization member:  - When a user with a SCIM-provisioned external identity is removed from an organization, the account's metadata is preserved to allow the user to re-join the organization in the future.  - When inviting a user to join an organization, you can expect to see their external identity in the results before they accept the invitation, or if the invitation is cancelled (or never accepted).  - When a user is invited over SCIM, an external identity is created that matches with the invitee's email address. However, this identity is only linked to a user account when the user accepts the invitation by going through SAML SSO.The returned list of external identities can include an entry for a `null` user. These are unlinked SAML identities that are created when a user goes through the following Single Sign-On (SSO) process but does not sign in to their GitHub Enterprise Cloud account after completing SSO:1. The user is granted access by the IdP and is not a member of the GitHub Enterprise Cloud organization.1. The user attempts to access the GitHub Enterprise Cloud organization and initiates the SAML SSO process, and is not currently signed in to their GitHub Enterprise Cloud account.1. After successfully authenticating with the SAML SSO IdP, the `null` external identity entry is created and the user is prompted to sign in to their GitHub Enterprise Cloud account:   - If the user signs in, their GitHub Enterprise Cloud account is linked to this entry.   - If the user does not sign in (or does not create a new account when prompted), they are not added to the GitHub Enterprise Cloud organization, and the external identity `null` entry remains in place.
type V2OrganizationsItemUsersRequestBuilderGetQueryParameters struct {
    // Used for pagination: the number of results to return.
    Count *int32 `uriparametername:"count"`
    // Filters results using the equals query parameter operator (`eq`). You can filter results that are equal to `id`, `userName`, `emails`, and `externalId`. For example, to search for an identity with the `userName` Octocat, you would use this query:`?filter=userName%20eq%20\"Octocat\"`.To filter results for the identity with the email `octocat@github.com`, you would use this query:`?filter=emails%20eq%20\"octocat@github.com\"`.
    Filter *string `uriparametername:"filter"`
    // Used for pagination: the index of the first result to return.
    StartIndex *int32 `uriparametername:"startIndex"`
}
// ByScim_user_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.scim.v2.organizations.item.Users.item collection
// returns a *V2OrganizationsItemUsersWithScim_user_ItemRequestBuilder when successful
func (m *V2OrganizationsItemUsersRequestBuilder) ByScim_user_id(scim_user_id string)(*V2OrganizationsItemUsersWithScim_user_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if scim_user_id != "" {
        urlTplParams["scim_user_id"] = scim_user_id
    }
    return NewV2OrganizationsItemUsersWithScim_user_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV2OrganizationsItemUsersRequestBuilderInternal instantiates a new V2OrganizationsItemUsersRequestBuilder and sets the default values.
func NewV2OrganizationsItemUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2OrganizationsItemUsersRequestBuilder) {
    m := &V2OrganizationsItemUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/organizations/{org}/Users{?count*,filter*,startIndex*}", pathParameters),
    }
    return m
}
// NewV2OrganizationsItemUsersRequestBuilder instantiates a new V2OrganizationsItemUsersRequestBuilder and sets the default values.
func NewV2OrganizationsItemUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2OrganizationsItemUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2OrganizationsItemUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves a paginated list of all provisioned organization members, including pending invitations. If you provide the `filter` parameter, the resources for all matching provisions members are returned.When a user with a SAML-provisioned external identity leaves (or is removed from) an organization, the account's metadata is immediately removed. However, the returned list of user accounts might not always match the organization or enterprise member list you see on GitHub Enterprise Cloud. This can happen in certain cases where an external identity associated with an organization will not match an organization member:  - When a user with a SCIM-provisioned external identity is removed from an organization, the account's metadata is preserved to allow the user to re-join the organization in the future.  - When inviting a user to join an organization, you can expect to see their external identity in the results before they accept the invitation, or if the invitation is cancelled (or never accepted).  - When a user is invited over SCIM, an external identity is created that matches with the invitee's email address. However, this identity is only linked to a user account when the user accepts the invitation by going through SAML SSO.The returned list of external identities can include an entry for a `null` user. These are unlinked SAML identities that are created when a user goes through the following Single Sign-On (SSO) process but does not sign in to their GitHub Enterprise Cloud account after completing SSO:1. The user is granted access by the IdP and is not a member of the GitHub Enterprise Cloud organization.1. The user attempts to access the GitHub Enterprise Cloud organization and initiates the SAML SSO process, and is not currently signed in to their GitHub Enterprise Cloud account.1. After successfully authenticating with the SAML SSO IdP, the `null` external identity entry is created and the user is prompted to sign in to their GitHub Enterprise Cloud account:   - If the user signs in, their GitHub Enterprise Cloud account is linked to this entry.   - If the user does not sign in (or does not create a new account when prompted), they are not added to the GitHub Enterprise Cloud organization, and the external identity `null` entry remains in place.
// returns a ScimUserListable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 403 status code
// returns a ScimError error when the service returns a 404 status code
// returns a ScimError error when the service returns a 429 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/scim/scim#list-scim-provisioned-identities
func (m *V2OrganizationsItemUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2OrganizationsItemUsersRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimUserListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "429": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimUserListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimUserListable), nil
}
// Post provisions organization membership for a user, and sends an activation email to the email address. If the user was previously a member of the organization, the invitation will reinstate any former privileges that the user had. For more information about reinstating former members, see "[Reinstating a former member of your organization](https://docs.github.com/enterprise-cloud@latest//organizations/managing-membership-in-your-organization/reinstating-a-former-member-of-your-organization)."
// returns a ScimUserable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 403 status code
// returns a ScimError error when the service returns a 404 status code
// returns a ScimError error when the service returns a 409 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/scim/scim#provision-and-invite-a-scim-user
func (m *V2OrganizationsItemUsersRequestBuilder) Post(ctx context.Context, body V2OrganizationsItemUsersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimUserable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "409": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateScimUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ScimUserable), nil
}
// ToGetRequestInformation retrieves a paginated list of all provisioned organization members, including pending invitations. If you provide the `filter` parameter, the resources for all matching provisions members are returned.When a user with a SAML-provisioned external identity leaves (or is removed from) an organization, the account's metadata is immediately removed. However, the returned list of user accounts might not always match the organization or enterprise member list you see on GitHub Enterprise Cloud. This can happen in certain cases where an external identity associated with an organization will not match an organization member:  - When a user with a SCIM-provisioned external identity is removed from an organization, the account's metadata is preserved to allow the user to re-join the organization in the future.  - When inviting a user to join an organization, you can expect to see their external identity in the results before they accept the invitation, or if the invitation is cancelled (or never accepted).  - When a user is invited over SCIM, an external identity is created that matches with the invitee's email address. However, this identity is only linked to a user account when the user accepts the invitation by going through SAML SSO.The returned list of external identities can include an entry for a `null` user. These are unlinked SAML identities that are created when a user goes through the following Single Sign-On (SSO) process but does not sign in to their GitHub Enterprise Cloud account after completing SSO:1. The user is granted access by the IdP and is not a member of the GitHub Enterprise Cloud organization.1. The user attempts to access the GitHub Enterprise Cloud organization and initiates the SAML SSO process, and is not currently signed in to their GitHub Enterprise Cloud account.1. After successfully authenticating with the SAML SSO IdP, the `null` external identity entry is created and the user is prompted to sign in to their GitHub Enterprise Cloud account:   - If the user signs in, their GitHub Enterprise Cloud account is linked to this entry.   - If the user does not sign in (or does not create a new account when prompted), they are not added to the GitHub Enterprise Cloud organization, and the external identity `null` entry remains in place.
// returns a *RequestInformation when successful
func (m *V2OrganizationsItemUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2OrganizationsItemUsersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    return requestInfo, nil
}
// ToPostRequestInformation provisions organization membership for a user, and sends an activation email to the email address. If the user was previously a member of the organization, the invitation will reinstate any former privileges that the user had. For more information about reinstating former members, see "[Reinstating a former member of your organization](https://docs.github.com/enterprise-cloud@latest//organizations/managing-membership-in-your-organization/reinstating-a-former-member-of-your-organization)."
// returns a *RequestInformation when successful
func (m *V2OrganizationsItemUsersRequestBuilder) ToPostRequestInformation(ctx context.Context, body V2OrganizationsItemUsersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V2OrganizationsItemUsersRequestBuilder when successful
func (m *V2OrganizationsItemUsersRequestBuilder) WithUrl(rawUrl string)(*V2OrganizationsItemUsersRequestBuilder) {
    return NewV2OrganizationsItemUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
