package main

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

type sendmailCmd struct {
	Cmd string
}

var _ mailSender = (*sendmailCmd)(nil)

func (c *sendmailCmd) Close() error {
	return nil
}

func (c *sendmailCmd) SendMail(ctx context.Context, from string, to []string, data io.Reader) error {
	params := []string{"-i"}
	if from != "" {
		params = append(params, "-f", from)
	}
	params = append(params, to...)

	shParams := append([]string{"-c", c.Cmd + ` "$@"`, "-"}, params...)
	cmd := exec.CommandContext(ctx, "sh", shParams...)
	cmd.Stdin = data
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sendmail command failed: %v", err)
	}

	return nil
}
