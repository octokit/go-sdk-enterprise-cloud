package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScimUserList sCIM User List
type ScimUserList struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The itemsPerPage property
    itemsPerPage *int32
    // The Resources property
    resources []ScimUserable
    // SCIM schema used.
    schemas []string
    // The startIndex property
    startIndex *int32
    // The totalResults property
    totalResults *int32
}
// NewScimUserList instantiates a new ScimUserList and sets the default values.
func NewScimUserList()(*ScimUserList) {
    m := &ScimUserList{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScimUserListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScimUserListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScimUserList(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScimUserList) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScimUserList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateScimUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScimUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScimUserable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["schemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetItemsPerPage gets the itemsPerPage property value. The itemsPerPage property
// returns a *int32 when successful
func (m *ScimUserList) GetItemsPerPage()(*int32) {
    return m.itemsPerPage
}
// GetResources gets the Resources property value. The Resources property
// returns a []ScimUserable when successful
func (m *ScimUserList) GetResources()([]ScimUserable) {
    return m.resources
}
// GetSchemas gets the schemas property value. SCIM schema used.
// returns a []string when successful
func (m *ScimUserList) GetSchemas()([]string) {
    return m.schemas
}
// GetStartIndex gets the startIndex property value. The startIndex property
// returns a *int32 when successful
func (m *ScimUserList) GetStartIndex()(*int32) {
    return m.startIndex
}
// GetTotalResults gets the totalResults property value. The totalResults property
// returns a *int32 when successful
func (m *ScimUserList) GetTotalResults()(*int32) {
    return m.totalResults
}
// Serialize serializes information the current object
func (m *ScimUserList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteCollectionOfStringValues("schemas", m.GetSchemas())
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
func (m *ScimUserList) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItemsPerPage sets the itemsPerPage property value. The itemsPerPage property
func (m *ScimUserList) SetItemsPerPage(value *int32)() {
    m.itemsPerPage = value
}
// SetResources sets the Resources property value. The Resources property
func (m *ScimUserList) SetResources(value []ScimUserable)() {
    m.resources = value
}
// SetSchemas sets the schemas property value. SCIM schema used.
func (m *ScimUserList) SetSchemas(value []string)() {
    m.schemas = value
}
// SetStartIndex sets the startIndex property value. The startIndex property
func (m *ScimUserList) SetStartIndex(value *int32)() {
    m.startIndex = value
}
// SetTotalResults sets the totalResults property value. The totalResults property
func (m *ScimUserList) SetTotalResults(value *int32)() {
    m.totalResults = value
}
type ScimUserListable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemsPerPage()(*int32)
    GetResources()([]ScimUserable)
    GetSchemas()([]string)
    GetStartIndex()(*int32)
    GetTotalResults()(*int32)
    SetItemsPerPage(value *int32)()
    SetResources(value []ScimUserable)()
    SetSchemas(value []string)()
    SetStartIndex(value *int32)()
    SetTotalResults(value *int32)()
}
