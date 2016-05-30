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

func TestDim(t *testing.T) {
	dev := &i2cDevice{address: bpi.DefaultAddress, t: t}
	b := bpi.NewBrightPI(dev, bpi.DefaultAddress)
	if b.Dim(bpi.WhiteAll, bpi.MaxDim) != nil {
		t.Error("Dimming the lights returned an error")
	}
	for i, reg := range []byte{1, 3, 4, 6} {
		if dev.regs[reg+1] != bpi.MaxDim {
			t.Error("Incorrect Dim value for light", i,
				"expected", bpi.MaxDim, "got", dev.regs[reg+1])
		}
	}
}

func TestGain(t *testing.T) {
	dev := &i2cDevice{address: bpi.DefaultAddress, t: t}
	b := bpi.NewBrightPI(dev, bpi.DefaultAddress)
	if b.Gain(bpi.MaxGain) != nil {
		t.Error("Setting lights gain returned an error")
	}
	if dev.regs[9] != bpi.MaxGain {
		t.Error("Invalid Gain value, expected", bpi.MaxGain, "got", dev.regs[9])
	}
}

func TestPower(t *testing.T) {
	dev := &i2cDevice{address: bpi.DefaultAddress, t: t}
	b := bpi.NewBrightPI(dev, bpi.DefaultAddress)
	if b.Power(bpi.IRAll) != nil {
		t.Error("Setting lights power returned an error")
	}
	if dev.regs[0] != 165 {
		t.Error("Invalid power value", dev.regs[0])
	}
	if b.Sleep() != nil {
		t.Error("Setting sleep mode failed")
	}
	if dev.regs[0] != 0 {
		t.Error("Expected Sleep() to shut down power, but got", dev.regs[0])
	}
}
