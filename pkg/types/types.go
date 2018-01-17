package types

// QoS data structure
type QoS struct {
	MinIOPS   int64 `json:"minIOPS,omitempty"`
	MaxIOPS   int64 `json:"maxIOPS,omitempty"`
	BurstIOPS int64 `json:"burstIOPS,omitempty"`
	BurstTime int64 `json:"-"`
}

// VolumePair data structure
type VolumePair struct {
	ClusterPairID    int64  `json:"clusterPairID"`
	RemoteVolumeID   int64  `json:"remoteVolumeID"`
	RemoteSliceID    int64  `json:"remoteSliceID"`
	RemoteVolumeName string `json:"remoteVolumeName"`
	VolumePairUUID   string `json:"volumePairUUID"`
}

// Volume data structure
type Volume struct {
	VolumeID           int64        `json:"volumeID"`
	Name               string       `json:"name"`
	AccountID          int64        `json:"accountID"`
	CreateTime         string       `json:"createTime"`
	Status             string       `json:"status"`
	Access             string       `json:"access"`
	Enable512e         bool         `json:"enable512e"`
	Iqn                string       `json:"iqn"`
	ScsiEUIDeviceID    string       `json:"scsiEUIDeviceID"`
	ScsiNAADeviceID    string       `json:"scsiNAADeviceID"`
	Qos                QoS          `json:"qos"`
	VolumeAccessGroups []int64      `json:"volumeAccessGroups"`
	VolumePairs        []VolumePair `json:"volumePairs"`
	DeleteTime         string       `json:"deleteTime"`
	PurgeTime          string       `json:"purgeTime"`
	SliceCount         int64        `json:"sliceCount"`
	TotalSize          int64        `json:"totalSize"`
	BlockSize          int64        `json:"blockSize"`
	VirtualVolumeID    string       `json:"virtualVolumeID"`
	Attributes         interface{}  `json:"attributes"`
}

// Snapshot data structure
type Snapshot struct {
	SnapshotID int64       `json:"snapshotID"`
	VolumeID   int64       `json:"volumeID"`
	Name       string      `json:"name"`
	Checksum   string      `json:"checksum"`
	Status     string      `json:"status"`
	TotalSize  int64       `json:"totalSize"`
	GroupID    int64       `json:"groupID"`
	CreateTime string      `json:"createTime"`
	Attributes interface{} `json:"attributes"`
}

// VolumeAccessGroup data structure
type VolumeAccessGroup struct {
	Initiators     []string    `json:"initiators"`
	Attributes     interface{} `json:"attributes"`
	DeletedVolumes []int64     `json:"deletedVolumes"`
	Name           string      `json:"name"`
	VAGID          int64       `json:"volumeAccessGroupID"`
	Volumes        []int64     `json:"volumes"`
}

// Account data structure
type Account struct {
	AccountID       int64       `json:"accountID,omitempty"`
	Username        string      `json:"username,omitempty"`
	Status          string      `json:"status,omitempty"`
	Volumes         []int64     `json:"volumes,omitempty"`
	InitiatorSecret string      `json:"initiatorSecret,omitempty"`
	TargetSecret    string      `json:"targetSecret,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

// ClusterCapacity data structure
type ClusterCapacity struct {
	ActiveBlockSpace             int64  `json:"activeBlockSpace"`
	ActiveSessions               int64  `json:"activeSessions"`
	AverageIOPS                  int64  `json:"averageIOPS"`
	ClusterRecentIOSize          int64  `json:"clusterRecentIOSize"`
	CurrentIOPS                  int64  `json:"currentIOPS"`
	MaxIOPS                      int64  `json:"maxIOPS"`
	MaxOverProvisionableSpace    int64  `json:"maxOverProvisionableSpace"`
	MaxProvisionedSpace          int64  `json:"maxProvisionedSpace"`
	MaxUsedMetadataSpace         int64  `json:"maxUsedMetadataSpace"`
	MaxUsedSpace                 int64  `json:"maxUsedSpace"`
	NonZeroBlocks                int64  `json:"nonZeroBlocks"`
	PeakActiveSessions           int64  `json:"peakActiveSessions"`
	PeakIOPS                     int64  `json:"peakIOPS"`
	ProvisionedSpace             int64  `json:"provisionedSpace"`
	Timestamp                    string `json:"timestamp"`
	TotalOps                     int64  `json:"totalOps"`
	UniqueBlocks                 int64  `json:"uniqueBlocks"`
	UniqueBlocksUsedSpace        int64  `json:"uniqueBlocksUsedSpace"`
	UsedMetadataSpace            int64  `json:"usedMetadataSpace"`
	UsedMetadataSpaceInSnapshots int64  `json:"usedMetadataSpaceInSnapshots"`
	UsedSpace                    int64  `json:"usedSpace"`
	ZeroBlocks                   int64  `json:"zeroBlocks"`
}

// ClusterHardwareInfo data structure
type ClusterHardwareInfo struct {
	Drives interface{} `json:"drives"`
	Nodes  interface{} `json:"nodes"`
}
