package entity

type OrderQueue struct {
	Orders []*Order
}

//Less - compara os preços
func (oq *OrderQueue) Less(i, j int) bool {
	 return oq.Orders[i].Price < oq.Orders[j].Price
}

//Swap - inverte as posições
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

//Len - tamanho/qunatidade de ordens
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

//Push - adiciona
func (oq *OrderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

//Pop - remove o último elemento
func (oq *OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}