package root


type Service struct {

}
func NewService() *Service {
	return &Service{}
}
func (s *Service) LoadSession() error {
	return nil
}