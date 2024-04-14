package data

// ID is Base Id property For Entity.
type Id interface {}

// Entity is Base Entity Interface.
type Entity[ID Id] interface {
	//GetId() ID
}

// GormEntityId is Generic Gorm Entity ID Interface.
type GormEntityId interface {
	Id
}

// GormEntity Generic Gorm Entity Interface.
type GormEntity[ID GormEntityId] interface {
	Entity[ID]
	//TableName() string
}

func ConvertToID(param interface{}) Id {
	switch param := param.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, string:
		return Id(param)
	default:
		panic("Invalid ID type")
	}
}