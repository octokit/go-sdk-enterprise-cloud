package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\settings\billing\cost-centers\{cost_center_id}
type ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilderInternal instantiates a new ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder and sets the default values.
func NewItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder) {
    m := &ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}", pathParameters),
    }
    return m
}
// NewItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder instantiates a new ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder and sets the default values.
func NewItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Resource the resource property
// returns a *ItemSettingsBillingCostCentersItemResourceRequestBuilder when successful
func (m *ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder) Resource()(*ItemSettingsBillingCostCentersItemResourceRequestBuilder) {
    return NewItemSettingsBillingCostCentersItemResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
