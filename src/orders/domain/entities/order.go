package entities

type Order struct {
	Id       int32
	Actor    string
	Product  int32
	Quantity int32
}

func NewOrder(id int32, actor string, product int32, quantity int32) *Order {
	return &Order{Id: id, Actor: actor, Product: product, Quantity: quantity}
}

func (o *Order) GetActor() string {
	return o.Actor
}

func (o *Order) SetActor(actor string) {
	o.Actor = actor
}

func (o *Order) SetProduct(product int32) {
	o.Product = product
}