// Code generated by bpf2go; DO NOT EDIT.
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package bpf2go

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadCgroupSock returns the embedded CollectionSpec for CgroupSock.
func LoadCgroupSock() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_CgroupSockBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load CgroupSock: %w", err)
	}

	return spec, err
}

// LoadCgroupSockObjects loads CgroupSock and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *CgroupSockObjects
//     *CgroupSockPrograms
//     *CgroupSockMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadCgroupSockObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadCgroupSock()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// CgroupSockSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type CgroupSockSpecs struct {
	CgroupSockProgramSpecs
	CgroupSockMapSpecs
}

// CgroupSockSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type CgroupSockProgramSpecs struct {
	SockConnect4 *ebpf.ProgramSpec `ebpf:"sock_connect4"`
}

// CgroupSockMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type CgroupSockMapSpecs struct {
	Cluster      *ebpf.MapSpec `ebpf:"cluster"`
	Filter       *ebpf.MapSpec `ebpf:"filter"`
	FilterChain  *ebpf.MapSpec `ebpf:"filter_chain"`
	Listener     *ebpf.MapSpec `ebpf:"listener"`
	TailCallCtx  *ebpf.MapSpec `ebpf:"tail_call_ctx"`
	TailCallProg *ebpf.MapSpec `ebpf:"tail_call_prog"`
}

// CgroupSockObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type CgroupSockObjects struct {
	CgroupSockPrograms
	CgroupSockMaps
}

func (o *CgroupSockObjects) Close() error {
	return _CgroupSockClose(
		&o.CgroupSockPrograms,
		&o.CgroupSockMaps,
	)
}

// CgroupSockMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type CgroupSockMaps struct {
	Cluster      *ebpf.Map `ebpf:"cluster"`
	Filter       *ebpf.Map `ebpf:"filter"`
	FilterChain  *ebpf.Map `ebpf:"filter_chain"`
	Listener     *ebpf.Map `ebpf:"listener"`
	TailCallCtx  *ebpf.Map `ebpf:"tail_call_ctx"`
	TailCallProg *ebpf.Map `ebpf:"tail_call_prog"`
}

func (m *CgroupSockMaps) Close() error {
	return _CgroupSockClose(
		m.Cluster,
		m.Filter,
		m.FilterChain,
		m.Listener,
		m.TailCallCtx,
		m.TailCallProg,
	)
}

// CgroupSockPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type CgroupSockPrograms struct {
	SockConnect4 *ebpf.Program `ebpf:"sock_connect4"`
}

func (p *CgroupSockPrograms) Close() error {
	return _CgroupSockClose(
		p.SockConnect4,
	)
}

func _CgroupSockClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed cgroupsock_bpfel.o
var _CgroupSockBytes []byte
