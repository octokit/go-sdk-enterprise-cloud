package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\deployment_protection_rule
type ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Deployment_protection_rulePostRequestBody composed type wrapper for classes i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
type Deployment_protection_rulePostRequestBody struct {
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
    deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
    deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
    reviewCustomGatesCommentRequired i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
    reviewCustomGatesStateRequired i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
}
// NewDeployment_protection_rulePostRequestBody instantiates a new Deployment_protection_rulePostRequestBody and sets the default values.
func NewDeployment_protection_rulePostRequestBody()(*Deployment_protection_rulePostRequestBody) {
    m := &Deployment_protection_rulePostRequestBody{
    }
    return m
}
// CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeployment_protection_rulePostRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable); ok {
                result.SetReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable); ok {
                result.SetReviewCustomGatesStateRequired(cast)
            }
        }
    }
    return result, nil
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
// returns a ReviewCustomGatesCommentRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
// returns a ReviewCustomGatesStateRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Deployment_protection_rulePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Deployment_protection_rulePostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
// returns a ReviewCustomGatesCommentRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesCommentRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable) {
    return m.reviewCustomGatesCommentRequired
}
// GetReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
// returns a ReviewCustomGatesStateRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesStateRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable) {
    return m.reviewCustomGatesStateRequired
}
// Serialize serializes information the current object
func (m *Deployment_protection_rulePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired = value
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired = value
}
// SetReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesCommentRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable)() {
    m.reviewCustomGatesCommentRequired = value
}
// SetReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesStateRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable)() {
    m.reviewCustomGatesStateRequired = value
}
type Deployment_protection_rulePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable)
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable)
    GetReviewCustomGatesCommentRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable)
    GetReviewCustomGatesStateRequired()(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable)
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable)()
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable)()
    SetReviewCustomGatesCommentRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesCommentRequiredable)()
    SetReviewCustomGatesStateRequired(value i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ReviewCustomGatesStateRequiredable)()
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal instantiates a new ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    m := &ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/deployment_protection_rule", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder instantiates a new ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/enterprise-cloud@latest//actions/deployment/targeting-different-environments/using-environments-for-deployment)."> [!NOTE]> GitHub Apps can only review their own custom deployment protection rules. To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/actions/workflow-runs#review-custom-deployment-protection-rules-for-a-workflow-run
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) Post(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/enterprise-cloud@latest//actions/deployment/targeting-different-environments/using-environments-for-deployment)."> [!NOTE]> GitHub Apps can only review their own custom deployment protection rules. To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) ToPostRequestInformation(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder when successful
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
