package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsEnterprisePermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permissions policy that controls the actions and reusable workflows that are allowed to run.
    allowed_actions *AllowedActions
    // The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions.
    enabled_organizations *EnabledOrganizations
    // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
    selected_actions_url *string
    // The API URL to use to get or set the selected organizations that are allowed to run GitHub Actions, when `enabled_organizations` is set to `selected`.
    selected_organizations_url *string
}
// NewActionsEnterprisePermissions instantiates a new ActionsEnterprisePermissions and sets the default values.
func NewActionsEnterprisePermissions()(*ActionsEnterprisePermissions) {
    m := &ActionsEnterprisePermissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsEnterprisePermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsEnterprisePermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsEnterprisePermissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsEnterprisePermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedActions gets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
// returns a *AllowedActions when successful
func (m *ActionsEnterprisePermissions) GetAllowedActions()(*AllowedActions) {
    return m.allowed_actions
}
// GetEnabledOrganizations gets the enabled_organizations property value. The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions.
// returns a *EnabledOrganizations when successful
func (m *ActionsEnterprisePermissions) GetEnabledOrganizations()(*EnabledOrganizations) {
    return m.enabled_organizations
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsEnterprisePermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowedActions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedActions(val.(*AllowedActions))
        }
        return nil
    }
    res["enabled_organizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnabledOrganizations)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledOrganizations(val.(*EnabledOrganizations))
        }
        return nil
    }
    res["selected_actions_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectedActionsUrl(val)
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
    return res
}
// GetSelectedActionsUrl gets the selected_actions_url property value. The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
// returns a *string when successful
func (m *ActionsEnterprisePermissions) GetSelectedActionsUrl()(*string) {
    return m.selected_actions_url
}
// GetSelectedOrganizationsUrl gets the selected_organizations_url property value. The API URL to use to get or set the selected organizations that are allowed to run GitHub Actions, when `enabled_organizations` is set to `selected`.
// returns a *string when successful
func (m *ActionsEnterprisePermissions) GetSelectedOrganizationsUrl()(*string) {
    return m.selected_organizations_url
}
// Serialize serializes information the current object
func (m *ActionsEnterprisePermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedActions() != nil {
        cast := (*m.GetAllowedActions()).String()
        err := writer.WriteStringValue("allowed_actions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnabledOrganizations() != nil {
        cast := (*m.GetEnabledOrganizations()).String()
        err := writer.WriteStringValue("enabled_organizations", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("selected_actions_url", m.GetSelectedActionsUrl())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionsEnterprisePermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedActions sets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
func (m *ActionsEnterprisePermissions) SetAllowedActions(value *AllowedActions)() {
    m.allowed_actions = value
}
// SetEnabledOrganizations sets the enabled_organizations property value. The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions.
func (m *ActionsEnterprisePermissions) SetEnabledOrganizations(value *EnabledOrganizations)() {
    m.enabled_organizations = value
}
// SetSelectedActionsUrl sets the selected_actions_url property value. The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
func (m *ActionsEnterprisePermissions) SetSelectedActionsUrl(value *string)() {
    m.selected_actions_url = value
}
// SetSelectedOrganizationsUrl sets the selected_organizations_url property value. The API URL to use to get or set the selected organizations that are allowed to run GitHub Actions, when `enabled_organizations` is set to `selected`.
func (m *ActionsEnterprisePermissions) SetSelectedOrganizationsUrl(value *string)() {
    m.selected_organizations_url = value
}
type ActionsEnterprisePermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedActions()(*AllowedActions)
    GetEnabledOrganizations()(*EnabledOrganizations)
    GetSelectedActionsUrl()(*string)
    GetSelectedOrganizationsUrl()(*string)
    SetAllowedActions(value *AllowedActions)()
    SetEnabledOrganizations(value *EnabledOrganizations)()
    SetSelectedActionsUrl(value *string)()
    SetSelectedOrganizationsUrl(value *string)()
}
