package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i7b34850208ccb9073f2f73e5c9bddf683ac1a17af4b3a31c72745011f99720c4 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/users/item/projects"
)

// ItemProjectsRequestBuilder builds and executes requests for operations under \users\{username}\projects
type ItemProjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemProjectsRequestBuilderGetQueryParameters > [!WARNING]> **Closing down notice:** Projects (classic) is being deprecated in favor of the new Projects experience.> See the [changelog](https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/) for more information.
type ItemProjectsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Indicates the state of the projects to return.
    State *i7b34850208ccb9073f2f73e5c9bddf683ac1a17af4b3a31c72745011f99720c4.GetStateQueryParameterType `uriparametername:"state"`
}
// NewItemProjectsRequestBuilderInternal instantiates a new ItemProjectsRequestBuilder and sets the default values.
func NewItemProjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemProjectsRequestBuilder) {
    m := &ItemProjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/projects{?page*,per_page*,state*}", pathParameters),
    }
    return m
}
// NewItemProjectsRequestBuilder instantiates a new ItemProjectsRequestBuilder and sets the default values.
func NewItemProjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemProjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemProjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Closing down notice:** Projects (classic) is being deprecated in favor of the new Projects experience.> See the [changelog](https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/) for more information.
// Deprecated: 
// returns a []Projectable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/projects/projects#list-user-projects
func (m *ItemProjectsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemProjectsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Projectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateProjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Projectable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Projectable)
        }
    }
    return val, nil
}
// ToGetRequestInformation > [!WARNING]> **Closing down notice:** Projects (classic) is being deprecated in favor of the new Projects experience.> See the [changelog](https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/) for more information.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemProjectsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemProjectsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemProjectsRequestBuilder when successful
func (m *ItemProjectsRequestBuilder) WithUrl(rawUrl string)(*ItemProjectsRequestBuilder) {
    return NewItemProjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
