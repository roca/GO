package worker

// CORE NOTE: On startup or when reorganizing the chain, the node needs to be
// in sync with the rest of the network. This includes the mempool and
// blockchain database. This operation needs to finish before the node can
// participate in the network.

// Sync updates the peer list, mempool and blocks.
func (w *Worker) Sync() {
	w.evHandler("worker: sync: started")
	defer w.evHandler("worker: sync: completed")

}
