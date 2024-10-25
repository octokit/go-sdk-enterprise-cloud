package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemCode_security_and_analysisRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\code_security_and_analysis
type ItemCode_security_and_analysisRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCode_security_and_analysisRequestBuilderInternal instantiates a new ItemCode_security_and_analysisRequestBuilder and sets the default values.
func NewItemCode_security_and_analysisRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCode_security_and_analysisRequestBuilder) {
    m := &ItemCode_security_and_analysisRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/code_security_and_analysis", pathParameters),
    }
    return m
}
// NewItemCode_security_and_analysisRequestBuilder instantiates a new ItemCode_security_and_analysisRequestBuilder and sets the default values.
func NewItemCode_security_and_analysisRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCode_security_and_analysisRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCode_security_and_analysisRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Closing down notice:** The ability to fetch code security and analysis settings for an enterprise is closing down. Please use [code security configurations](https://docs.github.com/enterprise-cloud@latest//rest/code-security/configurations) instead. For more information, see the [changelog](https://github.blog/changelog/2024-09-27-upcoming-replacement-of-enterprise-code-security-enablement-ui-and-apis).Gets code security and analysis settings for the specified enterprise.The authenticated user must be an administrator of the enterprise in order to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// Deprecated: 
// returns a EnterpriseSecurityAnalysisSettingsable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/code-security-and-analysis#get-code-security-and-analysis-features-for-an-enterprise
func (m *ItemCode_security_and_analysisRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.EnterpriseSecurityAnalysisSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateEnterpriseSecurityAnalysisSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.EnterpriseSecurityAnalysisSettingsable), nil
}
// Patch > [!WARNING]> **Closing down notice:** The ability to update code security and analysis settings for an enterprise is closing down. Please use [code security configurations](https://docs.github.com/enterprise-cloud@latest//rest/code-security/configurations) instead. For more information, see the [changelog](https://github.blog/changelog/2024-09-27-upcoming-replacement-of-enterprise-code-security-enablement-ui-and-apis).Updates the settings for advanced security, Dependabot alerts, secret scanning, and push protection for new repositories in an enterprise.The authenticated user must be an administrator of the enterprise to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// Deprecated: 
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/code-security-and-analysis#update-code-security-and-analysis-features-for-an-enterprise
func (m *ItemCode_security_and_analysisRequestBuilder) Patch(ctx context.Context, body ItemCode_security_and_analysisPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation > [!WARNING]> **Closing down notice:** The ability to fetch code security and analysis settings for an enterprise is closing down. Please use [code security configurations](https://docs.github.com/enterprise-cloud@latest//rest/code-security/configurations) instead. For more information, see the [changelog](https://github.blog/changelog/2024-09-27-upcoming-replacement-of-enterprise-code-security-enablement-ui-and-apis).Gets code security and analysis settings for the specified enterprise.The authenticated user must be an administrator of the enterprise in order to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemCode_security_and_analysisRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation > [!WARNING]> **Closing down notice:** The ability to update code security and analysis settings for an enterprise is closing down. Please use [code security configurations](https://docs.github.com/enterprise-cloud@latest//rest/code-security/configurations) instead. For more information, see the [changelog](https://github.blog/changelog/2024-09-27-upcoming-replacement-of-enterprise-code-security-enablement-ui-and-apis).Updates the settings for advanced security, Dependabot alerts, secret scanning, and push protection for new repositories in an enterprise.The authenticated user must be an administrator of the enterprise to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemCode_security_and_analysisRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemCode_security_and_analysisPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemCode_security_and_analysisRequestBuilder when successful
func (m *ItemCode_security_and_analysisRequestBuilder) WithUrl(rawUrl string)(*ItemCode_security_and_analysisRequestBuilder) {
    return NewItemCode_security_and_analysisRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
