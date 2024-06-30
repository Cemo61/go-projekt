type Person struct {
	Name   string
	Age    int
	Hobbys []string
}

func main() {
	p := Person{
		Name:   "Yasin Yazgan",
		Age:    34,
		Hobbys: []string{"math", "programming", "language"},
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
