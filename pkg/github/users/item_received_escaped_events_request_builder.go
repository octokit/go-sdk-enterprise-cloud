package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemReceived_eventsRequestBuilder builds and executes requests for operations under \users\{username}\received_events
type ItemReceived_eventsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemReceived_eventsRequestBuilderGetQueryParameters these are events that you've received by watching repositories and following users. If you are authenticated as thegiven user, you will see private events. Otherwise, you'll only see public events.> [!NOTE]> This API is not built to serve real-time use cases. Depending on the time of day, event latency can be anywhere from 30s to 6h.
type ItemReceived_eventsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemReceived_eventsRequestBuilderInternal instantiates a new ItemReceived_eventsRequestBuilder and sets the default values.
func NewItemReceived_eventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReceived_eventsRequestBuilder) {
    m := &ItemReceived_eventsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/received_events{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemReceived_eventsRequestBuilder instantiates a new ItemReceived_eventsRequestBuilder and sets the default values.
func NewItemReceived_eventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReceived_eventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReceived_eventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get these are events that you've received by watching repositories and following users. If you are authenticated as thegiven user, you will see private events. Otherwise, you'll only see public events.> [!NOTE]> This API is not built to serve real-time use cases. Depending on the time of day, event latency can be anywhere from 30s to 6h.
// returns a []Eventable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/activity/events#list-events-received-by-the-authenticated-user
func (m *ItemReceived_eventsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemReceived_eventsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateEventFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Eventable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.Eventable)
        }
    }
    return val, nil
}
// Public the public property
// returns a *ItemReceived_eventsPublicRequestBuilder when successful
func (m *ItemReceived_eventsRequestBuilder) Public()(*ItemReceived_eventsPublicRequestBuilder) {
    return NewItemReceived_eventsPublicRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation these are events that you've received by watching repositories and following users. If you are authenticated as thegiven user, you will see private events. Otherwise, you'll only see public events.> [!NOTE]> This API is not built to serve real-time use cases. Depending on the time of day, event latency can be anywhere from 30s to 6h.
// returns a *RequestInformation when successful
func (m *ItemReceived_eventsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemReceived_eventsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemReceived_eventsRequestBuilder when successful
func (m *ItemReceived_eventsRequestBuilder) WithUrl(rawUrl string)(*ItemReceived_eventsRequestBuilder) {
    return NewItemReceived_eventsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
