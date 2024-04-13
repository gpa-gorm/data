package data

// ID is Base Id property For Entity.
type Id interface {
	comparable
}

// Entity is Base Entity Interface.
type Entity[ID Id] interface {
	GetId() ID
}

// GormEntityId is Generic Gorm Entity ID Interface.
type GormEntityId interface {
	Id
}

// GormEntity Generic Gorm Entity Interface.
type GormEntity[ID GormEntityId] interface {
	Entity[ID]
	TableName() string
}

// AnyId is Helper Interface for handling ID type as Any.
type AnyId interface{}
