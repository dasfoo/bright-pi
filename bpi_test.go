package bpi_test

import (
	"testing"

	"github.com/dasfoo/bright-pi"
	"github.com/dasfoo/i2c"
)

type i2cDevice struct {
	regs    [256]byte
	address byte
	t       *testing.T
}

func (d *i2cDevice) Close() error {
	d.address = 0
	return nil
}

func (d *i2cDevice) WriteByteToReg(addr, reg, value byte) error {
	if addr == 0 || addr != d.address {
		d.t.Error("Invalid address", addr, "expected", d.address)
	}
	d.regs[reg] = value
	return nil
}

func (d *i2cDevice) ReadByteFromReg(addr, reg byte) (byte, error) {
	d.t.Error("Unexpected read (1) from", addr, ":", reg)
	return 0, nil
}

func (d *i2cDevice) ReadWordFromReg(addr, reg byte) (uint16, error) {
	d.t.Error("Unexpected read (2) from", addr, ":", reg)
	return 0, nil
}

func (d *i2cDevice) ReadSliceFromReg(addr, reg byte, data []byte) (int, error) {
	d.t.Error("Unexpected read of ", len(data), "bytes from ", addr, ":", reg)
	return 0, nil
}

func (d *i2cDevice) SetLogger(_ i2c.Logger) {
}

func TestNewBPi(t *testing.T) {
	dev := &i2cDevice{address: bpi.DefaultAddress, t: t}
	bpi.NewBrightPI(dev, bpi.DefaultAddress)
}
