package github

import (
    "context"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i0e5325290c7e23cc5414cfd9f36afcd9381c4c829b153c16d3602d8a09a76097 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/assignments"
    i16880d3d9bd87a455cb21c63b579f516725525172d4ed43e76c1dcaac227fccb "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/applications"
    i1d5a943e46ca4f66cd96631b86159e0ef2457fcd852b66dacdb3c858df33182f "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/octocat"
    i1e25f49239426bd2e69db174edad2b458de62223f6832633c4998e12a9822fd0 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/user"
    i1e3443d9c5c69a04268b19683e154a6bbfd9a9d5686c4991f80bac8da6456e47 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/markdown"
    i1eb6a2fb0597aa7d41d07a09a94d44536dcc47a43c6ef4a1a956d1da25922e34 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/feeds"
    i254b88d85fa7ce5a32692c895c370345be9a48bb510e34140ed0c8bd2d58e9ea "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/issues"
    i2f58036685f6c1f423db5a62a4598667c8697cf896e332d260bbb3401974d900 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/organizations"
    i301f9521617e10f16fb5803f8394b4a480faef6d52fba9b682e61ff12e94fe8f "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/codes_of_conduct"
    i4529ca8f37d8f2b18d8f9373b5cd997766265c2303aac50efff298d9d09ced95 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/appmanifests"
    i46fa65f2fec1c7d8df8b30be6e00b3463d79c7958500cb3d5f9d302af25efc73 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/licenses"
    i5c334645ff8e9a239fc143c168dbd2092f4e549380d4efbf777031dde1cda9df "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/enterprises"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i66dcdb7ca1167d24b7d506e0bbd1c8498fe9d30e0e57362b58349dd6372d006b "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/installation"
    i66f747cc7e39f2550b27f76f1deff6b4eb7ef9fe9508531a28f5daa16f21bd0c "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos"
    i6bf3c616c968bb26cf2a84862431cdac244fdb101616c891cb1ffe3bb76bb2e5 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/emojis"
    i7a2ac49b9b6f48e07dbfe8e254384c6396f35f44d01fec4711a22ac9a846ab3b "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/scim"
    i7d8b5c61b014f2f88333c6c173317637dad387c73a9c94ad43a7dade0b698a94 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/gitignore"
    i8269710af64afe801d4aaf314420eb5a9431ca6c92099cba46ab7d619772625a "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/apps"
    i8b33911a23a371a3dd3402d12b368b9c76c3ed37d700eadf51c8f1dd2562fe36 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/events"
    i8c3efe8d3f9bba02f3585224c19bc147b1c29223805959ad6e1883c376334208 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repositories"
    i90003368361fb7e2691b82c9f74809cb3242042523904b39d0f52be6ba92cd7d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/gists"
    i99d59e0a5895ad617ab4b55390645d2cf8383766f4f778a81418377d4d39beed "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/networks"
    ia5bb264fddfde83ccb9d99948319a2698ced77553c1949540db583fa7308e6af "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/notifications"
    ia6c97dde936cb5002ff38a8156751d4474a56efe1498de27839265eb17c35e08 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/advisories"
    ia78ddb5870800bf7cef21851ecf22ff5e1dd993fd9fd4ad35d61f54673b524e8 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/teams"
    iabe36f6f42302b9f7d51550f4ff0b71c3bd55f8a80a38c0b53c57cd5cc5b72d8 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/marketplace_listing"
    iace6d2df5b046b2621ed8938a2a7fb06adee52ac123cbb2c016a2118a97499d8 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/meta"
    iad80f7bb8538b815146b487f10b1d16c4e083c86d3018f0a8f7382b939f4e634 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/zen"
    ic8f695ba1188395d15025a672763a6fdf96121c1a19790a0cbc49f0422e9b21f "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/classrooms"
    icf6d9662ff8797fed51869213f4ae097e019afa5b10437df7f2be842bafb7983 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/enterpriseinstallation"
    id8aedfc95cc545f984105f50cf0ac5fdd0d8e5ab227e4a886d53228f435f8aca "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs"
    ie3b347c690ae9117f3a4da4f0e042e20664ce00896fa0a94cdd22b3a98748b55 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/users"
    ie4198ced272c261cf44838663bf9483c5a9a9df8fe18cc7420e64a31cd4f7bf8 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/rate_limit"
    iee46eecc378acacb4378b35f14d2d75f182b1e9ad75681c6ca7ba5a8e05c3fef "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/projects"
    iee5528fa1e80ae7995219d33e93e20141e576831230b8e485aa5202568338b41 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/app"
    if21e1a09c882b708cc2897e803a9731133fa5a954e4da837faea5a6611051822 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/search"
    ifc535ad0dacf3b171683de34966458c4e8276e5eb0e62ac8f374245a8faad8f8 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/versions"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Advisories the advisories property
// returns a *AdvisoriesRequestBuilder when successful
func (m *ApiClient) Advisories()(*ia6c97dde936cb5002ff38a8156751d4474a56efe1498de27839265eb17c35e08.AdvisoriesRequestBuilder) {
    return ia6c97dde936cb5002ff38a8156751d4474a56efe1498de27839265eb17c35e08.NewAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// App the app property
// returns a *AppRequestBuilder when successful
func (m *ApiClient) App()(*iee5528fa1e80ae7995219d33e93e20141e576831230b8e485aa5202568338b41.AppRequestBuilder) {
    return iee5528fa1e80ae7995219d33e93e20141e576831230b8e485aa5202568338b41.NewAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Applications the applications property
// returns a *ApplicationsRequestBuilder when successful
func (m *ApiClient) Applications()(*i16880d3d9bd87a455cb21c63b579f516725525172d4ed43e76c1dcaac227fccb.ApplicationsRequestBuilder) {
    return i16880d3d9bd87a455cb21c63b579f516725525172d4ed43e76c1dcaac227fccb.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManifests the appManifests property
// returns a *AppManifestsRequestBuilder when successful
func (m *ApiClient) AppManifests()(*i4529ca8f37d8f2b18d8f9373b5cd997766265c2303aac50efff298d9d09ced95.AppManifestsRequestBuilder) {
    return i4529ca8f37d8f2b18d8f9373b5cd997766265c2303aac50efff298d9d09ced95.NewAppManifestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
// returns a *AppsRequestBuilder when successful
func (m *ApiClient) Apps()(*i8269710af64afe801d4aaf314420eb5a9431ca6c92099cba46ab7d619772625a.AppsRequestBuilder) {
    return i8269710af64afe801d4aaf314420eb5a9431ca6c92099cba46ab7d619772625a.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments the assignments property
// returns a *AssignmentsRequestBuilder when successful
func (m *ApiClient) Assignments()(*i0e5325290c7e23cc5414cfd9f36afcd9381c4c829b153c16d3602d8a09a76097.AssignmentsRequestBuilder) {
    return i0e5325290c7e23cc5414cfd9f36afcd9381c4c829b153c16d3602d8a09a76097.NewAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Classrooms the classrooms property
// returns a *ClassroomsRequestBuilder when successful
func (m *ApiClient) Classrooms()(*ic8f695ba1188395d15025a672763a6fdf96121c1a19790a0cbc49f0422e9b21f.ClassroomsRequestBuilder) {
    return ic8f695ba1188395d15025a672763a6fdf96121c1a19790a0cbc49f0422e9b21f.NewClassroomsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codes_of_conduct the codes_of_conduct property
// returns a *Codes_of_conductRequestBuilder when successful
func (m *ApiClient) Codes_of_conduct()(*i301f9521617e10f16fb5803f8394b4a480faef6d52fba9b682e61ff12e94fe8f.Codes_of_conductRequestBuilder) {
    return i301f9521617e10f16fb5803f8394b4a480faef6d52fba9b682e61ff12e94fe8f.NewCodes_of_conductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://api.github.com")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Emojis the emojis property
// returns a *EmojisRequestBuilder when successful
func (m *ApiClient) Emojis()(*i6bf3c616c968bb26cf2a84862431cdac244fdb101616c891cb1ffe3bb76bb2e5.EmojisRequestBuilder) {
    return i6bf3c616c968bb26cf2a84862431cdac244fdb101616c891cb1ffe3bb76bb2e5.NewEmojisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnterpriseInstallation the enterpriseInstallation property
// returns a *EnterpriseInstallationRequestBuilder when successful
func (m *ApiClient) EnterpriseInstallation()(*icf6d9662ff8797fed51869213f4ae097e019afa5b10437df7f2be842bafb7983.EnterpriseInstallationRequestBuilder) {
    return icf6d9662ff8797fed51869213f4ae097e019afa5b10437df7f2be842bafb7983.NewEnterpriseInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enterprises the enterprises property
// returns a *EnterprisesRequestBuilder when successful
func (m *ApiClient) Enterprises()(*i5c334645ff8e9a239fc143c168dbd2092f4e549380d4efbf777031dde1cda9df.EnterprisesRequestBuilder) {
    return i5c334645ff8e9a239fc143c168dbd2092f4e549380d4efbf777031dde1cda9df.NewEnterprisesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
// returns a *EventsRequestBuilder when successful
func (m *ApiClient) Events()(*i8b33911a23a371a3dd3402d12b368b9c76c3ed37d700eadf51c8f1dd2562fe36.EventsRequestBuilder) {
    return i8b33911a23a371a3dd3402d12b368b9c76c3ed37d700eadf51c8f1dd2562fe36.NewEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Feeds the feeds property
// returns a *FeedsRequestBuilder when successful
func (m *ApiClient) Feeds()(*i1eb6a2fb0597aa7d41d07a09a94d44536dcc47a43c6ef4a1a956d1da25922e34.FeedsRequestBuilder) {
    return i1eb6a2fb0597aa7d41d07a09a94d44536dcc47a43c6ef4a1a956d1da25922e34.NewFeedsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get Hypermedia links to resources accessible in GitHub's REST API
// returns a Rootable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/meta/meta#github-api-root
func (m *ApiClient) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Rootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateRootFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Rootable), nil
}
// Gists the gists property
// returns a *GistsRequestBuilder when successful
func (m *ApiClient) Gists()(*i90003368361fb7e2691b82c9f74809cb3242042523904b39d0f52be6ba92cd7d.GistsRequestBuilder) {
    return i90003368361fb7e2691b82c9f74809cb3242042523904b39d0f52be6ba92cd7d.NewGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gitignore the gitignore property
// returns a *GitignoreRequestBuilder when successful
func (m *ApiClient) Gitignore()(*i7d8b5c61b014f2f88333c6c173317637dad387c73a9c94ad43a7dade0b698a94.GitignoreRequestBuilder) {
    return i7d8b5c61b014f2f88333c6c173317637dad387c73a9c94ad43a7dade0b698a94.NewGitignoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
// returns a *InstallationRequestBuilder when successful
func (m *ApiClient) Installation()(*i66dcdb7ca1167d24b7d506e0bbd1c8498fe9d30e0e57362b58349dd6372d006b.InstallationRequestBuilder) {
    return i66dcdb7ca1167d24b7d506e0bbd1c8498fe9d30e0e57362b58349dd6372d006b.NewInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
// returns a *IssuesRequestBuilder when successful
func (m *ApiClient) Issues()(*i254b88d85fa7ce5a32692c895c370345be9a48bb510e34140ed0c8bd2d58e9ea.IssuesRequestBuilder) {
    return i254b88d85fa7ce5a32692c895c370345be9a48bb510e34140ed0c8bd2d58e9ea.NewIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Licenses the licenses property
// returns a *LicensesRequestBuilder when successful
func (m *ApiClient) Licenses()(*i46fa65f2fec1c7d8df8b30be6e00b3463d79c7958500cb3d5f9d302af25efc73.LicensesRequestBuilder) {
    return i46fa65f2fec1c7d8df8b30be6e00b3463d79c7958500cb3d5f9d302af25efc73.NewLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Markdown the markdown property
// returns a *MarkdownRequestBuilder when successful
func (m *ApiClient) Markdown()(*i1e3443d9c5c69a04268b19683e154a6bbfd9a9d5686c4991f80bac8da6456e47.MarkdownRequestBuilder) {
    return i1e3443d9c5c69a04268b19683e154a6bbfd9a9d5686c4991f80bac8da6456e47.NewMarkdownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Marketplace_listing the marketplace_listing property
// returns a *Marketplace_listingRequestBuilder when successful
func (m *ApiClient) Marketplace_listing()(*iabe36f6f42302b9f7d51550f4ff0b71c3bd55f8a80a38c0b53c57cd5cc5b72d8.Marketplace_listingRequestBuilder) {
    return iabe36f6f42302b9f7d51550f4ff0b71c3bd55f8a80a38c0b53c57cd5cc5b72d8.NewMarketplace_listingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Meta the meta property
// returns a *MetaRequestBuilder when successful
func (m *ApiClient) Meta()(*iace6d2df5b046b2621ed8938a2a7fb06adee52ac123cbb2c016a2118a97499d8.MetaRequestBuilder) {
    return iace6d2df5b046b2621ed8938a2a7fb06adee52ac123cbb2c016a2118a97499d8.NewMetaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Networks the networks property
// returns a *NetworksRequestBuilder when successful
func (m *ApiClient) Networks()(*i99d59e0a5895ad617ab4b55390645d2cf8383766f4f778a81418377d4d39beed.NetworksRequestBuilder) {
    return i99d59e0a5895ad617ab4b55390645d2cf8383766f4f778a81418377d4d39beed.NewNetworksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
// returns a *NotificationsRequestBuilder when successful
func (m *ApiClient) Notifications()(*ia5bb264fddfde83ccb9d99948319a2698ced77553c1949540db583fa7308e6af.NotificationsRequestBuilder) {
    return ia5bb264fddfde83ccb9d99948319a2698ced77553c1949540db583fa7308e6af.NewNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Octocat the octocat property
// returns a *OctocatRequestBuilder when successful
func (m *ApiClient) Octocat()(*i1d5a943e46ca4f66cd96631b86159e0ef2457fcd852b66dacdb3c858df33182f.OctocatRequestBuilder) {
    return i1d5a943e46ca4f66cd96631b86159e0ef2457fcd852b66dacdb3c858df33182f.NewOctocatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
// returns a *OrganizationsRequestBuilder when successful
func (m *ApiClient) Organizations()(*i2f58036685f6c1f423db5a62a4598667c8697cf896e332d260bbb3401974d900.OrganizationsRequestBuilder) {
    return i2f58036685f6c1f423db5a62a4598667c8697cf896e332d260bbb3401974d900.NewOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
// returns a *OrgsRequestBuilder when successful
func (m *ApiClient) Orgs()(*id8aedfc95cc545f984105f50cf0ac5fdd0d8e5ab227e4a886d53228f435f8aca.OrgsRequestBuilder) {
    return id8aedfc95cc545f984105f50cf0ac5fdd0d8e5ab227e4a886d53228f435f8aca.NewOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
// returns a *ProjectsRequestBuilder when successful
func (m *ApiClient) Projects()(*iee46eecc378acacb4378b35f14d2d75f182b1e9ad75681c6ca7ba5a8e05c3fef.ProjectsRequestBuilder) {
    return iee46eecc378acacb4378b35f14d2d75f182b1e9ad75681c6ca7ba5a8e05c3fef.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rate_limit the rate_limit property
// returns a *Rate_limitRequestBuilder when successful
func (m *ApiClient) Rate_limit()(*ie4198ced272c261cf44838663bf9483c5a9a9df8fe18cc7420e64a31cd4f7bf8.Rate_limitRequestBuilder) {
    return ie4198ced272c261cf44838663bf9483c5a9a9df8fe18cc7420e64a31cd4f7bf8.NewRate_limitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
// returns a *ReposRequestBuilder when successful
func (m *ApiClient) Repos()(*i66f747cc7e39f2550b27f76f1deff6b4eb7ef9fe9508531a28f5daa16f21bd0c.ReposRequestBuilder) {
    return i66f747cc7e39f2550b27f76f1deff6b4eb7ef9fe9508531a28f5daa16f21bd0c.NewReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
// returns a *RepositoriesRequestBuilder when successful
func (m *ApiClient) Repositories()(*i8c3efe8d3f9bba02f3585224c19bc147b1c29223805959ad6e1883c376334208.RepositoriesRequestBuilder) {
    return i8c3efe8d3f9bba02f3585224c19bc147b1c29223805959ad6e1883c376334208.NewRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Scim the scim property
// returns a *ScimRequestBuilder when successful
func (m *ApiClient) Scim()(*i7a2ac49b9b6f48e07dbfe8e254384c6396f35f44d01fec4711a22ac9a846ab3b.ScimRequestBuilder) {
    return i7a2ac49b9b6f48e07dbfe8e254384c6396f35f44d01fec4711a22ac9a846ab3b.NewScimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
// returns a *SearchRequestBuilder when successful
func (m *ApiClient) Search()(*if21e1a09c882b708cc2897e803a9731133fa5a954e4da837faea5a6611051822.SearchRequestBuilder) {
    return if21e1a09c882b708cc2897e803a9731133fa5a954e4da837faea5a6611051822.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
// returns a *TeamsRequestBuilder when successful
func (m *ApiClient) Teams()(*ia78ddb5870800bf7cef21851ecf22ff5e1dd993fd9fd4ad35d61f54673b524e8.TeamsRequestBuilder) {
    return ia78ddb5870800bf7cef21851ecf22ff5e1dd993fd9fd4ad35d61f54673b524e8.NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get Hypermedia links to resources accessible in GitHub's REST API
// returns a *RequestInformation when successful
func (m *ApiClient) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// User the user property
// returns a *UserRequestBuilder when successful
func (m *ApiClient) User()(*i1e25f49239426bd2e69db174edad2b458de62223f6832633c4998e12a9822fd0.UserRequestBuilder) {
    return i1e25f49239426bd2e69db174edad2b458de62223f6832633c4998e12a9822fd0.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
// returns a *UsersRequestBuilder when successful
func (m *ApiClient) Users()(*ie3b347c690ae9117f3a4da4f0e042e20664ce00896fa0a94cdd22b3a98748b55.UsersRequestBuilder) {
    return ie3b347c690ae9117f3a4da4f0e042e20664ce00896fa0a94cdd22b3a98748b55.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions the versions property
// returns a *VersionsRequestBuilder when successful
func (m *ApiClient) Versions()(*ifc535ad0dacf3b171683de34966458c4e8276e5eb0e62ac8f374245a8faad8f8.VersionsRequestBuilder) {
    return ifc535ad0dacf3b171683de34966458c4e8276e5eb0e62ac8f374245a8faad8f8.NewVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Zen the zen property
// returns a *ZenRequestBuilder when successful
func (m *ApiClient) Zen()(*iad80f7bb8538b815146b487f10b1d16c4e083c86d3018f0a8f7382b939f4e634.ZenRequestBuilder) {
    return iad80f7bb8538b815146b487f10b1d16c4e083c86d3018f0a8f7382b939f4e634.NewZenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
