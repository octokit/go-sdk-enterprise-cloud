package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AdvancedSecurityActiveCommittersRepository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The advanced_security_committers property
    advanced_security_committers *int32
    // The advanced_security_committers_breakdown property
    advanced_security_committers_breakdown []AdvancedSecurityActiveCommittersUserable
    // The name property
    name *string
}
// NewAdvancedSecurityActiveCommittersRepository instantiates a new AdvancedSecurityActiveCommittersRepository and sets the default values.
func NewAdvancedSecurityActiveCommittersRepository()(*AdvancedSecurityActiveCommittersRepository) {
    m := &AdvancedSecurityActiveCommittersRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAdvancedSecurityActiveCommittersRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAdvancedSecurityActiveCommittersRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdvancedSecurityActiveCommittersRepository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AdvancedSecurityActiveCommittersRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurityCommitters gets the advanced_security_committers property value. The advanced_security_committers property
// returns a *int32 when successful
func (m *AdvancedSecurityActiveCommittersRepository) GetAdvancedSecurityCommitters()(*int32) {
    return m.advanced_security_committers
}
// GetAdvancedSecurityCommittersBreakdown gets the advanced_security_committers_breakdown property value. The advanced_security_committers_breakdown property
// returns a []AdvancedSecurityActiveCommittersUserable when successful
func (m *AdvancedSecurityActiveCommittersRepository) GetAdvancedSecurityCommittersBreakdown()([]AdvancedSecurityActiveCommittersUserable) {
    return m.advanced_security_committers_breakdown
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AdvancedSecurityActiveCommittersRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security_committers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurityCommitters(val)
        }
        return nil
    }
    res["advanced_security_committers_breakdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdvancedSecurityActiveCommittersUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdvancedSecurityActiveCommittersUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AdvancedSecurityActiveCommittersUserable)
                }
            }
            m.SetAdvancedSecurityCommittersBreakdown(res)
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
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AdvancedSecurityActiveCommittersRepository) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *AdvancedSecurityActiveCommittersRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("advanced_security_committers", m.GetAdvancedSecurityCommitters())
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedSecurityCommittersBreakdown() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdvancedSecurityCommittersBreakdown()))
        for i, v := range m.GetAdvancedSecurityCommittersBreakdown() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("advanced_security_committers_breakdown", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdvancedSecurityActiveCommittersRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurityCommitters sets the advanced_security_committers property value. The advanced_security_committers property
func (m *AdvancedSecurityActiveCommittersRepository) SetAdvancedSecurityCommitters(value *int32)() {
    m.advanced_security_committers = value
}
// SetAdvancedSecurityCommittersBreakdown sets the advanced_security_committers_breakdown property value. The advanced_security_committers_breakdown property
func (m *AdvancedSecurityActiveCommittersRepository) SetAdvancedSecurityCommittersBreakdown(value []AdvancedSecurityActiveCommittersUserable)() {
    m.advanced_security_committers_breakdown = value
}
// SetName sets the name property value. The name property
func (m *AdvancedSecurityActiveCommittersRepository) SetName(value *string)() {
    m.name = value
}
type AdvancedSecurityActiveCommittersRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurityCommitters()(*int32)
    GetAdvancedSecurityCommittersBreakdown()([]AdvancedSecurityActiveCommittersUserable)
    GetName()(*string)
    SetAdvancedSecurityCommitters(value *int32)()
    SetAdvancedSecurityCommittersBreakdown(value []AdvancedSecurityActiveCommittersUserable)()
    SetName(value *string)()
}
