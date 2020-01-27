// +build all azuresqlfirewall

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlFirewallRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	//rgName := tc.resourceGroupName
	sqlServerName := "t-sqlfwrule-test-srv" + helpers.RandomString(10)

	randomName := helpers.RandomString(10)
	sqlFirewallRuleName := "t-fwrule-dev-" + randomName

	// Create the SqlFirewallRule object and expect the Reconcile to be created
	sqlFirewallRuleInstance := &azurev1alpha1.AzureSqlFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFirewallRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlFirewallRuleSpec{
			ResourceGroup:  "t-rg-fake-srv" + helpers.RandomString(10),
			Server:         sqlServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	err := tc.k8sClient.Create(ctx, sqlFirewallRuleInstance)
	assert.Equal(nil, err, "create sqlfirewallrule in k8s")

	sqlFirewallRuleNamespacedName := types.NamespacedName{Name: sqlFirewallRuleName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return helpers.HasFinalizer(sqlFirewallRuleInstance, azureSQLFirewallRuleFinalizerName)
	}, tc.timeout, tc.retry, "wait for firewallrule to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return strings.Contains(sqlFirewallRuleInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for firewallrule to have rg not found error")

	err = tc.k8sClient.Delete(ctx, sqlFirewallRuleInstance)
	assert.Equal(nil, err, "delete sqlfirewallrule in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for firewallrule to be gone from k8s")

}