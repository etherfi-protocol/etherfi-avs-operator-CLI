package eigenda

import "github.com/urfave/cli/v3"

var EigenDACmd = &cli.Command{
	Name:  "eigenda",
	Usage: "various actions related to managing eigenDA operators",
	Commands: []*cli.Command{
		EigenDARegisterCmd,
		EigenDAPrepareRegistrationCmd,
	},
}
