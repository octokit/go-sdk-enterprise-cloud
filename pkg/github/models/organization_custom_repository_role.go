package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationCustomRepositoryRole custom repository roles created by organization owners
type OrganizationCustomRepositoryRole struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The system role from which this role inherits permissions.
    base_role *OrganizationCustomRepositoryRole_base_role
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A short description about who this role is for or what permissions it grants.
    description *string
    // The unique identifier of the custom role.
    id *int32
    // The name of the custom role.
    name *string
    // A GitHub user.
    organization SimpleUserable
    // A list of additional permissions included in this role.
    permissions []string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewOrganizationCustomRepositoryRole instantiates a new OrganizationCustomRepositoryRole and sets the default values.
func NewOrganizationCustomRepositoryRole()(*OrganizationCustomRepositoryRole) {
    m := &OrganizationCustomRepositoryRole{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationCustomRepositoryRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationCustomRepositoryRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationCustomRepositoryRole(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationCustomRepositoryRole) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBaseRole gets the base_role property value. The system role from which this role inherits permissions.
// returns a *OrganizationCustomRepositoryRole_base_role when successful
func (m *OrganizationCustomRepositoryRole) GetBaseRole()(*OrganizationCustomRepositoryRole_base_role) {
    return m.base_role
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *OrganizationCustomRepositoryRole) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. A short description about who this role is for or what permissions it grants.
// returns a *string when successful
func (m *OrganizationCustomRepositoryRole) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationCustomRepositoryRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["base_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationCustomRepositoryRole_base_role)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseRole(val.(*OrganizationCustomRepositoryRole_base_role))
        }
        return nil
    }
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(SimpleUserable))
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier of the custom role.
// returns a *int32 when successful
func (m *OrganizationCustomRepositoryRole) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the custom role.
// returns a *string when successful
func (m *OrganizationCustomRepositoryRole) GetName()(*string) {
    return m.name
}
// GetOrganization gets the organization property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *OrganizationCustomRepositoryRole) GetOrganization()(SimpleUserable) {
    return m.organization
}
// GetPermissions gets the permissions property value. A list of additional permissions included in this role.
// returns a []string when successful
func (m *OrganizationCustomRepositoryRole) GetPermissions()([]string) {
    return m.permissions
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *OrganizationCustomRepositoryRole) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *OrganizationCustomRepositoryRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBaseRole() != nil {
        cast := (*m.GetBaseRole()).String()
        err := writer.WriteStringValue("base_role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetPermissions() != nil {
        err := writer.WriteCollectionOfStringValues("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
func (m *OrganizationCustomRepositoryRole) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBaseRole sets the base_role property value. The system role from which this role inherits permissions.
func (m *OrganizationCustomRepositoryRole) SetBaseRole(value *OrganizationCustomRepositoryRole_base_role)() {
    m.base_role = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *OrganizationCustomRepositoryRole) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. A short description about who this role is for or what permissions it grants.
func (m *OrganizationCustomRepositoryRole) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The unique identifier of the custom role.
func (m *OrganizationCustomRepositoryRole) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the custom role.
func (m *OrganizationCustomRepositoryRole) SetName(value *string)() {
    m.name = value
}
// SetOrganization sets the organization property value. A GitHub user.
func (m *OrganizationCustomRepositoryRole) SetOrganization(value SimpleUserable)() {
    m.organization = value
}
// SetPermissions sets the permissions property value. A list of additional permissions included in this role.
func (m *OrganizationCustomRepositoryRole) SetPermissions(value []string)() {
    m.permissions = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *OrganizationCustomRepositoryRole) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type OrganizationCustomRepositoryRoleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseRole()(*OrganizationCustomRepositoryRole_base_role)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetOrganization()(SimpleUserable)
    GetPermissions()([]string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetBaseRole(value *OrganizationCustomRepositoryRole_base_role)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetOrganization(value SimpleUserable)()
    SetPermissions(value []string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
