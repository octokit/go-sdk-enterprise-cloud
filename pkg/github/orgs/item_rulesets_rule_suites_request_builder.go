package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i833ac8e90c82c8f666c7cd677c3d2548c7d68fbc156daf09cce37f63eedac29b "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/rulesets/rulesuites"
)

// ItemRulesetsRuleSuitesRequestBuilder builds and executes requests for operations under \orgs\{org}\rulesets\rule-suites
type ItemRulesetsRuleSuitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRulesetsRuleSuitesRequestBuilderGetQueryParameters lists suites of rule evaluations at the organization level.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/enterprise-cloud@latest//organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
type ItemRulesetsRuleSuitesRequestBuilderGetQueryParameters struct {
    // The handle for the GitHub user account to filter on. When specified, only rule evaluations triggered by this actor will be returned.
    Actor_name *string `uriparametername:"actor_name"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The name of the ref. Cannot contain wildcard characters. Optionally prefix with `refs/heads/` to limit to branches or `refs/tags/` to limit to tags. Omit the prefix to search across all refs. When specified, only rule evaluations triggered for this ref will be returned.
    Ref *string `uriparametername:"ref"`
    // The name of the repository to filter on.
    Repository_name *string `uriparametername:"repository_name"`
    // The rule results to filter on. When specified, only suites with this result will be returned.
    Rule_suite_result *i833ac8e90c82c8f666c7cd677c3d2548c7d68fbc156daf09cce37f63eedac29b.GetRule_suite_resultQueryParameterType `uriparametername:"rule_suite_result"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    Time_period *i833ac8e90c82c8f666c7cd677c3d2548c7d68fbc156daf09cce37f63eedac29b.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// ByRule_suite_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.orgs.item.rulesets.ruleSuites.item collection
// returns a *ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder when successful
func (m *ItemRulesetsRuleSuitesRequestBuilder) ByRule_suite_id(rule_suite_id int32)(*ItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["rule_suite_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(rule_suite_id), 10)
    return NewItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRulesetsRuleSuitesRequestBuilderInternal instantiates a new ItemRulesetsRuleSuitesRequestBuilder and sets the default values.
func NewItemRulesetsRuleSuitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRulesetsRuleSuitesRequestBuilder) {
    m := &ItemRulesetsRuleSuitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/rulesets/rule-suites{?actor_name*,page*,per_page*,ref*,repository_name*,rule_suite_result*,time_period*}", pathParameters),
    }
    return m
}
// NewItemRulesetsRuleSuitesRequestBuilder instantiates a new ItemRulesetsRuleSuitesRequestBuilder and sets the default values.
func NewItemRulesetsRuleSuitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRulesetsRuleSuitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRulesetsRuleSuitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists suites of rule evaluations at the organization level.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/enterprise-cloud@latest//organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
// returns a []RuleSuitesable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/rule-suites#list-organization-rule-suites
func (m *ItemRulesetsRuleSuitesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemRulesetsRuleSuitesRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RuleSuitesable, error) {
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
// ToGetRequestInformation lists suites of rule evaluations at the organization level.For more information, see "[Managing rulesets for repositories in your organization](https://docs.github.com/enterprise-cloud@latest//organizations/managing-organization-settings/managing-rulesets-for-repositories-in-your-organization#viewing-insights-for-rulesets)."
// returns a *RequestInformation when successful
func (m *ItemRulesetsRuleSuitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemRulesetsRuleSuitesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRulesetsRuleSuitesRequestBuilder when successful
func (m *ItemRulesetsRuleSuitesRequestBuilder) WithUrl(rawUrl string)(*ItemRulesetsRuleSuitesRequestBuilder) {
    return NewItemRulesetsRuleSuitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
