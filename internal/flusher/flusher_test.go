package flusher

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozonva/ova-travel-api/internal/mocks"
	"github.com/ozonva/ova-travel-api/internal/travel"
)

var _ = Describe("the strings package", func() {
	var (
		mockCtrl  *gomock.Controller
		mockThing *mocks.MockRepo
		slice     []travel.Trip
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockThing = mocks.NewMockRepo(mockCtrl)

		one := travel.Trip{1, "LA", "NY"}
		two := travel.Trip{1, "LA", "NY"}
		three := travel.Trip{1, "PA", "TG"}

		slice = []travel.Trip{one, two, three}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("flusher.Flush()", func() {
		When("flusher can handle all values", func() {
			It("returns empty slice", func() {
				one := travel.Trip{1, "LA", "NY"}
				two := travel.Trip{1, "LA", "NY"}
				three := travel.Trip{1, "PA", "TG"}

				slice := []travel.Trip{one, two, three}

				mockThing.
					EXPECT().
					AddEntities(gomock.Any()).
					Return(nil).
					AnyTimes()

				flusher := NewFlusher(1, mockThing)

				err := flusher.Flush(slice)
				Expect(err).To(Equal(nil))
			})
		})

		When("flusher can't handle all values", func() {
			It("returns empty slice", func() {
				mockThing.
					EXPECT().
					AddEntities(gomock.Any()).
					Return(errors.New("can't save data")).
					AnyTimes()

				flusher := NewFlusher(1, mockThing)

				err := flusher.Flush(slice)
				Expect(err).To(Not(Equal(nil)))
			})
		})
	})
})
