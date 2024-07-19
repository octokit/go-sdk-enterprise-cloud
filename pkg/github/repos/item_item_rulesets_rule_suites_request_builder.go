package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i89e73d086efc86ebad688074e3b9472802ac5a9b821caafcdcb3506fd794d57f "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos/item/item/rulesets/rulesuites"
)

// ItemItemRulesetsRuleSuitesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\rulesets\rule-suites
type ItemItemRulesetsRuleSuitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemRulesetsRuleSuitesRequestBuilderGetQueryParameters lists suites of rule evaluations at the repository level.For more information, see "[Managing rulesets for a repository](https://docs.github.com/enterprise-cloud@latest//repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/managing-rulesets-for-a-repository#viewing-insights-for-rulesets)."
type ItemItemRulesetsRuleSuitesRequestBuilderGetQueryParameters struct {
    // The handle for the GitHub user account to filter on. When specified, only rule evaluations triggered by this actor will be returned.
    Actor_name *string `uriparametername:"actor_name"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The name of the ref. Cannot contain wildcard characters. Optionally prefix with `refs/heads/` to limit to branches or `refs/tags/` to limit to tags. Omit the prefix to search across all refs. When specified, only rule evaluations triggered for this ref will be returned.
    Ref *string `uriparametername:"ref"`
    // The rule results to filter on. When specified, only suites with this result will be returned.
    Rule_suite_result *i89e73d086efc86ebad688074e3b9472802ac5a9b821caafcdcb3506fd794d57f.GetRule_suite_resultQueryParameterType `uriparametername:"rule_suite_result"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    Time_period *i89e73d086efc86ebad688074e3b9472802ac5a9b821caafcdcb3506fd794d57f.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// ByRule_suite_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.repos.item.item.rulesets.ruleSuites.item collection
// returns a *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder when successful
func (m *ItemItemRulesetsRuleSuitesRequestBuilder) ByRule_suite_id(rule_suite_id int32)(*ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["rule_suite_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(rule_suite_id), 10)
    return NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemRulesetsRuleSuitesRequestBuilderInternal instantiates a new ItemItemRulesetsRuleSuitesRequestBuilder and sets the default values.
func NewItemItemRulesetsRuleSuitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRuleSuitesRequestBuilder) {
    m := &ItemItemRulesetsRuleSuitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/rulesets/rule-suites{?actor_name*,page*,per_page*,ref*,rule_suite_result*,time_period*}", pathParameters),
    }
    return m
}
// NewItemItemRulesetsRuleSuitesRequestBuilder instantiates a new ItemItemRulesetsRuleSuitesRequestBuilder and sets the default values.
func NewItemItemRulesetsRuleSuitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRuleSuitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesetsRuleSuitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists suites of rule evaluations at the repository level.For more information, see "[Managing rulesets for a repository](https://docs.github.com/enterprise-cloud@latest//repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/managing-rulesets-for-a-repository#viewing-insights-for-rulesets)."
// returns a []RuleSuitesable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/repos/rule-suites#list-repository-rule-suites
func (m *ItemItemRulesetsRuleSuitesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemRulesetsRuleSuitesRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RuleSuitesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateRuleSuitesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RuleSuitesable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RuleSuitesable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists suites of rule evaluations at the repository level.For more information, see "[Managing rulesets for a repository](https://docs.github.com/enterprise-cloud@latest//repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/managing-rulesets-for-a-repository#viewing-insights-for-rulesets)."
// returns a *RequestInformation when successful
func (m *ItemItemRulesetsRuleSuitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemRulesetsRuleSuitesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemRulesetsRuleSuitesRequestBuilder when successful
func (m *ItemItemRulesetsRuleSuitesRequestBuilder) WithUrl(rawUrl string)(*ItemItemRulesetsRuleSuitesRequestBuilder) {
    return NewItemItemRulesetsRuleSuitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
