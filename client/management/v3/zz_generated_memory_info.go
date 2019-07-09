package client

const (
	MemoryInfoType             = "memoryInfo"
	MemoryInfoFieldMemTotalKiB = "memTotalKiB"
)

type MemoryInfo struct {
	MemTotalKiB int64 `json:"memTotalKiB,omitempty" yaml:"mem_total_ki_b,omitempty"`
}
