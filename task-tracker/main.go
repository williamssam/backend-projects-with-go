package main

func main() {
	tasks := Tasks{}
	storage := createStorage[Tasks]("tasks.json")
	storage.load(&tasks)

	cmd := createCommandFlags()
	cmd.execute(&tasks)
	storage.save(tasks)
}
