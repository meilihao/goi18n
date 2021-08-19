package goi18n

type Elem struct {
	Key string
	Map map[string]string
}

// var (
// 	Cluster = struct {
// 		A *Elem
// 	}{
// 		A: &Elem{
// 			Key: "CLuster.A",
// 			Map: map[string]string{
// 				"a":  `1`,
// 				"b2": `1`,
// 			},
// 		},
// 	}
// )

// func a() {
// 	fmt.Println(Cluster.A)
// }
