package monitor

import (
    "github.com/hako/durafmt"
    "time"
)

// scans the map of peers with their reported block height, and
// then returns the highest reported block height as well as a
// list of peers (more specifically their name) that reported
// exactly this height.
func max(blockHeightMap map[string]uint32) (uint32, []string) {
    var maxV uint32 = 0
    maxHeightPeersMap := make(map[uint32][]string)
    for key, value := range blockHeightMap {
        if value >= maxV {
            list, found := maxHeightPeersMap[value]
            if found {
                maxHeightPeersMap[value] = append(list, key)
            } else {
                maxHeightPeersMap[value] = []string{key}
            }
            maxV = value
        }
    }
    peers, _ := maxHeightPeersMap[maxV]
    return maxV, peers
}

// transforms the given up time into a human readable string.
func getHumanReadableUpTime(upTime time.Duration) string {
    fmt := durafmt.Parse(upTime)
    if fmt != nil {
        return fmt.String()
    } else {
        return "NA"
    }
}
