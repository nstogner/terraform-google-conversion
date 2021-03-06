// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
)

func GetSourceRepoRepositoryCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	if obj, err := GetSourceRepoRepositoryApiObject(d, config); err == nil {
		return Asset{
			Name: fmt.Sprintf("//sourcerepo.googleapis.com/%s", obj["selfLink"]),
			Type: "google.sourcerepo.Repository",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/sourcerepo/v1/rest",
				DiscoveryName:        "Repository",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetSourceRepoRepositoryApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandSourceRepoRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return obj, nil
}

func expandSourceRepoRepositoryName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
}
