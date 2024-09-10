package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i49bef2d41e85df11635b5f54fc591277b54ce82be042bb3486664ed17d3e94d5 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/migrations/item"
)

// ItemMigrationsWithMigration_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}
type ItemMigrationsWithMigration_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
type ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters struct {
    // Exclude attributes from the API response to improve performance
    Exclude []i49bef2d41e85df11635b5f54fc591277b54ce82be042bb3486664ed17d3e94d5.GetExcludeQueryParameterType `uriparametername:"exclude"`
}
// Archive the archive property
// returns a *ItemMigrationsItemArchiveRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Archive()(*ItemMigrationsItemArchiveRequestBuilder) {
    return NewItemMigrationsItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMigrationsWithMigration_ItemRequestBuilderInternal instantiates a new ItemMigrationsWithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    m := &ItemMigrationsWithMigration_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/migrations/{migration_id}{?exclude*}", pathParameters),
    }
    return m
}
// NewItemMigrationsWithMigration_ItemRequestBuilder instantiates a new ItemMigrationsWithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsWithMigration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
// returns a Migrationable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/migrations/orgs#get-an-organization-migration-status
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Migrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateMigrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Migrationable), nil
}
// Repos the repos property
// returns a *ItemMigrationsItemReposRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repos()(*ItemMigrationsItemReposRequestBuilder) {
    return NewItemMigrationsItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
// returns a *ItemMigrationsItemRepositoriesRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repositories()(*ItemMigrationsItemRepositoriesRequestBuilder) {
    return NewItemMigrationsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
// returns a *RequestInformation when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMigrationsWithMigration_ItemRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    return NewItemMigrationsWithMigration_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
