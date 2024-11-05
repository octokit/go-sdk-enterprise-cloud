package user

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i46bc095d8d4b8696bfce1ad7fb3e2e8188d3a31309d0f172088364f4ad1f8d1c "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/user/packages/item/item/versions"
)

// PackagesItemItemVersionsRequestBuilder builds and executes requests for operations under \user\packages\{package_type}\{package_name}\versions
type PackagesItemItemVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PackagesItemItemVersionsRequestBuilderGetQueryParameters lists package versions for a package owned by the authenticated user.OAuth app tokens and personal access tokens (classic) need the `read:packages` scope to use this endpoint. For more information, see "[About permissions for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
type PackagesItemItemVersionsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The state of the package, either active or deleted.
    State *i46bc095d8d4b8696bfce1ad7fb3e2e8188d3a31309d0f172088364f4ad1f8d1c.GetStateQueryParameterType `uriparametername:"state"`
}
// ByPackage_version_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.user.packages.item.item.versions.item collection
// returns a *PackagesItemItemVersionsWithPackage_version_ItemRequestBuilder when successful
func (m *PackagesItemItemVersionsRequestBuilder) ByPackage_version_id(package_version_id int32)(*PackagesItemItemVersionsWithPackage_version_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["package_version_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(package_version_id), 10)
    return NewPackagesItemItemVersionsWithPackage_version_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPackagesItemItemVersionsRequestBuilderInternal instantiates a new PackagesItemItemVersionsRequestBuilder and sets the default values.
func NewPackagesItemItemVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesItemItemVersionsRequestBuilder) {
    m := &PackagesItemItemVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/packages/{package_type}/{package_name}/versions{?page*,per_page*,state*}", pathParameters),
    }
    return m
}
// NewPackagesItemItemVersionsRequestBuilder instantiates a new PackagesItemItemVersionsRequestBuilder and sets the default values.
func NewPackagesItemItemVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesItemItemVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPackagesItemItemVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists package versions for a package owned by the authenticated user.OAuth app tokens and personal access tokens (classic) need the `read:packages` scope to use this endpoint. For more information, see "[About permissions for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// returns a []PackageVersionable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/packages/packages#list-package-versions-for-a-package-owned-by-the-authenticated-user
func (m *PackagesItemItemVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PackagesItemItemVersionsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackageVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreatePackageVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackageVersionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackageVersionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists package versions for a package owned by the authenticated user.OAuth app tokens and personal access tokens (classic) need the `read:packages` scope to use this endpoint. For more information, see "[About permissions for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// returns a *RequestInformation when successful
func (m *PackagesItemItemVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PackagesItemItemVersionsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PackagesItemItemVersionsRequestBuilder when successful
func (m *PackagesItemItemVersionsRequestBuilder) WithUrl(rawUrl string)(*PackagesItemItemVersionsRequestBuilder) {
    return NewPackagesItemItemVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
