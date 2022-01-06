package app

type (
	Repo interface {
	}

	Application struct {
		repo Repo
	}
)

func New(repo Repo) *Application {
	return &Application{
		repo: repo,
	}
}
