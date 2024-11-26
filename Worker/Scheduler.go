package Worker

import (
	"2024_akutansi_project/Dependencies"
	"2024_akutansi_project/Routes/Di"
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"os"
	"strconv"
	"strings"
	"time"
)

func InitScheduler(deps *Dependencies.Dependency) {
	// Push Notification Scheduler
	if func() bool {
		scheduler := os.Getenv("USE_SCHEDULER_PUSH_NOTIFICATION")
		b, _ := strconv.ParseBool(scheduler)
		return b
	}() {
		go func(ctx context.Context) {
			jobHandler := func() {
				fmt.Println("Start [JOB] :: Send Push Notification")

				service := Di.DIWorker(deps.DB)

				if err := service.SendPushNotification(ctx); err != nil {
					fmt.Printf("Failed [JOB] :: Send Push Notification, got err := %v\n", err)
					return
				}

				fmt.Println("Success [JOB] :: Send Push Notification")
			}

			// Retrieve interval and unit from environment variables
			intervalStr := os.Getenv("SCHEDULER_INTERVAL") // E.g., "1"
			unit := os.Getenv("SCHEDULER_UNIT")            // E.g., "minute"

			interval, err := strconv.Atoi(intervalStr)
			if err != nil || interval <= 0 {
				fmt.Println("Invalid or missing SCHEDULER_INTERVAL, defaulting to 1 minute")
				interval = 1
				unit = "minute"
			}

			sh := gocron.NewScheduler(time.Local)

			// Configure job based on unit
			var scheduleErr error
			switch strings.ToLower(unit) {
			case "second":
				_, scheduleErr = sh.Every(interval).Second().Do(jobHandler)
			case "minute":
				_, scheduleErr = sh.Every(interval).Minute().Do(jobHandler)
			case "hour":
				_, scheduleErr = sh.Every(interval).Hour().Do(jobHandler)
			case "day":
				_, scheduleErr = sh.Every(interval).Day().Do(jobHandler)
			default:
				fmt.Println("Invalid SCHEDULER_UNIT, defaulting to minutes")
				_, scheduleErr = sh.Every(interval).Minute().Do(jobHandler)
			}

			if scheduleErr != nil {
				fmt.Printf("Failed to schedule job: %v\n", scheduleErr)
				return
			}

			sh.StartAsync()
		}(context.Background())
	}
}
