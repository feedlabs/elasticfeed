package population

type SocietyController struct {
	people map[string]*HumanController
}

func (this *SocietyController) Init() {
	this.people = make(map[string]*HumanController)
}

func NewSocietyController() *SocietyController {
	return &SocietyController{nil}
}
