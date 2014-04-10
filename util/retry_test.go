package util

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

const TEST_RETRIES = 10
const TEST_SLEEP = 100 * time.Millisecond
const TRIES_TIL_PASS = 5

func TestRetriesUsedUp(t *testing.T) {
	Convey("When retrying a function that never succeeds", t, func() {

		failingFunc := RetriableFunc(
			func() error {
				return RetriableError{errors.New("something went wrong!")}
			},
		)

		start := time.Now()
		retryFail, err := Retry(failingFunc, TEST_RETRIES, TEST_SLEEP)
		end := time.Now()

		Convey("calling it with Retry should return an error", func() {
			So(err, ShouldNotBeNil)
		})
		Convey("the 'retried till failure' flag should be true", func() {
			So(retryFail, ShouldBeTrue)
		})
		Convey("Time spent doing Retry() should be total time sleeping", func() {
			So(end, ShouldHappenOnOrAfter, start.Add((TEST_RETRIES-1)*TEST_SLEEP))
		})
	})
}

func TestRetryUntilSuccess(t *testing.T) {
	Convey("When retrying a function that succeeds after 3 tries", t, func() {

		tryCounter := TRIES_TIL_PASS
		retryPassingFunc := RetriableFunc(
			func() error {
				tryCounter--
				if tryCounter <= 0 {
					return nil
				}
				return RetriableError{errors.New("something went wrong!")}
			},
		)

		start := time.Now()
		retryFail, err := Retry(retryPassingFunc, TEST_RETRIES, TEST_SLEEP)
		end := time.Now()

		Convey("calling it with Retry should not return any error", func() {
			So(err, ShouldBeNil)
		})
		Convey("the 'retried till failure' flag should be false", func() {
			So(retryFail, ShouldBeFalse)
		})
		Convey("time spent should be retry sleep * attempts needed to pass", func() {
			So(end, ShouldHappenOnOrAfter, start.Add((TRIES_TIL_PASS-1)*TEST_SLEEP))
			So(end, ShouldHappenBefore, start.Add((TEST_RETRIES-1)*TEST_SLEEP))
		})

	})
}

func TestNonRetriableFailure(t *testing.T) {
	Convey("When retrying a func that returns non-retriable err", t, func() {

		failingFuncNoRetry := RetriableFunc(
			func() error {
				return errors.New("something went wrong!")
			},
		)

		retryFail, err := Retry(failingFuncNoRetry, TEST_RETRIES, TEST_SLEEP)

		Convey("calling it with Retry should return an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("the 'retried till failure' flag should be false", func() {
			So(retryFail, ShouldBeFalse)
		})
	})
}
