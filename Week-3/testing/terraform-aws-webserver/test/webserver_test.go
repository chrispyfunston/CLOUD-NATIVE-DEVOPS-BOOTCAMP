package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformWebserverExample(t *testing.T) {

	//The values to pass into the Terraform CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		//The path to where the example Terraform code is located
		TerraformDir: "../examples/webserver",

		//Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"region":     "us-west-2",
			"servername": "testwebserver",
		},
	})

	//Run a terraform init and apply with Terraform options.
	terraform.InitAndApply(t, terraformOptions)

	//Run a Terraform Destroy at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	//Get the output of the example code
	publicIp := terraform.output(t, terraformOptions, "public_ip")

	url := fmt.Sprintf("http://%s:8080", publicIp)

	http_helper.HttpGetWithRetry(t, url, nil, 200, "I made a Terraform Module!", 30, 5*time.Second)
}
