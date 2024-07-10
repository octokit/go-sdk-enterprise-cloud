package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A human-readable name for a security group.
    displayName *string
    // A unique identifier for the resource as defined by the provisioning client.
    externalId *string
    // The group members.
    members []GroupResponse_membersable
    // The URIs that are used to indicate the namespaces of the SCIM schemas.
    schemas []GroupResponse_schemas
}
// NewGroupResponse instantiates a new GroupResponse and sets the default values.
func NewGroupResponse()(*GroupResponse) {
    m := &GroupResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GroupResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A human-readable name for a security group.
// returns a *string when successful
func (m *GroupResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. A unique identifier for the resource as defined by the provisioning client.
// returns a *string when successful
func (m *GroupResponse) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupResponse_membersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupResponse_membersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupResponse_membersable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseGroupResponse_schemas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupResponse_schemas, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*GroupResponse_schemas))
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. The group members.
// returns a []GroupResponse_membersable when successful
func (m *GroupResponse) GetMembers()([]GroupResponse_membersable) {
    return m.members
}
// GetSchemas gets the schemas property value. The URIs that are used to indicate the namespaces of the SCIM schemas.
// returns a []GroupResponse_schemas when successful
func (m *GroupResponse) GetSchemas()([]GroupResponse_schemas) {
    return m.schemas
}
// Serialize serializes information the current object
func (m *GroupResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        err := writer.WriteCollectionOfStringValues("schemas", SerializeGroupResponse_schemas(m.GetSchemas()))
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
func (m *GroupResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A human-readable name for a security group.
func (m *GroupResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. A unique identifier for the resource as defined by the provisioning client.
func (m *GroupResponse) SetExternalId(value *string)() {
    m.externalId = value
}
// SetMembers sets the members property value. The group members.
func (m *GroupResponse) SetMembers(value []GroupResponse_membersable)() {
    m.members = value
}
// SetSchemas sets the schemas property value. The URIs that are used to indicate the namespaces of the SCIM schemas.
func (m *GroupResponse) SetSchemas(value []GroupResponse_schemas)() {
    m.schemas = value
}
type GroupResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetMembers()([]GroupResponse_membersable)
    GetSchemas()([]GroupResponse_schemas)
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetMembers(value []GroupResponse_membersable)()
    SetSchemas(value []GroupResponse_schemas)()
}
