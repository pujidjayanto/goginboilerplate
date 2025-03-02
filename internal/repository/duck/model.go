package duck

type Duck struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Age  int
}
