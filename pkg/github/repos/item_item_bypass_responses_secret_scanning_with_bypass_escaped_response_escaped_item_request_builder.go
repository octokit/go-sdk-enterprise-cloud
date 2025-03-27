package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-responses\secret-scanning\{bypass_response_id}
type ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilderInternal instantiates a new ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder and sets the default values.
func NewItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder) {
    m := &ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-responses/secret-scanning/{bypass_response_id}", pathParameters),
    }
    return m
}
// NewItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder instantiates a new ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder and sets the default values.
func NewItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete dissmiss a response given to a bypass request for secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/secret-scanning/delegated-bypass#dismiss-a-response-on-a-bypass-request-for-secret-scanning
func (m *ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation dissmiss a response given to a bypass request for secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder when successful
func (m *ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder) {
    return NewItemItemBypassResponsesSecretScanningWithBypass_response_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
