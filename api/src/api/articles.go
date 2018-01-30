package articles

import "time"

type News struct {
	Title     string
	Content   string
	Author    string
	Published bool
	Date      time.Time
}
