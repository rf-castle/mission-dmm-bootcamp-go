package object

type (
	StatusId = int64

	// Status status
	Status struct {
		ID        StatusId  `json:"id"`
		AccountId AccountID `json:"-" db:"account_id"`
		Account   *Account  `json:"account"`
		Content   string    `json:"content"`
		CreateAt  DateTime  `json:"create_at,omitempty" db:"create_at"`
		// Todo: Media Attachments
	}
	TimeLineFilter struct {
		OnlyMedia bool     `db:"only_media"`
		MaxId     StatusId `db:"max_id"`
		SinceId   StatusId `db:"since_id"`
		Limit     int64    `db:"limit"`
	}
)
