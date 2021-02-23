package RESTstudent

// StudentsStorage represents a unified way of accessing Student data.
type StudentsStorage interface {
	Add(s Student) error
	Count() int
	Get(key string) (Student, bool)
	GetAll() []Student
}

/*
Student represents the main persistent data structure.
It is of the form:
{
	"name": 	<value>, 	e.g. "Tom"
	"age": 		<value>		e.g. 21
	"student_id":<value>		e.g. "id0"
}
*/
type Student struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	StudentID string `json:"student_id"`
}

/*
studentsDB is the handle to students in-memory storage.
*/
type studentsDB struct {
	students map[string]Student
}

/*
Init initializes the StudentsStorage.
Note: see how we have abstracted the underlying in-memory storage,
and hidden it behind an interface.
*/
func InitStudentsStorage() StudentsStorage {
	db := studentsDB{}
	db.students = make(map[string]Student)
	return &db
}

/*
Add adds new students to the storage.
*/
func (db *studentsDB) Add(s Student) error {
	db.students[s.StudentID] = s
	return nil
}

/*
Count returns the current count of the students in in-memory storage.
*/
func (db *studentsDB) Count() int {
	return len(db.students)
}

/*
Get returns a student with a given ID or empty student struct.
*/
func (db *studentsDB) Get(keyID string) (Student, bool) {
	s, ok := db.students[keyID]
	return s, ok
}

/*
GetAll returns all the students as slice
*/
func (db *studentsDB) GetAll() []Student {
	all := make([]Student, 0, db.Count())
	for _, s := range db.students {
		all = append(all, s)
	}
	return all
}
