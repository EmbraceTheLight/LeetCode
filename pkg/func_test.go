package pkg_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"lc/pkg"
	"testing"
)

func TestCreateSlices(t *testing.T) {
	intSli := pkg.CreateSlice[int]()
	require.IsTypef(t, intSli, []int{}, "int slice should be created")
	fmt.Println(intSli)

	float64Sli := pkg.CreateSlice[float64]()
	require.IsTypef(t, float64Sli, []float64{}, "float64 slice should be created")
	fmt.Println(float64Sli)

	string32Sli := pkg.CreateSlice[string]()
	require.IsTypef(t, string32Sli, []string{}, "string slice should be created")
	fmt.Println(string32Sli)
}
