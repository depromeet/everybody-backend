package dto

type SendFeedbackRequest struct {
	// Title 은 제목
	Title string `json:"title"`
	// Content 는 피드백 내용
	Content string `json:"content"`
	// 별점. 5점 만점
	StarRate *uint `json:"star_rate"`
}
