package productstorage

type sqlStore struct {
	//db *gorm.DB
	db string
}

func NewSQLStore(db string) *sqlStore {
	return &sqlStore{db: db}
}
