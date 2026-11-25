package main

import "fmt"

type Students struct {
	stus []Student
}
type Student struct {
	name string
	age  int
	id   string
}

func (s *Students) add(stu Student) {
	s.stus = append(s.stus, stu)
}
func (s *Students) delete(id string) {
	for i, v := range s.stus {
		if v.id == id {
			s.stus = append(s.stus[:i], s.stus[i+1:]...)
		}
	}
}
func (s *Students) find() []Student {
	return s.stus
}
func (s *Students) update(id string, student Student) {
	for i, v := range s.stus {
		if v.id == id {
			s.stus[i] = student
		}
	}

}
func main() {
	ss := Students{stus: make([]Student, 0)}
	ss.add(Student{
		name: "jindou",
		age:  11,
		id:   "1",
	})

	for index, value := range ss.find() {
		fmt.Println(index+1, value)
	}

	ss.add(Student{
		name: "yg",
		age:  28,
		id:   "2",
	})

	for index, value := range ss.find() {
		fmt.Println(index+1, value)
	}

	ss.update("2", Student{
		name: "yangguang",
		age:  28,
		id:   "2",
	})

	for index, value := range ss.find() {
		fmt.Println(index+1, value)
	}

	ss.delete("2")

	for index, value := range ss.find() {
		fmt.Println(index+1, value)
	}
}
