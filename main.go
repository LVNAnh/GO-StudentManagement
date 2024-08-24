package main

import (
	"fmt"
)

// Định nghĩa struct Student
type Student struct {
	ID     string
	Name   string
	Gender string // "Male" hoặc "Female"
	Age    int
}

// Khởi tạo map để lưu trữ các sinh viên
var students = make(map[string]Student)

// Thêm sinh viên mới vào map
func addStudent(student Student) {
	students[student.ID] = student
	fmt.Println("Thêm sinh viên thành công:", student.Name)
}

// Tìm kiếm sinh viên theo ID
func getStudentByID(id string) (Student, bool) {
	student, exists := students[id]
	return student, exists
}

// Sửa thông tin sinh viên theo ID
func updateStudent(id string, updatedStudent Student) bool {
	if _, exists := students[id]; exists {
		students[id] = updatedStudent
		fmt.Println("Sửa thông tin sinh viên thành công:", updatedStudent.Name)
		return true
	}
	fmt.Println("Không tìm thấy sinh viên với ID:", id)
	return false
}

// Xóa sinh viên theo ID
func deleteStudent(id string) bool {
	if _, exists := students[id]; exists {
		delete(students, id)
		fmt.Println("Xóa sinh viên thành công với ID:", id)
		return true
	}
	fmt.Println("Không tìm thấy sinh viên với ID:", id)
	return false
}

// Xuất danh sách sinh viên theo giới tính
func getStudentsByGender(gender string) []Student {
	var result []Student
	for _, student := range students {
		if student.Gender == gender {
			result = append(result, student)
		}
	}
	return result
}

func main() {
	// Thêm một số sinh viên mẫu
	addStudent(Student{ID: "S001", Name: "Lê Vũ Nhật Anh", Gender: "Male", Age: 27})
	addStudent(Student{ID: "S002", Name: "Trương Phạm Minh Uy", Gender: "Male", Age: 25})
	addStudent(Student{ID: "S003", Name: "Trần Thị B", Gender: "Female", Age: 20})
	addStudent(Student{ID: "S004", Name: "Nguyễn Thị C", Gender: "Female", Age: 22})

	// Tìm kiếm sinh viên
	student, found := getStudentByID("S002")
	if found {
		fmt.Println("Tìm thấy sinh viên:", student)
	} else {
		fmt.Println("Không tìm thấy sinh viên")
	}

	// Sửa thông tin sinh viên
	updateStudent("S002", Student{ID: "S002", Name: "Trương Phạm Minh Uy Updated", Gender: "Female", Age: 25})

	// Xóa sinh viên
	deleteStudent("S003")

	// Xuất danh sách sinh viên là nam
	maleStudents := getStudentsByGender("Male")
	fmt.Println("Danh sách sinh viên nam:")
	for _, student := range maleStudents {
		fmt.Println(student)
	}

	// Xuất danh sách sinh viên là nữ
	femaleStudents := getStudentsByGender("Female")
	fmt.Println("Danh sách sinh viên nữ:")
	for _, student := range femaleStudents {
		fmt.Println(student)
	}
}
