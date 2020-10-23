# config
Simple ini like config loader

```go
	cfg, err := SetFile("config.ini")
	if err != nil {
		panic(err)
	}
	ini, err := cfg.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(ini["Host"])
