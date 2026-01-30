package converter

import "strings"


type TransactionDataImporterOptions struct {
	payeeAsTag         bool
	payeeAsDescription bool
	memberAsTag        bool
	projectAsTag       bool
	merchantAsTag      bool
}


var DefaultImporterOptions = TransactionDataImporterOptions{
	payeeAsTag:         false,
	payeeAsDescription: false,
	memberAsTag:        false,
	projectAsTag:       false,
	merchantAsTag:      false,
}


func (o TransactionDataImporterOptions) IsPayeeAsTag() bool {
	return o.payeeAsTag
}


func (o TransactionDataImporterOptions) IsPayeeAsDescription() bool {
	return o.payeeAsDescription
}


func (o TransactionDataImporterOptions) IsMemberAsTag() bool {
	return o.memberAsTag
}


func (o TransactionDataImporterOptions) IsProjectAsTag() bool {
	return o.projectAsTag
}


func (o TransactionDataImporterOptions) IsMerchantAsTag() bool {
	return o.merchantAsTag
}


func (o TransactionDataImporterOptions) WithPayeeAsTag() TransactionDataImporterOptions {
	cloned := o.Clone()
	cloned.payeeAsTag = true
	return cloned
}


func (o TransactionDataImporterOptions) WithPayeeAsDescription() TransactionDataImporterOptions {
	cloned := o.Clone()
	cloned.payeeAsDescription = true
	return cloned
}


func (o TransactionDataImporterOptions) WithMemberAsTag() TransactionDataImporterOptions {
	cloned := o.Clone()
	cloned.memberAsTag = true
	return cloned
}


func (o TransactionDataImporterOptions) WithProjectAsTag() TransactionDataImporterOptions {
	cloned := o.Clone()
	cloned.projectAsTag = true
	return cloned
}


func (o TransactionDataImporterOptions) WithMerchantAsTag() TransactionDataImporterOptions {
	cloned := o.Clone()
	cloned.merchantAsTag = true
	return cloned
}


func (o TransactionDataImporterOptions) Clone() TransactionDataImporterOptions {
	return TransactionDataImporterOptions{
		payeeAsTag:         o.payeeAsTag,
		payeeAsDescription: o.payeeAsDescription,
		memberAsTag:        o.memberAsTag,
		projectAsTag:       o.projectAsTag,
		merchantAsTag:      o.merchantAsTag,
	}
}


func ParseImporterOptions(s string) TransactionDataImporterOptions {
	options := TransactionDataImporterOptions{}

	if s == "" {
		return options
	}

	for _, option := range strings.Split(s, ",") {
		switch option {
		case "payeeAsTag":
			options.payeeAsTag = true
		case "payeeAsDescription":
			options.payeeAsDescription = true
		case "memberAsTag":
			options.memberAsTag = true
		case "projectAsTag":
			options.projectAsTag = true
		case "merchantAsTag":
			options.merchantAsTag = true
		}
	}

	return options
}
