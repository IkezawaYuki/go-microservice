package services

type reposService struct {}

type reposServiceInterface interface {
}

var (
	RepositoryService reposServiceInterface
)

func init(){
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(request interface{})(interface{}, error){

}