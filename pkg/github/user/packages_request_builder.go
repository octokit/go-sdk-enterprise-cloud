package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i16df432fa09f3e82ff14962a2a572c893ae306937b5c66a816b790042b70958d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/user/packages"
)

// PackagesRequestBuilder builds and executes requests for operations under \user\packages
type PackagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PackagesRequestBuilderGetQueryParameters lists packages owned by the authenticated user within the user's namespace.OAuth app tokens and personal access tokens (classic) need the `read:packages` scope to use this endpoint. For more information, see "[About permissions for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
type PackagesRequestBuilderGetQueryParameters struct {
    // The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    Package_type *i16df432fa09f3e82ff14962a2a572c893ae306937b5c66a816b790042b70958d.GetPackage_typeQueryParameterType `uriparametername:"package_type"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The selected visibility of the packages.  This parameter is optional and only filters an existing result set.The `internal` visibility is only supported for GitHub Packages registries that allow for granular permissions. For other ecosystems `internal` is synonymous with `private`.For the list of GitHub Packages registries that support granular permissions, see "[About permissions for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//packages/learn-github-packages/about-permissions-for-github-packages#granular-permissions-for-userorganization-scoped-packages)."
    Visibility *i16df432fa09f3e82ff14962a2a572c893ae306937b5c66a816b790042b70958d.GetVisibilityQueryParameterType `uriparametername:"visibility"`
}
// ByPackage_type gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.user.packages.item collection
// returns a *PackagesWithPackage_typeItemRequestBuilder when successful
func (m *PackagesRequestBuilder) ByPackage_type(package_type string)(*PackagesWithPackage_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if package_type != "" {
        urlTplParams["package_type"] = package_type
    }
    return NewPackagesWithPackage_typeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPackagesRequestBuilderInternal instantiates a new PackagesRequestBuilder and sets the default values.
func NewPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesRequestBuilder) {
    m := &PackagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/packages?package_type={package_type}{&page*,per_page*,visibility*}", pathParameters),
    }
    return m
}
// NewPackagesRequestBuilder instantiates a new PackagesRequestBuilder and sets the default values.
func NewPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists packages owned by the authenticated user within the user's namespace.OAuth app tokens and personal access tokens (classic) need the `read:packages` scope to use this endpoint. For more information, see "[About permissions for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// returns a []PackageEscapedable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/packages/packages#list-packages-for-the-authenticated-users-namespace
func (m *PackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PackagesRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackageEscapedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreatePackageEscapedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackageEscapedable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PackageEscapedable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists packages owned by the authenticated user within the user's namespace.OAuth app tokens and personal access tokens (classic) need the `read:packages` scope to use this endpoint. For more information, see "[About permissions for GitHub Packages](https://docs.github.com/enterprise-cloud@latest//packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// returns a *RequestInformation when successful
func (m *PackagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PackagesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PackagesRequestBuilder when successful
func (m *PackagesRequestBuilder) WithUrl(rawUrl string)(*PackagesRequestBuilder) {
    return NewPackagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
