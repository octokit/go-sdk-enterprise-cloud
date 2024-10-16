package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemHooksItemDeliveriesItemAttemptsRequestBuilder builds and executes requests for operations under \orgs\{org}\hooks\{hook_id}\deliveries\{delivery_id}\attempts
type ItemHooksItemDeliveriesItemAttemptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal instantiates a new ItemHooksItemDeliveriesItemAttemptsRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    m := &ItemHooksItemDeliveriesItemAttemptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}/attempts", pathParameters),
    }
    return m
}
// NewItemHooksItemDeliveriesItemAttemptsRequestBuilder instantiates a new ItemHooksItemDeliveriesItemAttemptsRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesItemAttemptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post you must be an organization owner or have the "Manage organization webhooks" permission to use this endpoint.OAuth app tokens and personal access tokens (classic) need `admin:org_hook` scope. OAuth apps cannot list, view, or edit webhooksthat they did not create and users cannot list, view, or edit webhooks that were created by OAuth apps.
// returns a ItemHooksItemDeliveriesItemAttemptsPostResponseable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/webhooks#redeliver-a-delivery-for-an-organization-webhook
func (m *ItemHooksItemDeliveriesItemAttemptsRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemHooksItemDeliveriesItemAttemptsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemHooksItemDeliveriesItemAttemptsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemHooksItemDeliveriesItemAttemptsPostResponseable), nil
}
// ToPostRequestInformation you must be an organization owner or have the "Manage organization webhooks" permission to use this endpoint.OAuth app tokens and personal access tokens (classic) need `admin:org_hook` scope. OAuth apps cannot list, view, or edit webhooksthat they did not create and users cannot list, view, or edit webhooks that were created by OAuth apps.
// returns a *RequestInformation when successful
func (m *ItemHooksItemDeliveriesItemAttemptsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemHooksItemDeliveriesItemAttemptsRequestBuilder when successful
func (m *ItemHooksItemDeliveriesItemAttemptsRequestBuilder) WithUrl(rawUrl string)(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    return NewItemHooksItemDeliveriesItemAttemptsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
