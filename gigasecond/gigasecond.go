
package gigasecond


import "time"


// AddGigasecond returns the time when someone has lived 1 billion seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
	//return t
}
