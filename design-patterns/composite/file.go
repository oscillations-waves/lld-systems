package main

type File struct {
	name string
}

func (f *File) search(keyword string) {
	println("Searching for keyword", keyword, "in file", f.name)
}

func (f *File) getName() string {
	return f.name
}
