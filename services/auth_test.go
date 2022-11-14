package services_test

import (
	s "city/services"
	m "city/services/mock"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Auth", func() {
	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	When("signIn is called", func() {
		Context("sendString", func() {
			It("should be called with Sing param", func() {
				defer mockCtrl.Finish()

				m := m.NewMocksendString(mockCtrl)
				m.
					EXPECT().
					SendString("Sing").
					AnyTimes()

				s.SingIn(m)
			})
		})
	})

})
