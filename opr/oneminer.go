// Copyright (c) of parts are held by the various contributors (see the CLA)
// Licensed under the MIT License. See LICENSE file in the project root for full license information.
package opr

import (
	"fmt"
	"github.com/FactomProject/factom"
	"github.com/pegnet/pegnet/common"
	"github.com/zpatrick/go-config"
)

func OneMiner(verbose bool, config *config.Config, monitor *common.FactomdMonitor, grader *Grader, miner int) {
	alert := monitor.GetAlert()
	gAlert := grader.GetAlert()

	mining := false
	var opr *OraclePriceRecord
	var err error
	for {
		fds := <-alert
		common.Logf("miner", "Alert: miner%02d dbht %d minute %d", miner, fds.Dbht, fds.Minute)
		switch fds.Minute {
		case 1:
			if !mining {
				mining = true
				opr, err = NewOpr(miner, fds.Dbht, config, gAlert)
				if err != nil {
					panic(fmt.Sprintf("Error creating an OPR.  Likely a config file issue: %v\n", err))
				}
				go opr.Mine(verbose)
			}
		case 9:
			if mining {
				opr.StopMining <- 0

				common.Do(func() {
					data, ok := opr.Entry.MarshalBinary()
					if ok != nil {
						panic(fmt.Sprint("Can't json marshal the opr: ", ok))
					}
					ostr := opr.String()
					common.Logf("OPR-Rec", "OPR:        miner%02d entrySize %d", miner, len(data))
					common.Logf("OPR-Rec", "OPR Record: miner%02d %s", miner, ostr)
				})

				mining = false

				writeMiningRecord(opr)
			}
		}
	}
}

func writeMiningRecord(opr *OraclePriceRecord) {

	var err1, err2 error
	for i := 0; i < 100; i++ {
		if i == 0 || err1 != nil {
			_, err1 = factom.CommitEntry(opr.Entry, opr.EC)
		}
		if i == 0 || err2 != nil {
			_, err2 = factom.RevealEntry(opr.Entry)
		}
		if err1 == nil && err2 == nil {
			break
		}
	}
}
