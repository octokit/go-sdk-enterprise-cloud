package organizations

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Organization_ItemRequestBuilder builds and executes requests for operations under \organizations\{organization_-id}
type Organization_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewOrganization_ItemRequestBuilderInternal instantiates a new Organization_ItemRequestBuilder and sets the default values.
func NewOrganization_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Organization_ItemRequestBuilder) {
    m := &Organization_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organizations/{organization_%2Did}", pathParameters),
    }
    return m
}
// NewOrganization_ItemRequestBuilder instantiates a new Organization_ItemRequestBuilder and sets the default values.
func NewOrganization_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Organization_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganization_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custom_roles the custom_roles property
// returns a *ItemCustom_rolesRequestBuilder when successful
func (m *Organization_ItemRequestBuilder) Custom_roles()(*ItemCustom_rolesRequestBuilder) {
    return NewItemCustom_rolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
// returns a *ItemSettingsRequestBuilder when successful
func (m *Organization_ItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
