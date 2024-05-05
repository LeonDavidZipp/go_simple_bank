package db

import (
	"context"
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/LeonDavidZipp/go_simple_bank/util"
)

func testTransferTx(t *testing.T) {
	store := NewStore(testDB)
}