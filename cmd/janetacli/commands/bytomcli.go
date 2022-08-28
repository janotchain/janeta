package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/janotchain/janeta/util"
)

// janetacli usage template
var usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletDisable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

  available with wallet enable:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletEnable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// commandError is an error used to signal different error situations in command handling.
type commandError struct {
	s         string
	userError bool
}

func (c commandError) Error() string {
	return c.s
}

func (c commandError) isUserError() bool {
	return c.userError
}

func newUserError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: true}
}

func newSystemError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: false}
}

func newSystemErrorF(format string, a ...interface{}) commandError {
	return commandError{s: fmt.Sprintf(format, a...), userError: false}
}

// Catch some of the obvious user errors from Cobra.
// We don't want to show the usage message for every error.
// The below may be to generic. Time will show.
var userErrorRegexp = regexp.MustCompile("argument|flag|shorthand")

func isUserError(err error) bool {
	if cErr, ok := err.(commandError); ok && cErr.isUserError() {
		return true
	}

	return userErrorRegexp.MatchString(err.Error())
}

// janetacliCmd is janetacli's root command.
// Every other command attached to janetacliCmd is a child command to it.
var janetacliCmd = &cobra.Command{
	Use:   "janetacli",
	Short: "janetacli is a commond line client for janeta core (a.k.a. janetad)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.SetUsageTemplate(usageTemplate)
			cmd.Usage()
		}
	},
}

// Execute adds all child commands to the root command janetacliCmd and sets flags appropriately.
func Execute() {

	AddCommands()
	AddTemplateFunc()

	if _, err := janetacliCmd.ExecuteC(); err != nil {
		os.Exit(util.ErrLocalExe)
	}
}

// AddCommands adds child commands to the root command janetacliCmd.
func AddCommands() {
	janetacliCmd.AddCommand(createAccessTokenCmd)
	janetacliCmd.AddCommand(listAccessTokenCmd)
	janetacliCmd.AddCommand(deleteAccessTokenCmd)
	janetacliCmd.AddCommand(checkAccessTokenCmd)

	janetacliCmd.AddCommand(createAccountCmd)
	janetacliCmd.AddCommand(deleteAccountCmd)
	janetacliCmd.AddCommand(listAccountsCmd)
	janetacliCmd.AddCommand(updateAccountAliasCmd)
	janetacliCmd.AddCommand(createAccountReceiverCmd)
	janetacliCmd.AddCommand(listAddressesCmd)
	janetacliCmd.AddCommand(validateAddressCmd)
	janetacliCmd.AddCommand(listPubKeysCmd)

	janetacliCmd.AddCommand(createAssetCmd)
	janetacliCmd.AddCommand(getAssetCmd)
	janetacliCmd.AddCommand(listAssetsCmd)
	janetacliCmd.AddCommand(updateAssetAliasCmd)

	janetacliCmd.AddCommand(getTransactionCmd)
	janetacliCmd.AddCommand(listTransactionsCmd)

	janetacliCmd.AddCommand(getUnconfirmedTransactionCmd)
	janetacliCmd.AddCommand(listUnconfirmedTransactionsCmd)
	janetacliCmd.AddCommand(decodeRawTransactionCmd)

	janetacliCmd.AddCommand(listUnspentOutputsCmd)
	janetacliCmd.AddCommand(listBalancesCmd)

	janetacliCmd.AddCommand(rescanWalletCmd)
	janetacliCmd.AddCommand(walletInfoCmd)

	janetacliCmd.AddCommand(buildTransactionCmd)
	janetacliCmd.AddCommand(signTransactionCmd)
	janetacliCmd.AddCommand(submitTransactionCmd)
	janetacliCmd.AddCommand(estimateTransactionGasCmd)

	janetacliCmd.AddCommand(getBlockCountCmd)
	janetacliCmd.AddCommand(getBlockHashCmd)
	janetacliCmd.AddCommand(getBlockCmd)
	janetacliCmd.AddCommand(getBlockHeaderCmd)

	janetacliCmd.AddCommand(createKeyCmd)
	janetacliCmd.AddCommand(deleteKeyCmd)
	janetacliCmd.AddCommand(listKeysCmd)
	janetacliCmd.AddCommand(updateKeyAliasCmd)
	janetacliCmd.AddCommand(resetKeyPwdCmd)
	janetacliCmd.AddCommand(checkKeyPwdCmd)

	janetacliCmd.AddCommand(signMsgCmd)
	janetacliCmd.AddCommand(verifyMsgCmd)
	janetacliCmd.AddCommand(decodeProgCmd)

	janetacliCmd.AddCommand(createTransactionFeedCmd)
	janetacliCmd.AddCommand(listTransactionFeedsCmd)
	janetacliCmd.AddCommand(deleteTransactionFeedCmd)
	janetacliCmd.AddCommand(getTransactionFeedCmd)
	janetacliCmd.AddCommand(updateTransactionFeedCmd)

	janetacliCmd.AddCommand(netInfoCmd)
	janetacliCmd.AddCommand(gasRateCmd)

	janetacliCmd.AddCommand(versionCmd)
}

// AddTemplateFunc adds usage template to the root command janetacliCmd.
func AddTemplateFunc() {
	walletEnableCmd := []string{
		createAccountCmd.Name(),
		listAccountsCmd.Name(),
		deleteAccountCmd.Name(),
		updateAccountAliasCmd.Name(),
		createAccountReceiverCmd.Name(),
		listAddressesCmd.Name(),
		validateAddressCmd.Name(),
		listPubKeysCmd.Name(),

		createAssetCmd.Name(),
		getAssetCmd.Name(),
		listAssetsCmd.Name(),
		updateAssetAliasCmd.Name(),

		createKeyCmd.Name(),
		deleteKeyCmd.Name(),
		listKeysCmd.Name(),
		resetKeyPwdCmd.Name(),
		checkKeyPwdCmd.Name(),
		signMsgCmd.Name(),

		buildTransactionCmd.Name(),
		signTransactionCmd.Name(),

		getTransactionCmd.Name(),
		listTransactionsCmd.Name(),
		listUnspentOutputsCmd.Name(),
		listBalancesCmd.Name(),

		rescanWalletCmd.Name(),
		walletInfoCmd.Name(),
	}

	cobra.AddTemplateFunc("WalletEnable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return true
			}
		}
		return false
	})

	cobra.AddTemplateFunc("WalletDisable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return false
			}
		}
		return true
	})
}
