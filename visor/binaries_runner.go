// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package visor

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
	"syscall"

	"code.vegaprotocol.io/vega/visor/config"
	"code.vegaprotocol.io/vega/visor/utils"

	"golang.org/x/sync/errgroup"
)

type BinariesRunner struct {
	mut        sync.RWMutex
	running    map[string]*exec.Cmd
	binsFolder string
}

func NewBinariesRunner(binsFolder string) *BinariesRunner {
	return &BinariesRunner{
		binsFolder: binsFolder,
		running:    map[string]*exec.Cmd{},
	}
}

func (r *BinariesRunner) Run(ctx context.Context, binaries []config.BinaryConfig) chan error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, bin := range binaries {
		bin := bin

		eg.Go(func() error {
			binPath := path.Join(r.binsFolder, bin.Path)
			if err := utils.EnsureBinary(binPath); err != nil {
				return fmt.Errorf("failed to locate binary %s %v: %w", binPath, bin.Args, err)
			}

			cmd := exec.CommandContext(ctx, binPath, bin.Args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Start(); err != nil {
				return fmt.Errorf("failed to execute binary %s %v: %w", binPath, bin.Args, err)
			}

			// Ensures that if one binary failes all of them are killed
			go func() {
				<-ctx.Done()
				if cmd.Process == nil {
					return
				}

				log.Printf("killing binary %s %v", binPath, bin.Args)
				if err := cmd.Process.Kill(); err != nil {
					log.Printf("failed to kill binary %s %v: %s", binPath, bin.Args, err)
				}
			}()

			r.mut.Lock()
			r.running[binPath] = cmd
			r.mut.Unlock()

			defer func() {
				r.mut.Lock()
				delete(r.running, binPath)
				r.mut.Unlock()
			}()

			if err := cmd.Wait(); err != nil {
				return fmt.Errorf("failed to execute binary %s %v: %w", binPath, bin.Args, err)
			}

			return nil
		})
	}

	errChan := make(chan error)

	go func() {
		err := eg.Wait()
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func (r *BinariesRunner) signal(signal syscall.Signal) error {
	r.mut.RLock()
	defer r.mut.RUnlock()

	var err error
	for _, c := range r.running {
		err = c.Process.Signal(signal)
		if err != nil {
			log.Printf("failed to signal running binary %s: %s", c.Path, err)
		}
	}

	return err
}

func (r *BinariesRunner) Stop() error {
	return r.signal(syscall.SIGTERM)
}

func (r *BinariesRunner) Kill() error {
	return r.signal(syscall.SIGKILL)
}
