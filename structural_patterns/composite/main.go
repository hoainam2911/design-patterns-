package main

import "fmt"

type Component interface {
	search(keyword string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Đang tìm từ khóa '%s' trong file '%s'\n", keyword, f.name)
}

type Folder struct {
	name       string
	components []Component 
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Đang tìm từ khóa '%s' trong folder '%s'\n", keyword, f.name)
	for _, component := range f.components {
		component.search(keyword) 
	}
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}
	folder1.add(file1) 

	folder2 := &Folder{name: "Folder2"}
	folder2.add(file2)   
	folder2.add(file3)   
	folder2.add(folder1) 

	folder2.search("hoa hồng")
}