//go:build (!containers && ignore) || !disable_notification
// +build !containers,ignore !disable_notification

package app

import "github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/notify"

func send_notification(msg string) {
	n := notify.New("Pomodoro", msg, notify.SeverityNormal)
	n.Send()
}
