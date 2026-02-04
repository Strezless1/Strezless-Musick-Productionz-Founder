// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/strezless-musick-nexus-metadata-cli/internal/apiquery"
	"github.com/stainless-sdks/strezless-musick-nexus-metadata-cli/internal/requestflag"
	"github.com/stainless-sdks/strezless-musick-nexus-metadata-go"
	"github.com/stainless-sdks/strezless-musick-nexus-metadata-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var userCreate = cli.Command{
	Name:    "create",
	Usage:   "This can only be done by the logged in user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "username",
			BodyPath: "username",
		},
		&requestflag.Flag[int64]{
			Name:     "user-status",
			Usage:    "User Status",
			BodyPath: "userStatus",
		},
	},
	Action:          handleUserCreate,
	HideHelpCommand: true,
}

var userRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get user by user name",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "username",
			Required: true,
		},
	},
	Action:          handleUserRetrieve,
	HideHelpCommand: true,
}

var userUpdate = cli.Command{
	Name:    "update",
	Usage:   "This can only be done by the logged in user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "existing-username",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "username",
			BodyPath: "username",
		},
		&requestflag.Flag[int64]{
			Name:     "user-status",
			Usage:    "User Status",
			BodyPath: "userStatus",
		},
	},
	Action:          handleUserUpdate,
	HideHelpCommand: true,
}

var userDelete = cli.Command{
	Name:    "delete",
	Usage:   "This can only be done by the logged in user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "username",
			Required: true,
		},
	},
	Action:          handleUserDelete,
	HideHelpCommand: true,
}

var userCreateWithList = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-with-list",
	Usage:   "Creates list of users with given input array",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleUserCreateWithList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[int64]{
			Name:       "body.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.first-name",
			InnerField: "firstName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.last-name",
			InnerField: "lastName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.password",
			InnerField: "password",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.phone",
			InnerField: "phone",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.username",
			InnerField: "username",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "body.user-status",
			Usage:      "User Status",
			InnerField: "userStatus",
		},
	},
})

var userLogin = cli.Command{
	Name:    "login",
	Usage:   "Logs user into the system",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "password",
			Usage:     "The password for login in clear text",
			QueryPath: "password",
		},
		&requestflag.Flag[string]{
			Name:      "username",
			Usage:     "The user name for login",
			QueryPath: "username",
		},
	},
	Action:          handleUserLogin,
	HideHelpCommand: true,
}

var userLogout = cli.Command{
	Name:            "logout",
	Usage:           "Logs out current logged in user session",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUserLogout,
	HideHelpCommand: true,
}

func handleUserCreate(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.UserNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.User.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "user create", obj, format, transform)
}

func handleUserRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("username") && len(unusedArgs) > 0 {
		cmd.Set("username", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.User.Get(ctx, cmd.Value("username").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "user retrieve", obj, format, transform)
}

func handleUserUpdate(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("existing-username") && len(unusedArgs) > 0 {
		cmd.Set("existing-username", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.UserUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.User.Update(
		ctx,
		cmd.Value("existing-username").(string),
		params,
		options...,
	)
}

func handleUserDelete(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("username") && len(unusedArgs) > 0 {
		cmd.Set("username", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.User.Delete(ctx, cmd.Value("username").(string), options...)
}

func handleUserCreateWithList(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.UserNewWithListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.User.NewWithList(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "user create-with-list", obj, format, transform)
}

func handleUserLogin(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.UserLoginParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.User.Login(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "user login", obj, format, transform)
}

func handleUserLogout(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.User.Logout(ctx, options...)
}
