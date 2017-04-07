package elasticsearch

var NilCluster = Cluster{}

type Cluster struct {
	ClusterName	string			`json:"cluster_name"`
	Nodes		map[string]Node 	`json:"nodes"`
	Host		string
}

type Node struct {
	Name		string		`json:"name"`
	Host 		string 		`json:"host"`
	Jvm		JvmInfo		`json:"jvm"`
}

type JvmInfo struct {
	Timestamp 	int64		`json:"timestamp"`
	Mem   		MemInfo		`json:"mem"`
}

type MemInfo struct {
	HeapUsedInBytes		int64 			`json:"heap_used_in_bytes"`
	HeapUsedPercent		int 			`json:"heap_used_percent"`
	HeapMaxInBytes		int64			`json:"heap_max_in_bytes"`
	NoneHeapUsedInBytes	int64 			`json:"none_heap_used_in_bytes"`
	Pools			map[string] PoolInfo	`json:"pools"`
}

//type Pool struct {
//	Young		PoolInfo		`json:"young"`
//	Survivor	PoolInfo		`json:"survivor"`
//	Old		PoolInfo		`json:"old"`
//}

type PoolInfo struct {
	UsedInBytes		int64 			`json:"used_in_bytes"`
	MaxInBytes		int64 			`json:"max_in_bytes"`
}
