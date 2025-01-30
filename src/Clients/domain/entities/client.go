package entities

type Client struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Name      string `json:"name"`
    Direccion string `json:"direccion"`
}