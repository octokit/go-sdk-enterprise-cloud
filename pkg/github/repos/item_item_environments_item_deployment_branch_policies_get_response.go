package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

type ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The branch_policies property
    branch_policies []i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DeploymentBranchPolicyable
    // The number of deployment branch policies for the environment.
    total_count *int32
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse instantiates a new ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse()(*ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) {
    m := &ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranchPolicies gets the branch_policies property value. The branch_policies property
// returns a []DeploymentBranchPolicyable when successful
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetBranchPolicies()([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DeploymentBranchPolicyable) {
    return m.branch_policies
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch_policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateDeploymentBranchPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DeploymentBranchPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DeploymentBranchPolicyable)
                }
            }
            m.SetBranchPolicies(res)
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
// GetTotalCount gets the total_count property value. The number of deployment branch policies for the environment.
// returns a *int32 when successful
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBranchPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBranchPolicies()))
        for i, v := range m.GetBranchPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("branch_policies", cast)
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
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranchPolicies sets the branch_policies property value. The branch_policies property
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) SetBranchPolicies(value []i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DeploymentBranchPolicyable)() {
    m.branch_policies = value
}
// SetTotalCount sets the total_count property value. The number of deployment branch policies for the environment.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranchPolicies()([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DeploymentBranchPolicyable)
    GetTotalCount()(*int32)
    SetBranchPolicies(value []i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DeploymentBranchPolicyable)()
    SetTotalCount(value *int32)()
}
