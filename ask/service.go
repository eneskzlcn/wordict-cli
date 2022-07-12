package ask

import "log"

type Service struct {

}
func NewService() * Service {
	return &Service{}
}
func(s *Service) Ask(questionType string) error {
	log.Printf("I have asked you a word with type %s",questionType)
	return nil
}
