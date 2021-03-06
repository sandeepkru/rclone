// Test AzureBlob filesystem interface
package azureblob_test

import (
	"testing"

	"github.com/ncw/rclone/backend/azureblob"
	"github.com/ncw/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestAzureBlob:",
		NilObject:  (*azureblob.Object)(nil),
	})
}
