package iogetter

import (
	"bufio"
	"os"

	"github.com/Alex-Wolf-7/Kisa/plog"
)

type IOListener interface {
	PreInput()
	HandleInput(string) (error, bool)
	EndLoop()
}

type IOPublisher struct {
	reader *bufio.Reader
	ioListener IOListener
}

func NewIOPublisher() *IOPublisher {
	reader := bufio.NewReader(os.Stdin)

	iop := &IOPublisher{
		reader: reader,
		ioListener: nil,
	}

	go iop.Loop()

	return iop
}

func (iop *IOPublisher) Loop() {
	var input string
	var err error
	for {
		if iop.ioListener != nil {
			iop.ioListener.PreInput()
		}
		
		input, err = iop.reader.ReadString('\n')
		if err != nil {
			plog.ErrorfWithBackup("unable to read input: retrying", "unable to read input: %s", err.Error())
			err = nil
		}

		if iop.ioListener != nil {
			err, quit := iop.ioListener.HandleInput(input)
			if err != nil {
				plog.ErrorfWithBackup("error handling command: please try again", "error handling input: %s", err.Error())
			}
			if quit {
				break
			}
		} else {
			continue
		}
	}

	iop.ioListener.EndLoop()
}

func (iop *IOPublisher) Listen(ioListener IOListener) {
	iop.ioListener = ioListener
	iop.ioListener.PreInput()
}

func (iop *IOPublisher) RemoveListener() {
	iop.ioListener = nil
}