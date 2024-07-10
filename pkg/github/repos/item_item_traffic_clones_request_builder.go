package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i4d9c2a1d9d9a5c0921332f4a1fdd0cc3c7676d7aff750f24006f853283ee3665 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos/item/item/traffic/clones"
)

// ItemItemTrafficClonesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\traffic\clones
type ItemItemTrafficClonesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTrafficClonesRequestBuilderGetQueryParameters get the total number of clones and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
type ItemItemTrafficClonesRequestBuilderGetQueryParameters struct {
    // The time frame to display results for.
    Per *i4d9c2a1d9d9a5c0921332f4a1fdd0cc3c7676d7aff750f24006f853283ee3665.GetPerQueryParameterType `uriparametername:"per"`
}
// NewItemItemTrafficClonesRequestBuilderInternal instantiates a new ItemItemTrafficClonesRequestBuilder and sets the default values.
func NewItemItemTrafficClonesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficClonesRequestBuilder) {
    m := &ItemItemTrafficClonesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/traffic/clones{?per*}", pathParameters),
    }
    return m
}
// NewItemItemTrafficClonesRequestBuilder instantiates a new ItemItemTrafficClonesRequestBuilder and sets the default values.
func NewItemItemTrafficClonesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficClonesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficClonesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the total number of clones and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
// returns a CloneTrafficable when successful
// returns a BasicError error when the service returns a 403 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/metrics/traffic#get-repository-clones
func (m *ItemItemTrafficClonesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemTrafficClonesRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CloneTrafficable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateCloneTrafficFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CloneTrafficable), nil
}
// ToGetRequestInformation get the total number of clones and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
// returns a *RequestInformation when successful
func (m *ItemItemTrafficClonesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemTrafficClonesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemTrafficClonesRequestBuilder when successful
func (m *ItemItemTrafficClonesRequestBuilder) WithUrl(rawUrl string)(*ItemItemTrafficClonesRequestBuilder) {
    return NewItemItemTrafficClonesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
