package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Group struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A human-readable name for a security group.
    displayName *string
    // A unique identifier for the resource as defined by the provisioning client.
    externalId *string
    // The group members.
    members []Group_membersable
    // The URIs that are used to indicate the namespaces of the SCIM schemas.
    schemas []Group_schemas
}
// NewGroup instantiates a new Group and sets the default values.
func NewGroup()(*Group) {
    m := &Group{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Group) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A human-readable name for a security group.
// returns a *string when successful
func (m *Group) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. A unique identifier for the resource as defined by the provisioning client.
// returns a *string when successful
func (m *Group) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Group) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateGroup_membersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Group_membersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Group_membersable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseGroup_schemas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Group_schemas, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*Group_schemas))
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. The group members.
// returns a []Group_membersable when successful
func (m *Group) GetMembers()([]Group_membersable) {
    return m.members
}
// GetSchemas gets the schemas property value. The URIs that are used to indicate the namespaces of the SCIM schemas.
// returns a []Group_schemas when successful
func (m *Group) GetSchemas()([]Group_schemas) {
    return m.schemas
}
// Serialize serializes information the current object
func (m *Group) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteCollectionOfStringValues("schemas", SerializeGroup_schemas(m.GetSchemas()))
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
func (m *Group) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A human-readable name for a security group.
func (m *Group) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. A unique identifier for the resource as defined by the provisioning client.
func (m *Group) SetExternalId(value *string)() {
    m.externalId = value
}
// SetMembers sets the members property value. The group members.
func (m *Group) SetMembers(value []Group_membersable)() {
    m.members = value
}
// SetSchemas sets the schemas property value. The URIs that are used to indicate the namespaces of the SCIM schemas.
func (m *Group) SetSchemas(value []Group_schemas)() {
    m.schemas = value
}
type Groupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetMembers()([]Group_membersable)
    GetSchemas()([]Group_schemas)
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetMembers(value []Group_membersable)()
    SetSchemas(value []Group_schemas)()
}
