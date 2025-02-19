package entities

type Ios struct {
	ID     int    `json:"id_ios" db:"id_ios"`
	Space  int    `json:"data_storage" db:"data_storage"`
	Model  string `json:"model" db:"model"`
	Batery int    `json:"batery" db:"batery"`
}
