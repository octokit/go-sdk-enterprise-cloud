package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

type ItemActionsRunnerGroupsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The runner_groups property
    runner_groups []i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RunnerGroupsEnterpriseable
    // The total_count property
    total_count *float64
}
// NewItemActionsRunnerGroupsGetResponse instantiates a new ItemActionsRunnerGroupsGetResponse and sets the default values.
func NewItemActionsRunnerGroupsGetResponse()(*ItemActionsRunnerGroupsGetResponse) {
    m := &ItemActionsRunnerGroupsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRunnerGroupsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsRunnerGroupsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnerGroupsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsRunnerGroupsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsRunnerGroupsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["runner_groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateRunnerGroupsEnterpriseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RunnerGroupsEnterpriseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RunnerGroupsEnterpriseable)
                }
            }
            m.SetRunnerGroups(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
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
// GetRunnerGroups gets the runner_groups property value. The runner_groups property
// returns a []RunnerGroupsEnterpriseable when successful
func (m *ItemActionsRunnerGroupsGetResponse) GetRunnerGroups()([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RunnerGroupsEnterpriseable) {
    return m.runner_groups
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *float64 when successful
func (m *ItemActionsRunnerGroupsGetResponse) GetTotalCount()(*float64) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemActionsRunnerGroupsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRunnerGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRunnerGroups()))
        for i, v := range m.GetRunnerGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("runner_groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("total_count", m.GetTotalCount())
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
func (m *ItemActionsRunnerGroupsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRunnerGroups sets the runner_groups property value. The runner_groups property
func (m *ItemActionsRunnerGroupsGetResponse) SetRunnerGroups(value []i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RunnerGroupsEnterpriseable)() {
    m.runner_groups = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemActionsRunnerGroupsGetResponse) SetTotalCount(value *float64)() {
    m.total_count = value
}
type ItemActionsRunnerGroupsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRunnerGroups()([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RunnerGroupsEnterpriseable)
    GetTotalCount()(*float64)
    SetRunnerGroups(value []i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.RunnerGroupsEnterpriseable)()
    SetTotalCount(value *float64)()
}
