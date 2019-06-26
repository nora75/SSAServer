// Code generated by goa v3.0.2, DO NOT EDIT.
//
// SSAServer HTTP client CLI support package
//
// Command:
// $ goa gen SSAServer/design

package cli

import (
	ssaserverc "SSAServer/gen/http/ssa_server/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `ssa-server (register|login|change-group|delete-user|save-data|return-data-list|pick-up-data)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` ssa-server register --body '{
      "group_id": "group-bzuiphjjgdas",
      "mail": "hoge@hoge.com",
      "password": "passW0rd",
      "user_name": "平野竜也"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	sSAServerSaveDataEncoderFn ssaserverc.SSAServerSaveDataEncoderFunc,
	sSAServerPickUpDataEncoderFn ssaserverc.SSAServerPickUpDataEncoderFunc,
) (goa.Endpoint, interface{}, error) {
	var (
		sSAServerFlags = flag.NewFlagSet("ssa-server", flag.ContinueOnError)

		sSAServerRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		sSAServerRegisterBodyFlag = sSAServerRegisterFlags.String("body", "REQUIRED", "")

		sSAServerLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		sSAServerLoginBodyFlag = sSAServerLoginFlags.String("body", "REQUIRED", "")

		sSAServerChangeGroupFlags      = flag.NewFlagSet("change-group", flag.ExitOnError)
		sSAServerChangeGroupBodyFlag   = sSAServerChangeGroupFlags.String("body", "REQUIRED", "")
		sSAServerChangeGroupUserIDFlag = sSAServerChangeGroupFlags.String("userid", "REQUIRED", "User ID")

		sSAServerDeleteUserFlags      = flag.NewFlagSet("delete-user", flag.ExitOnError)
		sSAServerDeleteUserBodyFlag   = sSAServerDeleteUserFlags.String("body", "REQUIRED", "")
		sSAServerDeleteUserUserIDFlag = sSAServerDeleteUserFlags.String("userid", "REQUIRED", "User ID")

		sSAServerSaveDataFlags       = flag.NewFlagSet("save-data", flag.ExitOnError)
		sSAServerSaveDataBodyFlag    = sSAServerSaveDataFlags.String("body", "REQUIRED", "")
		sSAServerSaveDataGroupIDFlag = sSAServerSaveDataFlags.String("groupid", "REQUIRED", "Group ID")

		sSAServerReturnDataListFlags       = flag.NewFlagSet("return-data-list", flag.ExitOnError)
		sSAServerReturnDataListBodyFlag    = sSAServerReturnDataListFlags.String("body", "REQUIRED", "")
		sSAServerReturnDataListGroupIDFlag = sSAServerReturnDataListFlags.String("groupid", "REQUIRED", "Group ID")

		sSAServerPickUpDataFlags        = flag.NewFlagSet("pick-up-data", flag.ExitOnError)
		sSAServerPickUpDataBodyFlag     = sSAServerPickUpDataFlags.String("body", "REQUIRED", "")
		sSAServerPickUpDataGroupIDFlag  = sSAServerPickUpDataFlags.String("groupid", "REQUIRED", "Group ID")
		sSAServerPickUpDataDataTypeFlag = sSAServerPickUpDataFlags.String("data-type", "REQUIRED", "Data type")
	)
	sSAServerFlags.Usage = sSAServerUsage
	sSAServerRegisterFlags.Usage = sSAServerRegisterUsage
	sSAServerLoginFlags.Usage = sSAServerLoginUsage
	sSAServerChangeGroupFlags.Usage = sSAServerChangeGroupUsage
	sSAServerDeleteUserFlags.Usage = sSAServerDeleteUserUsage
	sSAServerSaveDataFlags.Usage = sSAServerSaveDataUsage
	sSAServerReturnDataListFlags.Usage = sSAServerReturnDataListUsage
	sSAServerPickUpDataFlags.Usage = sSAServerPickUpDataUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "ssa-server":
			svcf = sSAServerFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "ssa-server":
			switch epn {
			case "register":
				epf = sSAServerRegisterFlags

			case "login":
				epf = sSAServerLoginFlags

			case "change-group":
				epf = sSAServerChangeGroupFlags

			case "delete-user":
				epf = sSAServerDeleteUserFlags

			case "save-data":
				epf = sSAServerSaveDataFlags

			case "return-data-list":
				epf = sSAServerReturnDataListFlags

			case "pick-up-data":
				epf = sSAServerPickUpDataFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "ssa-server":
			c := ssaserverc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = ssaserverc.BuildRegisterPayload(*sSAServerRegisterBodyFlag)
			case "login":
				endpoint = c.Login()
				data, err = ssaserverc.BuildLoginPayload(*sSAServerLoginBodyFlag)
			case "change-group":
				endpoint = c.ChangeGroup()
				data, err = ssaserverc.BuildChangeGroupPayload(*sSAServerChangeGroupBodyFlag, *sSAServerChangeGroupUserIDFlag)
			case "delete-user":
				endpoint = c.DeleteUser()
				data, err = ssaserverc.BuildDeleteUserPayload(*sSAServerDeleteUserBodyFlag, *sSAServerDeleteUserUserIDFlag)
			case "save-data":
				endpoint = c.SaveData(sSAServerSaveDataEncoderFn)
				data, err = ssaserverc.BuildSaveDataPayload(*sSAServerSaveDataBodyFlag, *sSAServerSaveDataGroupIDFlag)
			case "return-data-list":
				endpoint = c.ReturnDataList()
				data, err = ssaserverc.BuildReturnDataListPayload(*sSAServerReturnDataListBodyFlag, *sSAServerReturnDataListGroupIDFlag)
			case "pick-up-data":
				endpoint = c.PickUpData(sSAServerPickUpDataEncoderFn)
				data, err = ssaserverc.BuildPickUpDataPayload(*sSAServerPickUpDataBodyFlag, *sSAServerPickUpDataGroupIDFlag, *sSAServerPickUpDataDataTypeFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// ssa-serverUsage displays the usage of the ssa-server command and its
// subcommands.
func sSAServerUsage() {
	fmt.Fprintf(os.Stderr, `SSAアプリケーションのサーバーサイド
Usage:
    %s [globalflags] ssa-server COMMAND [flags]

COMMAND:
    register: SSAへの新規登録
    login: SSAへのログイン
    change-group: グループIDを変更する
    delete-user: 既存ユーザーの消去
    save-data: データをサーバーへ保存する
    return-data-list: データのリストを取得する
    pick-up-data: データをサーバーから取得する

Additional help:
    %s ssa-server COMMAND --help
`, os.Args[0], os.Args[0])
}
func sSAServerRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ssa-server register -body JSON

SSAへの新規登録
    -body JSON: 

Example:
    `+os.Args[0]+` ssa-server register --body '{
      "group_id": "group-bzuiphjjgdas",
      "mail": "hoge@hoge.com",
      "password": "passW0rd",
      "user_name": "平野竜也"
   }'
`, os.Args[0])
}

func sSAServerLoginUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ssa-server login -body JSON

SSAへのログイン
    -body JSON: 

Example:
    `+os.Args[0]+` ssa-server login --body '{
      "mail": "hoge@hoge.com",
      "password": "password"
   }'
`, os.Args[0])
}

func sSAServerChangeGroupUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ssa-server change-group -body JSON -userid INT

グループIDを変更する
    -body JSON: 
    -userid INT: User ID

Example:
    `+os.Args[0]+` ssa-server change-group --body '{
      "group_id": "group2",
      "password": "pass1234"
   }' --userid 123
`, os.Args[0])
}

func sSAServerDeleteUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ssa-server delete-user -body JSON -userid INT

既存ユーザーの消去
    -body JSON: 
    -userid INT: User ID

Example:
    `+os.Args[0]+` ssa-server delete-user --body '{
      "password": "pass12345"
   }' --userid 1342
`, os.Args[0])
}

func sSAServerSaveDataUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ssa-server save-data -body JSON -groupid STRING

データをサーバーへ保存する
    -body JSON: 
    -groupid STRING: Group ID

Example:
    `+os.Args[0]+` ssa-server save-data --body '{
      "Data": "Qui ut.",
      "Image": "Aliquam alias harum nisi.",
      "data_name": "Diary_312_2019-03-02_12-07-35",
      "data_type": 1,
      "image_name": "Image_2017-05-25-26-32",
      "title": "たいとる",
      "user_id": 28532
   }' --groupid "group1"
`, os.Args[0])
}

func sSAServerReturnDataListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ssa-server return-data-list -body JSON -groupid STRING

データのリストを取得する
    -body JSON: 
    -groupid STRING: Group ID

Example:
    `+os.Args[0]+` ssa-server return-data-list --body '{
      "user_id": 537829
   }' --groupid "group-agieoa"
`, os.Args[0])
}

func sSAServerPickUpDataUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ssa-server pick-up-data -body JSON -groupid STRING -data-type STRING

データをサーバーから取得する
    -body JSON: 
    -groupid STRING: Group ID
    -data-type STRING: Data type

Example:
    `+os.Args[0]+` ssa-server pick-up-data --body '{
      "data_name": "Record_12_2019-06-02_12-07-35",
      "user_id": 65
   }' --groupid "group-isg" --data-type "0"
`, os.Args[0])
}
