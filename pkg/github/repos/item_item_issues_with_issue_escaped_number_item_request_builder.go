package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemIssuesWithIssue_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\issues\{issue_number}
type ItemItemIssuesWithIssue_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Assignees the assignees property
// returns a *ItemItemIssuesItemAssigneesRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Assignees()(*ItemItemIssuesItemAssigneesRequestBuilder) {
    return NewItemItemIssuesItemAssigneesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
// returns a *ItemItemIssuesItemCommentsRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Comments()(*ItemItemIssuesItemCommentsRequestBuilder) {
    return NewItemItemIssuesItemCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal instantiates a new ItemItemIssuesWithIssue_numberItemRequestBuilder and sets the default values.
func NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    m := &ItemItemIssuesWithIssue_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/issues/{issue_number}", pathParameters),
    }
    return m
}
// NewItemItemIssuesWithIssue_numberItemRequestBuilder instantiates a new ItemItemIssuesWithIssue_numberItemRequestBuilder and sets the default values.
func NewItemItemIssuesWithIssue_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
// returns a *ItemItemIssuesItemEventsRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Events()(*ItemItemIssuesItemEventsRequestBuilder) {
    return NewItemItemIssuesItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the API returns a [`301 Moved Permanently` status](https://docs.github.com/enterprise-cloud@latest//rest/guides/best-practices-for-using-the-rest-api#follow-redirects) if the issue was[transferred](https://docs.github.com/enterprise-cloud@latest//articles/transferring-an-issue-to-another-repository/) to another repository. Ifthe issue was transferred to or deleted from a repository where the authenticated user lacks read access, the APIreturns a `404 Not Found` status. If the issue was deleted from a repository where the authenticated user has readaccess, the API returns a `410 Gone` status. To receive webhook events for transferred and deleted issues, subscribeto the [`issues`](https://docs.github.com/enterprise-cloud@latest//webhooks/event-payloads/#issues) webhook.> [!NOTE]> GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For this reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull request id, use the "[List pull requests](https://docs.github.com/enterprise-cloud@latest//rest/pulls/pulls#list-pull-requests)" endpoint.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a Issueable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 410 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/issues/issues#get-an-issue
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Issueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "410": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Issueable), nil
}
// Labels the labels property
// returns a *ItemItemIssuesItemLabelsRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Labels()(*ItemItemIssuesItemLabelsRequestBuilder) {
    return NewItemItemIssuesItemLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lock the lock property
// returns a *ItemItemIssuesItemLockRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Lock()(*ItemItemIssuesItemLockRequestBuilder) {
    return NewItemItemIssuesItemLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch issue owners and users with push access or Triage role can edit an issue.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a Issueable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 410 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a Issue503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/issues/issues#update-an-issue
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Patch(ctx context.Context, body ItemItemIssuesItemWithIssue_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Issueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "410": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
        "503": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateIssue503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Issueable), nil
}
// Reactions the reactions property
// returns a *ItemItemIssuesItemReactionsRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Reactions()(*ItemItemIssuesItemReactionsRequestBuilder) {
    return NewItemItemIssuesItemReactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sub_issue the sub_issue property
// returns a *ItemItemIssuesItemSub_issueRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Sub_issue()(*ItemItemIssuesItemSub_issueRequestBuilder) {
    return NewItemItemIssuesItemSub_issueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sub_issues the sub_issues property
// returns a *ItemItemIssuesItemSub_issuesRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Sub_issues()(*ItemItemIssuesItemSub_issuesRequestBuilder) {
    return NewItemItemIssuesItemSub_issuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Timeline the timeline property
// returns a *ItemItemIssuesItemTimelineRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Timeline()(*ItemItemIssuesItemTimelineRequestBuilder) {
    return NewItemItemIssuesItemTimelineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the API returns a [`301 Moved Permanently` status](https://docs.github.com/enterprise-cloud@latest//rest/guides/best-practices-for-using-the-rest-api#follow-redirects) if the issue was[transferred](https://docs.github.com/enterprise-cloud@latest//articles/transferring-an-issue-to-another-repository/) to another repository. Ifthe issue was transferred to or deleted from a repository where the authenticated user lacks read access, the APIreturns a `404 Not Found` status. If the issue was deleted from a repository where the authenticated user has readaccess, the API returns a `410 Gone` status. To receive webhook events for transferred and deleted issues, subscribeto the [`issues`](https://docs.github.com/enterprise-cloud@latest//webhooks/event-payloads/#issues) webhook.> [!NOTE]> GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For this reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull request id, use the "[List pull requests](https://docs.github.com/enterprise-cloud@latest//rest/pulls/pulls#list-pull-requests)" endpoint.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a *RequestInformation when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation issue owners and users with push access or Triage role can edit an issue.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a *RequestInformation when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemIssuesItemWithIssue_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemIssuesWithIssue_numberItemRequestBuilder when successful
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    return NewItemItemIssuesWithIssue_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
