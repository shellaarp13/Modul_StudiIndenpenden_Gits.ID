package app

type Models struct {
	Model interface{}
}

func RegisterModels() []Models {
	return []Models{
		{Model: models.User{}},
		{Model: models.Address{}},
	}
}
