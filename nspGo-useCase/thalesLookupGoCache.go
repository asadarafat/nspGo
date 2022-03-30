package nspgousecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
)

func ThalesLookupWithGoCache(minIteration int, maxIteration int) {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	// c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	// fmt.Println(c.Items())

	// dataBase := `{
	// 		"pwLabel01": "1000001",
	// 		"pwLabel02": "1000002",
	// 		"pwLabel03": "1000003",
	// 		"pwLabel04": "1000004",
	// 		"pwLabel05": "1000005"
	// }`
	// fmt.Println(dataBase)

	// bytesDataBase := []byte(dataBase)
	// var jOut map[string]string
	// json.Unmarshal(bytesDataBase, &jOut)
	// // fmt.Println(jOut["pwLabel01"])

	// var cacheItem map[string]cache.Item

	// cNewForm := cache.NewFrom(5*time.Minute, 10*time.Minute, cacheItem)

	// fmt.Println(cNewForm.Items())

	// init DB for PW labels
	log.Info("Initialize DB for PW labels")
	minPwLabelId := 1
	maxPwLabelId := 1000000
	barInitDbPwLabel := progressbar.Default(int64((maxPwLabelId + 1) - minPwLabelId))
	for i := minPwLabelId; i <= maxPwLabelId; i++ {
		pwLabelId := "pwLabel" + strconv.Itoa(1000000+i)
		pwLabelIdValue := 1000000 + i
		c.Set(pwLabelId, pwLabelIdValue, cache.NoExpiration)
		barInitDbPwLabel.Add(1)
	}

	// init DB for MPLS labels
	log.Info("Initialize DB for MPLS labels")
	minMplsLabelId := 1
	maxMplsLabelId := 350000
	barInitDbMplsLabel := progressbar.Default(int64((maxMplsLabelId + 1) - minMplsLabelId))
	for i := minMplsLabelId; i <= maxMplsLabelId; i++ {
		MplsLabelId := "mplsLabel" + strconv.Itoa(1000000+i)
		MplsLabelValue := 1000000 + i
		c.Set(MplsLabelId, MplsLabelValue, cache.NoExpiration)
		barInitDbMplsLabel.Add(1)
	}

	// init DB for Uplink feeder id to geenerate IP Address and VxLan Network Identifier (VNI)
	log.Info("Initialize DB for Uplink feeder id to generate IP Address and VxLan Network Identifier (VNI)")
	minUplinkFeederId := 1
	maxUplinkFeederId := 15000
	barUplinkFeederId := progressbar.Default(int64((maxUplinkFeederId + 1) - minUplinkFeederId))
	for i := minUplinkFeederId; i <= maxUplinkFeederId; i++ {
		UplinkFeederId := "IpVniId" + strconv.Itoa(1000000+i)
		UplinkFeederValue := 1000000 + i
		c.Set(UplinkFeederId, UplinkFeederValue, cache.NoExpiration)
		barUplinkFeederId.Add(1)
	}

	// init DB for Uplink feeder carrier id to geenerate IP Address and VxLan Network Identifier (VNI)
	log.Info("Initialize DB for Uplink feeder carrier id to generate IP Address and VxLan Network Identifier (VNI)")
	minUplinkFeederCarrierId := 1
	maxUplinkFeederCarrierId := 15000
	barUplinkFeederCarrierId := progressbar.Default(int64((maxUplinkFeederCarrierId + 1) - minUplinkFeederCarrierId))
	for i := minUplinkFeederId; i <= maxUplinkFeederId; i++ {
		UplinkFeederCarrierId := "UplinkFeederCarrierId" + strconv.Itoa(1000000+i)
		minUplinkFeederCarrierValue := 1000000 + i
		c.Set(UplinkFeederCarrierId, minUplinkFeederCarrierValue, cache.NoExpiration)
		barUplinkFeederCarrierId.Add(1)
	}

	// Look-up on the couple <Downlink feeder id, Downlink feeder carrier id> to retrieve  eth interface and VLAN Id in a table of 38400 entries (300 sat x 4 if x 32 carriers)
	log.Info("Initialize DB for Downlink Feeder id to generate Ethernet Port and VLAN Id - Feeder")
	minDownlinkFeederId := 1
	maxDownlinkFeederId := 38400
	barDownlinkFeederId := progressbar.Default(int64((maxDownlinkFeederId + 1) - minDownlinkFeederId))
	for i := minDownlinkFeederId; i <= maxDownlinkFeederId; i++ {
		DownlinkFeederId := "DownlikFeederId" + strconv.Itoa(1000000+i)
		DownlinkFeederValue := 1000000 + i
		c.Set(DownlinkFeederId, DownlinkFeederValue, cache.NoExpiration)
		barDownlinkFeederId.Add(1)
	}

	// Look-up on the couple <Downlink feeder id, Downlink feeder carrier id> to retrieve  eth interface and VLAN Id in a table of 38400 entries (300 sat x 4 if x 32 carriers)
	log.Info("Initialize DB for Downlink Feeder Carrier id to generate Ethernet Port and VLAN Id - Feeder")
	minDownlinkFeederCarrierId := 1
	maxDownlinkFeederCarrierId := 38400
	barDownlinkFeederCarrierId := progressbar.Default(int64((maxDownlinkFeederCarrierId + 1) - minDownlinkFeederCarrierId))
	for i := minDownlinkFeederCarrierId; i <= maxDownlinkFeederCarrierId; i++ {
		DownlinkFeederCarrierId := "DownlikFeederCarrierId" + strconv.Itoa(1000000+i)
		DownlinkFeederCarrierValue := 1000000 + i
		c.Set(DownlinkFeederCarrierId, DownlinkFeederCarrierValue, cache.NoExpiration)
		barDownlinkFeederCarrierId.Add(1)
	}

	// init DB for Logical Beam Id to generte eth interface and VLAN Id
	log.Info("Initialize DB for Logical Beam Id to generte Eth interface and VLAN Id")
	minLogicalBeamId := 1
	maxLogicalBeamId := 10000
	barLogicalBeamId := progressbar.Default(int64((maxLogicalBeamId + 1) - minLogicalBeamId))
	for i := minLogicalBeamId; i <= maxLogicalBeamId; i++ {
		LogicalBeamId := "LogicalBeamId" + strconv.Itoa(1000000+i)
		LogicalBeamValue := 1000000 + i
		c.Set(LogicalBeamId, LogicalBeamValue, cache.NoExpiration)
		barLogicalBeamId.Add(1)
	}

	// init DB for Downlink User Carrier Id to generte eth interface and VLAN Id
	log.Info("Initialize DB for Downlink User Carrier to generte Eth interface and VLAN Id")
	minDownlinkUserCarrierId := 1
	maxDownlinkUserCarrierId := 10000
	barDownlinkUserCarrierId := progressbar.Default(int64((maxDownlinkUserCarrierId + 1) - minDownlinkUserCarrierId))
	for i := minDownlinkUserCarrierId; i <= maxDownlinkUserCarrierId; i++ {
		DownlinkUserCarrierId := "LogicalBeamId" + strconv.Itoa(1000000+i)
		DownlinkUserCarrierValue := 1000000 + i
		c.Set(DownlinkUserCarrierId, DownlinkUserCarrierValue, cache.NoExpiration)
		barDownlinkUserCarrierId.Add(1)
	}

	// Lookup Performed

	forwardConstId := 0
	returnConstId := 0
	// uplinkFeederId := 0
	// uplinkFeederCarrierId := 0
	// downlinkFeederId := 0
	// downlinkFeederCarrierId := 0
	// logicalBeamId := 0
	// downlinkUserCarrierId := 0

	// 	·      A service id
	// ·       A forward const id
	// ·       A return const id
	// ·       An Uplink feeder id
	// ·       An Uplink feeder carrier id
	// ·       A Downlink feeder id
	// ·       A Downlink feeder carrier id
	// ·       A logical beam id
	// ·       A Downlink user carrier id

	log.Info("Lookup Interation Start")
	listOfLookupTimePerIteration := []time.Duration{}
	barLookUpProgress := progressbar.Default(int64((maxPwLabelId + 1) - minPwLabelId))
	for i := minIteration; i <= maxIteration; i++ {
		startLookupTime := time.Now()

		pwLabelId := "pwLabel" + strconv.Itoa(1000000+i)
		// foo, found := c.Get(pwLabelId) // If IO enabled, it will add more execution time
		// if found {
		// 	fmt.Println(foo)
		// }

		// Look-up on <service id> to retrieve  2 PW labels in a table of 1 M (instead of 3,66M previously) entries
		if forwardConstId == 0 {
			c.Get(pwLabelId)
			c.Get(pwLabelId)

			// If const id <> 0, Look-up on <forward const id> to retrieve 1 MPLS Transport Label in a table of 350 000 entries (220 sat x 220 sat x  7 paths)
			if i > 350000 {
				k := 350000
				if forwardConstId == 0 {
					MplsLabelId := "mplsLabel" + strconv.Itoa(1000000+k)
					c.Get(MplsLabelId)
				}
			}

			// If const id <> 0, Look-up on <return const id> to retrieve 1 MPLS Transport Label in a table of 350 000 entries (220 sat x 220 sat x  7 paths)
			if i > 350000 {
				k := 350000
				if returnConstId == 0 {
					MplsLabelId := "mplsLabel" + strconv.Itoa(1000000+k)
					c.Get(MplsLabelId)
				}
			}

			// Look-up on the couple <Uplink feeder id, Uplink feeder carrier id> to retrieve the IP@ and VxLan Network Identifier (VNI) in a table of 15000 entries (max 15000 carriers in the system)
			if i > 15000 {
				l := 15000
				UplinkFeederId := "UplinkFeederId" + strconv.Itoa(1000000+l)
				c.Get(UplinkFeederId)

				UplinkFeederCarrierId := "UplinkFeederCarrierId" + strconv.Itoa(1000000+l)
				c.Get(UplinkFeederCarrierId)
			}

			// Look-up on the couple <Downlink feeder id, Downlink feeder carrier id> to retrieve  eth interface and VLAN Id in a table of 38400 entries (300 sat x 4 if x 32 carriers)
			if i > 38400 {
				m := 38400
				DownlinkFeederId := "DownlinkFeederId" + strconv.Itoa(1000000+m)
				c.Get(DownlinkFeederId)

				DownlinkFeederCarrierId := "DownlinkFeederCarrierId" + strconv.Itoa(1000000+m)
				c.Get(DownlinkFeederCarrierId)
			}

			// Look-up on the couple <logical beam id, Downlink user carrier id> to retrieve eth interface and VLAN Id in a table of 10000 entries (24 ts x 24 beam x 16 carriers)
			if i > 10000 {
				n := 10000
				LogicalBeamId := "LogicalBeamId" + strconv.Itoa(1000000+n)
				c.Get(LogicalBeamId)

				DownlinkUserCarrierId := "DownlinkUserCarrierId" + strconv.Itoa(1000000+n)
				c.Get(DownlinkUserCarrierId)
			}

			listOfLookupTimePerIteration = append(listOfLookupTimePerIteration, (time.Since(startLookupTime)))
			barLookUpProgress.Add(1)

		}
	}
	var totalElapsedMidPoint time.Duration
	for w := 0; w <= len(listOfLookupTimePerIteration)/2; w++ {
		totalElapsedMidPoint += listOfLookupTimePerIteration[w]
	}

	//Find Total Elapsed Time
	var totalElapsed time.Duration
	for _, v := range listOfLookupTimePerIteration {
		totalElapsed += v
	}

	//Find Min Max in Total Elapsed List
	var max time.Duration = listOfLookupTimePerIteration[0]
	var min time.Duration = listOfLookupTimePerIteration[0]
	for _, value := range listOfLookupTimePerIteration {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	// Find Total Elapsed Time at MidPoint

	log.Info("Min Elapsed Time Per Iteration (second): ", min.Seconds())
	log.Info("Max Elapsed Time Per Iteration(second): ", max.Seconds())
	log.Info("Total Iteration: ", maxIteration-minIteration+1)
	log.Info("Total Elapsed Time MidPoint at "+strconv.Itoa(len(listOfLookupTimePerIteration)/2)+" Iteration (seconds): ", totalElapsedMidPoint)
	log.Info("Total Elapsed Time(seconds): ", totalElapsed)
	// fmt.Println(c.Items())
}
