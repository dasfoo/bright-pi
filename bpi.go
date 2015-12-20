package bpi

// Code examples taken from
// https://www.pi-supply.com/bright-pi-v1-0-code-examples/

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

type BrightPI struct {
	i2c     embd.I2CBus
	bpiAddr byte
}

const DEFAULT_BPI_ADDR = 0x70

func NewBrightPI(i2c embd.I2CBus, bpiAddr byte) *BrightPI {
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

func (p *BrightPI) All(white bool, ir bool) {
	var cmd byte
	if white {
		cmd += 0xa5
	}
	if ir {
		cmd += 0x5a
	}
	p.i2c.WriteByteToReg(p.bpiAddr, 0x00, cmd)
}
