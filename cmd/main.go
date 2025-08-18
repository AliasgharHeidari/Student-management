package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Student struct {
	Name         string
	Id           int
	Age          int
	Major        string
	Status       string
	SignUpTime   time.Time
	AverageScore int
}

var students = map[int]Student{}

func memUsage() {
	fmt.Println("--------------------------- Memory usage -------------------------------")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc : %.2f MiB\n", float64(m.Alloc)/1024/1024)
	fmt.Printf("Total Alloc : %.2f MiB\n", float64(m.TotalAlloc)/1024/1024)
	fmt.Printf("Sys : %.2f MiB\n", float64(m.Sys)/1024/1024)
	fmt.Printf("NumGC : %v\n", m.NumGC)
	fmt.Println("-------------------------------------------------------------------------")
}

func registerStudent() {
	loadStudents()
	fmt.Println("-------------------- Student sign up ------------------")
	rand.Seed(time.Now().UnixNano())
	var s Student
	fmt.Print("Enter student name: ")
	fmt.Scan(&s.Name)

	for {
		var ageInput string
		fmt.Print("Enter student age: ")
		fmt.Scan(&ageInput)
		age, err := strconv.Atoi(ageInput)
		if err != nil {
			fmt.Println("Invalid age. Please enter a number.")
			continue
		}
		s.Age = age
		break
	}

	for {
		s.Id = rand.Intn(90000) + 10000
		if _, exists := students[s.Id]; !exists {
			break
		}
	}

	var major int
	for {
		fmt.Println("To select student major, enter a number from 1 to 3.")
		fmt.Println("1.Mathematics")
		fmt.Println("2.Science")
		fmt.Println("3.Liberal Arts")
		fmt.Printf("Enter the number : ")
		fmt.Scan(&major)
		switch major {
		case 1:
			fmt.Println("you chose Mathematics")
			s.Major = "Mathematics"
		case 2:
			fmt.Println("you chose Science")
			s.Major = "Science"
		case 3:
			fmt.Println("you chose Libral Arts")
			s.Major = "Libral-Art"
		default:
			fmt.Println("Invalid choice, Please enter a number from 1-3")
			continue
		}
		break
	}

	s.SignUpTime = time.Now()
	students[s.Id] = s

	fmt.Println("--------------------------- successfully registered ------------------------")
	fmt.Println("Name          : ", s.Name)
	fmt.Println("Age           : ", s.Age)
	fmt.Println("Major         : ", s.Major)
	fmt.Println("Student ID    : ", s.Id)
	fmt.Println("Registration  : ", s.SignUpTime.Format("2006-01-02 15:04:05"))
	fmt.Println("----------------------------------------------------------------------------")
	saveStudents()
}

func searchStudent() {
	loadStudents()
	fmt.Print("Enter student ID to search: ")
	var input string
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	student, ok := students[id]
	if ok {
		fmt.Println("------------------------- Student Found -------------------------------")
		fmt.Printf("Name               : %s\n", student.Name)
		fmt.Printf("Age                : %d\n", student.Age)
		fmt.Printf("major:             : %s\n", student.Major)
		fmt.Printf("ID                 : %d\n", student.Id)
		fmt.Printf("AvrageScore        : %.2d\n", student.AverageScore)
		fmt.Printf("Acceptance status  : %s\n", student.Status)
		fmt.Printf("SignUpTime         : %s\n", student.SignUpTime.Format("2006-01-02 15:04:05"))
		fmt.Println("-----------------------------------------------------------------------")
	} else {
		fmt.Println("Student not found. Make sure you've entered the correct ID.")
	}
}

func changeNameByID() {
	loadStudents()
	fmt.Print("Enter student ID to change name: ")
	var input string
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	student, exists := students[id]
	if !exists {
		fmt.Println("Student not found.")
		return
	}

	fmt.Print("Enter the new name: ")
	fmt.Scan(&student.Name)

	students[id] = student
	fmt.Println("Name successfully changed to", student.Name)
	saveStudents()
}

func studentList() {
	loadStudents()
	if len(students) == 0 {
		fmt.Println("No students registered yet.")
		return
	}

	fmt.Println("------------------------------------ Student List ------------------------------------")
	for _, student := range students {
		fmt.Printf("Name: %-12s  |  Age: %-3d  |  Major: %-12s | ID: %-5d | Status:%-10s\n", student.Name, student.Age, student.Major, student.Id, student.Status)
	}
	fmt.Println("--------------------------------------------------------------------------------------")
}

func averageScore() {
	loadStudents()
	var identify string
	var err error
	var ID_int int

	for {
		fmt.Printf("Please enter student ID: ")
		fmt.Scan(&identify)
		ID_int, err = strconv.Atoi(identify)
		if err != nil {
			fmt.Println("Invalid ID, Please enter a number")
			continue
		}
		break
	}

	student, ok := students[ID_int]
	if !ok {
		fmt.Println("Student not found")
		return
	}

	fmt.Println("Identification complete")

	var score int
	for {
		fmt.Printf("Please enter the student AverageScore: ")
		fmt.Scan(&score)

		if score < 0 || score > 100 {
			fmt.Println("Invalid Score, it must be 0-100")
			continue
		}

		fmt.Println("-----Score successfully updated-----")
		student.AverageScore = score
		students[ID_int] = student
		break
	}

	if score < 75 {
		student.Status = "Rejected"
		students[ID_int] = student
	} else {
		student.Status = "Accepted"
		students[ID_int] = student
	}
	saveStudents()
}

func deleteStudent() {
	loadStudents()
	var input string
	fmt.Printf("Enter the student ID : ")
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID, please enter a number")
		return
	}

	if _, ok := students[id]; ok {
		delete(students, id)
		fmt.Println("Student successfuly deleted")

	} else {
		fmt.Println("Student not found")

	}
	saveStudents()
}

func saveStudents() {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		fmt.Println("Saving error : ", err)
	}

	os.WriteFile("students.json", data, 0644)
}

func loadStudents() {
	file, err := os.ReadFile("students.json")
	if err != nil {
		fmt.Println("Loading error : ", err)
		return
	}

	json.Unmarshal(file, &students)

}

func main() {
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Register Student")
		fmt.Println("2 - Search Student")
		fmt.Println("3 - Student List")
		fmt.Println("4 - Change Name")
		fmt.Println("5 - Avrage Score")
		fmt.Println("6 - Delete Student")
		fmt.Println("7 - Memory Usage")
		fmt.Println("8 - Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			registerStudent()
		case 2:
			searchStudent()
		case 3:
			studentList()
		case 4:
			changeNameByID()
		case 5:
			averageScore()
		case 6:
			deleteStudent()
		case 7:
			memUsage()
		case 8:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
