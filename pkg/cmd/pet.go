// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/apiquery"
	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/binaryparam"
	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/requestflag"
	"github.com/stainless-sdks/strezless-musick-nexus-metadata-go"
	"github.com/stainless-sdks/strezless-musick-nexus-metadata-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var petCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Add a new pet to the store",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "photo-url",
			Required: true,
			BodyPath: "photoUrls",
		},
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "category",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "pet status in the store",
			BodyPath: "status",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handlePetCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"category": {
		&requestflag.InnerFlag[int64]{
			Name:       "category.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "category.name",
			InnerField: "name",
		},
	},
	"tag": {
		&requestflag.InnerFlag[int64]{
			Name:       "tag.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tag.name",
			InnerField: "name",
		},
	},
})

var petRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a single pet",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
	},
	Action:          handlePetRetrieve,
	HideHelpCommand: true,
}

var petUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing pet by Id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "photo-url",
			Required: true,
			BodyPath: "photoUrls",
		},
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "category",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "pet status in the store",
			BodyPath: "status",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handlePetUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"category": {
		&requestflag.InnerFlag[int64]{
			Name:       "category.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "category.name",
			InnerField: "name",
		},
	},
	"tag": {
		&requestflag.InnerFlag[int64]{
			Name:       "tag.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tag.name",
			InnerField: "name",
		},
	},
})

var petDelete = cli.Command{
	Name:    "delete",
	Usage:   "delete a pet",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
	},
	Action:          handlePetDelete,
	HideHelpCommand: true,
}

var petFindByStatus = cli.Command{
	Name:    "find-by-status",
	Usage:   "Multiple status values can be provided with comma separated strings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Status values that need to be considered for filter",
			Default:   "available",
			QueryPath: "status",
		},
	},
	Action:          handlePetFindByStatus,
	HideHelpCommand: true,
}

var petFindByTags = cli.Command{
	Name:    "find-by-tags",
	Usage:   "Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3\nfor testing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Tags to filter by",
			QueryPath: "tags",
		},
	},
	Action:          handlePetFindByTags,
	HideHelpCommand: true,
}

var petUpdateWithForm = cli.Command{
	Name:    "update-with-form",
	Usage:   "Updates a pet in the store with form data",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Name of pet that needs to be updated",
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Status of pet that needs to be updated",
			QueryPath: "status",
		},
	},
	Action:          handlePetUpdateWithForm,
	HideHelpCommand: true,
}

var petUploadImage = cli.Command{
	Name:    "upload-image",
	Usage:   "uploads an image",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "body",
			Required:  true,
			BodyRoot:  true,
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:      "additional-metadata",
			Usage:     "Additional Metadata",
			QueryPath: "additionalMetadata",
		},
	},
	Action:          handlePetUploadImage,
	HideHelpCommand: true,
}

func handlePetCreate(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.PetNewParams{}

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
	_, err = client.Pet.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pet create", obj, format, explicitFormat, transform)
}

func handlePetRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
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
	_, err = client.Pet.Get(ctx, cmd.Value("pet-id").(int64), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pet retrieve", obj, format, explicitFormat, transform)
}

func handlePetUpdate(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.PetUpdateParams{}

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
	_, err = client.Pet.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pet update", obj, format, explicitFormat, transform)
}

func handlePetDelete(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
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

	return client.Pet.Delete(ctx, cmd.Value("pet-id").(int64), options...)
}

func handlePetFindByStatus(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.PetFindByStatusParams{}

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
	_, err = client.Pet.FindByStatus(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pet find-by-status", obj, format, explicitFormat, transform)
}

func handlePetFindByTags(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.PetFindByTagsParams{}

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
	_, err = client.Pet.FindByTags(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pet find-by-tags", obj, format, explicitFormat, transform)
}

func handlePetUpdateWithForm(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := strezlessmusicknexusmetadata.PetUpdateWithFormParams{}

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

	return client.Pet.UpdateWithForm(
		ctx,
		cmd.Value("pet-id").(int64),
		params,
		options...,
	)
}

func handlePetUploadImage(ctx context.Context, cmd *cli.Command) error {
	client := strezlessmusicknexusmetadata.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("body") && len(unusedArgs) > 0 {
		cmd.Set("body", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	bodyReader, stdinInUse, err := binaryparam.FileOrStdin(os.Stdin, cmd.Value("body").(string))
	if err != nil {
		return fmt.Errorf("Failed on param '%s': %w", "body", err)
	}
	defer bodyReader.Close()

	params := strezlessmusicknexusmetadata.PetUploadImageParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationOctetStream,
		stdinInUse,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pet.UploadImage(
		ctx,
		cmd.Value("pet-id").(int64),
		bodyReader,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pet upload-image", obj, format, explicitFormat, transform)
}
