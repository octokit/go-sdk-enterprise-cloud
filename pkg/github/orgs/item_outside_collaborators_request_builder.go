package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i9bc30b236570142d88004fd1b034ec56dc99a2ea976f2b6fea9218e37de0e887 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/outside_collaborators"
)

// ItemOutside_collaboratorsRequestBuilder builds and executes requests for operations under \orgs\{org}\outside_collaborators
type ItemOutside_collaboratorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutside_collaboratorsRequestBuilderGetQueryParameters list all users who are outside collaborators of an organization.
type ItemOutside_collaboratorsRequestBuilderGetQueryParameters struct {
    // Filter the list of outside collaborators. `2fa_disabled` means that only outside collaborators without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled will be returned.
    Filter *i9bc30b236570142d88004fd1b034ec56dc99a2ea976f2b6fea9218e37de0e887.GetFilterQueryParameterType `uriparametername:"filter"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByUsername gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.orgs.item.outside_collaborators.item collection
// returns a *ItemOutside_collaboratorsWithUsernameItemRequestBuilder when successful
func (m *ItemOutside_collaboratorsRequestBuilder) ByUsername(username string)(*ItemOutside_collaboratorsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemOutside_collaboratorsWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOutside_collaboratorsRequestBuilderInternal instantiates a new ItemOutside_collaboratorsRequestBuilder and sets the default values.
func NewItemOutside_collaboratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutside_collaboratorsRequestBuilder) {
    m := &ItemOutside_collaboratorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/outside_collaborators{?filter*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemOutside_collaboratorsRequestBuilder instantiates a new ItemOutside_collaboratorsRequestBuilder and sets the default values.
func NewItemOutside_collaboratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutside_collaboratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutside_collaboratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all users who are outside collaborators of an organization.
// returns a []SimpleUserable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/outside-collaborators#list-outside-collaborators-for-an-organization
func (m *ItemOutside_collaboratorsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOutside_collaboratorsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateSimpleUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SimpleUserable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list all users who are outside collaborators of an organization.
// returns a *RequestInformation when successful
func (m *ItemOutside_collaboratorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOutside_collaboratorsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOutside_collaboratorsRequestBuilder when successful
func (m *ItemOutside_collaboratorsRequestBuilder) WithUrl(rawUrl string)(*ItemOutside_collaboratorsRequestBuilder) {
    return NewItemOutside_collaboratorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
