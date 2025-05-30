package worker

import "blockchain/foundation/blockchain/peer"

// CORE NOTE: The p2p network is managed by this goroutine. There is
// a single node that is considered the origin node. The defaults in
// main.go represent the origin node. That node must be running first.
// All new peer nodes connect to the origin node to identify all other
// peers on the network. The topology is all nodes having a connection
// to all other nodes. If a node does not respond to a network call,
// they are removed from the peer list until the next peer operation.

// addNewPeers takes the list of known peers and makes sure they are included
// in the nodes list of know peers.
func (w *Worker) addNewPeers(knownPeers []peer.Peer) error {
	w.evHandler("worker: runPeerUpdatesOperation: addNewPeers: started")
	defer w.evHandler("worker: runPeerUpdatesOperation: addNewPeers: completed")

	for _, peer := range knownPeers {

		// Don't add this running node to the known peer list.
		if peer.Match(w.state.Host()) {
			continue
		}

		// Only log when the peer is new.
		if w.state.AddKnownPeer(peer) {
			w.evHandler("worker: runPeerUpdatesOperation: addNewPeers: add peer nodes: adding peer-node %s", peer.Host)
		}
	}

	return nil
}
