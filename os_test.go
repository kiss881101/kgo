package kgo

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestOS_Pwd(t *testing.T) {
	res := KOS.Pwd()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Pwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Pwd()
	}
}

func TestOS_Getcwd_Chdir(t *testing.T) {
	var ori, res string
	var err error

	ori, err = KOS.Getcwd()
	assert.Nil(t, err)
	assert.NotEmpty(t, ori)

	//切换目录
	err = KOS.Chdir(dirTdat)
	assert.Nil(t, err)

	res, err = KOS.Getcwd()
	assert.Nil(t, err)

	//返回原来目录
	err = KOS.Chdir(ori)
	assert.Nil(t, err)
	assert.Equal(t, KFile.AbsPath(res), KFile.AbsPath(dirTdat))
}

func BenchmarkOS_Getcwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Getcwd()
	}
}

func BenchmarkOS_Chdir(b *testing.B) {
	b.ResetTimer()
	dir := KOS.Pwd()
	for i := 0; i < b.N; i++ {
		_ = KOS.Chdir(dir)
	}
}

func TestOS_LocalIP(t *testing.T) {
	res, err := KOS.LocalIP()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_LocalIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.LocalIP()
	}
}

func TestOS_OutboundIP(t *testing.T) {
	res, err := KOS.OutboundIP()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_OutboundIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.OutboundIP()
	}
}

func TestOS_IsPublicIPv4(t *testing.T) {
	var res bool

	res = KOS.IsPublicIPv4(net.ParseIP(localIp))
	assert.False(t, res)

	res = KOS.IsPublicIPv4(net.ParseIP(lanIp))
	assert.False(t, res)

	res = KOS.IsPublicIPv4(net.ParseIP(googleIpv4))
	assert.True(t, res)

	res = KOS.IsPublicIPv4(net.ParseIP(googleIpv6))
	assert.False(t, res)
}

func BenchmarkOS_IsPublicIPv4(b *testing.B) {
	b.ResetTimer()
	ip := net.ParseIP(publicIp1)
	for i := 0; i < b.N; i++ {
		KOS.IsPublicIPv4(ip)
	}
}

func TestOS_GetIPs(t *testing.T) {
	res := KOS.GetIPs()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_GetIPs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetIPs()
	}
}

func TestOS_GetMacAddrs(t *testing.T) {
	res := KOS.GetMacAddrs()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_GetMacAddrs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetMacAddrs()
	}
}

func TestOS_Hostname_GetIpByHostname(t *testing.T) {
	var res string
	var err error

	res, err = KOS.Hostname()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetIpByHostname(res)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetIpByHostname(tesIp2)
	assert.Empty(t, res)

	res, err = KOS.GetIpByHostname(strHello)
	assert.NotNil(t, err)
}

func BenchmarkOS_Hostname(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Hostname()
	}
}

func BenchmarkOS_GetIpByHostname(b *testing.B) {
	b.ResetTimer()
	host, _ := KOS.Hostname()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetIpByHostname(host)
	}
}

func TestOS_GetIpsByDomain(t *testing.T) {
	var res []string
	var err error

	res, err = KOS.GetIpsByDomain(tesDomain30)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetIpsByDomain(strHello)
	assert.NotNil(t, err)
}

func BenchmarkOS_GetIpsByDomain(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetIpsByDomain(tesDomain30)
	}
}

func TestOS_GetHostByIp(t *testing.T) {
	var res string
	var err error

	res, err = KOS.GetHostByIp(localIp)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetHostByIp(strHello)
	assert.NotNil(t, err)
	assert.Empty(t, res)
}

func BenchmarkOS_GetHostByIp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetHostByIp(localIp)
	}
}
