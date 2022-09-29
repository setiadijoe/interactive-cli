package resolver

type (
	solver func(input string) error

	promptContent struct {
		errorMsg string
		label    string
	}
)
