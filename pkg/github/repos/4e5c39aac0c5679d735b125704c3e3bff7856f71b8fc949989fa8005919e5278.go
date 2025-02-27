package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-requests\secret-scanning\{bypass_request_number}
type ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilderInternal instantiates a new ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder and sets the default values.
func NewItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) {
    m := &ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-requests/secret-scanning/{bypass_request_number}", pathParameters),
    }
    return m
}
// NewItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder instantiates a new ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder and sets the default values.
func NewItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a specific request to bypass secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a SecretScanningBypassRequestable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/secret-scanning/delegated-bypass#get-a-bypass-request-for-secret-scanning
func (m *ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SecretScanningBypassRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateSecretScanningBypassRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SecretScanningBypassRequestable), nil
}
// Patch approve or deny a request to bypass secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/secret-scanning/delegated-bypass#review-a-bypass-request-for-secret-scanning
func (m *ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) Patch(ctx context.Context, body ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchResponseable), nil
}
// ToGetRequestInformation gets a specific request to bypass secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation approve or deny a request to bypass secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemBypassRequestsSecretScanningItemWithBypass_request_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder when successful
func (m *ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) {
    return NewItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
