package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemDependabotSecretsItemRepositoriesPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the `visibility` is set to `selected`. You can add and remove individual repositories using the [Set selected repositories for an organization secret](https://docs.github.com/enterprise-cloud@latest//rest/dependabot/secrets#set-selected-repositories-for-an-organization-secret) and [Remove selected repository from an organization secret](https://docs.github.com/enterprise-cloud@latest//rest/dependabot/secrets#remove-selected-repository-from-an-organization-secret) endpoints.
    selected_repository_ids []int32
}
// NewItemDependabotSecretsItemRepositoriesPutRequestBody instantiates a new ItemDependabotSecretsItemRepositoriesPutRequestBody and sets the default values.
func NewItemDependabotSecretsItemRepositoriesPutRequestBody()(*ItemDependabotSecretsItemRepositoriesPutRequestBody) {
    m := &ItemDependabotSecretsItemRepositoriesPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemDependabotSecretsItemRepositoriesPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemDependabotSecretsItemRepositoriesPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDependabotSecretsItemRepositoriesPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemDependabotSecretsItemRepositoriesPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemDependabotSecretsItemRepositoriesPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["selected_repository_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetSelectedRepositoryIds(res)
        }
        return nil
    }
    return res
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the `visibility` is set to `selected`. You can add and remove individual repositories using the [Set selected repositories for an organization secret](https://docs.github.com/enterprise-cloud@latest//rest/dependabot/secrets#set-selected-repositories-for-an-organization-secret) and [Remove selected repository from an organization secret](https://docs.github.com/enterprise-cloud@latest//rest/dependabot/secrets#remove-selected-repository-from-an-organization-secret) endpoints.
// returns a []int32 when successful
func (m *ItemDependabotSecretsItemRepositoriesPutRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// Serialize serializes information the current object
func (m *ItemDependabotSecretsItemRepositoriesPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_repository_ids", m.GetSelectedRepositoryIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemDependabotSecretsItemRepositoriesPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the `visibility` is set to `selected`. You can add and remove individual repositories using the [Set selected repositories for an organization secret](https://docs.github.com/enterprise-cloud@latest//rest/dependabot/secrets#set-selected-repositories-for-an-organization-secret) and [Remove selected repository from an organization secret](https://docs.github.com/enterprise-cloud@latest//rest/dependabot/secrets#remove-selected-repository-from-an-organization-secret) endpoints.
func (m *ItemDependabotSecretsItemRepositoriesPutRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
type ItemDependabotSecretsItemRepositoriesPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedRepositoryIds()([]int32)
    SetSelectedRepositoryIds(value []int32)()
}
