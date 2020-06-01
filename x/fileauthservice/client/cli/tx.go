package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	utils "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/linbeier/authsys/x/fileauthservice/internal/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	fileauthserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	fileauthserviceTxCmd.AddCommand(flags.PostCommands(
		// TODO: Add tx based commands
		// GetCmd<Action>(cdc)
		GetCmdSetFiles(cdc),
		GetCmdTransFiles(cdc),
	)...)

	return fileauthserviceTxCmd
}

// Example:
//
// GetCmd<Action> is the CLI command for doing <Action>
// func GetCmd<Action>(cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "/* Describe your action cmd */",
// 		Short: "/* Provide a short description on the cmd */",
// 		Args:  cobra.ExactArgs(2), // Does your request require arguments
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

// 			msg := types.NewMsg<Action>(/* Action params */)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}
// }

func GetCmdSetFiles(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "setfile [filenames] [filehash] [authority]",
		Short: "set file to have your account in store! authority should be like rwt or r or rw",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			types.Timestart = time.Now()
			fmt.Printf("time is: %s\n", types.Timestart.String())
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgSetFileAuth(args[0], args[1], cliCtx.GetFromAddress(), cliCtx.GetFromAddress(), args[2])

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			fp, err0 := os.OpenFile(fmt.Sprintf("/home/lin/go/src/github.com/linbeier/authsys/data.json"), os.O_CREATE|os.O_RDWR, 0755)
			if err0 != nil {
				fmt.Printf("%s", err0)
			}
			defer fp.Close()

			var keymap map[string][]byte
			keymap = make(map[string][]byte)
			keymapbyte, err0 := ioutil.ReadAll(fp)

			if err0 != nil {
				fmt.Printf("error with read\n")
			}
			if !bytes.Equal(nil, keymapbyte) {
				codec.Cdc.MustUnmarshalJSON(keymapbyte, &keymap)
			}

			keymap[cliCtx.GetFromName()] = cliCtx.GetFromAddress().Bytes()

			keymapbyte = cdc.MustMarshalJSON(keymap)
			//清空文件，并写入新的map
			os.Truncate("/home/lin/go/src/github.com/linbeier/authsys/data.json", 0)
			fp.Seek(0, 0)

			_, err0 = fp.Write(keymapbyte)
			if err0 != nil {
				fmt.Printf("%s", err0)
			}

			// fptest, err2 := os.OpenFile("/home/lin/go/src/github.com/linbeier/authsys/testtime", os.O_APPEND, 0666)
			// if err2 != nil {
			// 	fmt.Printf("%s\n", err2)
			// }
			// fmt.Printf("fortest fortest\n")
			// n, err2 := fptest.Write(fmt.Sprintf("查询权限信息所使用时间为：%v\n", time.Since(timestart)))
			// if err2 != nil {
			// 	fmt.Printf("%s\n", err2)
			// }
			// defer fptest.Close()

			filepath := "./testtime"
			file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}
			defer file.Close()
			write := bufio.NewWriter(file)
			write.WriteString(fmt.Sprintf("SET交易权限1信息所使用时间为：%v\n", time.Since(types.Timestart)))
			write.Flush()
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdTransFiles(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transfile [toAccount] [filenames] [filehash] [origin] [authority]",
		Short: "transfer file's authority to other account. in msg, authority field represent which authority you want to transfer. in store, you should have the authority and 't' in your authority field",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			types.Timestart = time.Now()
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			fp, err0 := os.OpenFile(fmt.Sprintf("/home/lin/go/src/github.com/linbeier/authsys/data.json"), os.O_CREATE|os.O_RDWR, 0755)
			if err0 != nil {
				fmt.Printf("%s", err0)
			}
			defer fp.Close()

			var keymap map[string][]byte
			keymap = make(map[string][]byte)
			keymapbyte, err0 := ioutil.ReadAll(fp)
			//fmt.Printf("map byte: %s\n", keymapbyte)
			if err0 != nil {
				fmt.Printf("error with read\n")
			}
			if !bytes.Equal(nil, keymapbyte) {
				codec.Cdc.MustUnmarshalJSON(keymapbyte, &keymap)
			}
			var toaccount sdk.AccAddress
			toaccount = keymap[args[0]]
			// codec.Cdc.MustUnmarshalJSON(keymap[args[0]], &toaccount)

			var origin sdk.AccAddress
			origin = keymap[args[3]]
			// codec.Cdc.MustUnmarshalJSON(keymap[args[3]], &origin)

			msg := types.NewMsgTransFileAuth(args[1], args[2], cliCtx.GetFromAddress(), toaccount, origin, args[4])

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// fptest, err2 := os.OpenFile("/home/lin/go/src/github.com/linbeier/authsys/testtime", os.O_APPEND, 0666)
			// if err2 != nil {
			// 	fmt.Printf("%s", err0)
			// }
			// fmt.Printf("fortest fortest\n")
			// fptest.Write(fmt.Sprintf("查询权限信息所使用时间为：%v\n", time.Since(timestart)))
			// defer fptest.Close()
			filepath := "./testtime"
			file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}
			defer file.Close()
			write := bufio.NewWriter(file)
			write.WriteString(fmt.Sprintf("trans交易权限信息1所使用时间为：%v\n", time.Since(types.Timestart)))
			write.Flush()

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
