package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodespacesSecretsItemWithSecret_namePutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get the public key for the authenticated user](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#get-public-key-for-the-authenticated-user) endpoint.
    encrypted_value *string
    // ID of the key you used to encrypt the secret.
    key_id *string
    // An array of repository ids that can access the user secret. You can manage the list of selected repositories using the [List selected repositories for a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#list-selected-repositories-for-a-user-secret), [Set selected repositories for a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#set-selected-repositories-for-a-user-secret), and [Remove a selected repository from a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#remove-a-selected-repository-from-a-user-secret) endpoints.
    selected_repository_ids []int32
}
// NewCodespacesSecretsItemWithSecret_namePutRequestBody instantiates a new CodespacesSecretsItemWithSecret_namePutRequestBody and sets the default values.
func NewCodespacesSecretsItemWithSecret_namePutRequestBody()(*CodespacesSecretsItemWithSecret_namePutRequestBody) {
    m := &CodespacesSecretsItemWithSecret_namePutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesSecretsItemWithSecret_namePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespacesSecretsItemWithSecret_namePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesSecretsItemWithSecret_namePutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncryptedValue gets the encrypted_value property value. Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get the public key for the authenticated user](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#get-public-key-for-the-authenticated-user) endpoint.
// returns a *string when successful
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) GetEncryptedValue()(*string) {
    return m.encrypted_value
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encrypted_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedValue(val)
        }
        return nil
    }
    res["key_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
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
// GetKeyId gets the key_id property value. ID of the key you used to encrypt the secret.
// returns a *string when successful
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) GetKeyId()(*string) {
    return m.key_id
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. An array of repository ids that can access the user secret. You can manage the list of selected repositories using the [List selected repositories for a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#list-selected-repositories-for-a-user-secret), [Set selected repositories for a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#set-selected-repositories-for-a-user-secret), and [Remove a selected repository from a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#remove-a-selected-repository-from-a-user-secret) endpoints.
// returns a []int32 when successful
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// Serialize serializes information the current object
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encrypted_value", m.GetEncryptedValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key_id", m.GetKeyId())
        if err != nil {
            return err
        }
    }
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
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncryptedValue sets the encrypted_value property value. Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get the public key for the authenticated user](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#get-public-key-for-the-authenticated-user) endpoint.
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) SetEncryptedValue(value *string)() {
    m.encrypted_value = value
}
// SetKeyId sets the key_id property value. ID of the key you used to encrypt the secret.
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) SetKeyId(value *string)() {
    m.key_id = value
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. An array of repository ids that can access the user secret. You can manage the list of selected repositories using the [List selected repositories for a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#list-selected-repositories-for-a-user-secret), [Set selected repositories for a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#set-selected-repositories-for-a-user-secret), and [Remove a selected repository from a user secret](https://docs.github.com/enterprise-cloud@latest//rest/codespaces/secrets#remove-a-selected-repository-from-a-user-secret) endpoints.
func (m *CodespacesSecretsItemWithSecret_namePutRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
type CodespacesSecretsItemWithSecret_namePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedValue()(*string)
    GetKeyId()(*string)
    GetSelectedRepositoryIds()([]int32)
    SetEncryptedValue(value *string)()
    SetKeyId(value *string)()
    SetSelectedRepositoryIds(value []int32)()
}
