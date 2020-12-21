package beamer

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
)

type Beamer struct {
	device string
	baud   uint
}

func NewBeamer(device string, baud uint) *Beamer {
	return &Beamer{
		device,
		baud,
	}
}

func (b *Beamer) openSendCmdAndClose(cmd string) error {
	rawPort, err := openUart(b.device, b.baud)	
	if err != nil {
		return err
	}
	_, err = rawPort.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}

func (b *Beamer) On() error {
	return b.openSendCmdAndClose("\r*pow=on#\r")
}

func (b *Beamer) Off() error {
	return b.openSendCmdAndClose("\r*pow=off#\r")
}

func (b *Beamer) VolumeUp() error {
	return b.openSendCmdAndClose("\r*vol=+#\r")
}

func (b *Beamer) VolumeDown() error {
	return b.openSendCmdAndClose("\r*vol=-#\r")
}

func (b *Beamer) BrightnessUp(n uint) error {
	err := b.openSendCmdAndClose("\r*bri=+#\r")
	if err == nil {
		if n > 0 {
			return b.BrightnessUp(n - 1)
		}
	}
	return err
}

func (b *Beamer) BrightnessDown(n uint) error {
	err := b.openSendCmdAndClose("\r*bri=-#\r")
	if err == nil {
		if n > 0 {
			return b.BrightnessUp(n - 1)
		}
	}
	return err
}

func openUart(device string, baud uint) (io.ReadWriteCloser, error){
		return serial.Open(serial.OpenOptions{
			PortName:        device,
			BaudRate:        baud,
			DataBits:        8,
			StopBits:        1,
			MinimumReadSize: 4,
		})
}
