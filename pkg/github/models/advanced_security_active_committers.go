package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AdvancedSecurityActiveCommitters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total number of GitHub Advanced Security licences required if all repositories were to enable GitHub Advanced Security
    maximum_advanced_security_committers *int32
    // The total number of GitHub Advanced Security licences purchased
    purchased_advanced_security_committers *int32
    // The repositories property
    repositories []AdvancedSecurityActiveCommittersRepositoryable
    // The total_advanced_security_committers property
    total_advanced_security_committers *int32
    // The total_count property
    total_count *int32
}
// NewAdvancedSecurityActiveCommitters instantiates a new AdvancedSecurityActiveCommitters and sets the default values.
func NewAdvancedSecurityActiveCommitters()(*AdvancedSecurityActiveCommitters) {
    m := &AdvancedSecurityActiveCommitters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAdvancedSecurityActiveCommittersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAdvancedSecurityActiveCommittersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdvancedSecurityActiveCommitters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AdvancedSecurityActiveCommitters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AdvancedSecurityActiveCommitters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maximum_advanced_security_committers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAdvancedSecurityCommitters(val)
        }
        return nil
    }
    res["purchased_advanced_security_committers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurchasedAdvancedSecurityCommitters(val)
        }
        return nil
    }
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdvancedSecurityActiveCommittersRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdvancedSecurityActiveCommittersRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AdvancedSecurityActiveCommittersRepositoryable)
                }
            }
            m.SetRepositories(res)
        }
        return nil
    }
    res["total_advanced_security_committers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAdvancedSecurityCommitters(val)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetMaximumAdvancedSecurityCommitters gets the maximum_advanced_security_committers property value. The total number of GitHub Advanced Security licences required if all repositories were to enable GitHub Advanced Security
// returns a *int32 when successful
func (m *AdvancedSecurityActiveCommitters) GetMaximumAdvancedSecurityCommitters()(*int32) {
    return m.maximum_advanced_security_committers
}
// GetPurchasedAdvancedSecurityCommitters gets the purchased_advanced_security_committers property value. The total number of GitHub Advanced Security licences purchased
// returns a *int32 when successful
func (m *AdvancedSecurityActiveCommitters) GetPurchasedAdvancedSecurityCommitters()(*int32) {
    return m.purchased_advanced_security_committers
}
// GetRepositories gets the repositories property value. The repositories property
// returns a []AdvancedSecurityActiveCommittersRepositoryable when successful
func (m *AdvancedSecurityActiveCommitters) GetRepositories()([]AdvancedSecurityActiveCommittersRepositoryable) {
    return m.repositories
}
// GetTotalAdvancedSecurityCommitters gets the total_advanced_security_committers property value. The total_advanced_security_committers property
// returns a *int32 when successful
func (m *AdvancedSecurityActiveCommitters) GetTotalAdvancedSecurityCommitters()(*int32) {
    return m.total_advanced_security_committers
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *AdvancedSecurityActiveCommitters) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *AdvancedSecurityActiveCommitters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maximum_advanced_security_committers", m.GetMaximumAdvancedSecurityCommitters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("purchased_advanced_security_committers", m.GetPurchasedAdvancedSecurityCommitters())
        if err != nil {
            return err
        }
    }
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_advanced_security_committers", m.GetTotalAdvancedSecurityCommitters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *AdvancedSecurityActiveCommitters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaximumAdvancedSecurityCommitters sets the maximum_advanced_security_committers property value. The total number of GitHub Advanced Security licences required if all repositories were to enable GitHub Advanced Security
func (m *AdvancedSecurityActiveCommitters) SetMaximumAdvancedSecurityCommitters(value *int32)() {
    m.maximum_advanced_security_committers = value
}
// SetPurchasedAdvancedSecurityCommitters sets the purchased_advanced_security_committers property value. The total number of GitHub Advanced Security licences purchased
func (m *AdvancedSecurityActiveCommitters) SetPurchasedAdvancedSecurityCommitters(value *int32)() {
    m.purchased_advanced_security_committers = value
}
// SetRepositories sets the repositories property value. The repositories property
func (m *AdvancedSecurityActiveCommitters) SetRepositories(value []AdvancedSecurityActiveCommittersRepositoryable)() {
    m.repositories = value
}
// SetTotalAdvancedSecurityCommitters sets the total_advanced_security_committers property value. The total_advanced_security_committers property
func (m *AdvancedSecurityActiveCommitters) SetTotalAdvancedSecurityCommitters(value *int32)() {
    m.total_advanced_security_committers = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *AdvancedSecurityActiveCommitters) SetTotalCount(value *int32)() {
    m.total_count = value
}
type AdvancedSecurityActiveCommittersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumAdvancedSecurityCommitters()(*int32)
    GetPurchasedAdvancedSecurityCommitters()(*int32)
    GetRepositories()([]AdvancedSecurityActiveCommittersRepositoryable)
    GetTotalAdvancedSecurityCommitters()(*int32)
    GetTotalCount()(*int32)
    SetMaximumAdvancedSecurityCommitters(value *int32)()
    SetPurchasedAdvancedSecurityCommitters(value *int32)()
    SetRepositories(value []AdvancedSecurityActiveCommittersRepositoryable)()
    SetTotalAdvancedSecurityCommitters(value *int32)()
    SetTotalCount(value *int32)()
}
