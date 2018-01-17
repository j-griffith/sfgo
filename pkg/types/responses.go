package types

// ListVolumesResult tbd
type ListVolumesResult struct {
	ID     int `json:"id"`
	Result struct {
		Volumes []Volume `json:"volumes"`
	} `json:"result"`
}

// CreateVolumeResult tbd
type CreateVolumeResult struct {
	ID     int `json:"id"`
	Result struct {
		VolumeID int64 `json:"volumeID"`
	} `json:"result"`
}

type CloneVolumeResult struct {
	Id     int `json:"id"`
	Result struct {
		CloneID     int64 `json:"cloneID"`
		VolumeID    int64 `json:"volumeID"`
		AsyncHandle int64 `json:"asyncHandle"`
	} `json:"result"`
}

type CreateSnapshotResult struct {
	Id     int `json:"id"`
	Result struct {
		SnapshotID int64  `json:"snapshotID"`
		Checksum   string `json:"checksum"`
	} `json:"result"`
}

type ListSnapshotsResult struct {
	ID     int `json:"id"`
	Result struct {
		Snapshots []Snapshot `json:"snapshots"`
	} `json:"result"`
}

type RollbackToSnapshotResult struct {
	ID     int `json:"id"`
	Result struct {
		Checksum   string `json:"checksum"`
		SnapshotID int64  `json:"snapshotID"`
	} `json:"result"`
}

// CreateVolumeAccessGroupResult tbd
type CreateVolumeAccessGroupResult struct {
	ID     int `json:"id"`
	Result struct {
		VagID int64 `json:"volumeAccessGroupID"`
	} `json:"result"`
}

// ListVolumesAccessGroupsResult tbd
type ListVolumesAccessGroupsResult struct {
	ID     int `json:"id"`
	Result struct {
		Vags []VolumeAccessGroup `json:"volumeAccessGroups"`
	} `json:"result"`
}

// EmptyResponse tbd
type EmptyResponse struct {
	ID     int `json:"id"`
	Result struct {
	} `json:"result"`
}

// GetAccountResult tbd
type GetAccountResult struct {
	ID     int `json:"id"`
	Result struct {
		Account Account `json:"account"`
	} `json:"result"`
}

// AddAccountResult tbd
type AddAccountResult struct {
	ID     int `json:"id"`
	Result struct {
		AccountID int64 `json:"accountID"`
	} `json:"result"`
}

type GetClusterCapacityResult struct {
	Id     int `json:"id"`
	Result struct {
		ClusterCapacity ClusterCapacity `json:"clusterCapacity"`
	} `json:"result"`
}

type GetClusterHardwareInfoResult struct {
	Id     int `json:"id"`
	Result struct {
		ClusterHardwareInfo ClusterHardwareInfo `json:"clusterHardwareInfo"`
	} `json:"result"`
}

type ModifyVolumeResult struct {
	Volume Volume `json:"volume,omitempty"`
	Curve  QoS    `json:"curve,omitempty"`
}
