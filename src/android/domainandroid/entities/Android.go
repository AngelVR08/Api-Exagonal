package entities

type Android struct {
	ID    int    `json:"id_android" db:"id_android"`
	Brand string `json:"brand" db:"brand"`
	Model string `json:"model" db:"model"`
	Space int    `json:"data_storage" db:"data_storage"`
}
