package orgs

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemExternalGroupRequestBuilder builds and executes requests for operations under \orgs\{org}\external-group
type ItemExternalGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByGroup_id gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.orgs.item.externalGroup.item collection
// returns a *ItemExternalGroupWithGroup_ItemRequestBuilder when successful
func (m *ItemExternalGroupRequestBuilder) ByGroup_id(group_id int32)(*ItemExternalGroupWithGroup_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["group_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(group_id), 10)
    return NewItemExternalGroupWithGroup_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemExternalGroupRequestBuilderInternal instantiates a new ItemExternalGroupRequestBuilder and sets the default values.
func NewItemExternalGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExternalGroupRequestBuilder) {
    m := &ItemExternalGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/external-group", pathParameters),
    }
    return m
}
// NewItemExternalGroupRequestBuilder instantiates a new ItemExternalGroupRequestBuilder and sets the default values.
func NewItemExternalGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExternalGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemExternalGroupRequestBuilderInternal(urlParams, requestAdapter)
}
