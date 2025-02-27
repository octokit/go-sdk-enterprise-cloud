package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\properties\schema\organizations\{org}\{custom_property_name}\promote
type ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilderInternal instantiates a new ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder) {
    m := &ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/properties/schema/organizations/{org}/{custom_property_name}/promote", pathParameters),
    }
    return m
}
// NewItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder instantiates a new ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder and sets the default values.
func NewItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilderInternal(urlParams, requestAdapter)
}
// Put > [!NOTE]> This endpoint is in public preview and is subject to change.Promotes an existing organization custom property to an enterprise.To use this endpoint, the authenticated user must be an administrator for the enterprise.
// returns a CustomPropertyable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/custom-properties#promote-a-custom-property-to-an-enterprise
func (m *ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder) Put(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CustomPropertyable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateCustomPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CustomPropertyable), nil
}
// ToPutRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Promotes an existing organization custom property to an enterprise.To use this endpoint, the authenticated user must be an administrator for the enterprise.
// returns a *RequestInformation when successful
func (m *ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder when successful
func (m *ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder) WithUrl(rawUrl string)(*ItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder) {
    return NewItemPropertiesSchemaOrganizationsItemItemPromoteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
