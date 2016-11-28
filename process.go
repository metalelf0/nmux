package nmux

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/neovim/go-client/nvim"
	"github.com/metalelf0/nmux/screen"
)

const (
	MsgRequest = 0
	MsgReply
	MsgNotification
)

var ErrFirstArgString = errors.New("first item must be a string")
var ErrEmpty = errors.New("empty")

// TODO: Multiple nvim process management.
var procs struct {
	mu    sync.Mutex
	id    int
	procs []*Process
}

func addProcess(proc *Process) {
	procs.mu.Lock()
	defer procs.mu.Unlock()
	procs.procs = append(procs.procs, proc)
}

func removeProcess(proc *Process) {
	procs.mu.Lock()
	defer procs.mu.Unlock()

	curProcs := procs.procs
	procs.procs = procs.procs[0:0]

	for _, p := range curProcs {
		if p != proc {
			procs.procs = append(procs.procs, p)
		}
	}
}

type Process struct {
	*screen.Screen
	ID      int
	Deadman <-chan int
	nvim    *nvim.Nvim
}

type msg struct {
	code uint8
	name string
	args []interface{}
}

func NewProcess(cwd string, width, height int) (*Process, error) {
	n, err := nvim.NewEmbedded(&nvim.EmbedOptions{
		Env: os.Environ(),
		Dir: cwd,
		Logf: func(msg string, args ...interface{}) {
			log.Println("Embedded Log:", msg, args)
		},
	})

	if err != nil {
		return nil, err
	}

	procs.id++
	deadman := make(chan int)
	proc := &Process{
		ID:      procs.id,
		nvim:    n,
		Deadman: deadman,
		Screen:  screen.NewScreen(width, height),
	}

	addProcess(proc)

	go func() {
		if err := proc.nvim.Serve(); err != nil {
			fmt.Println("RPC Err:", err)
		}
		proc.nvim = nil
		removeProcess(proc)
		close(deadman)
		fmt.Println("Dead")
	}()

	n.RegisterHandler("redraw", proc.Screen.RedrawHandler)

	if err := n.AttachUI(proc.Size.X, proc.Size.Y, map[string]interface{}{
		"rgb":                true,
		"popupmenu_external": false,
	}); err != nil {
		return nil, err
	}

	return proc, nil
}

func (p *Process) IsRunning() bool {
	return p.nvim != nil
}

func (p *Process) Input(keys string) (int, error) {
	return p.nvim.Input(keys)
}

func (p *Process) Resize(w, h int) error {
	return p.nvim.TryResizeUI(w, h)
}
