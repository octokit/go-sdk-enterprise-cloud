package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerStatisticsPackages_ecosystems struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total number of packages in an ecosystem that have been created in the 24 hours prior to `collection_date` for a GHES installation
    daily_create_count *int32
    // The total number of packages in an ecosystem that have been deleted in the 24 hours prior to `collection_date` for a GHES installation
    daily_delete_count *int32
    // The total number of packages in an ecosystem that have been downloaded in the 24 hours prior to `collection_date` for a GHES installation
    daily_download_count *int32
    // The total number of packages in an ecosystem that have been updated in the 24 hours prior to `collection_date` for a GHES installation
    daily_update_count *int32
    // Shows if a package system is enabled, disabled, or read-only in a GHES installation
    enabled *ServerStatisticsPackages_ecosystems_enabled
    // The total number of internal packages in a package ecosystem in a GHES installation
    internal_packages_count *int32
    // The name of the package ecosystem
    name *ServerStatisticsPackages_ecosystems_name
    // The total number of organization packages in a package ecosystem in a GHES installation
    organization_packages_count *int32
    // The total number of private packages in a package ecosystem in a GHES installation
    private_packages_count *int32
    // The total number of public packages in a package ecosystem in a GHES installation
    public_packages_count *int32
    // The total number of published packages in a package ecosystem in a GHES installation
    published_packages_count *int32
    // The total number of user packages in a package ecosystem in a GHES installation
    user_packages_count *int32
}
// NewServerStatisticsPackages_ecosystems instantiates a new ServerStatisticsPackages_ecosystems and sets the default values.
func NewServerStatisticsPackages_ecosystems()(*ServerStatisticsPackages_ecosystems) {
    m := &ServerStatisticsPackages_ecosystems{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerStatisticsPackages_ecosystemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerStatisticsPackages_ecosystemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerStatisticsPackages_ecosystems(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerStatisticsPackages_ecosystems) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDailyCreateCount gets the daily_create_count property value. The total number of packages in an ecosystem that have been created in the 24 hours prior to `collection_date` for a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetDailyCreateCount()(*int32) {
    return m.daily_create_count
}
// GetDailyDeleteCount gets the daily_delete_count property value. The total number of packages in an ecosystem that have been deleted in the 24 hours prior to `collection_date` for a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetDailyDeleteCount()(*int32) {
    return m.daily_delete_count
}
// GetDailyDownloadCount gets the daily_download_count property value. The total number of packages in an ecosystem that have been downloaded in the 24 hours prior to `collection_date` for a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetDailyDownloadCount()(*int32) {
    return m.daily_download_count
}
// GetDailyUpdateCount gets the daily_update_count property value. The total number of packages in an ecosystem that have been updated in the 24 hours prior to `collection_date` for a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetDailyUpdateCount()(*int32) {
    return m.daily_update_count
}
// GetEnabled gets the enabled property value. Shows if a package system is enabled, disabled, or read-only in a GHES installation
// returns a *ServerStatisticsPackages_ecosystems_enabled when successful
func (m *ServerStatisticsPackages_ecosystems) GetEnabled()(*ServerStatisticsPackages_ecosystems_enabled) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerStatisticsPackages_ecosystems) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["daily_create_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDailyCreateCount(val)
        }
        return nil
    }
    res["daily_delete_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDailyDeleteCount(val)
        }
        return nil
    }
    res["daily_download_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDailyDownloadCount(val)
        }
        return nil
    }
    res["daily_update_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDailyUpdateCount(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServerStatisticsPackages_ecosystems_enabled)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val.(*ServerStatisticsPackages_ecosystems_enabled))
        }
        return nil
    }
    res["internal_packages_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalPackagesCount(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServerStatisticsPackages_ecosystems_name)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val.(*ServerStatisticsPackages_ecosystems_name))
        }
        return nil
    }
    res["organization_packages_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPackagesCount(val)
        }
        return nil
    }
    res["private_packages_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivatePackagesCount(val)
        }
        return nil
    }
    res["public_packages_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicPackagesCount(val)
        }
        return nil
    }
    res["published_packages_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedPackagesCount(val)
        }
        return nil
    }
    res["user_packages_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPackagesCount(val)
        }
        return nil
    }
    return res
}
// GetInternalPackagesCount gets the internal_packages_count property value. The total number of internal packages in a package ecosystem in a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetInternalPackagesCount()(*int32) {
    return m.internal_packages_count
}
// GetName gets the name property value. The name of the package ecosystem
// returns a *ServerStatisticsPackages_ecosystems_name when successful
func (m *ServerStatisticsPackages_ecosystems) GetName()(*ServerStatisticsPackages_ecosystems_name) {
    return m.name
}
// GetOrganizationPackagesCount gets the organization_packages_count property value. The total number of organization packages in a package ecosystem in a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetOrganizationPackagesCount()(*int32) {
    return m.organization_packages_count
}
// GetPrivatePackagesCount gets the private_packages_count property value. The total number of private packages in a package ecosystem in a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetPrivatePackagesCount()(*int32) {
    return m.private_packages_count
}
// GetPublicPackagesCount gets the public_packages_count property value. The total number of public packages in a package ecosystem in a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetPublicPackagesCount()(*int32) {
    return m.public_packages_count
}
// GetPublishedPackagesCount gets the published_packages_count property value. The total number of published packages in a package ecosystem in a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetPublishedPackagesCount()(*int32) {
    return m.published_packages_count
}
// GetUserPackagesCount gets the user_packages_count property value. The total number of user packages in a package ecosystem in a GHES installation
// returns a *int32 when successful
func (m *ServerStatisticsPackages_ecosystems) GetUserPackagesCount()(*int32) {
    return m.user_packages_count
}
// Serialize serializes information the current object
func (m *ServerStatisticsPackages_ecosystems) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("daily_create_count", m.GetDailyCreateCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("daily_delete_count", m.GetDailyDeleteCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("daily_download_count", m.GetDailyDownloadCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("daily_update_count", m.GetDailyUpdateCount())
        if err != nil {
            return err
        }
    }
    if m.GetEnabled() != nil {
        cast := (*m.GetEnabled()).String()
        err := writer.WriteStringValue("enabled", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("internal_packages_count", m.GetInternalPackagesCount())
        if err != nil {
            return err
        }
    }
    if m.GetName() != nil {
        cast := (*m.GetName()).String()
        err := writer.WriteStringValue("name", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("organization_packages_count", m.GetOrganizationPackagesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("private_packages_count", m.GetPrivatePackagesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("public_packages_count", m.GetPublicPackagesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("published_packages_count", m.GetPublishedPackagesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("user_packages_count", m.GetUserPackagesCount())
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
func (m *ServerStatisticsPackages_ecosystems) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDailyCreateCount sets the daily_create_count property value. The total number of packages in an ecosystem that have been created in the 24 hours prior to `collection_date` for a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetDailyCreateCount(value *int32)() {
    m.daily_create_count = value
}
// SetDailyDeleteCount sets the daily_delete_count property value. The total number of packages in an ecosystem that have been deleted in the 24 hours prior to `collection_date` for a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetDailyDeleteCount(value *int32)() {
    m.daily_delete_count = value
}
// SetDailyDownloadCount sets the daily_download_count property value. The total number of packages in an ecosystem that have been downloaded in the 24 hours prior to `collection_date` for a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetDailyDownloadCount(value *int32)() {
    m.daily_download_count = value
}
// SetDailyUpdateCount sets the daily_update_count property value. The total number of packages in an ecosystem that have been updated in the 24 hours prior to `collection_date` for a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetDailyUpdateCount(value *int32)() {
    m.daily_update_count = value
}
// SetEnabled sets the enabled property value. Shows if a package system is enabled, disabled, or read-only in a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetEnabled(value *ServerStatisticsPackages_ecosystems_enabled)() {
    m.enabled = value
}
// SetInternalPackagesCount sets the internal_packages_count property value. The total number of internal packages in a package ecosystem in a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetInternalPackagesCount(value *int32)() {
    m.internal_packages_count = value
}
// SetName sets the name property value. The name of the package ecosystem
func (m *ServerStatisticsPackages_ecosystems) SetName(value *ServerStatisticsPackages_ecosystems_name)() {
    m.name = value
}
// SetOrganizationPackagesCount sets the organization_packages_count property value. The total number of organization packages in a package ecosystem in a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetOrganizationPackagesCount(value *int32)() {
    m.organization_packages_count = value
}
// SetPrivatePackagesCount sets the private_packages_count property value. The total number of private packages in a package ecosystem in a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetPrivatePackagesCount(value *int32)() {
    m.private_packages_count = value
}
// SetPublicPackagesCount sets the public_packages_count property value. The total number of public packages in a package ecosystem in a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetPublicPackagesCount(value *int32)() {
    m.public_packages_count = value
}
// SetPublishedPackagesCount sets the published_packages_count property value. The total number of published packages in a package ecosystem in a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetPublishedPackagesCount(value *int32)() {
    m.published_packages_count = value
}
// SetUserPackagesCount sets the user_packages_count property value. The total number of user packages in a package ecosystem in a GHES installation
func (m *ServerStatisticsPackages_ecosystems) SetUserPackagesCount(value *int32)() {
    m.user_packages_count = value
}
type ServerStatisticsPackages_ecosystemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDailyCreateCount()(*int32)
    GetDailyDeleteCount()(*int32)
    GetDailyDownloadCount()(*int32)
    GetDailyUpdateCount()(*int32)
    GetEnabled()(*ServerStatisticsPackages_ecosystems_enabled)
    GetInternalPackagesCount()(*int32)
    GetName()(*ServerStatisticsPackages_ecosystems_name)
    GetOrganizationPackagesCount()(*int32)
    GetPrivatePackagesCount()(*int32)
    GetPublicPackagesCount()(*int32)
    GetPublishedPackagesCount()(*int32)
    GetUserPackagesCount()(*int32)
    SetDailyCreateCount(value *int32)()
    SetDailyDeleteCount(value *int32)()
    SetDailyDownloadCount(value *int32)()
    SetDailyUpdateCount(value *int32)()
    SetEnabled(value *ServerStatisticsPackages_ecosystems_enabled)()
    SetInternalPackagesCount(value *int32)()
    SetName(value *ServerStatisticsPackages_ecosystems_name)()
    SetOrganizationPackagesCount(value *int32)()
    SetPrivatePackagesCount(value *int32)()
    SetPublicPackagesCount(value *int32)()
    SetPublishedPackagesCount(value *int32)()
    SetUserPackagesCount(value *int32)()
}
