package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemAuditLogStreamsWithStream_ItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\audit-log\streams\{stream_id}
type ItemAuditLogStreamsWithStream_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAuditLogStreamsWithStream_ItemRequestBuilderInternal instantiates a new ItemAuditLogStreamsWithStream_ItemRequestBuilder and sets the default values.
func NewItemAuditLogStreamsWithStream_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogStreamsWithStream_ItemRequestBuilder) {
    m := &ItemAuditLogStreamsWithStream_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/audit-log/streams/{stream_id}", pathParameters),
    }
    return m
}
// NewItemAuditLogStreamsWithStream_ItemRequestBuilder instantiates a new ItemAuditLogStreamsWithStream_ItemRequestBuilder and sets the default values.
func NewItemAuditLogStreamsWithStream_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogStreamsWithStream_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuditLogStreamsWithStream_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes an existing audit log stream configuration for an enterprise. When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/audit-log#delete-an-audit-log-streaming-configuration-for-an-enterprise
func (m *ItemAuditLogStreamsWithStream_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get lists one audit log stream configuration via a stream ID.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a GetAuditLogStreamConfigable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/audit-log#list-one-audit-log-streaming-configuration-via-a-stream-id
func (m *ItemAuditLogStreamsWithStream_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateGetAuditLogStreamConfigFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigable), nil
}
// Put updates an existing audit log stream configuration for an enterprise.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a GetAuditLogStreamConfigable when successful
// returns a ItemAuditLogStreamsItemGetAuditLogStreamConfig422Error error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/audit-log#update-an-existing-audit-log-stream-configuration
func (m *ItemAuditLogStreamsWithStream_ItemRequestBuilder) Put(ctx context.Context, body ItemAuditLogStreamsItemWithStream_PutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": CreateItemAuditLogStreamsItemGetAuditLogStreamConfig422ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateGetAuditLogStreamConfigFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigable), nil
}
// ToDeleteRequestInformation deletes an existing audit log stream configuration for an enterprise. When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a *RequestInformation when successful
func (m *ItemAuditLogStreamsWithStream_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation lists one audit log stream configuration via a stream ID.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a *RequestInformation when successful
func (m *ItemAuditLogStreamsWithStream_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation updates an existing audit log stream configuration for an enterprise.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a *RequestInformation when successful
func (m *ItemAuditLogStreamsWithStream_ItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemAuditLogStreamsItemWithStream_PutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuditLogStreamsWithStream_ItemRequestBuilder when successful
func (m *ItemAuditLogStreamsWithStream_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuditLogStreamsWithStream_ItemRequestBuilder) {
    return NewItemAuditLogStreamsWithStream_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
