package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    ib37e34869e811df582ca7d33a04e21ffb04b609da48f184ddcb796e1b622cac1 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos/item/item/activity"
)

// ItemItemActivityRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\activity
type ItemItemActivityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActivityRequestBuilderGetQueryParameters lists a detailed history of changes to a repository, such as pushes, merges, force pushes, and branch changes, and associates these changes with commits and users.For more information about viewing repository activity,see "[Viewing activity and data for your repository](https://docs.github.com/enterprise-cloud@latest//repositories/viewing-activity-and-data-for-your-repository)."
type ItemItemActivityRequestBuilderGetQueryParameters struct {
    // The activity type to filter by.For example, you can choose to filter by "force_push", to see all force pushes to the repository.
    Activity_type *ib37e34869e811df582ca7d33a04e21ffb04b609da48f184ddcb796e1b622cac1.GetActivity_typeQueryParameterType `uriparametername:"activity_type"`
    // The GitHub username to use to filter by the actor who performed the activity.
    Actor *string `uriparametername:"actor"`
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    Direction *ib37e34869e811df582ca7d33a04e21ffb04b609da48f184ddcb796e1b622cac1.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The Git reference for the activities you want to list.The `ref` for a branch can be formatted either as `refs/heads/BRANCH_NAME` or `BRANCH_NAME`, where `BRANCH_NAME` is the name of your branch.
    Ref *string `uriparametername:"ref"`
    // The time period to filter by.For example, `day` will filter for activity that occurred in the past 24 hours, and `week` will filter for activity that occurred in the past 7 days (168 hours).
    Time_period *ib37e34869e811df582ca7d33a04e21ffb04b609da48f184ddcb796e1b622cac1.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// NewItemItemActivityRequestBuilderInternal instantiates a new ItemItemActivityRequestBuilder and sets the default values.
func NewItemItemActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActivityRequestBuilder) {
    m := &ItemItemActivityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/activity{?activity_type*,actor*,after*,before*,direction*,per_page*,ref*,time_period*}", pathParameters),
    }
    return m
}
// NewItemItemActivityRequestBuilder instantiates a new ItemItemActivityRequestBuilder and sets the default values.
func NewItemItemActivityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActivityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActivityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists a detailed history of changes to a repository, such as pushes, merges, force pushes, and branch changes, and associates these changes with commits and users.For more information about viewing repository activity,see "[Viewing activity and data for your repository](https://docs.github.com/enterprise-cloud@latest//repositories/viewing-activity-and-data-for-your-repository)."
// returns a []Activityable when successful
// returns a ValidationErrorSimple error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/repos/repos#list-repository-activities
func (m *ItemItemActivityRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActivityRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Activityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Activityable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Activityable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists a detailed history of changes to a repository, such as pushes, merges, force pushes, and branch changes, and associates these changes with commits and users.For more information about viewing repository activity,see "[Viewing activity and data for your repository](https://docs.github.com/enterprise-cloud@latest//repositories/viewing-activity-and-data-for-your-repository)."
// returns a *RequestInformation when successful
func (m *ItemItemActivityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActivityRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActivityRequestBuilder when successful
func (m *ItemItemActivityRequestBuilder) WithUrl(rawUrl string)(*ItemItemActivityRequestBuilder) {
    return NewItemItemActivityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
