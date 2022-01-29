package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/cosmonaut/scavenge/x/scavenge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCommitSolution() *cobra.Command {
	cmd := &cobra.Command{
		// pass a solution as the only argument
		Use:   "commit-solution [solution]",
		Short: "Broadcast message commit-solution",
		// Set the number of arguments to 1
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			solution := args[0]
			// find a hash for the solution
			solutionHash := sha256.Sum256([]byte(solution))
			// convert the solution hash to a string
			solutionHashString := hex.EncodeToString(solutionHash[:])
			// convert a scavenger address to a string
			var scavenger = clientCtx.GetFromAddress().String()

			// find the hash of the solution and scavenger address
			var solutionScavengerHash = sha256.Sum256([]byte(solution + scavenger))

			// convert hash to string
			var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])

			// Create a new message
			msg := types.NewMsgCommitSolution(clientCtx.GetFromAddress().String(), string(solutionHashString), string(solutionScavengerHashString))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// broadcast the transasction with the messsage to the blockchain
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
