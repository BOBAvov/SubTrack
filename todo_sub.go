package sub_track

type Subscription struct {
	Id          int    `json:"id" db:"id"`
	Userid      string `json:"user_id" binding:"required,uuid" db:"user_id"`
	ServiceName string `json:"service_name" binding:"required" db:"service_name"`
	Price       int    `json:"price" binding:"required" db:"price"`
	StartDate   string `json:"start_date" binding:"required" db:"start_date"`
	EndDate     string `json:"end_date" db:"end_date"`
}

type SubscriptionUpdate struct {
	Price   int    `json:"price" db:"price"`
	EndDate string `json:"end_date" db:"end_date"`
}
