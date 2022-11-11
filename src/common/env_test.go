package common

import (
	"fmt"
	"os"
	"testing"
)

func TestGetOSLookupEnv(t *testing.T) {
	t.Run("collection of 3 string", func(t *testing.T) {
		os.Setenv("PORT", "3000")
		os.Setenv("PROJECT", "ticketing")
		os.Setenv("ENV", "dev")
		mocks := []string{"PORT", "PROJECT", "ENV"}
		got, _ := getOSLookupEnv(mocks)
		wants := map[string]string{"PORT": "3000", "PROJECT": "ticketing", "ENV": "dev"}
		fmt.Println(got)
		if got["PORT"] != wants["PORT"] {
			t.Errorf("got %q want %q given, %q", got, wants, "3000")
		}
	})
}
