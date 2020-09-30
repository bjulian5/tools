package imports

import (
	"os"
	"testing"

	"github.com/bjulian5/tools/internal/testenv"
)

func TestMain(m *testing.M) {
	testenv.ExitIfSmallMachine()
	os.Exit(m.Run())
}
