package component_test

import (
	"os"
	"testing"

	"github.com/utkarsh-josh/hdfc/testutil"
)

var testObj *testutil.TestUtil

func TestMain(m *testing.M) {
	testObj = testutil.InitTestInfra()
	status := m.Run()
	os.Exit(status)
}
