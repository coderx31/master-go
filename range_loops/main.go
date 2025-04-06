package main

/** (in Go, everything we assign is a copy)
- if we assign the result of a function returning a struct, it performs a copy of that struct
- if we assign the result if a function returning a pointer, it performs a copy of that memory address
*/

type account struct {
	balance float32
}

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

//func (s *Store) storeCustomers(customers []Customer) {
//	for _, customer := range customers {
//		s.m[customer.ID] = &customer
//	}
//}

//func (s *Store) storeCustomers(customers []Customer) {
//	for _, customer := range customers {
//		current := customer
//		s.m[current.ID] = &current
//	}
//}

func (s *Store) storeCustomers(customers []Customer) {
	for i := range customers {
		customer := &customers[i]
		s.m[customer.ID] = customer
	}
}

func main() {
	//accounts := []account{
	//	{balance: 100},
	//	{balance: 200},
	//	{balance: 300},
	//}
	//
	////for _, a := range accounts {
	////	a.balance += 1000
	////}
	////
	////fmt.Println(accounts)
	//
	//for i := range accounts {
	//	accounts[i].balance += 1000
	//}
	//
	//fmt.Println(accounts)

	//s := []int{0, 1, 2}
	//for range s { // here expression evaluated once , then added to a temporary variable
	//	s = append(s, 10)
	//}

	//fmt.Println(s)

	// this will never stoped
	//ss := []int{0,1,2}
	//for i:=0; i<len(ss); i++ {
	//	ss = append(ss, 10)
	//}
	//
	//fmt.Println(ss)

	//a := [3]int{0, 1, 2}
	//for i, v := range a {
	//	a[2] = 10
	//	if i == 2 {
	//		fmt.Println(v)
	//	}
	//}

	//a := [3]int{0, 1, 2}
	//for i := 0; i < len(a); i++ {
	//	a[2] = 10
	//	if i == 2 {
	//		fmt.Println(a[i])
	//	}
	//}

	//a := [3]int{0, 1, 2}
	//for i := range a {
	//	a[2] = 10
	//	if i == 2 {
	//		fmt.Println(a[2])
	//	}
	//}

	//s := Store{}
	//s.m = make(map[string]*Customer)
	//
	//s.storeCustomers([]Customer{
	//	{ID: "1", Balance: 10},
	//	{ID: "2", Balance: -10},
	//	{ID: "3", Balance: 0},
	//})
	//
	//for k, v := range s.m {
	//	fmt.Printf("key: %s | value: %v \n", k, v)
	//}

}
