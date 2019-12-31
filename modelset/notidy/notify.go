package notidy

import (
	"context"
	"fmt"
	"os"

	"github.com/valyala/fastjson"

	"github.com/tsingson/goums/pkg/vtils"
)

// Notify listen
func (s *NotifyDbo) Notify(ctx context.Context) {
	conn, err := s.Acquire(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		os.Exit(1)
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, SqlListenTerminal)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error listening to chat channel:", err)
		os.Exit(1)
	}

	for {
		notification, er1 := conn.Conn().WaitForNotification(context.Background())
		if er1 != nil {
			continue
		}
		var p fastjson.Parser

		v, er2 := p.Parse(notification.Payload)
		if er2 == nil {
			fmt.Println("PID: ", notification.PID,
				"Channel: ", notification.Channel,
				"ID: ", vtils.B2S(v.GetStringBytes("id")),
				"Serial: ", vtils.B2S(v.GetStringBytes("Serial")),
				"Active: ", v.GetBool("Active"),
				"role: ", vtils.B2S(v.GetStringBytes("role")),
				"status: ", v.GetInt("status"),
				"expiration: ", vtils.B2S(v.GetStringBytes("expiration")),
				"op: ", vtils.B2S(v.GetStringBytes("tg_op")))
		}
	}
}

// Notify listen
func (s *NotifyDbo) UnifiedNotify(ctx context.Context) {
	conn, err := s.Acquire(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		os.Exit(1)
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `listen ums_notify`)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error listening to chat channel:", err)
		os.Exit(1)
	}

	for {
		notification, er1 := conn.Conn().WaitForNotification(context.Background())
		if er1 != nil {
			continue
		}
		var p fastjson.Parser

		v, er2 := p.Parse(notification.Payload)
		if er2 == nil {
			// TODO: parse
			table := vtils.B2S(v.GetStringBytes("table"))
			switch table {
			case "terminal":
			// TODO:
			case "apktype":
			// TODO
			default:

			}
			fmt.Println("PID: ", notification.PID,
				"Channel: ", notification.Channel,
				"table: ", vtils.B2S(v.GetStringBytes("table")),
				"data: ", vtils.B2S(v.GetStringBytes("data")),
			)
		}

	}
}
