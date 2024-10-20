/*


Testing HostNetworksManagementAPIService

*/

package fusion

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_HostNetworksManagementAPIService(t *testing.T) {

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)

	t.Run("Test HostNetworksManagementAPIService CreateNetwork", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HostNetworksManagementAPI.CreateNetwork(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostNetworksManagementAPIService DeletePortforward", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vmnet string
		var protocol string
		var port int32

		httpRes, err := apiClient.HostNetworksManagementAPI.DeletePortforward(context.Background(), vmnet, protocol, port).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostNetworksManagementAPIService GetAllNetworks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HostNetworksManagementAPI.GetAllNetworks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostNetworksManagementAPIService GetMACToIPs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vmnet string

		resp, httpRes, err := apiClient.HostNetworksManagementAPI.GetMACToIPs(context.Background(), vmnet).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostNetworksManagementAPIService GetPortforwards", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vmnet string

		resp, httpRes, err := apiClient.HostNetworksManagementAPI.GetPortforwards(context.Background(), vmnet).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostNetworksManagementAPIService UpdateMacToIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vmnet string
		var mac string

		resp, httpRes, err := apiClient.HostNetworksManagementAPI.UpdateMacToIP(context.Background(), vmnet, mac).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostNetworksManagementAPIService UpdatePortforward", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vmnet string
		var protocol string
		var port int32

		resp, httpRes, err := apiClient.HostNetworksManagementAPI.UpdatePortforward(context.Background(), vmnet, protocol, port).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
