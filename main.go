package main

func main() {
	bench := Benchmark{
		Exec: "./li.ss",
		Args: "./queen6.lsp",
		Entries: []*Entry{
			{
				Config: newCacheConfig(
					"il1:256:16:1:l",
					"",
					"dl1:256:16:1:l",
					"",
					"",
					"",
				),
			},
			{
				Config: newCacheConfig(
					"il1:128:32:1:l",
					"",
					"dl1:128:32:1:l",
					"",
					"",
					"",
				),
			},
			{
				Config: newCacheConfig(
					"il1:64:64:1:l",
					"",
					"dl1:64:64:1:l",
					"",
					"",
					"",
				),
			},
			{
				Config: newCacheConfig(
					"il1:32:128:1:l",
					"",
					"dl1:32:128:1:l",
					"",
					"",
					"",
				),
			},
			{
				Config: newCacheConfig(
					"il1:32:128:1:l",
					"",
					"dl1:32:128:1:l",
					"",
					"",
					"",
				),
			},
			{
				Config: newCacheConfig(
					"il1:64:32:2:l",
					"",
					"dl1:64:32:2:l",
					"",
					"",
					"",
				),
			},
			{
				Config: newCacheConfig(
					"il1:32:64:2:l",
					"",
					"dl1:32:64:2:l",
					"",
					"",
					"",
				),
			},
			{
				Config: newCacheConfig(
					"il1:16:128:2:l",
					"",
					"dl1:16:128:2:l",
					"",
					"",
					"",
				),
			},
		},
	}
	bench.Run()
	bench.ShowResults()
}
