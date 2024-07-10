package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationCustomRepositoryRoleUpdateSchema struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The system role from which this role inherits permissions.
    base_role *OrganizationCustomRepositoryRoleUpdateSchema_base_role
    // A short description about who this role is for or what permissions it grants.
    description *string
    // The name of the custom role.
    name *string
    // A list of additional permissions included in this role.
    permissions []string
}
// NewOrganizationCustomRepositoryRoleUpdateSchema instantiates a new OrganizationCustomRepositoryRoleUpdateSchema and sets the default values.
func NewOrganizationCustomRepositoryRoleUpdateSchema()(*OrganizationCustomRepositoryRoleUpdateSchema) {
    m := &OrganizationCustomRepositoryRoleUpdateSchema{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationCustomRepositoryRoleUpdateSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationCustomRepositoryRoleUpdateSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationCustomRepositoryRoleUpdateSchema(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationCustomRepositoryRoleUpdateSchema) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBaseRole gets the base_role property value. The system role from which this role inherits permissions.
// returns a *OrganizationCustomRepositoryRoleUpdateSchema_base_role when successful
func (m *OrganizationCustomRepositoryRoleUpdateSchema) GetBaseRole()(*OrganizationCustomRepositoryRoleUpdateSchema_base_role) {
    return m.base_role
}
// GetDescription gets the description property value. A short description about who this role is for or what permissions it grants.
// returns a *string when successful
func (m *OrganizationCustomRepositoryRoleUpdateSchema) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationCustomRepositoryRoleUpdateSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["base_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationCustomRepositoryRoleUpdateSchema_base_role)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseRole(val.(*OrganizationCustomRepositoryRoleUpdateSchema_base_role))
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
    return res
}
// GetName gets the name property value. The name of the custom role.
// returns a *string when successful
func (m *OrganizationCustomRepositoryRoleUpdateSchema) GetName()(*string) {
    return m.name
}
// GetPermissions gets the permissions property value. A list of additional permissions included in this role.
// returns a []string when successful
func (m *OrganizationCustomRepositoryRoleUpdateSchema) GetPermissions()([]string) {
    return m.permissions
}
// Serialize serializes information the current object
func (m *OrganizationCustomRepositoryRoleUpdateSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBaseRole() != nil {
        cast := (*m.GetBaseRole()).String()
        err := writer.WriteStringValue("base_role", &cast)
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
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationCustomRepositoryRoleUpdateSchema) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBaseRole sets the base_role property value. The system role from which this role inherits permissions.
func (m *OrganizationCustomRepositoryRoleUpdateSchema) SetBaseRole(value *OrganizationCustomRepositoryRoleUpdateSchema_base_role)() {
    m.base_role = value
}
// SetDescription sets the description property value. A short description about who this role is for or what permissions it grants.
func (m *OrganizationCustomRepositoryRoleUpdateSchema) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of the custom role.
func (m *OrganizationCustomRepositoryRoleUpdateSchema) SetName(value *string)() {
    m.name = value
}
// SetPermissions sets the permissions property value. A list of additional permissions included in this role.
func (m *OrganizationCustomRepositoryRoleUpdateSchema) SetPermissions(value []string)() {
    m.permissions = value
}
type OrganizationCustomRepositoryRoleUpdateSchemaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseRole()(*OrganizationCustomRepositoryRoleUpdateSchema_base_role)
    GetDescription()(*string)
    GetName()(*string)
    GetPermissions()([]string)
    SetBaseRole(value *OrganizationCustomRepositoryRoleUpdateSchema_base_role)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetPermissions(value []string)()
}
