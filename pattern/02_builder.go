package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Плюсы:
Позволяет создавать объекты пошагово;
Позволяет создавать разные объекты, использя один код;
Изолирует код, реализующий конструирование объекта;

Минусы:
алгоритм создания сложного объекта не должен зависеть от того, из каких частей состоит объект и как они стыкуются между собой;
процесс конструирования должен обеспечивать различные представления конструируемого объекта.
*/

type User struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

type UserBuilder struct {
	User
}

func (ub *UserBuilder) Build() User {
	return ub.User
}

func (ub *UserBuilder) Name(name string) *UserBuilder {
	ub.User.Name = name
	return ub
}

func (ub *UserBuilder) Role(role string) *UserBuilder {
	if role == "manager" {
		ub.User.MinSalary = 20000
		ub.User.MaxSalary = 60000
	}
	ub.User.Role = role
	return ub
}
