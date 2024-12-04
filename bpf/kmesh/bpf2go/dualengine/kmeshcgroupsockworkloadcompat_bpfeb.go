// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package dualengine

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshCgroupSockWorkloadCompatBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshCgroupSockWorkloadCompatBuf struct{ Data [40]int8 }

type KmeshCgroupSockWorkloadCompatKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshCgroupSockWorkloadCompatManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshCgroupSockWorkloadCompatOperationUsageData struct {
	StartTime     uint64
	EndTime       uint64
	PidTgid       uint64
	OperationType uint32
	_             [4]byte
}

type KmeshCgroupSockWorkloadCompatOperationUsageKey struct {
	SocketCookie  uint64
	OperationType uint32
	_             [4]byte
}

type KmeshCgroupSockWorkloadCompatSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshCgroupSockWorkloadCompat returns the embedded CollectionSpec for KmeshCgroupSockWorkloadCompat.
func LoadKmeshCgroupSockWorkloadCompat() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshCgroupSockWorkloadCompatBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshCgroupSockWorkloadCompat: %w", err)
	}

	return spec, err
}

// LoadKmeshCgroupSockWorkloadCompatObjects loads KmeshCgroupSockWorkloadCompat and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshCgroupSockWorkloadCompatObjects
//	*KmeshCgroupSockWorkloadCompatPrograms
//	*KmeshCgroupSockWorkloadCompatMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshCgroupSockWorkloadCompatObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshCgroupSockWorkloadCompat()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshCgroupSockWorkloadCompatSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadCompatSpecs struct {
	KmeshCgroupSockWorkloadCompatProgramSpecs
	KmeshCgroupSockWorkloadCompatMapSpecs
}

// KmeshCgroupSockWorkloadCompatSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadCompatProgramSpecs struct {
	CgroupConnect4Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect4_prog"`
	CgroupConnect6Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect6_prog"`
}

// KmeshCgroupSockWorkloadCompatMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadCompatMapSpecs struct {
	KmAuth         *ebpf.MapSpec `ebpf:"km_auth"`
	KmBackend      *ebpf.MapSpec `ebpf:"km_backend"`
	KmConfigmap    *ebpf.MapSpec `ebpf:"km_configmap"`
	KmDstinfo      *ebpf.MapSpec `ebpf:"km_dstinfo"`
	KmEndpoint     *ebpf.MapSpec `ebpf:"km_endpoint"`
	KmFrontend     *ebpf.MapSpec `ebpf:"km_frontend"`
	KmLogEvent     *ebpf.MapSpec `ebpf:"km_log_event"`
	KmManage       *ebpf.MapSpec `ebpf:"km_manage"`
	KmPerfInfo     *ebpf.MapSpec `ebpf:"km_perf_info"`
	KmPerfMap      *ebpf.MapSpec `ebpf:"km_perf_map"`
	KmService      *ebpf.MapSpec `ebpf:"km_service"`
	KmSockstorage  *ebpf.MapSpec `ebpf:"km_sockstorage"`
	KmTailcallprog *ebpf.MapSpec `ebpf:"km_tailcallprog"`
	KmTcpinfo      *ebpf.MapSpec `ebpf:"km_tcpinfo"`
	KmTmpbuf       *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmTuple        *ebpf.MapSpec `ebpf:"km_tuple"`
	KmWlpolicy     *ebpf.MapSpec `ebpf:"km_wlpolicy"`
	KmXdptailcall  *ebpf.MapSpec `ebpf:"km_xdptailcall"`
	KmeshMap1600   *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.MapSpec `ebpf:"kmesh_map64"`
}

// KmeshCgroupSockWorkloadCompatObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadCompatObjects struct {
	KmeshCgroupSockWorkloadCompatPrograms
	KmeshCgroupSockWorkloadCompatMaps
}

func (o *KmeshCgroupSockWorkloadCompatObjects) Close() error {
	return _KmeshCgroupSockWorkloadCompatClose(
		&o.KmeshCgroupSockWorkloadCompatPrograms,
		&o.KmeshCgroupSockWorkloadCompatMaps,
	)
}

// KmeshCgroupSockWorkloadCompatMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadCompatMaps struct {
	KmAuth         *ebpf.Map `ebpf:"km_auth"`
	KmBackend      *ebpf.Map `ebpf:"km_backend"`
	KmConfigmap    *ebpf.Map `ebpf:"km_configmap"`
	KmDstinfo      *ebpf.Map `ebpf:"km_dstinfo"`
	KmEndpoint     *ebpf.Map `ebpf:"km_endpoint"`
	KmFrontend     *ebpf.Map `ebpf:"km_frontend"`
	KmLogEvent     *ebpf.Map `ebpf:"km_log_event"`
	KmManage       *ebpf.Map `ebpf:"km_manage"`
	KmPerfInfo     *ebpf.Map `ebpf:"km_perf_info"`
	KmPerfMap      *ebpf.Map `ebpf:"km_perf_map"`
	KmService      *ebpf.Map `ebpf:"km_service"`
	KmSockstorage  *ebpf.Map `ebpf:"km_sockstorage"`
	KmTailcallprog *ebpf.Map `ebpf:"km_tailcallprog"`
	KmTcpinfo      *ebpf.Map `ebpf:"km_tcpinfo"`
	KmTmpbuf       *ebpf.Map `ebpf:"km_tmpbuf"`
	KmTuple        *ebpf.Map `ebpf:"km_tuple"`
	KmWlpolicy     *ebpf.Map `ebpf:"km_wlpolicy"`
	KmXdptailcall  *ebpf.Map `ebpf:"km_xdptailcall"`
	KmeshMap1600   *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.Map `ebpf:"kmesh_map64"`
}

func (m *KmeshCgroupSockWorkloadCompatMaps) Close() error {
	return _KmeshCgroupSockWorkloadCompatClose(
		m.KmAuth,
		m.KmBackend,
		m.KmConfigmap,
		m.KmDstinfo,
		m.KmEndpoint,
		m.KmFrontend,
		m.KmLogEvent,
		m.KmManage,
		m.KmPerfInfo,
		m.KmPerfMap,
		m.KmService,
		m.KmSockstorage,
		m.KmTailcallprog,
		m.KmTcpinfo,
		m.KmTmpbuf,
		m.KmTuple,
		m.KmWlpolicy,
		m.KmXdptailcall,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
	)
}

// KmeshCgroupSockWorkloadCompatPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadCompatPrograms struct {
	CgroupConnect4Prog *ebpf.Program `ebpf:"cgroup_connect4_prog"`
	CgroupConnect6Prog *ebpf.Program `ebpf:"cgroup_connect6_prog"`
}

func (p *KmeshCgroupSockWorkloadCompatPrograms) Close() error {
	return _KmeshCgroupSockWorkloadCompatClose(
		p.CgroupConnect4Prog,
		p.CgroupConnect6Prog,
	)
}

func _KmeshCgroupSockWorkloadCompatClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshcgroupsockworkloadcompat_bpfeb.o
var _KmeshCgroupSockWorkloadCompatBytes []byte
