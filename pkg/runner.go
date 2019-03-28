package root

type Runner struct {
	Id               string `json:"id"`
	OverallPosition  int32  `json:"overall_position"`
	GenderPosition   int32  `json:"gender_position"`
	CategoryPosition int32  `json:"category_position"`
	Category         string `json:"category"`
	Name             string `json:name`
	ChiName          string `json:chi_name`
	OfficalTime      int32  `json:"offical_time"`
	ChipTime         int32  `json:"chip_time"`
	Fivekm           int32  `json: "Fivekm"`
	FivekmToFinish   int32  `json: "Fivekm_to_finish"`
}

type RunnerService interface {
	GetRunnerByTimeRange(lowerBound int32, upperBound int32) (*Runner, error)
}
