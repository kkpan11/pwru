// SPDX-License-Identifier: Apache-2.0
// Copyright 2021-2022 Authors of Cilium

//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || ppc64le || riscv64 || wasm
// +build 386 amd64 arm arm64 loong64 mips64le ppc64le riscv64 wasm

package byteorder

import (
	"encoding/binary"
	"math/bits"
)

var Native binary.ByteOrder = binary.LittleEndian

func HostToNetwork16(u uint16) uint16 { return bits.ReverseBytes16(u) }
func HostToNetwork32(u uint32) uint32 { return bits.ReverseBytes32(u) }
func NetworkToHost16(u uint16) uint16 { return bits.ReverseBytes16(u) }
func NetworkToHost32(u uint32) uint32 { return bits.ReverseBytes32(u) }
