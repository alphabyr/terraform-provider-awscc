// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssmcontacts_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSSMContactsContactDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SSMContacts::Contact", "awscc_ssmcontacts_contact", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSSSMContactsContactDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SSMContacts::Contact", "awscc_ssmcontacts_contact", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}