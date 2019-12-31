package notidy

/**
func TestTerminalDbo_Notify(t *testing.T) {
	ctx := context.Background()
	terminalDbo, err := NewTerminalDbo(ctx, cfg, log)
	assert.NoError(t, err)
	in := &flatums.TerminalProfileT {
		SerialNumber:vtils.RandString(16),
		ActiveCode:vtils.RandString(16),
	}

	var userID int64
	userID, err = terminalDbo.Insert(ctx, in )

	assert.NoError(t, err)
	if err == nil {
		fmt.Println("id ",  userID)
	}

	terminalDbo.Notify(ctx)

	var c int64
	c, err = terminalDbo.Update(ctx, userID, true, 2, 2)

	assert.NoError(t, err)
	if err == nil {
		fmt.Println(c)
	}

	time.Sleep(5 * time.Second)
}

func TestTerminalDbo_UmsNotify(t *testing.T) {
	ctx := context.Background()
	terminalDbo, err := NewTerminalDbo(ctx, cfg, log)
	assert.NoError(t, err)
	terminalDbo.UmsNotify(ctx)
}

*/
