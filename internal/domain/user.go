package domain

import "time"

type User struct {
	Contact  string
	Name     string
	Timezone time.Location
}
