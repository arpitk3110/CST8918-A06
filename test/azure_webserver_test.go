package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"labelPrefix":    "arpitpatel",
			"admin_username": "azureadmin",
			"region":         "eastus2",
		},
		MaxRetries:         3,
		TimeBetweenRetries: 15 * time.Second,
	}

	defer func() {
		time.Sleep(60 * time.Second)
		terraform.Destroy(t, terraformOptions)
	}()

	terraform.InitAndApply(t, terraformOptions)
}