package viper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTMDBRepo(t *testing.T) {
	f, clo := newFile(t)
	defer clo()

	r, err := NewRepo(f.Name())
	require.Nil(t, err)

	apiToken, err := r.GetAPIToken()
	require.Nil(t, err)
	require.Empty(t, apiToken)

	want := "api_token"
	err = r.SetAPIToken(want)
	require.Nil(t, err)

	apiToken, err = r.GetAPIToken()
	require.Nil(t, err)
	require.Equal(t, want, apiToken)
}
