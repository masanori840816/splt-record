package dto

type ActionResult struct {
	Succeeded    bool   `json:"succeeded"`
	ErrorMessage string `json:"errorMessage"`
}

func GetSucceededResult() ActionResult {
	return ActionResult{
		Succeeded: true,
	}
}
func GetFailedResult(errorMessage string) ActionResult {
	return ActionResult{
		Succeeded:    false,
		ErrorMessage: errorMessage,
	}
}
