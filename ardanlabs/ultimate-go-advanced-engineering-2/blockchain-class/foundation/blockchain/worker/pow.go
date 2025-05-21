package worker

// CORE NOTE: The POW mining operation is managed by this function which runs on
// it's own goroutine. When a startMining signal is received (mainly because a
// wallet transaction was received) a block is created and then the POW operation
// starts. This operation can be cancelled if a proposed block is received and
// is validated.

// powOperations handles mining.
func (w *Worker) powOperations() {
	w.evHandler("worker: powOperations: G started")
	defer w.evHandler("worker: powOperations: G completed")

	for {
		select {
		case <-w.startMining:
			if !w.isShutdown() {
				w.runPowOperation()
			}
		case <-w.shut:
			w.evHandler("worker: powOperations: received shut signal")
			return
		}
	}
}
