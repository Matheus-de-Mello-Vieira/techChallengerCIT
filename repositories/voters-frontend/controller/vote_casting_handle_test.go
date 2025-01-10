package controller

import (
	"fmt"

	mocksdatamappers "bbb-voting/voting-commons/tests"
	"context"
	"net/http/httptest"
	"net/url"
	"strings"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("DashBoardController", func() {
	var (
		ctx        context.Context
		controller FrontendController
	)
	BeforeEach(func() {
		templatesPath = "../view/templates/"
		ctx = context.Background()
		controller = NewFrontendController(mocksdatamappers.MockedParticipantDataMapper{}, mocksdatamappers.MockedVotesDataMapper{}, ctx)
	})

	Describe("GetThoroughTotals", func() {
		It("Post cast Vote", func() {
			const participantID int = 1

			data := url.Values{}
			data.Set("id", fmt.Sprint(participantID))
			req := httptest.NewRequest("POST", "http://example.com/votes", strings.NewReader(data.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()

			oldLen := len(mocksdatamappers.MockedVotes)
			controller.VoteCastingHandler(w, req)

			newLen := len(mocksdatamappers.MockedVotes)

			resp := w.Result()
			Expect(resp.StatusCode).To(Equal(200))
			Expect(newLen).To(Equal(oldLen + 1))
			Expect(mocksdatamappers.MockedVotes[len(mocksdatamappers.MockedVotes)-1].Participant.ParticipantID).To(Equal(1))
		})
	})
})