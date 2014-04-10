package util

import (
	"time"
)

//RetriableError can be returned by any function called with Retry(),
//to indicate that it should be retried again after a sleep interval.
type RetriableError struct {
	Failure error
}

type RetriableFunc func() error

func (retriable RetriableError) Error() string {
	return retriable.Failure.Error()
}

//Retry will call attemptFunc up to maxTries until it returns nil,
//sleeping the specified amount of time between each call.
//The function can return an error to abort the retrying, or return
//RetriableError to allow the function to be called again.
func Retry(attemptFunc RetriableFunc, maxTries int, sleep time.Duration) (bool, error) {
	triesLeft := maxTries
	for {
		err := attemptFunc()
		if err == nil {
			//the attempt succeeded, so we return no error
			return false, nil
		}
		triesLeft--

		if retriableErr, ok := err.(RetriableError); ok {
			if triesLeft <= 0 {
				// used up all retry attempts, so return the failure.
				return true, retriableErr.Failure
			} else {
				// it's safe to retry this, so sleep for a moment and try again
				time.Sleep(sleep)
				continue
			}
		} else {
			//function returned err but it can't be retried - fail immediately
			return false, retriableErr.Failure
		}
	}
}
