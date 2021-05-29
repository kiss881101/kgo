// +build darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_Darwin_IsMac(t *testing.T) {
	res := KOS.IsMac()
	assert.True(t, res)
}

func BenchmarkOS_Darwin_IsMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsMac()
	}
}

func TestOS_Darwin_MemoryUsage(t *testing.T) {
	var used, free, total uint64

	used, free, total = KOS.MemoryUsage(true)
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Darwin_MemoryUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage(true)
	}
}

func TestOS_Darwin_CpuUsage(t *testing.T) {
	var user, idle, total uint64
	user, idle, total = KOS.CpuUsage()
	assert.GreaterOrEqual(t, int(user), 0)
	assert.GreaterOrEqual(t, int(idle), 0)
	assert.GreaterOrEqual(t, int(total), 0)
}

func BenchmarkOS_Darwin_CpuUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.CpuUsage()
	}
}

func TestOS_Darwin_DiskUsage(t *testing.T) {
	var used, free, total uint64
	used, free, total = KOS.DiskUsage("/")
	dumpPrint("-----------TestOS_Linux_DiskUsage:", used, free, total)
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Darwin_DiskUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.DiskUsage("/")
	}
}
