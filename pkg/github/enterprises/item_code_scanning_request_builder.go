package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCodeScanningRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\code-scanning
type ItemCodeScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Alerts the alerts property
// returns a *ItemCodeScanningAlertsRequestBuilder when successful
func (m *ItemCodeScanningRequestBuilder) Alerts()(*ItemCodeScanningAlertsRequestBuilder) {
    return NewItemCodeScanningAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCodeScanningRequestBuilderInternal instantiates a new ItemCodeScanningRequestBuilder and sets the default values.
func NewItemCodeScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeScanningRequestBuilder) {
    m := &ItemCodeScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/code-scanning", pathParameters),
    }
    return m
}
// NewItemCodeScanningRequestBuilder instantiates a new ItemCodeScanningRequestBuilder and sets the default values.
func NewItemCodeScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeScanningRequestBuilderInternal(urlParams, requestAdapter)
}
