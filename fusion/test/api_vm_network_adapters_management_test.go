/*


Testing VMNetworkAdaptersManagementAPIService

*/

package fusion

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_VMNetworkAdaptersManagementAPIService(t *testing.T) {

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)

	t.Run("Test VMNetworkAdaptersManagementAPIService CreateNICDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.VMNetworkAdaptersManagementAPI.CreateNICDevice(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMNetworkAdaptersManagementAPIService DeleteNICDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var index string

		httpRes, err := apiClient.VMNetworkAdaptersManagementAPI.DeleteNICDevice(context.Background(), id, index).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMNetworkAdaptersManagementAPIService GetAllNICDevices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.VMNetworkAdaptersManagementAPI.GetAllNICDevices(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMNetworkAdaptersManagementAPIService GetIPAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.VMNetworkAdaptersManagementAPI.GetIPAddress(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMNetworkAdaptersManagementAPIService GetNicInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.VMNetworkAdaptersManagementAPI.GetNicInfo(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMNetworkAdaptersManagementAPIService UpdateNICDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var index string

		resp, httpRes, err := apiClient.VMNetworkAdaptersManagementAPI.UpdateNICDevice(context.Background(), id, index).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
