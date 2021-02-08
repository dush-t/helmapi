package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// InstallRequest represents an install command
type InstallRequest struct {
	ChartName         string                 `json:"chartName"`
	ReleaseName       string                 `json:"releaseName"`
	PrivateChartsRepo string                 `json:"privateChartsRepo"`
	Values            map[string]interface{} `json:"values"`
	Flags             []string               `json:"flags"`
}

// Describe returns a string description (for printing) of
// the chart to be installed
func (ir *InstallRequest) String() string {
	prettyValues, _ := json.MarshalIndent(ir.Values, "", "  ")
	return fmt.Sprintf(`
--------------------
Chart:   %s
Release: %s
Values:  %s
Flags:   %s
--------------------
	`, ir.ChartName, ir.ReleaseName, string(prettyValues), strings.Join(ir.Flags, " "))
}

func serializeValues(prefix string, values map[string]interface{}) []string {
	result := []string{}

	for key, val := range values {
		fullKey := prefix + key
		switch val.(type) {
		case string:
			result = append(result, fullKey+"="+val.(string))
		case map[string]interface{}:
			newPrefix := fullKey + "."
			nestedVals := serializeValues(newPrefix, val.(map[string]interface{}))
			result = append(result, nestedVals...)
		}
	}

	return result
}

// GetValues flattens the values JSON and returns a slice of
// key=value pairs
func (ir *InstallRequest) GetValues() []string {
	return serializeValues("", ir.Values)
}

// Execute will install the chart as specified by the InstallRequest
func (ir *InstallRequest) Execute() error {
	app := "helm"
	args := []string{"install", ir.ReleaseName, ir.ChartName}

	// Constructing the --set argument
	values := []string{"--set", strings.Join(ir.GetValues(), ",")}
	args = append(args, values...)

	// Checking release name is not empty
	if len(ir.ReleaseName) == 0 {
		return fmt.Errorf("You cannot provide an empty release name")
	}

	// Checking Chart name is not empty
	if len(ir.ChartName) == 0 {
		return fmt.Errorf("You cannot provide an empty chart name")
	}

	flags := ir.Flags
	if len(ir.PrivateChartsRepo) != 0 {
		flags = append(flags, "--repo", ir.PrivateChartsRepo)
	}

	args = append(args, flags...)

	log.Println("Installing chart:")
	log.Println(ir.String())

	cmd := exec.Command(app, args...)
	if err := execute(cmd); err != nil {
		return err
	}

	log.Println("Installation successful")
	return nil
}

// DeleteRequest represents an uninstall command
type DeleteRequest struct {
	ReleaseName string `json:"releaseName"`
}

// Describe returns a string description (for printing) of
// the release to be uninstalled
func (dr *DeleteRequest) String() string {
	return fmt.Sprintf(`
--------------------
Release: %s
--------------------
	`, dr.ReleaseName)
}

// Execute will uninstall the chart as specified by the DeleteRequest
func (dr *DeleteRequest) Execute() error {
	app := "helm"

	if len(dr.ReleaseName) == 0 {
		return fmt.Errorf("You cannot provide an empty release name")
	}

	log.Println("Uninstalling release:")
	log.Println(dr.String())

	cmd := exec.Command(app, "uninstall", dr.ReleaseName)
	if err := execute(cmd); err != nil {
		return err
	}

	log.Println("Uninstallation Successful")
	return nil
}
