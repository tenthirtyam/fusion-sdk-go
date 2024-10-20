/*


Testing VMSharedFoldersManagementAPIService

*/

package fusion

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_VMSharedFoldersManagementAPIService(t *testing.T) {

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)

	t.Run("Test VMSharedFoldersManagementAPIService CreateSharedFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.VMSharedFoldersManagementAPI.CreateSharedFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMSharedFoldersManagementAPIService DeleteSharedFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var folderId string

		httpRes, err := apiClient.VMSharedFoldersManagementAPI.DeleteSharedFolder(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMSharedFoldersManagementAPIService GetAllSharedFolders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.VMSharedFoldersManagementAPI.GetAllSharedFolders(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMSharedFoldersManagementAPIService UpdataSharedFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.VMSharedFoldersManagementAPI.UpdataSharedFolder(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
