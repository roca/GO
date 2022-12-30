package notify

import "os/exec"

var command = exec.Command

func (n *Notify) Send() error {
	notifyCmdName := "powershell.exe"

	notifyCmd, err := exec.LookPath(notifyCmdName)
	if err != nil {
		return err
	}

	psScript := fmt.Sprintf(`Add-Type -AssemblyName System.Windows.Forms
	  $notify = New-Object System.Windows.Forms.NotifyIcon
	  $notify.BalloonTipIcon = %q
	  $notify.BalloonTipTitle = %q
	  $notify.BalloonTipText = %q
	  $notify.Visible = $true
	  $notify.ShowBalloonTip(10000)`,
	  n.severity,n.title,n.message
	)

	args := []string{}

	args = append(args, psScript)

	notifyCommand := command(notifyCmd, args...)
	return notifyCommand.Run()
}
