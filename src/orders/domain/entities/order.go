package entities

type Order struct {
	Id       int32
	Actor    int32
	Product  int32
	Quantity int32
}

func NewOrder(id int32, actor int32, product int32, quantity int32) *Order {
	return &Order{Id: id, Actor: actor, Product: product, Quantity: quantity}
}

func (o *Order) GetActor() int32 { // Ahora devuelve int32
	return o.Actor
}

func (o *Order) SetActor(actor int32) { // Ahora recibe int32
	o.Actor = actor
}

func (o *Order) SetProduct(product int32) {
	o.Product = product
}
