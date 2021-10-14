/*Package hash supports hash function
*
	 hash 方法
*
*/
package hash

import (
	"github.com/spaolacci/murmur3"
)

var (

	// DefualtTotalWeight is the defualt total weight each layer
	DefualtTotalWeight uint32 = 100
)

// Hash is to calculate the hash value by hashkey and layerID, then modulo total.
func Hash(hashkey, layerID string, total uint32) uint32 {

	// get total weight
	if total == 0 {
		total = DefualtTotalWeight
	}

	// set seed
	var seed uint32 = total

	// [1,total]
	return murmur3.Sum32WithSeed([]byte(hashkey+layerID), seed)%total + 1
}
