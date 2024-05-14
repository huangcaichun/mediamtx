// Package h265conf contains a h265 configuration parser.
package h265conf

import "fmt"

const (
	NAL_TYPE_VPS = 32
	NAL_TYPE_SPS = 33
	NAL_TYPE_PPS = 34
)

// HEVCConf is a RTMP H265 configuration.
type Conf struct {
	SPS           []byte
	PPS           []byte
	VPS           []byte
	NalUnitLength int
}

// Unmarshal 从字节序列解码 HEVCConf 结构
func (c *Conf) Unmarshal(buf []byte) error {
	if len(buf) < 24 {
		return fmt.Errorf("invalid size 1")
	}

	pos := 22
	if buf[pos] != 3 {
		return fmt.Errorf("num of arrays != 3 is unsupported")
	}
	pos++
	if buf[pos] != NAL_TYPE_VPS {
		return fmt.Errorf("vps type is unsupported")
	}
	pos++
	vpsCount := int(uint16(buf[pos])<<8 | uint16(buf[pos+1]))
	if vpsCount != 1 {
		return fmt.Errorf("vps count != 1 is unsupported")
	}
	pos += 2
	vpsLen := int(uint16(buf[pos])<<8 | uint16(buf[pos+1]))
	pos += 2
	if (len(buf) - pos) < vpsLen {
		return fmt.Errorf("invalid size 2")
	}
	c.VPS = buf[pos : pos+vpsLen]
	pos += vpsLen

	if buf[pos] != NAL_TYPE_SPS {
		return fmt.Errorf("sps type is unsupported")
	}
	pos++

	spsCount := int(uint16(buf[pos])<<8 | uint16(buf[pos+1]))
	pos += 2
	if spsCount != 1 {
		return fmt.Errorf("sps count != 1 is unsupported")
	}

	spsLen := int(uint16(buf[pos])<<8 | uint16(buf[pos+1]))
	pos += 2
	if (len(buf) - pos) < spsLen {
		return fmt.Errorf("invalid size 2")
	}

	c.SPS = buf[pos : pos+spsLen]
	pos += spsLen

	if (len(buf) - pos) < 3 {
		return fmt.Errorf("invalid size 3")
	}

	if buf[pos] != NAL_TYPE_PPS {
		return fmt.Errorf("pps type is unsupported")
	}
	pos++

	ppsCount := int(uint16(buf[pos])<<8 | uint16(buf[pos+1]))
	pos += 2
	if ppsCount != 1 {
		return fmt.Errorf("pps count != 1 is unsupported")
	}

	ppsLen := int(uint16(buf[pos])<<8 | uint16(buf[pos+1]))
	pos += 2
	if (len(buf) - pos) < ppsLen {
		return fmt.Errorf("invalid size")
	}

	c.PPS = buf[pos : pos+ppsLen]
	return nil
}

// Marshal 将 HEVCConf 结构编码为字节序列
func (c Conf) Marshal() ([]byte, error) {
	var buf []byte
	//todo 需要实现
	return buf, nil
}
