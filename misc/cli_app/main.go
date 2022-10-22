package main

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/secrets"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
	"github.com/oracle/oci-go-sdk/vault"
)

const region string = "us-phoenix-1"

func initializeConfigurationProvider() common.ConfigurationProvider {
	privateKey := os.Getenv("pem")
	tenancyOCID := os.Getenv("tenancyOCID")
	userOCID := os.Getenv("userOCID")
	fingerprint := os.Getenv("fingerprint")
	configurationProvider := common.NewRawConfigurationProvider(tenancyOCID, userOCID, region, fingerprint, string(privateKey), nil)
	return configurationProvider
}

func listDomains(configurationProvider common.ConfigurationProvider) {
	c, err := identity.NewIdentityClientWithConfigurationProvider(configurationProvider)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tenancyOCID := os.Getenv("tenancyOCID")

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: common.String(string(tenancyOCID)),
	}

	r, err := c.ListAvailabilityDomains(context.Background(), request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nList of available domains: %v\n", r.Items)

}

func ExampleGetSecret(configurationProvider common.ConfigurationProvider) {
	// Use authentication provider for vault client
	client, err := vault.NewVaultsClientWithConfigurationProvider(configurationProvider)
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := vault.GetSecretRequest{
		SecretId: common.String(os.Getenv("secret_ocid"))}

	// Send the request using the service client
	resp, err := client.GetSecret(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func ExampleGetSecretBundle(configurationProvider common.ConfigurationProvider) {
	client, err := secrets.NewSecretsClientWithConfigurationProvider(configurationProvider)
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).
	secretID := os.Getenv("secret_ocid")
	fmt.Println(secretID)

	req := secrets.GetSecretBundleRequest{
		SecretId: common.String(secretID),
		Stage:    secrets.GetSecretBundleStageLatest,
	}

	// Send the request using the service client
	resp, err := client.GetSecretBundle(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	var content string
	base64Details, ok := resp.SecretBundleContent.(secrets.Base64SecretBundleContentDetails)
	if ok {
		//log.Println("base64 content details", *base64Details.Content)
		content = *base64Details.Content
	}
	rawDecodedText, err := b64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(err)
	}

	//fmt.Println(rawDecodedText)
	fmt.Println(string(rawDecodedText))
	fmt.Printf("type of content is %T", content)
}

func main() {
	fmt.Println("Simple Application")
	configurationProvider := initializeConfigurationProvider()
	//listDomains(configurationProvider) // just a good test that your config is right
	spacer()
	ExampleGetSecret(configurationProvider)
	spacer()
	ExampleGetSecretBundle(configurationProvider)
}

func spacer() {
	fmt.Println()
	fmt.Println()
}
