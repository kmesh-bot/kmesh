// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package dualengine

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshXDPAuthCompatBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshXDPAuthCompatBuf struct{ Data [40]int8 }

type KmeshXDPAuthCompatKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshXDPAuthCompatManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshXDPAuthCompatSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshXDPAuthCompat returns the embedded CollectionSpec for KmeshXDPAuthCompat.
func LoadKmeshXDPAuthCompat() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshXDPAuthCompatBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshXDPAuthCompat: %w", err)
	}

	return spec, err
}

// LoadKmeshXDPAuthCompatObjects loads KmeshXDPAuthCompat and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshXDPAuthCompatObjects
//	*KmeshXDPAuthCompatPrograms
//	*KmeshXDPAuthCompatMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshXDPAuthCompatObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshXDPAuthCompat()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshXDPAuthCompatSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshXDPAuthCompatSpecs struct {
	KmeshXDPAuthCompatProgramSpecs
	KmeshXDPAuthCompatMapSpecs
}

// KmeshXDPAuthCompatSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshXDPAuthCompatProgramSpecs struct {
	PolicyCheck            *ebpf.ProgramSpec `ebpf:"policy_check"`
	RuleCheck              *ebpf.ProgramSpec `ebpf:"rule_check"`
	XdpAuthz               *ebpf.ProgramSpec `ebpf:"xdp_authz"`
	XdpShutdownInUserspace *ebpf.ProgramSpec `ebpf:"xdp_shutdown_in_userspace"`
}

// KmeshXDPAuthCompatMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshXDPAuthCompatMapSpecs struct {
	KmAuth         *ebpf.MapSpec `ebpf:"km_auth"`
	KmAuthz        *ebpf.MapSpec `ebpf:"km_authz"`
	KmBackend      *ebpf.MapSpec `ebpf:"km_backend"`
	KmConfigmap    *ebpf.MapSpec `ebpf:"km_configmap"`
	KmEndpoint     *ebpf.MapSpec `ebpf:"km_endpoint"`
	KmFrontend     *ebpf.MapSpec `ebpf:"km_frontend"`
	KmLogEvent     *ebpf.MapSpec `ebpf:"km_log_event"`
	KmManage       *ebpf.MapSpec `ebpf:"km_manage"`
	KmService      *ebpf.MapSpec `ebpf:"km_service"`
	KmSockstorage  *ebpf.MapSpec `ebpf:"km_sockstorage"`
	KmTailcallprog *ebpf.MapSpec `ebpf:"km_tailcallprog"`
	KmTcargs       *ebpf.MapSpec `ebpf:"km_tcargs"`
	KmTmpbuf       *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmTuple        *ebpf.MapSpec `ebpf:"km_tuple"`
	KmWlpolicy     *ebpf.MapSpec `ebpf:"km_wlpolicy"`
	KmXdptailcall  *ebpf.MapSpec `ebpf:"km_xdptailcall"`
	KmeshMap1600   *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.MapSpec `ebpf:"kmesh_map64"`
}

// KmeshXDPAuthCompatObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshXDPAuthCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshXDPAuthCompatObjects struct {
	KmeshXDPAuthCompatPrograms
	KmeshXDPAuthCompatMaps
}

func (o *KmeshXDPAuthCompatObjects) Close() error {
	return _KmeshXDPAuthCompatClose(
		&o.KmeshXDPAuthCompatPrograms,
		&o.KmeshXDPAuthCompatMaps,
	)
}

// KmeshXDPAuthCompatMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshXDPAuthCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshXDPAuthCompatMaps struct {
	KmAuth         *ebpf.Map `ebpf:"km_auth"`
	KmAuthz        *ebpf.Map `ebpf:"km_authz"`
	KmBackend      *ebpf.Map `ebpf:"km_backend"`
	KmConfigmap    *ebpf.Map `ebpf:"km_configmap"`
	KmEndpoint     *ebpf.Map `ebpf:"km_endpoint"`
	KmFrontend     *ebpf.Map `ebpf:"km_frontend"`
	KmLogEvent     *ebpf.Map `ebpf:"km_log_event"`
	KmManage       *ebpf.Map `ebpf:"km_manage"`
	KmService      *ebpf.Map `ebpf:"km_service"`
	KmSockstorage  *ebpf.Map `ebpf:"km_sockstorage"`
	KmTailcallprog *ebpf.Map `ebpf:"km_tailcallprog"`
	KmTcargs       *ebpf.Map `ebpf:"km_tcargs"`
	KmTmpbuf       *ebpf.Map `ebpf:"km_tmpbuf"`
	KmTuple        *ebpf.Map `ebpf:"km_tuple"`
	KmWlpolicy     *ebpf.Map `ebpf:"km_wlpolicy"`
	KmXdptailcall  *ebpf.Map `ebpf:"km_xdptailcall"`
	KmeshMap1600   *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.Map `ebpf:"kmesh_map64"`
}

func (m *KmeshXDPAuthCompatMaps) Close() error {
	return _KmeshXDPAuthCompatClose(
		m.KmAuth,
		m.KmAuthz,
		m.KmBackend,
		m.KmConfigmap,
		m.KmEndpoint,
		m.KmFrontend,
		m.KmLogEvent,
		m.KmManage,
		m.KmService,
		m.KmSockstorage,
		m.KmTailcallprog,
		m.KmTcargs,
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

// KmeshXDPAuthCompatPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshXDPAuthCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshXDPAuthCompatPrograms struct {
	PolicyCheck            *ebpf.Program `ebpf:"policy_check"`
	RuleCheck              *ebpf.Program `ebpf:"rule_check"`
	XdpAuthz               *ebpf.Program `ebpf:"xdp_authz"`
	XdpShutdownInUserspace *ebpf.Program `ebpf:"xdp_shutdown_in_userspace"`
}

func (p *KmeshXDPAuthCompatPrograms) Close() error {
	return _KmeshXDPAuthCompatClose(
		p.PolicyCheck,
		p.RuleCheck,
		p.XdpAuthz,
		p.XdpShutdownInUserspace,
	)
}

func _KmeshXDPAuthCompatClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshxdpauthcompat_bpfel.o
var _KmeshXDPAuthCompatBytes []byte
