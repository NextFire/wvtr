package data

type IStorable interface {
	Store(*Inventory, int)
	GetName() string
}

func (s *Weapon) Store(i *Inventory, number int) {
	i.AddWeapon(s)
}

func (s *Weapon) GetName() string {
	return s.Name
}

func (s *Armor) Store(i *Inventory, number int) {
	i.AddArmor(s)
}

func (s *Armor) GetName() string {
	return s.Name
}

func (s *Omamori) Store(i *Inventory, number int) {
	i.AddOmamori(s)
}

func (s *Omamori) GetName() string {
	return s.Name
}

func (s *Currency) Store(i *Inventory, number int) {
	i.AddCurrency(s, number)
}

func (s *Currency) GetName() string {
	return s.Name
}
