package orm

type Comment struct {
	Model
	ArticleId string
	UserId    string
	Pid       string
	Content   string
	Email     string
	Url       string
	Ip        string
	Status    bool
	Agent     string
}
