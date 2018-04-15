package types

// ListVolumes general List with options, should replace ListVolumesForAccount and ListActiveVolumes
type ListVolumesRequest struct {
	Accounts              []int64 `json:"accounts,omitempty"`
	IncludeVirtualVolumes bool    `json:"includeVirtualVolumes,omitempty"`
	IsPaired              bool    `json:"isPaired,omitempty"`
	Limit                 int64   `json:"limit,omitempty"`
	StartVolumeID         int64   `json:"startVolumeID,omitempty"`
	VolumeIDs             []int64 `json:"volumeIDs,omitempty"`
	VolumeName            string  `json:"volumeName,omitempty"`
	VolumeStatus          string  `json:"volumeStatus,omitempty"` // creating, snapshotting, active or deleted
}

// ListVolumesForAccountRequest tbd
type ListVolumesForAccountRequest struct {
	AccountID int64 `json:"accountID"`
}

// ListActiveVolumesRequest tbd
type ListActiveVolumesRequest struct {
	StartVolumeID int64 `json:"startVolumeID"`
	Limit         int64 `json:"limit"`
}

// CreateVolumeRequest tbd
type CreateVolumeRequest struct {
	Name       string      `json:"name"`
	AccountID  int64       `json:"accountID"`
	TotalSize  int64       `json:"totalSize"`
	Enable512e bool        `json:"enable512e"`
	Qos        QoS         `json:"qos,omitempty"`
	Attributes interface{} `json:"attributes"`
}

// DeleteVolumeRequest tbd
type DeleteVolumeRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type CloneVolumeRequest struct {
	VolumeID     int64       `json:"volumeID"`
	Name         string      `json:"name"`
	NewAccountID int64       `json:"newAccountID"`
	NewSize      int64       `json:"newSize"`
	Access       string      `json:"access"`
	SnapshotID   int64       `json:"snapshotID"`
	Attributes   interface{} `json:"attributes"`
}

type CreateSnapshotRequest struct {
	VolumeID                int64       `json:"volumeID"`
	SnapshotID              int64       `json:"snapshotID"`
	Name                    string      `json:"name"`
	EnableRemoteReplication bool        `json:"enableRemoteReplication"`
	Retention               string      `json:"retention"`
	Attributes              interface{} `json:"attributes"`
}

type ListSnapshotsRequest struct {
	VolumeID int64 `json:"volumeID"`
}

type RollbackToSnapshotRequest struct {
	VolumeID         int64       `json:"volumeID"`
	SnapshotID       int64       `json:"snapshotID"`
	SaveCurrentState bool        `json:"saveCurrentState"`
	Name             string      `json:"name"`
	Attributes       interface{} `json:"attributes"`
}

type DeleteSnapshotRequest struct {
	SnapshotID int64 `json:"snapshotID"`
}

// AddVolumesToVolumeAccessGroupRequest tbd
type AddVolumesToVolumeAccessGroupRequest struct {
	VolumeAccessGroupID int64   `json:"volumeAccessGroupID"`
	Volumes             []int64 `json:"volumes"`
}

// CreateVolumeAccessGroupRequest tbd
type CreateVolumeAccessGroupRequest struct {
	Name       string   `json:"name"`
	Volumes    []int64  `json:"volumes,omitempty"`
	Initiators []string `json:"initiators,omitempty"`
}

// AddInitiatorsToVolumeAccessGroupRequest tbd
type AddInitiatorsToVolumeAccessGroupRequest struct {
	Initiators []string `json:"initiators"`
	VAGID      int64    `json:"volumeAccessGroupID"`
}

// ListVolumeAccessGroupsRequest tbd
type ListVolumeAccessGroupsRequest struct {
	StartVAGID int64 `json:"startVolumeAccessGroupID,omitempty"`
	Limit      int64 `json:"limit,omitempty"`
}

// GetAccountByNameRequest tbd
type GetAccountByNameRequest struct {
	Name string `json:"username"`
}

// GetAccountByIDRequest tbd
type GetAccountByIDRequest struct {
	AccountID int64 `json:"accountID"`
}

// AddAccountRequest tbd
type AddAccountRequest struct {
	Username        string      `json:"username"`
	InitiatorSecret string      `json:"initiatorSecret,omitempty"`
	TargetSecret    string      `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type GetClusterCapacityRequest struct {
}

type ModifyVolumeRequest struct {
	VolumeID   int64       `json:"volumeID"`
	AccountID  int64       `json:"accountID,omitempty"`
	Access     string      `json:"access,omitempty"`
	Qos        QoS         `json:"qos,omitempty"`
	TotalSize  int64       `json:"totalSize,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}
