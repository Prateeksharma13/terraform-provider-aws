//go:generate go run ../../generate/tags/main.go -ServiceTagsSlice -UpdateTags
//go:generate go run ../../generate/listpages/main.go -ListOps=ListLicenseConfigurations,ListLicenseSpecificationsForResource,ListReceivedLicenses,ListDistributedGrants
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package licensemanager
