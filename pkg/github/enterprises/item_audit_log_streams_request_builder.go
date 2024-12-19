package enterprises

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemAuditLogStreamsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\audit-log\streams
type ItemAuditLogStreamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByStream_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.enterprises.item.auditLog.streams.item collection
// returns a *ItemAuditLogStreamsWithStream_ItemRequestBuilder when successful
func (m *ItemAuditLogStreamsRequestBuilder) ByStream_id(stream_id int32)(*ItemAuditLogStreamsWithStream_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["stream_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(stream_id), 10)
    return NewItemAuditLogStreamsWithStream_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuditLogStreamsRequestBuilderInternal instantiates a new ItemAuditLogStreamsRequestBuilder and sets the default values.
func NewItemAuditLogStreamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogStreamsRequestBuilder) {
    m := &ItemAuditLogStreamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/audit-log/streams", pathParameters),
    }
    return m
}
// NewItemAuditLogStreamsRequestBuilder instantiates a new ItemAuditLogStreamsRequestBuilder and sets the default values.
func NewItemAuditLogStreamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogStreamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuditLogStreamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the configured audit log streaming configurations for an enterprise.This only lists configured streams for supported providers.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a []GetAuditLogStreamConfigsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/audit-log#list-audit-log-stream-configurations-for-an-enterprise
func (m *ItemAuditLogStreamsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateGetAuditLogStreamConfigsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigsable)
        }
    }
    return val, nil
}
// Post creates an audit log streaming configuration for any of the supported streaming endpoints: Azure Blob Storage, Azure Event Hubs, Amazon S3, Splunk, Google Cloud Storage, Datadog.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a GetAuditLogStreamConfigable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/audit-log#create-an-audit-log-streaming-configuration-for-an-enterprise
func (m *ItemAuditLogStreamsRequestBuilder) Post(ctx context.Context, body ItemAuditLogStreamsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.GetAuditLogStreamConfigable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation lists the configured audit log streaming configurations for an enterprise.This only lists configured streams for supported providers.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a *RequestInformation when successful
func (m *ItemAuditLogStreamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates an audit log streaming configuration for any of the supported streaming endpoints: Azure Blob Storage, Azure Event Hubs, Amazon S3, Splunk, Google Cloud Storage, Datadog.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a *RequestInformation when successful
func (m *ItemAuditLogStreamsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAuditLogStreamsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuditLogStreamsRequestBuilder when successful
func (m *ItemAuditLogStreamsRequestBuilder) WithUrl(rawUrl string)(*ItemAuditLogStreamsRequestBuilder) {
    return NewItemAuditLogStreamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
