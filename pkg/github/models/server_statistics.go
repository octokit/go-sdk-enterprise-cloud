package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatistics struct {
    // Actions metrics that are included in the Server Statistics payload/export from GHES
    actions_stats ServerStatisticsActionsable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The collection_date property
    collection_date *string
    // The dormant_users property
    dormant_users ServerStatistics_dormant_usersable
    // The ghe_stats property
    ghe_stats ServerStatistics_ghe_statsable
    // The ghes_version property
    ghes_version *string
    // The github_connect property
    github_connect ServerStatistics_github_connectable
    // The host_name property
    host_name *string
    // Packages metrics that are included in the Server Statistics payload/export from GHES
    packages_stats ServerStatisticsPackagesable
    // The schema_version property
    schema_version *string
    // The server_id property
    server_id *string
}
// NewServerStatistics instantiates a new ServerStatistics and sets the default values.
func NewServerStatistics()(*ServerStatistics) {
    m := &ServerStatistics{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatistics(), nil
}
// GetActionsStats gets the actions_stats property value. Actions metrics that are included in the Server Statistics payload/export from GHES
// returns a ServerStatisticsActionsable when successful
func (m *ServerStatistics) GetActionsStats()(ServerStatisticsActionsable) {
    return m.actions_stats
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatistics) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCollectionDate gets the collection_date property value. The collection_date property
// returns a *string when successful
func (m *ServerStatistics) GetCollectionDate()(*string) {
    return m.collection_date
}
// GetDormantUsers gets the dormant_users property value. The dormant_users property
// returns a ServerStatistics_dormant_usersable when successful
func (m *ServerStatistics) GetDormantUsers()(ServerStatistics_dormant_usersable) {
    return m.dormant_users
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions_stats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatisticsActionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionsStats(val.(ServerStatisticsActionsable))
        }
        return nil
    }
    res["collection_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectionDate(val)
        }
        return nil
    }
    res["dormant_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_dormant_usersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDormantUsers(val.(ServerStatistics_dormant_usersable))
        }
        return nil
    }
    res["ghe_stats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_ghe_statsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGheStats(val.(ServerStatistics_ghe_statsable))
        }
        return nil
    }
    res["ghes_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGhesVersion(val)
        }
        return nil
    }
    res["github_connect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatistics_github_connectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubConnect(val.(ServerStatistics_github_connectable))
        }
        return nil
    }
    res["host_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    res["packages_stats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerStatisticsPackagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackagesStats(val.(ServerStatisticsPackagesable))
        }
        return nil
    }
    res["schema_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaVersion(val)
        }
        return nil
    }
    res["server_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerId(val)
        }
        return nil
    }
    return res
}
// GetGheStats gets the ghe_stats property value. The ghe_stats property
// returns a ServerStatistics_ghe_statsable when successful
func (m *ServerStatistics) GetGheStats()(ServerStatistics_ghe_statsable) {
    return m.ghe_stats
}
// GetGhesVersion gets the ghes_version property value. The ghes_version property
// returns a *string when successful
func (m *ServerStatistics) GetGhesVersion()(*string) {
    return m.ghes_version
}
// GetGithubConnect gets the github_connect property value. The github_connect property
// returns a ServerStatistics_github_connectable when successful
func (m *ServerStatistics) GetGithubConnect()(ServerStatistics_github_connectable) {
    return m.github_connect
}
// GetHostName gets the host_name property value. The host_name property
// returns a *string when successful
func (m *ServerStatistics) GetHostName()(*string) {
    return m.host_name
}
// GetPackagesStats gets the packages_stats property value. Packages metrics that are included in the Server Statistics payload/export from GHES
// returns a ServerStatisticsPackagesable when successful
func (m *ServerStatistics) GetPackagesStats()(ServerStatisticsPackagesable) {
    return m.packages_stats
}
// GetSchemaVersion gets the schema_version property value. The schema_version property
// returns a *string when successful
func (m *ServerStatistics) GetSchemaVersion()(*string) {
    return m.schema_version
}
// GetServerId gets the server_id property value. The server_id property
// returns a *string when successful
func (m *ServerStatistics) GetServerId()(*string) {
    return m.server_id
}
// Serialize serializes information the current object
func (m *ServerStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actions_stats", m.GetActionsStats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("collection_date", m.GetCollectionDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dormant_users", m.GetDormantUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ghes_version", m.GetGhesVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ghe_stats", m.GetGheStats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("github_connect", m.GetGithubConnect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("host_name", m.GetHostName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("packages_stats", m.GetPackagesStats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schema_version", m.GetSchemaVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("server_id", m.GetServerId())
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
// SetActionsStats sets the actions_stats property value. Actions metrics that are included in the Server Statistics payload/export from GHES
func (m *ServerStatistics) SetActionsStats(value ServerStatisticsActionsable)() {
    m.actions_stats = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServerStatistics) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCollectionDate sets the collection_date property value. The collection_date property
func (m *ServerStatistics) SetCollectionDate(value *string)() {
    m.collection_date = value
}
// SetDormantUsers sets the dormant_users property value. The dormant_users property
func (m *ServerStatistics) SetDormantUsers(value ServerStatistics_dormant_usersable)() {
    m.dormant_users = value
}
// SetGheStats sets the ghe_stats property value. The ghe_stats property
func (m *ServerStatistics) SetGheStats(value ServerStatistics_ghe_statsable)() {
    m.ghe_stats = value
}
// SetGhesVersion sets the ghes_version property value. The ghes_version property
func (m *ServerStatistics) SetGhesVersion(value *string)() {
    m.ghes_version = value
}
// SetGithubConnect sets the github_connect property value. The github_connect property
func (m *ServerStatistics) SetGithubConnect(value ServerStatistics_github_connectable)() {
    m.github_connect = value
}
// SetHostName sets the host_name property value. The host_name property
func (m *ServerStatistics) SetHostName(value *string)() {
    m.host_name = value
}
// SetPackagesStats sets the packages_stats property value. Packages metrics that are included in the Server Statistics payload/export from GHES
func (m *ServerStatistics) SetPackagesStats(value ServerStatisticsPackagesable)() {
    m.packages_stats = value
}
// SetSchemaVersion sets the schema_version property value. The schema_version property
func (m *ServerStatistics) SetSchemaVersion(value *string)() {
    m.schema_version = value
}
// SetServerId sets the server_id property value. The server_id property
func (m *ServerStatistics) SetServerId(value *string)() {
    m.server_id = value
}
type ServerStatisticsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionsStats()(ServerStatisticsActionsable)
    GetCollectionDate()(*string)
    GetDormantUsers()(ServerStatistics_dormant_usersable)
    GetGheStats()(ServerStatistics_ghe_statsable)
    GetGhesVersion()(*string)
    GetGithubConnect()(ServerStatistics_github_connectable)
    GetHostName()(*string)
    GetPackagesStats()(ServerStatisticsPackagesable)
    GetSchemaVersion()(*string)
    GetServerId()(*string)
    SetActionsStats(value ServerStatisticsActionsable)()
    SetCollectionDate(value *string)()
    SetDormantUsers(value ServerStatistics_dormant_usersable)()
    SetGheStats(value ServerStatistics_ghe_statsable)()
    SetGhesVersion(value *string)()
    SetGithubConnect(value ServerStatistics_github_connectable)()
    SetHostName(value *string)()
    SetPackagesStats(value ServerStatisticsPackagesable)()
    SetSchemaVersion(value *string)()
    SetServerId(value *string)()
}
