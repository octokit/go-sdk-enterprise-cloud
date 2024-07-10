package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

type ItemItemHooksItemConfigPatchRequestBody struct {
    // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
    content_type *string
    // The insecure_ssl property
    insecure_ssl i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.WebhookConfigInsecureSslable
    // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/enterprise-cloud@latest//webhooks/event-payloads/#delivery-headers).
    secret *string
    // The URL to which the payloads will be delivered.
    url *string
}
// NewItemItemHooksItemConfigPatchRequestBody instantiates a new ItemItemHooksItemConfigPatchRequestBody and sets the default values.
func NewItemItemHooksItemConfigPatchRequestBody()(*ItemItemHooksItemConfigPatchRequestBody) {
    m := &ItemItemHooksItemConfigPatchRequestBody{
    }
    return m
}
// CreateItemItemHooksItemConfigPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemHooksItemConfigPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemHooksItemConfigPatchRequestBody(), nil
}
// GetContentType gets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
// returns a *string when successful
func (m *ItemItemHooksItemConfigPatchRequestBody) GetContentType()(*string) {
    return m.content_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemHooksItemConfigPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["insecure_ssl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateWebhookConfigInsecureSslFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsecureSsl(val.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.WebhookConfigInsecureSslable))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetInsecureSsl gets the insecure_ssl property value. The insecure_ssl property
// returns a WebhookConfigInsecureSslable when successful
func (m *ItemItemHooksItemConfigPatchRequestBody) GetInsecureSsl()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.WebhookConfigInsecureSslable) {
    return m.insecure_ssl
}
// GetSecret gets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/enterprise-cloud@latest//webhooks/event-payloads/#delivery-headers).
// returns a *string when successful
func (m *ItemItemHooksItemConfigPatchRequestBody) GetSecret()(*string) {
    return m.secret
}
// GetUrl gets the url property value. The URL to which the payloads will be delivered.
// returns a *string when successful
func (m *ItemItemHooksItemConfigPatchRequestBody) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ItemItemHooksItemConfigPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("insecure_ssl", m.GetInsecureSsl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentType sets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *ItemItemHooksItemConfigPatchRequestBody) SetContentType(value *string)() {
    m.content_type = value
}
// SetInsecureSsl sets the insecure_ssl property value. The insecure_ssl property
func (m *ItemItemHooksItemConfigPatchRequestBody) SetInsecureSsl(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.WebhookConfigInsecureSslable)() {
    m.insecure_ssl = value
}
// SetSecret sets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/enterprise-cloud@latest//webhooks/event-payloads/#delivery-headers).
func (m *ItemItemHooksItemConfigPatchRequestBody) SetSecret(value *string)() {
    m.secret = value
}
// SetUrl sets the url property value. The URL to which the payloads will be delivered.
func (m *ItemItemHooksItemConfigPatchRequestBody) SetUrl(value *string)() {
    m.url = value
}
type ItemItemHooksItemConfigPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(*string)
    GetInsecureSsl()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.WebhookConfigInsecureSslable)
    GetSecret()(*string)
    GetUrl()(*string)
    SetContentType(value *string)()
    SetInsecureSsl(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.WebhookConfigInsecureSslable)()
    SetSecret(value *string)()
    SetUrl(value *string)()
}
