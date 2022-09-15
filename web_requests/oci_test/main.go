package main

import (
	"context"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
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

	fmt.Printf("List of available domains: %v", r.Items)

}

func main() {
	fmt.Println("Simple Application")
	configurationProvider := initializeConfigurationProvider()
	listDomains(configurationProvider)
}
