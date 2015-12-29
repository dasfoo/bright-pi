package bpi

// Code examples taken from
// https://www.pi-supply.com/bright-pi-v1-0-code-examples/

import (
	"github.com/dasfoo/i2c"
)

// BrightPI i2c wrapper. Built from code examples at:
// https://www.pi-supply.com/bright-pi-v1-0-code-examples/
type BrightPI struct {
	i2c     *i2c.Bus
	bpiAddr byte
}

// DefaultBPiAddress is a default BrightPI Address
const DefaultBPiAddress = 0x70

// NewBrightPI creates an instance of BrightPI and sets fields
func NewBrightPI(i2c *i2c.Bus, bpiAddr byte) *BrightPI {
	return &BrightPI{i2c: i2c, bpiAddr: bpiAddr}
}

/*
* Then you need to turn the gain up to full using:
*
* sudo i2cset -y 1 0x70 0x09 0x0f
*
* Then you need to turn brightness of individual LEDs up. For white:
*
* sudo i2cset -y 1 0x70 0x02 0x32
* sudo i2cset -y 1 0x70 0x04 0x32
* sudo i2cset -y 1 0x70 0x05 0x32
* sudo i2cset -y 1 0x70 0x07 0x32
*
* For IR:
*
* sudo i2cset -y 1 0x70 0x01 0x32
* sudo i2cset -y 1 0x70 0x03 0x32
* sudo i2cset -y 1 0x70 0x06 0x32
* sudo i2cset -y 1 0x70 0x08 0x32
* */

// All lights configured at once
func (p *BrightPI) All(white, ir bool) error {
	var cmd byte
	if white {
		cmd += 0xa5
	}
	if ir {
		cmd += 0x5a
	}
	return p.i2c.WriteByteToReg(p.bpiAddr, 0x00, cmd)
}
