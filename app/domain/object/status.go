package object

type (
	StatusId = int64

	// Status status
	Status struct {
		ID StatusId `json:"id"`
		AccountId AccountID `json:"-" db:"account_id"`
		Account *Account `json:"account"`
		Content string `json:"content"`
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
		// Todo: Media Attachments
	}
)