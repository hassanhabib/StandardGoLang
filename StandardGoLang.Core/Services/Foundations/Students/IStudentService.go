package Students

import "StandardGoLang/StandardGoLang.Core/Models/Students"

type IStudentService interface {
	AddStudent(student Students.Student) Students.Student
}
