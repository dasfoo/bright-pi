package bpi

import "github.com/dasfoo/i2c"

// BrightPI i2c wrapper. Built from code examples at:
// https://www.pi-supply.com/bright-pi-v1-0-code-examples/
// and register specification at:
// http://www.semtech.com/images/datasheet/sc620.pdf (page 14)
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

// Led color and position
const (
	WhiteTopLeft     byte = 1 << 1
	WhiteBottomLeft       = 1 << 3
	WhiteBottomRight      = 1 << 4
	WhiteTopRight         = 1 << 6
	WhiteAll              = WhiteTopLeft + WhiteBottomLeft + WhiteBottomRight + WhiteTopRight
	IRBottomLeft          = 1 << 0
	IRTopLeft             = 1 << 2
	IRTopRight            = 1 << 5
	IRBottomRight         = 1 << 7
	IRAll                 = IRTopLeft + IRBottomLeft + IRBottomRight + IRTopRight
)

// Max and default levels of Dim and Gain
const (
	MaxDim     = 0x3f
	DefaultDim = 0x01

	MaxGain     = 0x0f
	DefaultGain = 0x08
)

// Power setting for the specified LEDs (others are turned off)
func (p *BrightPI) Power(leds byte) error {
	return p.i2c.WriteByteToReg(p.bpiAddr, 0x00, leds)
}

// Dim one or multiple LEDs (value range 0-MaxDim, default DefaultDim)
func (p *BrightPI) Dim(leds, value byte) error {
	var i byte
	for i = 0; i < 8; i++ {
		if leds&(1<<i) > 0 {
			if err := p.i2c.WriteByteToReg(p.bpiAddr, i+1, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// Gain overall LEDs brightness (value range 0-MaxGain, default DefaultGain)
func (p *BrightPI) Gain(value byte) error {
	return p.i2c.WriteByteToReg(p.bpiAddr, 0x09, value)
}
