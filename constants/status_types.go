package constants

type Status string

const (
	SUCCESS Status = "success"
	FAILED  Status = "failed"
)

func (o Status) String() string {
	return string(o)
}
