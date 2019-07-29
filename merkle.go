package merkle

import "hash"

// BinaryTreeRoot returns the root of the binary Merkle tree for data.
func BinaryTreeRoot(data [][]byte, fn hash.Hash) []byte {
	n := len(data)
	if (n) == 0 {
		return nil
	}

	// If odd, append copy of last data item.
	if n&1 > 0 {
		data = append(data, data[n-1])
	}

	// Build the binary Merkle tree bottom-up.
	nodes := make([][]byte, treeSize(n))
	start := n - nextPowerOfTwo(uint64(n))
	for i := start; i < start+n; i++ {
		nodes[i] = fn.Sum(data[i-start])
	}
	for i := start - 1; i >= 0; i-- {
		lc := nodes[i*2+1]
		rc := nodes[i*2+2]
		if lc == nil && rc == nil {
			continue
		}
		if rc == nil {
			rc := make([]byte, len(lc))
			copy(rc, lc)
		}
		nodes[i] = fn.Sum(append(lc, rc...))
	}

	return nodes[0]
}

// isPowerOfTwo returns whether n is a power of two.
func isPowerOfTwo(n int) bool {
	return n > 0 && ((n & (n - 1)) == 0)
}

// nextPowerofTwo returns the next power of two greater than n.
func nextPowerOfTwo(n uint64) int {
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return int(n + 1)
}

// treeSize returns the number of nodes in the binary Merkle tree for n data
// items.
func treeSize(n int) int {
	return 2*nextPowerOfTwo(uint64(n)) - 1
}
