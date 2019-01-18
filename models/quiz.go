package models

// Quiz Information
type Quiz struct {
	ID      string     `json:"id"`
	Content []Question `json:"content"`
	// Reviews          []Review    `json:"reviews"`
}

func getQuizzes() []Quiz {
	allQuizzes := []Quiz{
		Quiz{
			ID: "someid",
			Content: []Question{
				Question{
					ID:           "question Id",
					QuestionText: "what is the world's tallest building",
					Options: []Option{
						Option{
							ID:         "option1",
							OptionText: "Burj Khalifa",
						},
						Option{
							ID:         "option2",
							OptionText: "World Trade Center",
						},
					},
				},
			},
		},
	}
	return allQuizzes
}

// GetQuiz gets a Quiz by id
func (quiz *Quiz) GetQuiz(id string) int {
	allQuizzes := getQuizzes()
	statusCode := 0
	for _, item := range allQuizzes {
		if id == item.ID {
			statusCode = 1
			*quiz = item
		}
	}
	return statusCode
}

// GetAllQuizzes func gets all the products in the database
func GetAllQuizzes() (int, []Quiz) {
	allQuizzes := getQuizzes()
	return 1, allQuizzes
}
