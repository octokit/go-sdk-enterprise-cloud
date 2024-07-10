package enterpriseinstallation

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithEnterprise_or_orgItemRequestBuilder builds and executes requests for operations under \enterprise-installation\{enterprise_or_org}
type WithEnterprise_or_orgItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithEnterprise_or_orgItemRequestBuilderInternal instantiates a new WithEnterprise_or_orgItemRequestBuilder and sets the default values.
func NewWithEnterprise_or_orgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterprise_or_orgItemRequestBuilder) {
    m := &WithEnterprise_or_orgItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprise-installation/{enterprise_or_org}", pathParameters),
    }
    return m
}
// NewWithEnterprise_or_orgItemRequestBuilder instantiates a new WithEnterprise_or_orgItemRequestBuilder and sets the default values.
func NewWithEnterprise_or_orgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterprise_or_orgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithEnterprise_or_orgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ServerStatistics the serverStatistics property
// returns a *ItemServerStatisticsRequestBuilder when successful
func (m *WithEnterprise_or_orgItemRequestBuilder) ServerStatistics()(*ItemServerStatisticsRequestBuilder) {
    return NewItemServerStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
