package saver

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"

	"github.com/ozonva/ova-travel-api/internal/mocks"
	"github.com/ozonva/ova-travel-api/internal/travel"
)

func barrier(t uint64) {
	barrier := time.NewTimer(time.Duration(t) * time.Second)
	<-barrier.C
}

var _ = Describe("saver package tests", func() {
	var (
		testTrip    travel.Trip
		mockCtrl    *gomock.Controller
		mockFlusher *mocks.MockFlusher
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(mockCtrl)

		testTrip = travel.Trip{1, "LA", "NY"}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("saver can instantiated with invalid paramenters", func() {
		When("0 capacity was passed", func() {
			It("returns error", func() {
				saver, err := NewSaver(0, 1, mockFlusher)
				Expect(err).To(Not(BeNil()))
				Expect(saver).To(BeNil())
			})
		})

		When("0 time limit was passed", func() {
			It("returns error", func() {
				saver, err := NewSaver(1, 0, mockFlusher)
				Expect(err).To(Not(BeNil()))
				Expect(saver).To(BeNil())
			})
		})
	})

	Context("saver.Save()", func() {
		When("saver riches capacity", func() {
			It("flushes the content", func() {
				mockFlusher.
					EXPECT().
					Flush(gomock.Any()).
					Return(nil).
					Times(1)

				capacity := uint(3)
				saver, err := NewSaver(capacity, 1000, mockFlusher)
				Expect(err).To(BeNil())
				initErr := saver.Init()
				Expect(initErr).To(BeNil())

				for i := uint(0); i < capacity+1; i += 1 {
					saver.Save(testTrip)
				}

				barrier(1)
			})
		})

		When("the timer is out", func() {
			It("flushes the content", func() {
				mockFlusher.
					EXPECT().
					Flush(gomock.Any()).
					Return(nil).
					Times(1)

				capacity := uint(5)
				timeLimit := uint64(1)
				saver, err := NewSaver(capacity, timeLimit, mockFlusher)
				Expect(err).To(BeNil())
				initErr := saver.Init()
				Expect(initErr).To(BeNil())

				for i := uint(0); i < capacity-1; i += 1 {
					saver.Save(testTrip)
				}

				barrier(timeLimit + 1)
			})
		})

		When("no limit is reached", func() {
			It("doesn't flush", func() {
				mockFlusher.
					EXPECT().
					Flush(gomock.Any()).
					Return(nil).
					Times(0)

				capacity := uint(1000)
				timeLimit := uint64(1000)
				saver, err := NewSaver(capacity, timeLimit, mockFlusher)
				Expect(err).To(BeNil())
				initErr := saver.Init()
				Expect(initErr).To(BeNil())

				for i := uint(0); i < 1; i += 1 {
					saver.Save(testTrip)
				}

				barrier(1)
			})
		})
	})

	Context("saver.Close()", func() {
		When("there are unsaved entities", func() {
			It("they are flushed ", func() {
				mockFlusher.
					EXPECT().
					Flush(gomock.Any()).
					Return(nil).
					Times(1)

				capacity := uint(1000)
				timeLimit := uint64(1000)
				saver, err := NewSaver(capacity, timeLimit, mockFlusher)
				Expect(err).To(BeNil())
				initErr := saver.Init()
				Expect(initErr).To(BeNil())

				for i := uint(0); i < 1; i += 1 {
					saver.Save(testTrip)
				}

				saver.Close()
				barrier(1)
			})
		})
	})
})
