// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SystemInfo system info
// swagger:model SystemInfo
type SystemInfo struct {

	// Hardware architecture of the host, as returned by the Go runtime
	// (`GOARCH`).
	//
	// A full list of possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	//
	Architecture string `json:"Architecture,omitempty"`

	// The driver to use for managing cgroups.
	//
	// Enum: [cgroupfs systemd]
	CgroupDriver string `json:"CgroupDriver,omitempty"`

	// containerd commit
	ContainerdCommit *Commit `json:"ContainerdCommit,omitempty"`

	// Total number of containers on the host.
	Containers int64 `json:"Containers,omitempty"`

	// Number of containers with status `"paused"`.
	//
	ContainersPaused int64 `json:"ContainersPaused,omitempty"`

	// Number of containers with status `"running"`.
	//
	ContainersRunning int64 `json:"ContainersRunning,omitempty"`

	// Number of containers with status `"stopped"`.
	//
	ContainersStopped int64 `json:"ContainersStopped,omitempty"`

	// Indicates if the daemon is running in debug-mode / with debug-level logging enabled.
	Debug bool `json:"Debug,omitempty"`

	// default registry can be defined by user.
	//
	DefaultRegistry string `json:"DefaultRegistry,omitempty"`

	// Name of the default OCI runtime that is used when starting containers.
	// The default can be overridden per-container at create time.
	//
	DefaultRuntime string `json:"DefaultRuntime,omitempty"`

	// Name of the storage driver in use.
	Driver string `json:"Driver,omitempty"`

	// Information specific to the storage driver, provided as
	// "label" / "value" pairs.
	//
	// This information is provided by the storage driver, and formatted
	// in a way consistent with the output of `pouch info` on the command
	// line.
	//
	// <p><br /></p>
	//
	// > **Note**: The information returned in this field, including the
	// > formatting of values and labels, should not be considered stable,
	// > and may change without notice.
	//
	DriverStatus [][]string `json:"DriverStatus"`

	// Indicates if experimental features are enabled on the daemon.
	//
	ExperimentalBuild bool `json:"ExperimentalBuild,omitempty"`

	// HTTP-proxy configured for the daemon. This value is obtained from the
	// [`HTTP_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable.
	//
	// Containers do not automatically inherit this configuration.
	//
	HTTPProxy string `json:"HttpProxy,omitempty"`

	// HTTPS-proxy configured for the daemon. This value is obtained from the
	// [`HTTPS_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable.
	//
	// Containers do not automatically inherit this configuration.
	//
	HTTPSProxy string `json:"HttpsProxy,omitempty"`

	// Unique identifier of the daemon.
	//
	// <p><br /></p>
	//
	// > **Note**: The format of the ID itself is not part of the API, and
	// > should not be considered stable.
	//
	ID string `json:"ID,omitempty"`

	// Total number of images on the host.
	//
	// Both _tagged_ and _untagged_ (dangling) images are counted.
	//
	Images int64 `json:"Images,omitempty"`

	// Address / URL of the index server that is used for image search,
	// and as a default for user authentication.
	//
	IndexServerAddress string `json:"IndexServerAddress,omitempty"`

	// Kernel version of the host.
	// On Linux, this information obtained from `uname`.
	//
	KernelVersion string `json:"KernelVersion,omitempty"`

	// User-defined labels (key/value metadata) as set on the daemon.
	//
	Labels []string `json:"Labels"`

	// List of addresses the pouchd listens on
	ListenAddresses []string `json:"ListenAddresses"`

	// Indicates if live restore is enabled.
	// If enabled, containers are kept running when the daemon is shutdown
	// or upon daemon start if running containers are detected.
	//
	LiveRestoreEnabled bool `json:"LiveRestoreEnabled,omitempty"`

	// The logging driver to use as a default for new containers.
	//
	LoggingDriver string `json:"LoggingDriver,omitempty"`

	// Indicates if lxcfs is enabled.
	//
	LxcfsEnabled bool `json:"LxcfsEnabled,omitempty"`

	// Total amount of physical memory available on the host, in kilobytes (kB).
	//
	MemTotal int64 `json:"MemTotal,omitempty"`

	// The number of logical CPUs usable by the daemon.
	//
	// The number of available CPUs is checked by querying the operating
	// system when the daemon starts. Changes to operating system CPU
	// allocation after the daemon is started are not reflected.
	//
	NCPU int64 `json:"NCPU,omitempty"`

	// Hostname of the host.
	Name string `json:"Name,omitempty"`

	// Generic type of the operating system of the host, as returned by the
	// Go runtime (`GOOS`).
	//
	// Currently returned value is "linux". A full list of
	// possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	//
	OSType string `json:"OSType,omitempty"`

	// Name of the host's operating system, for example: "Ubuntu 16.04.2 LTS".
	//
	OperatingSystem string `json:"OperatingSystem,omitempty"`

	// Root directory of persistent Pouch state.
	//
	// Defaults to `/var/lib/pouch` on Linux.
	//
	PouchRootDir string `json:"PouchRootDir,omitempty"`

	// registry config
	RegistryConfig *RegistryServiceConfig `json:"RegistryConfig,omitempty"`

	// runc commit
	RuncCommit *Commit `json:"RuncCommit,omitempty"`

	// List of [OCI compliant](https://github.com/opencontainers/runtime-spec)
	// runtimes configured on the daemon. Keys hold the "name" used to
	// reference the runtime.
	//
	// The Pouch daemon relies on an OCI compliant runtime (invoked via the
	// `containerd` daemon) as its interface to the Linux kernel namespaces,
	// cgroups, and SELinux.
	//
	// The default runtime is `runc`, and automatically configured. Additional
	// runtimes can be configured by the user and will be listed here.
	//
	Runtimes map[string]Runtime `json:"Runtimes,omitempty"`

	// List of security features that are enabled on the daemon, such as
	// apparmor, seccomp, SELinux, and user-namespaces (userns).
	//
	// Additional configuration options for each security feature may
	// be present, and are included as a comma-separated list of key/value
	// pairs.
	//
	SecurityOptions []string `json:"SecurityOptions"`

	// Version string of the daemon.
	//
	ServerVersion string `json:"ServerVersion,omitempty"`

	// The list of volume drivers which the pouchd supports
	//
	VolumeDrivers []string `json:"VolumeDrivers"`
}

// Validate validates this system info
func (m *SystemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCgroupDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerdCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuncCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var systemInfoTypeCgroupDriverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cgroupfs","systemd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		systemInfoTypeCgroupDriverPropEnum = append(systemInfoTypeCgroupDriverPropEnum, v)
	}
}

const (

	// SystemInfoCgroupDriverCgroupfs captures enum value "cgroupfs"
	SystemInfoCgroupDriverCgroupfs string = "cgroupfs"

	// SystemInfoCgroupDriverSystemd captures enum value "systemd"
	SystemInfoCgroupDriverSystemd string = "systemd"
)

// prop value enum
func (m *SystemInfo) validateCgroupDriverEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, systemInfoTypeCgroupDriverPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SystemInfo) validateCgroupDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.CgroupDriver) { // not required
		return nil
	}

	// value enum
	if err := m.validateCgroupDriverEnum("CgroupDriver", "body", m.CgroupDriver); err != nil {
		return err
	}

	return nil
}

func (m *SystemInfo) validateContainerdCommit(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerdCommit) { // not required
		return nil
	}

	if m.ContainerdCommit != nil {
		if err := m.ContainerdCommit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContainerdCommit")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) validateRegistryConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistryConfig) { // not required
		return nil
	}

	if m.RegistryConfig != nil {
		if err := m.RegistryConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegistryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) validateRuncCommit(formats strfmt.Registry) error {

	if swag.IsZero(m.RuncCommit) { // not required
		return nil
	}

	if m.RuncCommit != nil {
		if err := m.RuncCommit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RuncCommit")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) validateRuntimes(formats strfmt.Registry) error {

	if swag.IsZero(m.Runtimes) { // not required
		return nil
	}

	for k := range m.Runtimes {

		if err := validate.Required("Runtimes"+"."+k, "body", m.Runtimes[k]); err != nil {
			return err
		}
		if val, ok := m.Runtimes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemInfo) UnmarshalBinary(b []byte) error {
	var res SystemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
