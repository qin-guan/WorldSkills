// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/asset"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assettransferlog"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	assetFields := schema.Asset{}.Fields()
	_ = assetFields
	// assetDescAssetSN is the schema descriptor for AssetSN field.
	assetDescAssetSN := assetFields[4].Descriptor()
	// asset.AssetSNValidator is a validator for the "AssetSN" field. It is called by the builders before save.
	asset.AssetSNValidator = assetDescAssetSN.Validators[0].(func(int) error)
	assettransferlogFields := schema.AssetTransferLog{}.Fields()
	_ = assettransferlogFields
	// assettransferlogDescFromAssetSN is the schema descriptor for FromAssetSN field.
	assettransferlogDescFromAssetSN := assettransferlogFields[2].Descriptor()
	// assettransferlog.FromAssetSNValidator is a validator for the "FromAssetSN" field. It is called by the builders before save.
	assettransferlog.FromAssetSNValidator = assettransferlogDescFromAssetSN.Validators[0].(func(int) error)
	// assettransferlogDescToAssetSN is the schema descriptor for ToAssetSN field.
	assettransferlogDescToAssetSN := assettransferlogFields[3].Descriptor()
	// assettransferlog.ToAssetSNValidator is a validator for the "ToAssetSN" field. It is called by the builders before save.
	assettransferlog.ToAssetSNValidator = assettransferlogDescToAssetSN.Validators[0].(func(int) error)
}
