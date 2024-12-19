package organizations

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// OrganizationsRequestBuilder builds and executes requests for operations under \organizations
type OrganizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OrganizationsRequestBuilderGetQueryParameters lists all organizations, in the order that they were created.> [!NOTE]> Pagination is powered exclusively by the `since` parameter. Use the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers) to get the URL for the next page of organizations.
type OrganizationsRequestBuilderGetQueryParameters struct {
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // An organization ID. Only return organizations with an ID greater than this ID.
    Since *int32 `uriparametername:"since"`
}
// ByOrganization_Id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.organizations.item collection
// returns a *Organization_ItemRequestBuilder when successful
func (m *OrganizationsRequestBuilder) ByOrganization_Id(organization_Id string)(*Organization_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if organization_Id != "" {
        urlTplParams["organization_%2Did"] = organization_Id
    }
    return NewOrganization_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOrganizationsRequestBuilderInternal instantiates a new OrganizationsRequestBuilder and sets the default values.
func NewOrganizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsRequestBuilder) {
    m := &OrganizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organizations{?per_page*,since*}", pathParameters),
    }
    return m
}
// NewOrganizationsRequestBuilder instantiates a new OrganizationsRequestBuilder and sets the default values.
func NewOrganizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all organizations, in the order that they were created.> [!NOTE]> Pagination is powered exclusively by the `since` parameter. Use the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers) to get the URL for the next page of organizations.
// returns a []OrganizationSimpleable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/orgs#list-organizations
func (m *OrganizationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[OrganizationsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.OrganizationSimpleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateOrganizationSimpleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.OrganizationSimpleable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.OrganizationSimpleable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all organizations, in the order that they were created.> [!NOTE]> Pagination is powered exclusively by the `since` parameter. Use the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers) to get the URL for the next page of organizations.
// returns a *RequestInformation when successful
func (m *OrganizationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[OrganizationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OrganizationsRequestBuilder when successful
func (m *OrganizationsRequestBuilder) WithUrl(rawUrl string)(*OrganizationsRequestBuilder) {
    return NewOrganizationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
