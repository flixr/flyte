// Code generated by "enumer --type=CloudDeployment -json -yaml -trimprefix=CloudDeployment"; DO NOT EDIT.

package config

import (
	"encoding/json"
	"fmt"
)

const _CloudDeploymentName = "NoneAWSGCPSandboxLocal"

var _CloudDeploymentIndex = [...]uint8{0, 4, 7, 10, 17, 22}

func (i CloudDeployment) String() string {
	if i >= CloudDeployment(len(_CloudDeploymentIndex)-1) {
		return fmt.Sprintf("CloudDeployment(%d)", i)
	}
	return _CloudDeploymentName[_CloudDeploymentIndex[i]:_CloudDeploymentIndex[i+1]]
}

var _CloudDeploymentValues = []CloudDeployment{0, 1, 2, 3, 4}

var _CloudDeploymentNameToValueMap = map[string]CloudDeployment{
	_CloudDeploymentName[0:4]:   0,
	_CloudDeploymentName[4:7]:   1,
	_CloudDeploymentName[7:10]:  2,
	_CloudDeploymentName[10:17]: 3,
	_CloudDeploymentName[17:22]: 4,
}

// CloudDeploymentString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CloudDeploymentString(s string) (CloudDeployment, error) {
	if val, ok := _CloudDeploymentNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CloudDeployment values", s)
}

// CloudDeploymentValues returns all values of the enum
func CloudDeploymentValues() []CloudDeployment {
	return _CloudDeploymentValues
}

// IsACloudDeployment returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CloudDeployment) IsACloudDeployment() bool {
	for _, v := range _CloudDeploymentValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CloudDeployment
func (i CloudDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CloudDeployment
func (i *CloudDeployment) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CloudDeployment should be a string, got %s", data)
	}

	var err error
	*i, err = CloudDeploymentString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for CloudDeployment
func (i CloudDeployment) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CloudDeployment
func (i *CloudDeployment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CloudDeploymentString(s)
	return err
}