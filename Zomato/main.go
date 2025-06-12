package main

import (
	"time"

	"github.com/wangjia184/sortedset"
	"google.golang.org/protobuf/types/known/timestamppb"

	"fmt"
)

type Analytics struct {
	set               *sortedset.SortedSet // maintains all the delivery payment id
	paymentAmountById map[string]int       // key: time(s), value: payment amount
	outstandingAmt    int
	paidAmt           int // increases when there is a payment event
	lastPaidTimeSec   int // last time payment was made in seconds
}

func (a *Analytics) MakePaymentUptoTs(ts int64) {
	startTs := a.lastPaidTimeSec + 1
	endTs := ts
	tsRange := a.set.GetByScoreRange(sortedset.SCORE(startTs), sortedset.SCORE(endTs), nil)

	count := 0
	for _, v := range tsRange {
		// v.Score() is the timestamp in seconds
		id := (v.Value).(string)
		amt, ok := a.paymentAmountById[id]
		if !ok {
			fmt.Println("Payment ID not found:", id)
		}
		count += amt
	}

	// ts is in seconds
	a.paidAmt += a.paymentAmountById[fmt.Sprintf("%d", ts)]
	a.outstandingAmt -= a.paymentAmountById[fmt.Sprintf("%d", ts)]
	a.lastPaidTimeSec = int(ts)
	delete(a.paymentAmountById, fmt.Sprintf("%d", ts))
}

func main() {

	ts1 := timestamppb.New(time.Now().Add(-time.Second * 10)).Seconds
	ts2 := timestamppb.New(time.Now().Add(-time.Second * 8)).Seconds
	ts3 := timestamppb.New(time.Now().Add(-time.Second * 6)).Seconds
	ts4 := timestamppb.New(time.Now().Add(-time.Second * 5)).Seconds
	ts5 := timestamppb.New(time.Now().Add(-time.Second * 3)).Seconds

	set := sortedset.New()

	// fill in new node
	set.AddOrUpdate("a", sortedset.SCORE(ts1), "Kelly")
	set.AddOrUpdate("b", sortedset.SCORE(ts2), "Staley")
	set.AddOrUpdate("c", sortedset.SCORE(ts2), "Jordon")
	set.AddOrUpdate("d", sortedset.SCORE(ts2), "Park")
	set.AddOrUpdate("e", sortedset.SCORE(ts3), "Albert")
	set.AddOrUpdate("f", sortedset.SCORE(ts4), "Lyman")
	set.AddOrUpdate("g", sortedset.SCORE(ts5), "Singleton")
	set.AddOrUpdate("h", sortedset.SCORE(ts5), "Audrey")

	output := set.GetByScoreRange(sortedset.SCORE(ts3), sortedset.SCORE(ts5), nil)
	for _, v := range output {
		fmt.Printf("%d: %s\n", v.Score(), v.Value)
	}
}
