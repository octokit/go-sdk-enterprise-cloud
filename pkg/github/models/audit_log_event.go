package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuditLogEvent struct {
    // A unique identifier for an audit event.
    _document_id *string
    // The name of the action that was performed, for example `user.login` or `repo.create`.
    action *string
    // The active property
    active *bool
    // The active_was property
    active_was *bool
    // The actor who performed the action.
    actor *string
    // The id of the actor who performed the action.
    actor_id *int32
    // The actor_location property
    actor_location AuditLogEvent_actor_locationable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The username of the account being blocked.
    blocked_user *string
    // The business property
    business *string
    // The business_id property
    business_id *int32
    // The config property
    config []AuditLogEvent_configable
    // The config_was property
    config_was []AuditLogEvent_config_wasable
    // The content_type property
    content_type *string
    // The time the audit log event was recorded, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time).
    created_at *int32
    // The data property
    data AuditLogEvent_dataable
    // The deploy_key_fingerprint property
    deploy_key_fingerprint *string
    // The emoji property
    emoji *string
    // The events property
    events []AuditLogEvent_eventsable
    // The events_were property
    events_were []AuditLogEvent_events_wereable
    // The explanation property
    explanation *string
    // The fingerprint property
    fingerprint *string
    // The hook_id property
    hook_id *int32
    // The limited_availability property
    limited_availability *bool
    // The message property
    message *string
    // The name property
    name *string
    // The old_user property
    old_user *string
    // The openssh_public_key property
    openssh_public_key *string
    // The operation_type property
    operation_type *string
    // The org property
    org *string
    // The org_id property
    org_id *int32
    // The previous_visibility property
    previous_visibility *string
    // The read_only property
    read_only *bool
    // The name of the repository.
    repo *string
    // The name of the repository.
    repository *string
    // The repository_public property
    repository_public *bool
    // The target_login property
    target_login *string
    // The team property
    team *string
    // The time the audit log event occurred, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time).
    timestamp *int32
    // The type of protocol (for example, HTTP or SSH) used to transfer Git data.
    transport_protocol *int32
    // A human readable name for the protocol (for example, HTTP or SSH) used to transfer Git data.
    transport_protocol_name *string
    // The user that was affected by the action performed (if available).
    user *string
    // The user_id property
    user_id *int32
    // The repository visibility, for example `public` or `private`.
    visibility *string
}
// NewAuditLogEvent instantiates a new AuditLogEvent and sets the default values.
func NewAuditLogEvent()(*AuditLogEvent) {
    m := &AuditLogEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuditLogEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditLogEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogEvent(), nil
}
// GetAction gets the action property value. The name of the action that was performed, for example `user.login` or `repo.create`.
// returns a *string when successful
func (m *AuditLogEvent) GetAction()(*string) {
    return m.action
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *AuditLogEvent) GetActive()(*bool) {
    return m.active
}
// GetActiveWas gets the active_was property value. The active_was property
// returns a *bool when successful
func (m *AuditLogEvent) GetActiveWas()(*bool) {
    return m.active_was
}
// GetActor gets the actor property value. The actor who performed the action.
// returns a *string when successful
func (m *AuditLogEvent) GetActor()(*string) {
    return m.actor
}
// GetActorId gets the actor_id property value. The id of the actor who performed the action.
// returns a *int32 when successful
func (m *AuditLogEvent) GetActorId()(*int32) {
    return m.actor_id
}
// GetActorLocation gets the actor_location property value. The actor_location property
// returns a AuditLogEvent_actor_locationable when successful
func (m *AuditLogEvent) GetActorLocation()(AuditLogEvent_actor_locationable) {
    return m.actor_location
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuditLogEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlockedUser gets the blocked_user property value. The username of the account being blocked.
// returns a *string when successful
func (m *AuditLogEvent) GetBlockedUser()(*string) {
    return m.blocked_user
}
// GetBusiness gets the business property value. The business property
// returns a *string when successful
func (m *AuditLogEvent) GetBusiness()(*string) {
    return m.business
}
// GetBusinessId gets the business_id property value. The business_id property
// returns a *int32 when successful
func (m *AuditLogEvent) GetBusinessId()(*int32) {
    return m.business_id
}
// GetConfig gets the config property value. The config property
// returns a []AuditLogEvent_configable when successful
func (m *AuditLogEvent) GetConfig()([]AuditLogEvent_configable) {
    return m.config
}
// GetConfigWas gets the config_was property value. The config_was property
// returns a []AuditLogEvent_config_wasable when successful
func (m *AuditLogEvent) GetConfigWas()([]AuditLogEvent_config_wasable) {
    return m.config_was
}
// GetContentType gets the content_type property value. The content_type property
// returns a *string when successful
func (m *AuditLogEvent) GetContentType()(*string) {
    return m.content_type
}
// GetCreatedAt gets the created_at property value. The time the audit log event was recorded, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time).
// returns a *int32 when successful
func (m *AuditLogEvent) GetCreatedAt()(*int32) {
    return m.created_at
}
// GetData gets the data property value. The data property
// returns a AuditLogEvent_dataable when successful
func (m *AuditLogEvent) GetData()(AuditLogEvent_dataable) {
    return m.data
}
// GetDeployKeyFingerprint gets the deploy_key_fingerprint property value. The deploy_key_fingerprint property
// returns a *string when successful
func (m *AuditLogEvent) GetDeployKeyFingerprint()(*string) {
    return m.deploy_key_fingerprint
}
// GetDocumentId gets the _document_id property value. A unique identifier for an audit event.
// returns a *string when successful
func (m *AuditLogEvent) GetDocumentId()(*string) {
    return m._document_id
}
// GetEmoji gets the emoji property value. The emoji property
// returns a *string when successful
func (m *AuditLogEvent) GetEmoji()(*string) {
    return m.emoji
}
// GetEvents gets the events property value. The events property
// returns a []AuditLogEvent_eventsable when successful
func (m *AuditLogEvent) GetEvents()([]AuditLogEvent_eventsable) {
    return m.events
}
// GetEventsWere gets the events_were property value. The events_were property
// returns a []AuditLogEvent_events_wereable when successful
func (m *AuditLogEvent) GetEventsWere()([]AuditLogEvent_events_wereable) {
    return m.events_were
}
// GetExplanation gets the explanation property value. The explanation property
// returns a *string when successful
func (m *AuditLogEvent) GetExplanation()(*string) {
    return m.explanation
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuditLogEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_document_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentId(val)
        }
        return nil
    }
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val)
        }
        return nil
    }
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["active_was"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveWas(val)
        }
        return nil
    }
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val)
        }
        return nil
    }
    res["actor_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorId(val)
        }
        return nil
    }
    res["actor_location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditLogEvent_actor_locationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorLocation(val.(AuditLogEvent_actor_locationable))
        }
        return nil
    }
    res["blocked_user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockedUser(val)
        }
        return nil
    }
    res["business"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusiness(val)
        }
        return nil
    }
    res["business_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessId(val)
        }
        return nil
    }
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditLogEvent_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditLogEvent_configable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditLogEvent_configable)
                }
            }
            m.SetConfig(res)
        }
        return nil
    }
    res["config_was"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditLogEvent_config_wasFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditLogEvent_config_wasable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditLogEvent_config_wasable)
                }
            }
            m.SetConfigWas(res)
        }
        return nil
    }
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
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditLogEvent_dataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(AuditLogEvent_dataable))
        }
        return nil
    }
    res["deploy_key_fingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployKeyFingerprint(val)
        }
        return nil
    }
    res["emoji"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmoji(val)
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditLogEvent_eventsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditLogEvent_eventsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditLogEvent_eventsable)
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["events_were"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditLogEvent_events_wereFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditLogEvent_events_wereable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditLogEvent_events_wereable)
                }
            }
            m.SetEventsWere(res)
        }
        return nil
    }
    res["explanation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExplanation(val)
        }
        return nil
    }
    res["fingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprint(val)
        }
        return nil
    }
    res["hook_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHookId(val)
        }
        return nil
    }
    res["limited_availability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimitedAvailability(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["old_user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldUser(val)
        }
        return nil
    }
    res["openssh_public_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpensshPublicKey(val)
        }
        return nil
    }
    res["operation_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val)
        }
        return nil
    }
    res["org"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrg(val)
        }
        return nil
    }
    res["org_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrgId(val)
        }
        return nil
    }
    res["previous_visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousVisibility(val)
        }
        return nil
    }
    res["read_only"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnly(val)
        }
        return nil
    }
    res["repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepo(val)
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val)
        }
        return nil
    }
    res["repository_public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryPublic(val)
        }
        return nil
    }
    res["target_login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetLogin(val)
        }
        return nil
    }
    res["team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val)
        }
        return nil
    }
    res["@timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    res["transport_protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransportProtocol(val)
        }
        return nil
    }
    res["transport_protocol_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransportProtocolName(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val)
        }
        return nil
    }
    res["user_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
// GetFingerprint gets the fingerprint property value. The fingerprint property
// returns a *string when successful
func (m *AuditLogEvent) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetHookId gets the hook_id property value. The hook_id property
// returns a *int32 when successful
func (m *AuditLogEvent) GetHookId()(*int32) {
    return m.hook_id
}
// GetLimitedAvailability gets the limited_availability property value. The limited_availability property
// returns a *bool when successful
func (m *AuditLogEvent) GetLimitedAvailability()(*bool) {
    return m.limited_availability
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *AuditLogEvent) GetMessage()(*string) {
    return m.message
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AuditLogEvent) GetName()(*string) {
    return m.name
}
// GetOldUser gets the old_user property value. The old_user property
// returns a *string when successful
func (m *AuditLogEvent) GetOldUser()(*string) {
    return m.old_user
}
// GetOpensshPublicKey gets the openssh_public_key property value. The openssh_public_key property
// returns a *string when successful
func (m *AuditLogEvent) GetOpensshPublicKey()(*string) {
    return m.openssh_public_key
}
// GetOperationType gets the operation_type property value. The operation_type property
// returns a *string when successful
func (m *AuditLogEvent) GetOperationType()(*string) {
    return m.operation_type
}
// GetOrg gets the org property value. The org property
// returns a *string when successful
func (m *AuditLogEvent) GetOrg()(*string) {
    return m.org
}
// GetOrgId gets the org_id property value. The org_id property
// returns a *int32 when successful
func (m *AuditLogEvent) GetOrgId()(*int32) {
    return m.org_id
}
// GetPreviousVisibility gets the previous_visibility property value. The previous_visibility property
// returns a *string when successful
func (m *AuditLogEvent) GetPreviousVisibility()(*string) {
    return m.previous_visibility
}
// GetReadOnly gets the read_only property value. The read_only property
// returns a *bool when successful
func (m *AuditLogEvent) GetReadOnly()(*bool) {
    return m.read_only
}
// GetRepo gets the repo property value. The name of the repository.
// returns a *string when successful
func (m *AuditLogEvent) GetRepo()(*string) {
    return m.repo
}
// GetRepository gets the repository property value. The name of the repository.
// returns a *string when successful
func (m *AuditLogEvent) GetRepository()(*string) {
    return m.repository
}
// GetRepositoryPublic gets the repository_public property value. The repository_public property
// returns a *bool when successful
func (m *AuditLogEvent) GetRepositoryPublic()(*bool) {
    return m.repository_public
}
// GetTargetLogin gets the target_login property value. The target_login property
// returns a *string when successful
func (m *AuditLogEvent) GetTargetLogin()(*string) {
    return m.target_login
}
// GetTeam gets the team property value. The team property
// returns a *string when successful
func (m *AuditLogEvent) GetTeam()(*string) {
    return m.team
}
// GetTimestamp gets the @timestamp property value. The time the audit log event occurred, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time).
// returns a *int32 when successful
func (m *AuditLogEvent) GetTimestamp()(*int32) {
    return m.timestamp
}
// GetTransportProtocol gets the transport_protocol property value. The type of protocol (for example, HTTP or SSH) used to transfer Git data.
// returns a *int32 when successful
func (m *AuditLogEvent) GetTransportProtocol()(*int32) {
    return m.transport_protocol
}
// GetTransportProtocolName gets the transport_protocol_name property value. A human readable name for the protocol (for example, HTTP or SSH) used to transfer Git data.
// returns a *string when successful
func (m *AuditLogEvent) GetTransportProtocolName()(*string) {
    return m.transport_protocol_name
}
// GetUser gets the user property value. The user that was affected by the action performed (if available).
// returns a *string when successful
func (m *AuditLogEvent) GetUser()(*string) {
    return m.user
}
// GetUserId gets the user_id property value. The user_id property
// returns a *int32 when successful
func (m *AuditLogEvent) GetUserId()(*int32) {
    return m.user_id
}
// GetVisibility gets the visibility property value. The repository visibility, for example `public` or `private`.
// returns a *string when successful
func (m *AuditLogEvent) GetVisibility()(*string) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *AuditLogEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("active_was", m.GetActiveWas())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("actor_id", m.GetActorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("actor_location", m.GetActorLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blocked_user", m.GetBlockedUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("business", m.GetBusiness())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("business_id", m.GetBusinessId())
        if err != nil {
            return err
        }
    }
    if m.GetConfig() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfig()))
        for i, v := range m.GetConfig() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("config", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigWas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigWas()))
        for i, v := range m.GetConfigWas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("config_was", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deploy_key_fingerprint", m.GetDeployKeyFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emoji", m.GetEmoji())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEventsWere() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEventsWere()))
        for i, v := range m.GetEventsWere() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("events_were", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("explanation", m.GetExplanation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fingerprint", m.GetFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("hook_id", m.GetHookId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("limited_availability", m.GetLimitedAvailability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("old_user", m.GetOldUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("openssh_public_key", m.GetOpensshPublicKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operation_type", m.GetOperationType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("org", m.GetOrg())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("org_id", m.GetOrgId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previous_visibility", m.GetPreviousVisibility())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("read_only", m.GetReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repo", m.GetRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("repository_public", m.GetRepositoryPublic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_login", m.GetTargetLogin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("@timestamp", m.GetTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("transport_protocol", m.GetTransportProtocol())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("transport_protocol_name", m.GetTransportProtocolName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("user_id", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("_document_id", m.GetDocumentId())
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
// SetAction sets the action property value. The name of the action that was performed, for example `user.login` or `repo.create`.
func (m *AuditLogEvent) SetAction(value *string)() {
    m.action = value
}
// SetActive sets the active property value. The active property
func (m *AuditLogEvent) SetActive(value *bool)() {
    m.active = value
}
// SetActiveWas sets the active_was property value. The active_was property
func (m *AuditLogEvent) SetActiveWas(value *bool)() {
    m.active_was = value
}
// SetActor sets the actor property value. The actor who performed the action.
func (m *AuditLogEvent) SetActor(value *string)() {
    m.actor = value
}
// SetActorId sets the actor_id property value. The id of the actor who performed the action.
func (m *AuditLogEvent) SetActorId(value *int32)() {
    m.actor_id = value
}
// SetActorLocation sets the actor_location property value. The actor_location property
func (m *AuditLogEvent) SetActorLocation(value AuditLogEvent_actor_locationable)() {
    m.actor_location = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditLogEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlockedUser sets the blocked_user property value. The username of the account being blocked.
func (m *AuditLogEvent) SetBlockedUser(value *string)() {
    m.blocked_user = value
}
// SetBusiness sets the business property value. The business property
func (m *AuditLogEvent) SetBusiness(value *string)() {
    m.business = value
}
// SetBusinessId sets the business_id property value. The business_id property
func (m *AuditLogEvent) SetBusinessId(value *int32)() {
    m.business_id = value
}
// SetConfig sets the config property value. The config property
func (m *AuditLogEvent) SetConfig(value []AuditLogEvent_configable)() {
    m.config = value
}
// SetConfigWas sets the config_was property value. The config_was property
func (m *AuditLogEvent) SetConfigWas(value []AuditLogEvent_config_wasable)() {
    m.config_was = value
}
// SetContentType sets the content_type property value. The content_type property
func (m *AuditLogEvent) SetContentType(value *string)() {
    m.content_type = value
}
// SetCreatedAt sets the created_at property value. The time the audit log event was recorded, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time).
func (m *AuditLogEvent) SetCreatedAt(value *int32)() {
    m.created_at = value
}
// SetData sets the data property value. The data property
func (m *AuditLogEvent) SetData(value AuditLogEvent_dataable)() {
    m.data = value
}
// SetDeployKeyFingerprint sets the deploy_key_fingerprint property value. The deploy_key_fingerprint property
func (m *AuditLogEvent) SetDeployKeyFingerprint(value *string)() {
    m.deploy_key_fingerprint = value
}
// SetDocumentId sets the _document_id property value. A unique identifier for an audit event.
func (m *AuditLogEvent) SetDocumentId(value *string)() {
    m._document_id = value
}
// SetEmoji sets the emoji property value. The emoji property
func (m *AuditLogEvent) SetEmoji(value *string)() {
    m.emoji = value
}
// SetEvents sets the events property value. The events property
func (m *AuditLogEvent) SetEvents(value []AuditLogEvent_eventsable)() {
    m.events = value
}
// SetEventsWere sets the events_were property value. The events_were property
func (m *AuditLogEvent) SetEventsWere(value []AuditLogEvent_events_wereable)() {
    m.events_were = value
}
// SetExplanation sets the explanation property value. The explanation property
func (m *AuditLogEvent) SetExplanation(value *string)() {
    m.explanation = value
}
// SetFingerprint sets the fingerprint property value. The fingerprint property
func (m *AuditLogEvent) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetHookId sets the hook_id property value. The hook_id property
func (m *AuditLogEvent) SetHookId(value *int32)() {
    m.hook_id = value
}
// SetLimitedAvailability sets the limited_availability property value. The limited_availability property
func (m *AuditLogEvent) SetLimitedAvailability(value *bool)() {
    m.limited_availability = value
}
// SetMessage sets the message property value. The message property
func (m *AuditLogEvent) SetMessage(value *string)() {
    m.message = value
}
// SetName sets the name property value. The name property
func (m *AuditLogEvent) SetName(value *string)() {
    m.name = value
}
// SetOldUser sets the old_user property value. The old_user property
func (m *AuditLogEvent) SetOldUser(value *string)() {
    m.old_user = value
}
// SetOpensshPublicKey sets the openssh_public_key property value. The openssh_public_key property
func (m *AuditLogEvent) SetOpensshPublicKey(value *string)() {
    m.openssh_public_key = value
}
// SetOperationType sets the operation_type property value. The operation_type property
func (m *AuditLogEvent) SetOperationType(value *string)() {
    m.operation_type = value
}
// SetOrg sets the org property value. The org property
func (m *AuditLogEvent) SetOrg(value *string)() {
    m.org = value
}
// SetOrgId sets the org_id property value. The org_id property
func (m *AuditLogEvent) SetOrgId(value *int32)() {
    m.org_id = value
}
// SetPreviousVisibility sets the previous_visibility property value. The previous_visibility property
func (m *AuditLogEvent) SetPreviousVisibility(value *string)() {
    m.previous_visibility = value
}
// SetReadOnly sets the read_only property value. The read_only property
func (m *AuditLogEvent) SetReadOnly(value *bool)() {
    m.read_only = value
}
// SetRepo sets the repo property value. The name of the repository.
func (m *AuditLogEvent) SetRepo(value *string)() {
    m.repo = value
}
// SetRepository sets the repository property value. The name of the repository.
func (m *AuditLogEvent) SetRepository(value *string)() {
    m.repository = value
}
// SetRepositoryPublic sets the repository_public property value. The repository_public property
func (m *AuditLogEvent) SetRepositoryPublic(value *bool)() {
    m.repository_public = value
}
// SetTargetLogin sets the target_login property value. The target_login property
func (m *AuditLogEvent) SetTargetLogin(value *string)() {
    m.target_login = value
}
// SetTeam sets the team property value. The team property
func (m *AuditLogEvent) SetTeam(value *string)() {
    m.team = value
}
// SetTimestamp sets the @timestamp property value. The time the audit log event occurred, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time).
func (m *AuditLogEvent) SetTimestamp(value *int32)() {
    m.timestamp = value
}
// SetTransportProtocol sets the transport_protocol property value. The type of protocol (for example, HTTP or SSH) used to transfer Git data.
func (m *AuditLogEvent) SetTransportProtocol(value *int32)() {
    m.transport_protocol = value
}
// SetTransportProtocolName sets the transport_protocol_name property value. A human readable name for the protocol (for example, HTTP or SSH) used to transfer Git data.
func (m *AuditLogEvent) SetTransportProtocolName(value *string)() {
    m.transport_protocol_name = value
}
// SetUser sets the user property value. The user that was affected by the action performed (if available).
func (m *AuditLogEvent) SetUser(value *string)() {
    m.user = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *AuditLogEvent) SetUserId(value *int32)() {
    m.user_id = value
}
// SetVisibility sets the visibility property value. The repository visibility, for example `public` or `private`.
func (m *AuditLogEvent) SetVisibility(value *string)() {
    m.visibility = value
}
type AuditLogEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetActive()(*bool)
    GetActiveWas()(*bool)
    GetActor()(*string)
    GetActorId()(*int32)
    GetActorLocation()(AuditLogEvent_actor_locationable)
    GetBlockedUser()(*string)
    GetBusiness()(*string)
    GetBusinessId()(*int32)
    GetConfig()([]AuditLogEvent_configable)
    GetConfigWas()([]AuditLogEvent_config_wasable)
    GetContentType()(*string)
    GetCreatedAt()(*int32)
    GetData()(AuditLogEvent_dataable)
    GetDeployKeyFingerprint()(*string)
    GetDocumentId()(*string)
    GetEmoji()(*string)
    GetEvents()([]AuditLogEvent_eventsable)
    GetEventsWere()([]AuditLogEvent_events_wereable)
    GetExplanation()(*string)
    GetFingerprint()(*string)
    GetHookId()(*int32)
    GetLimitedAvailability()(*bool)
    GetMessage()(*string)
    GetName()(*string)
    GetOldUser()(*string)
    GetOpensshPublicKey()(*string)
    GetOperationType()(*string)
    GetOrg()(*string)
    GetOrgId()(*int32)
    GetPreviousVisibility()(*string)
    GetReadOnly()(*bool)
    GetRepo()(*string)
    GetRepository()(*string)
    GetRepositoryPublic()(*bool)
    GetTargetLogin()(*string)
    GetTeam()(*string)
    GetTimestamp()(*int32)
    GetTransportProtocol()(*int32)
    GetTransportProtocolName()(*string)
    GetUser()(*string)
    GetUserId()(*int32)
    GetVisibility()(*string)
    SetAction(value *string)()
    SetActive(value *bool)()
    SetActiveWas(value *bool)()
    SetActor(value *string)()
    SetActorId(value *int32)()
    SetActorLocation(value AuditLogEvent_actor_locationable)()
    SetBlockedUser(value *string)()
    SetBusiness(value *string)()
    SetBusinessId(value *int32)()
    SetConfig(value []AuditLogEvent_configable)()
    SetConfigWas(value []AuditLogEvent_config_wasable)()
    SetContentType(value *string)()
    SetCreatedAt(value *int32)()
    SetData(value AuditLogEvent_dataable)()
    SetDeployKeyFingerprint(value *string)()
    SetDocumentId(value *string)()
    SetEmoji(value *string)()
    SetEvents(value []AuditLogEvent_eventsable)()
    SetEventsWere(value []AuditLogEvent_events_wereable)()
    SetExplanation(value *string)()
    SetFingerprint(value *string)()
    SetHookId(value *int32)()
    SetLimitedAvailability(value *bool)()
    SetMessage(value *string)()
    SetName(value *string)()
    SetOldUser(value *string)()
    SetOpensshPublicKey(value *string)()
    SetOperationType(value *string)()
    SetOrg(value *string)()
    SetOrgId(value *int32)()
    SetPreviousVisibility(value *string)()
    SetReadOnly(value *bool)()
    SetRepo(value *string)()
    SetRepository(value *string)()
    SetRepositoryPublic(value *bool)()
    SetTargetLogin(value *string)()
    SetTeam(value *string)()
    SetTimestamp(value *int32)()
    SetTransportProtocol(value *int32)()
    SetTransportProtocolName(value *string)()
    SetUser(value *string)()
    SetUserId(value *int32)()
    SetVisibility(value *string)()
}
