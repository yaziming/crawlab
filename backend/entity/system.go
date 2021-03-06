package entity

type SystemInfo struct {
	ARCH        string       `json:"arch"`
	OS          string       `json:"os"`
	Hostname    string       `json:"host_name"`
	NumCpu      int          `json:"num_cpu"`
	Executables []Executable `json:"executables"`
}

type Executable struct {
	Path        string `json:"path"`
	FileName    string `json:"file_name"`
	DisplayName string `json:"display_name"`
}

type Lang struct {
	Name              string   `json:"name"`
	ExecutableName    string   `json:"executable_name"`
	ExecutablePaths   []string `json:"executable_paths"`
	DepExecutablePath string   `json:"dep_executable_path"`
	LockPath          string   `json:"lock_path"`
	InstallScript     string   `json:"install_script"`
	InstallStatus     string   `json:"install_status"`
	DepFileName       string   `json:"dep_file_name"`
	InstallDepArgs    string   `json:"install_dep_cmd"`
}

type Dependency struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Installed   bool   `json:"installed"`
}

type PackageJson struct {
	Dependencies map[string]string `json:"dependencies"`
}

type RedisMemoryStats struct {
	PeakAllocated    int64 `json:"peak_allocated"`
	TotalAllocated   int64 `json:"total_allocated"`
	StartupAllocated int64 `json:"startup_allocated"`
	OverheadTotal    int64 `json:"overhead_total"`
	KeysCount        int64 `json:"keys_count"`
	DatasetBytes     int64 `json:"dataset_bytes"`
}

type NodeStats struct {
	MemoryTotal        uint64  `json:"memory_total"`
	MemoryUsage        uint64  `json:"memory_usage"`
	MemoryUsagePercent float64 `json:"memory_usage_percent"`
	CpuUsagePercent    float64 `json:"cpu_usage_percent"`
	DiskTotal          uint64  `json:"disk_total"`
	DiskUsage          uint64  `json:"disk_usage"`
	DiskUsagePercent   float64 `json:"disk_usage_percent"`
}
