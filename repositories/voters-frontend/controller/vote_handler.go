package controller

import (
	"bbb-voting/voting-commons/domain"
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

// @Summary Get Rough Totals
// @Description Get rough totals
// @Tags totals votes
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]int
// @Router /votes/totals/rough [get]
func (controller *FrontendController) GetVotesRoughTotalsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	totalsMap, err := controller.participantRepository.GetRoughTotals(controller.ctx)
	if err != nil {
		handleInternalServerError(responseWriter, err)
		return
	}

	result := formatRoughTotals(totalsMap)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(result)
}
func formatRoughTotals(totalsMap map[domain.Participant]int) map[string]int {
	result := map[string]int{}

	for participant, vote := range totalsMap {
		result[participant.Name] = vote
	}

	return result
}

type postVoteBodyModel struct {
	ParticipantID int `json:"participant_id"`
}

// @Summary Post Vote
// @Description Cast a Vote
// @Tags votes
// @Accept  json
// @Produce  json
// @Body postVoteBodyModel
// @Success 201 {object} domain.Vote
// @Router /votes [post]
func (controller *FrontendController) PostVoteHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body := postVoteBodyModel{}
	loadBody(responseWriter, request, &body)

	participant, _ := controller.participantRepository.FindByID(controller.ctx, body.ParticipantID)

	if participant == nil {
		http.Error(responseWriter, "Participant not found", http.StatusNotFound)
		return
	}

	vote := domain.Vote{Participant: *participant, Timestamp: time.Now()}

	controller.voteRepository.SaveOne(controller.ctx, &vote)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(vote)
}

// @Summary Serve HTML rought total page
// @Description Responds with an HTML page with a rought total graph
// @Tags html
// @Produce html
// @Success 200 {string} string "HTML Content"
// @Router /pages/totals/rough [get]
func (controller *FrontendController) LoadRoughTotalPage(responseWriter http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFS(controller.embedTemplates, "rough_results.html")
	if err != nil {
		handleInternalServerError(responseWriter, err)
		return
	}

	err = tmpl.Execute(responseWriter, nil)
	if err != nil {
		handleInternalServerError(responseWriter, err)
	}
}
