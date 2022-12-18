package version

import (
	"bytes"
	"fmt"
)

type VersionInfo struct {
	Revision  string `json:"revision,omitempty"`
	Version   string `json:"version,omitempty"`
	BuildDate string `json:"build_date,omitempty"`
}

func GetVersion() *VersionInfo {
	ver := Version

	return &VersionInfo{
		Revision:  GitCommit,
		Version:   ver,
		BuildDate: BuildDate,
	}
}

func (c *VersionInfo) FullVersionNumber() string {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "v%s", c.Version)
	fmt.Fprintf(&versionString, " (%s)", c.Revision)

	if c.BuildDate != "" {
		fmt.Fprintf(&versionString, ", built %s", c.BuildDate)
	}

	return versionString.String()
}
