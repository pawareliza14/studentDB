//micro-project
//Creating student database 


package main
import (
	"fmt"
	"os"
	features "student-db/features"
)

func displayChoices() {
	fmt.Println("Select one of the choices")
	fmt.Println("1. Add a new Student")
	fmt.Println("2. Show all Students")
	fmt.Println("3. Update Student")
	fmt.Println("4. Delete Student")
	fmt.Println("5. Exit")
}

func main() {
	fmt.Println("Welcome to student DB")
	var class features.Class
	class.NewClass()

	for {
		displayChoices()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			class.AddStudent()
		case 2:
			class.ShowStudents()
		case 3:
			class.UpdateStudent()
		case 4:
			class.DeleteStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}

}