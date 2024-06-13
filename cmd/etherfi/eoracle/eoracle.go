package eoracle

import "github.com/urfave/cli/v3"

var EOracleCmd = &cli.Command{
	Name:  "eoracle",
	Usage: "various actions related to managing eOracle operators",
	Commands: []*cli.Command{
		//EOracleRegisterCmd,
		EOraclePrepareRegistrationCmd,
	},
}
