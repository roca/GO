// Package worker implements mining, peer updates, and transaction sharing for
// the blockchain.
package worker

import (
	"blockchain/foundation/blockchain/state"
	"sync"
)

// Worker manages the POW workflows for the blockchain.
type Worker struct {
	state        *state.State
	wg           sync.WaitGroup
	shut         chan struct{}
	startMining  chan bool
	cancelMining chan bool
	evHandler    state.EventHandler
}

// Run creates a worker, registers the worker with the state package, and
// starts up all the background processes.
func Run(st *state.State, evHandler state.EventHandler) {
	w := Worker{
		state:        st,
		shut:         make(chan struct{}),
		startMining:  make(chan bool, 1),
		cancelMining: make(chan bool, 1),
		evHandler:    evHandler,
	}

	// Register this worker with the state package.
	st.Worker = &w

	// Update this node before starting any support G's.
	w.Sync()

	// Load the set of operations we need to run.
	operations := []func(){
		w.powOperations,
	}

	// Set waitgroup to match the number of G's we need for the set
	// of operations we have.
	g := len(operations)
	w.wg.Add(g)

	// We don't want to return until we know all the G's are up and running.
	hasStarted := make(chan bool)

	// Start all the operational G's.
	for _, op := range operations {
		go func(op func()) {
			defer w.wg.Done()
			hasStarted <- true
			op()
		}(op)
	}

	// Wait for the G's to report they are running.
	for i := 0; i < g; i++ {
		<-hasStarted
	}
}

// =============================================================================
// These methods implement the state.Worker interface.

// Shutdown terminates the goroutine performing work.
func (w *Worker) Shutdown() {
	w.evHandler("worker: shutdown: started")
	defer w.evHandler("worker: shutdown: completed")

	w.evHandler("worker: shutdown: signal cancel mining")
	w.SignalCancelMining()

	w.evHandler("worker: shutdown: terminate goroutines")
	close(w.shut)
	w.wg.Wait()
}

// SignalStartMining starts a mining operation. If there is already a signal
// pending in the channel, just return since a mining operation will start.
func (w *Worker) SignalStartMining() {
	select {
	case w.startMining <- true:
	default:
	}
	w.evHandler("worker: SignalStartMining: mining signaled")
}

// SignalCancelMining signals the G executing the runMiningOperation function
// to stop immediately.
func (w *Worker) SignalCancelMining() {
	select {
	case w.cancelMining <- true:
	default:
	}
	w.evHandler("worker: SignalCancelMining: MINING: CANCEL: signaled")
}

// =============================================================================

// isShutdown is used to test if a shutdown has been signaled.
func (w *Worker) isShutdown() bool {
	select {
	case <-w.shut:
		return true
	default:
		return false
	}
}
