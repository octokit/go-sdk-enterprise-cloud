package organizations

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithOrganization_ItemRequestBuilder builds and executes requests for operations under \organizations\{organization_id}
type WithOrganization_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithOrganization_ItemRequestBuilderInternal instantiates a new WithOrganization_ItemRequestBuilder and sets the default values.
func NewWithOrganization_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrganization_ItemRequestBuilder) {
    m := &WithOrganization_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organizations/{organization_id}", pathParameters),
    }
    return m
}
// NewWithOrganization_ItemRequestBuilder instantiates a new WithOrganization_ItemRequestBuilder and sets the default values.
func NewWithOrganization_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrganization_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithOrganization_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custom_roles the custom_roles property
// returns a *ItemCustom_rolesRequestBuilder when successful
func (m *WithOrganization_ItemRequestBuilder) Custom_roles()(*ItemCustom_rolesRequestBuilder) {
    return NewItemCustom_rolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
