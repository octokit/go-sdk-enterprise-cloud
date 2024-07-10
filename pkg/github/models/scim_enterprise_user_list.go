package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScimEnterpriseUserList struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of objects per page
    itemsPerPage *int32
    // Information about each provisioned account.
    resources []ScimEnterpriseUserResponseable
    // The URIs that are used to indicate the namespaces of the list SCIM schemas.
    schemas []ScimEnterpriseUserList_schemas
    // A starting index for the returned page
    startIndex *int32
    // Number of results found
    totalResults *int32
}
// NewScimEnterpriseUserList instantiates a new ScimEnterpriseUserList and sets the default values.
func NewScimEnterpriseUserList()(*ScimEnterpriseUserList) {
    m := &ScimEnterpriseUserList{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimEnterpriseUserListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimEnterpriseUserListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimEnterpriseUserList(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimEnterpriseUserList) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimEnterpriseUserList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["itemsPerPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemsPerPage(val)
        }
        return nil
    }
    res["Resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScimEnterpriseUserResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimEnterpriseUserResponseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScimEnterpriseUserResponseable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseScimEnterpriseUserList_schemas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimEnterpriseUserList_schemas, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*ScimEnterpriseUserList_schemas))
                }
            }
            m.SetSchemas(res)
        }
        return nil
    }
    res["startIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartIndex(val)
        }
        return nil
    }
    res["totalResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalResults(val)
        }
        return nil
    }
    return res
}
// GetItemsPerPage gets the itemsPerPage property value. Number of objects per page
// returns a *int32 when successful
func (m *ScimEnterpriseUserList) GetItemsPerPage()(*int32) {
    return m.itemsPerPage
}
// GetResources gets the Resources property value. Information about each provisioned account.
// returns a []ScimEnterpriseUserResponseable when successful
func (m *ScimEnterpriseUserList) GetResources()([]ScimEnterpriseUserResponseable) {
    return m.resources
}
// GetSchemas gets the schemas property value. The URIs that are used to indicate the namespaces of the list SCIM schemas.
// returns a []ScimEnterpriseUserList_schemas when successful
func (m *ScimEnterpriseUserList) GetSchemas()([]ScimEnterpriseUserList_schemas) {
    return m.schemas
}
// GetStartIndex gets the startIndex property value. A starting index for the returned page
// returns a *int32 when successful
func (m *ScimEnterpriseUserList) GetStartIndex()(*int32) {
    return m.startIndex
}
// GetTotalResults gets the totalResults property value. Number of results found
// returns a *int32 when successful
func (m *ScimEnterpriseUserList) GetTotalResults()(*int32) {
    return m.totalResults
}
// Serialize serializes information the current object
func (m *ScimEnterpriseUserList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("itemsPerPage", m.GetItemsPerPage())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("Resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchemas() != nil {
        err := writer.WriteCollectionOfStringValues("schemas", SerializeScimEnterpriseUserList_schemas(m.GetSchemas()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("startIndex", m.GetStartIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalResults", m.GetTotalResults())
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
func (m *ScimEnterpriseUserList) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItemsPerPage sets the itemsPerPage property value. Number of objects per page
func (m *ScimEnterpriseUserList) SetItemsPerPage(value *int32)() {
    m.itemsPerPage = value
}
// SetResources sets the Resources property value. Information about each provisioned account.
func (m *ScimEnterpriseUserList) SetResources(value []ScimEnterpriseUserResponseable)() {
    m.resources = value
}
// SetSchemas sets the schemas property value. The URIs that are used to indicate the namespaces of the list SCIM schemas.
func (m *ScimEnterpriseUserList) SetSchemas(value []ScimEnterpriseUserList_schemas)() {
    m.schemas = value
}
// SetStartIndex sets the startIndex property value. A starting index for the returned page
func (m *ScimEnterpriseUserList) SetStartIndex(value *int32)() {
    m.startIndex = value
}
// SetTotalResults sets the totalResults property value. Number of results found
func (m *ScimEnterpriseUserList) SetTotalResults(value *int32)() {
    m.totalResults = value
}
type ScimEnterpriseUserListable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemsPerPage()(*int32)
    GetResources()([]ScimEnterpriseUserResponseable)
    GetSchemas()([]ScimEnterpriseUserList_schemas)
    GetStartIndex()(*int32)
    GetTotalResults()(*int32)
    SetItemsPerPage(value *int32)()
    SetResources(value []ScimEnterpriseUserResponseable)()
    SetSchemas(value []ScimEnterpriseUserList_schemas)()
    SetStartIndex(value *int32)()
    SetTotalResults(value *int32)()
}
