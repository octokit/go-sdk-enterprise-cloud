package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RunnerGroupsEnterprise struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allows_public_repositories property
    allows_public_repositories *bool
    // The default property
    defaultEscaped *bool
    // The hosted_runners_url property
    hosted_runners_url *string
    // The id property
    id *float64
    // The name property
    name *string
    // The identifier of a hosted compute network configuration.
    network_configuration_id *string
    // If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
    restricted_to_workflows *bool
    // The runners_url property
    runners_url *string
    // The selected_organizations_url property
    selected_organizations_url *string
    // List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
    selected_workflows []string
    // The visibility property
    visibility *string
    // If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
    workflow_restrictions_read_only *bool
}
// NewRunnerGroupsEnterprise instantiates a new RunnerGroupsEnterprise and sets the default values.
func NewRunnerGroupsEnterprise()(*RunnerGroupsEnterprise) {
    m := &RunnerGroupsEnterprise{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRunnerGroupsEnterpriseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunnerGroupsEnterpriseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunnerGroupsEnterprise(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RunnerGroupsEnterprise) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowsPublicRepositories gets the allows_public_repositories property value. The allows_public_repositories property
// returns a *bool when successful
func (m *RunnerGroupsEnterprise) GetAllowsPublicRepositories()(*bool) {
    return m.allows_public_repositories
}
// GetDefaultEscaped gets the default property value. The default property
// returns a *bool when successful
func (m *RunnerGroupsEnterprise) GetDefaultEscaped()(*bool) {
    return m.defaultEscaped
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RunnerGroupsEnterprise) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allows_public_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowsPublicRepositories(val)
        }
        return nil
    }
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEscaped(val)
        }
        return nil
    }
    res["hosted_runners_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostedRunnersUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
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
    res["network_configuration_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkConfigurationId(val)
        }
        return nil
    }
    res["restricted_to_workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictedToWorkflows(val)
        }
        return nil
    }
    res["runners_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnersUrl(val)
        }
        return nil
    }
    res["selected_organizations_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectedOrganizationsUrl(val)
        }
        return nil
    }
    res["selected_workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSelectedWorkflows(res)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    res["workflow_restrictions_read_only"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowRestrictionsReadOnly(val)
        }
        return nil
    }
    return res
}
// GetHostedRunnersUrl gets the hosted_runners_url property value. The hosted_runners_url property
// returns a *string when successful
func (m *RunnerGroupsEnterprise) GetHostedRunnersUrl()(*string) {
    return m.hosted_runners_url
}
// GetId gets the id property value. The id property
// returns a *float64 when successful
func (m *RunnerGroupsEnterprise) GetId()(*float64) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *RunnerGroupsEnterprise) GetName()(*string) {
    return m.name
}
// GetNetworkConfigurationId gets the network_configuration_id property value. The identifier of a hosted compute network configuration.
// returns a *string when successful
func (m *RunnerGroupsEnterprise) GetNetworkConfigurationId()(*string) {
    return m.network_configuration_id
}
// GetRestrictedToWorkflows gets the restricted_to_workflows property value. If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
// returns a *bool when successful
func (m *RunnerGroupsEnterprise) GetRestrictedToWorkflows()(*bool) {
    return m.restricted_to_workflows
}
// GetRunnersUrl gets the runners_url property value. The runners_url property
// returns a *string when successful
func (m *RunnerGroupsEnterprise) GetRunnersUrl()(*string) {
    return m.runners_url
}
// GetSelectedOrganizationsUrl gets the selected_organizations_url property value. The selected_organizations_url property
// returns a *string when successful
func (m *RunnerGroupsEnterprise) GetSelectedOrganizationsUrl()(*string) {
    return m.selected_organizations_url
}
// GetSelectedWorkflows gets the selected_workflows property value. List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
// returns a []string when successful
func (m *RunnerGroupsEnterprise) GetSelectedWorkflows()([]string) {
    return m.selected_workflows
}
// GetVisibility gets the visibility property value. The visibility property
// returns a *string when successful
func (m *RunnerGroupsEnterprise) GetVisibility()(*string) {
    return m.visibility
}
// GetWorkflowRestrictionsReadOnly gets the workflow_restrictions_read_only property value. If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
// returns a *bool when successful
func (m *RunnerGroupsEnterprise) GetWorkflowRestrictionsReadOnly()(*bool) {
    return m.workflow_restrictions_read_only
}
// Serialize serializes information the current object
func (m *RunnerGroupsEnterprise) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allows_public_repositories", m.GetAllowsPublicRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("default", m.GetDefaultEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hosted_runners_url", m.GetHostedRunnersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("id", m.GetId())
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
        err := writer.WriteStringValue("network_configuration_id", m.GetNetworkConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("restricted_to_workflows", m.GetRestrictedToWorkflows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("runners_url", m.GetRunnersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("selected_organizations_url", m.GetSelectedOrganizationsUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedWorkflows() != nil {
        err := writer.WriteCollectionOfStringValues("selected_workflows", m.GetSelectedWorkflows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("workflow_restrictions_read_only", m.GetWorkflowRestrictionsReadOnly())
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
func (m *RunnerGroupsEnterprise) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowsPublicRepositories sets the allows_public_repositories property value. The allows_public_repositories property
func (m *RunnerGroupsEnterprise) SetAllowsPublicRepositories(value *bool)() {
    m.allows_public_repositories = value
}
// SetDefaultEscaped sets the default property value. The default property
func (m *RunnerGroupsEnterprise) SetDefaultEscaped(value *bool)() {
    m.defaultEscaped = value
}
// SetHostedRunnersUrl sets the hosted_runners_url property value. The hosted_runners_url property
func (m *RunnerGroupsEnterprise) SetHostedRunnersUrl(value *string)() {
    m.hosted_runners_url = value
}
// SetId sets the id property value. The id property
func (m *RunnerGroupsEnterprise) SetId(value *float64)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *RunnerGroupsEnterprise) SetName(value *string)() {
    m.name = value
}
// SetNetworkConfigurationId sets the network_configuration_id property value. The identifier of a hosted compute network configuration.
func (m *RunnerGroupsEnterprise) SetNetworkConfigurationId(value *string)() {
    m.network_configuration_id = value
}
// SetRestrictedToWorkflows sets the restricted_to_workflows property value. If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
func (m *RunnerGroupsEnterprise) SetRestrictedToWorkflows(value *bool)() {
    m.restricted_to_workflows = value
}
// SetRunnersUrl sets the runners_url property value. The runners_url property
func (m *RunnerGroupsEnterprise) SetRunnersUrl(value *string)() {
    m.runners_url = value
}
// SetSelectedOrganizationsUrl sets the selected_organizations_url property value. The selected_organizations_url property
func (m *RunnerGroupsEnterprise) SetSelectedOrganizationsUrl(value *string)() {
    m.selected_organizations_url = value
}
// SetSelectedWorkflows sets the selected_workflows property value. List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
func (m *RunnerGroupsEnterprise) SetSelectedWorkflows(value []string)() {
    m.selected_workflows = value
}
// SetVisibility sets the visibility property value. The visibility property
func (m *RunnerGroupsEnterprise) SetVisibility(value *string)() {
    m.visibility = value
}
// SetWorkflowRestrictionsReadOnly sets the workflow_restrictions_read_only property value. If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
func (m *RunnerGroupsEnterprise) SetWorkflowRestrictionsReadOnly(value *bool)() {
    m.workflow_restrictions_read_only = value
}
type RunnerGroupsEnterpriseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowsPublicRepositories()(*bool)
    GetDefaultEscaped()(*bool)
    GetHostedRunnersUrl()(*string)
    GetId()(*float64)
    GetName()(*string)
    GetNetworkConfigurationId()(*string)
    GetRestrictedToWorkflows()(*bool)
    GetRunnersUrl()(*string)
    GetSelectedOrganizationsUrl()(*string)
    GetSelectedWorkflows()([]string)
    GetVisibility()(*string)
    GetWorkflowRestrictionsReadOnly()(*bool)
    SetAllowsPublicRepositories(value *bool)()
    SetDefaultEscaped(value *bool)()
    SetHostedRunnersUrl(value *string)()
    SetId(value *float64)()
    SetName(value *string)()
    SetNetworkConfigurationId(value *string)()
    SetRestrictedToWorkflows(value *bool)()
    SetRunnersUrl(value *string)()
    SetSelectedOrganizationsUrl(value *string)()
    SetSelectedWorkflows(value []string)()
    SetVisibility(value *string)()
    SetWorkflowRestrictionsReadOnly(value *bool)()
}
