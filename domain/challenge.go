package domain

type Challenge struct {
	Base
	Title          string
	Slug           string
	Description    string
	Tags           string
	Requirements   string
	Level          string
	ChallengeFiles []ChallengeFile
	Authors        []Author
}

func NewChallenge(title string) *Challenge {
	return &Challenge{
		Title: title,
	}
}
