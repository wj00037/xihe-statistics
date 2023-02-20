package app

const (
	ErrorEmptyGitLabProjectIdPage = "empty page of project id"
)

type errorEmptyGitLabProjectIdPage struct {
	error
}

func IsErrorEmptyProjectIdPage(err error) bool {
	_, ok := err.(errorEmptyGitLabProjectIdPage)

	return ok
}
