package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDownloadedGamesGetPaidStatusBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDownloadedGamesGetPaidStatusBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}
