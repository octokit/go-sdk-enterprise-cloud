package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemExternalGroupWithGroup_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\external-group\{group_id}
type ItemExternalGroupWithGroup_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemExternalGroupWithGroup_ItemRequestBuilderGetQueryParameters displays information about the specific group's usage.  Provides a list of the group's external members as well as a list of teams that this group is connected to.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
type ItemExternalGroupWithGroup_ItemRequestBuilderGetQueryParameters struct {
    // The page number of the "members" array results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page for the "members" array (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemExternalGroupWithGroup_ItemRequestBuilderInternal instantiates a new ItemExternalGroupWithGroup_ItemRequestBuilder and sets the default values.
func NewItemExternalGroupWithGroup_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExternalGroupWithGroup_ItemRequestBuilder) {
    m := &ItemExternalGroupWithGroup_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/external-group/{group_id}{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemExternalGroupWithGroup_ItemRequestBuilder instantiates a new ItemExternalGroupWithGroup_ItemRequestBuilder and sets the default values.
func NewItemExternalGroupWithGroup_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExternalGroupWithGroup_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemExternalGroupWithGroup_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get displays information about the specific group's usage.  Provides a list of the group's external members as well as a list of teams that this group is connected to.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a ExternalGroupable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/teams/external-groups#get-an-external-group
func (m *ItemExternalGroupWithGroup_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemExternalGroupWithGroup_ItemRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ExternalGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateExternalGroupFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ExternalGroupable), nil
}
// ToGetRequestInformation displays information about the specific group's usage.  Provides a list of the group's external members as well as a list of teams that this group is connected to.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-cloud@latest//github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemExternalGroupWithGroup_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemExternalGroupWithGroup_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemExternalGroupWithGroup_ItemRequestBuilder when successful
func (m *ItemExternalGroupWithGroup_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemExternalGroupWithGroup_ItemRequestBuilder) {
    return NewItemExternalGroupWithGroup_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
