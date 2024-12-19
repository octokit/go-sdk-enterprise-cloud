package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PushRuleBypassRequest a bypass request made by a user asking to be exempted from a push rule in this repository.
type PushRuleBypassRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date and time the bypass request was created.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Data describing the push rules that are being requested to be bypassed.
    data []PushRuleBypassRequest_dataable
    // The date and time the bypass request will expire.
    expires_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The URL to view the bypass request in a browser.
    html_url *string
    // The unique identifier of the bypass request.
    id *int32
    // The number uniquely identifying the bypass request within its repository.
    number *int32
    // The organization associated with the repository the bypass request is for.
    organization PushRuleBypassRequest_organizationable
    // The repository the bypass request is for.
    repository PushRuleBypassRequest_repositoryable
    // The type of request.
    request_type *string
    // The user who requested the bypass.
    requester PushRuleBypassRequest_requesterable
    // The comment the requester provided when creating the bypass request.
    requester_comment *string
    // The unique identifier for the request type of the bypass request. For example, a commit SHA.
    resource_identifier *string
    // The responses to the bypass request.
    responses []BypassResponseable
    // The status of the bypass request.
    status *PushRuleBypassRequest_status
    // The url property
    url *string
}
// NewPushRuleBypassRequest instantiates a new PushRuleBypassRequest and sets the default values.
func NewPushRuleBypassRequest()(*PushRuleBypassRequest) {
    m := &PushRuleBypassRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePushRuleBypassRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePushRuleBypassRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPushRuleBypassRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PushRuleBypassRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The date and time the bypass request was created.
// returns a *Time when successful
func (m *PushRuleBypassRequest) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetData gets the data property value. Data describing the push rules that are being requested to be bypassed.
// returns a []PushRuleBypassRequest_dataable when successful
func (m *PushRuleBypassRequest) GetData()([]PushRuleBypassRequest_dataable) {
    return m.data
}
// GetExpiresAt gets the expires_at property value. The date and time the bypass request will expire.
// returns a *Time when successful
func (m *PushRuleBypassRequest) GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PushRuleBypassRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePushRuleBypassRequest_dataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PushRuleBypassRequest_dataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PushRuleBypassRequest_dataable)
                }
            }
            m.SetData(res)
        }
        return nil
    }
    res["expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePushRuleBypassRequest_organizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(PushRuleBypassRequest_organizationable))
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePushRuleBypassRequest_repositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(PushRuleBypassRequest_repositoryable))
        }
        return nil
    }
    res["request_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val)
        }
        return nil
    }
    res["requester"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePushRuleBypassRequest_requesterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequester(val.(PushRuleBypassRequest_requesterable))
        }
        return nil
    }
    res["requester_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequesterComment(val)
        }
        return nil
    }
    res["resource_identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceIdentifier(val)
        }
        return nil
    }
    res["responses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBypassResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BypassResponseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BypassResponseable)
                }
            }
            m.SetResponses(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePushRuleBypassRequest_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*PushRuleBypassRequest_status))
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
// GetHtmlUrl gets the html_url property value. The URL to view the bypass request in a browser.
// returns a *string when successful
func (m *PushRuleBypassRequest) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The unique identifier of the bypass request.
// returns a *int32 when successful
func (m *PushRuleBypassRequest) GetId()(*int32) {
    return m.id
}
// GetNumber gets the number property value. The number uniquely identifying the bypass request within its repository.
// returns a *int32 when successful
func (m *PushRuleBypassRequest) GetNumber()(*int32) {
    return m.number
}
// GetOrganization gets the organization property value. The organization associated with the repository the bypass request is for.
// returns a PushRuleBypassRequest_organizationable when successful
func (m *PushRuleBypassRequest) GetOrganization()(PushRuleBypassRequest_organizationable) {
    return m.organization
}
// GetRepository gets the repository property value. The repository the bypass request is for.
// returns a PushRuleBypassRequest_repositoryable when successful
func (m *PushRuleBypassRequest) GetRepository()(PushRuleBypassRequest_repositoryable) {
    return m.repository
}
// GetRequester gets the requester property value. The user who requested the bypass.
// returns a PushRuleBypassRequest_requesterable when successful
func (m *PushRuleBypassRequest) GetRequester()(PushRuleBypassRequest_requesterable) {
    return m.requester
}
// GetRequesterComment gets the requester_comment property value. The comment the requester provided when creating the bypass request.
// returns a *string when successful
func (m *PushRuleBypassRequest) GetRequesterComment()(*string) {
    return m.requester_comment
}
// GetRequestType gets the request_type property value. The type of request.
// returns a *string when successful
func (m *PushRuleBypassRequest) GetRequestType()(*string) {
    return m.request_type
}
// GetResourceIdentifier gets the resource_identifier property value. The unique identifier for the request type of the bypass request. For example, a commit SHA.
// returns a *string when successful
func (m *PushRuleBypassRequest) GetResourceIdentifier()(*string) {
    return m.resource_identifier
}
// GetResponses gets the responses property value. The responses to the bypass request.
// returns a []BypassResponseable when successful
func (m *PushRuleBypassRequest) GetResponses()([]BypassResponseable) {
    return m.responses
}
// GetStatus gets the status property value. The status of the bypass request.
// returns a *PushRuleBypassRequest_status when successful
func (m *PushRuleBypassRequest) GetStatus()(*PushRuleBypassRequest_status) {
    return m.status
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PushRuleBypassRequest) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PushRuleBypassRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetData()))
        for i, v := range m.GetData() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("data", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expires_at", m.GetExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requester", m.GetRequester())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("requester_comment", m.GetRequesterComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("request_type", m.GetRequestType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resource_identifier", m.GetResourceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetResponses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResponses()))
        for i, v := range m.GetResponses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("responses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PushRuleBypassRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The date and time the bypass request was created.
func (m *PushRuleBypassRequest) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetData sets the data property value. Data describing the push rules that are being requested to be bypassed.
func (m *PushRuleBypassRequest) SetData(value []PushRuleBypassRequest_dataable)() {
    m.data = value
}
// SetExpiresAt sets the expires_at property value. The date and time the bypass request will expire.
func (m *PushRuleBypassRequest) SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expires_at = value
}
// SetHtmlUrl sets the html_url property value. The URL to view the bypass request in a browser.
func (m *PushRuleBypassRequest) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The unique identifier of the bypass request.
func (m *PushRuleBypassRequest) SetId(value *int32)() {
    m.id = value
}
// SetNumber sets the number property value. The number uniquely identifying the bypass request within its repository.
func (m *PushRuleBypassRequest) SetNumber(value *int32)() {
    m.number = value
}
// SetOrganization sets the organization property value. The organization associated with the repository the bypass request is for.
func (m *PushRuleBypassRequest) SetOrganization(value PushRuleBypassRequest_organizationable)() {
    m.organization = value
}
// SetRepository sets the repository property value. The repository the bypass request is for.
func (m *PushRuleBypassRequest) SetRepository(value PushRuleBypassRequest_repositoryable)() {
    m.repository = value
}
// SetRequester sets the requester property value. The user who requested the bypass.
func (m *PushRuleBypassRequest) SetRequester(value PushRuleBypassRequest_requesterable)() {
    m.requester = value
}
// SetRequesterComment sets the requester_comment property value. The comment the requester provided when creating the bypass request.
func (m *PushRuleBypassRequest) SetRequesterComment(value *string)() {
    m.requester_comment = value
}
// SetRequestType sets the request_type property value. The type of request.
func (m *PushRuleBypassRequest) SetRequestType(value *string)() {
    m.request_type = value
}
// SetResourceIdentifier sets the resource_identifier property value. The unique identifier for the request type of the bypass request. For example, a commit SHA.
func (m *PushRuleBypassRequest) SetResourceIdentifier(value *string)() {
    m.resource_identifier = value
}
// SetResponses sets the responses property value. The responses to the bypass request.
func (m *PushRuleBypassRequest) SetResponses(value []BypassResponseable)() {
    m.responses = value
}
// SetStatus sets the status property value. The status of the bypass request.
func (m *PushRuleBypassRequest) SetStatus(value *PushRuleBypassRequest_status)() {
    m.status = value
}
// SetUrl sets the url property value. The url property
func (m *PushRuleBypassRequest) SetUrl(value *string)() {
    m.url = value
}
type PushRuleBypassRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetData()([]PushRuleBypassRequest_dataable)
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetNumber()(*int32)
    GetOrganization()(PushRuleBypassRequest_organizationable)
    GetRepository()(PushRuleBypassRequest_repositoryable)
    GetRequester()(PushRuleBypassRequest_requesterable)
    GetRequesterComment()(*string)
    GetRequestType()(*string)
    GetResourceIdentifier()(*string)
    GetResponses()([]BypassResponseable)
    GetStatus()(*PushRuleBypassRequest_status)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetData(value []PushRuleBypassRequest_dataable)()
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetNumber(value *int32)()
    SetOrganization(value PushRuleBypassRequest_organizationable)()
    SetRepository(value PushRuleBypassRequest_repositoryable)()
    SetRequester(value PushRuleBypassRequest_requesterable)()
    SetRequesterComment(value *string)()
    SetRequestType(value *string)()
    SetResourceIdentifier(value *string)()
    SetResponses(value []BypassResponseable)()
    SetStatus(value *PushRuleBypassRequest_status)()
    SetUrl(value *string)()
}
