package calculations

import (
	// "fmt"
	log "github.com/sirupsen/logrus"
)

func Calculate(val int64, flag bool) int64 {
	if flag {
		log.WithFields(log.Fields{
			"Number":  val,
			"Logging": flag,
		}).Info("Starting calculations...")
		// fmt.Println("Starting calculations...\nCalculate ", val, "!")
		if val-1 > 0 {
			// fmt.Println("Caltulations complete!")
			log.Info("Calculations complete!")
			return val * Calculate(val-1, false)
		} else {
			return val
		}
	} else {
		if val-1 > 0 {
			return val * Calculate(val-1, false)
		} else {
			return val
		}
	}
}
