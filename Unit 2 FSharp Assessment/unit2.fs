module Unit2

open System
open Microsoft.Data.Sqlite

// Domain model: an F# record represents one student.
type Student =
    {
        Id: int
        Name: string
        Department: string
        Age: int
    }

let connectionString = "Data Source=unit2_students.db"

let openConnection () =
    let connection = new SqliteConnection(connectionString)
    connection.Open()
    connection

// Create the database table if it does not already exist.
let initializeDatabase () =
    use connection = openConnection ()
    use command = connection.CreateCommand()
    command.CommandText <-
        """
        CREATE TABLE IF NOT EXISTS Students (
            Id INTEGER PRIMARY KEY,
            Name TEXT NOT NULL,
            Department TEXT NOT NULL,
            Age INTEGER NOT NULL
        );
        """
    command.ExecuteNonQuery() |> ignore

// CREATE
let addStudent (student: Student) =
    use connection = openConnection ()
    use command = connection.CreateCommand()
    command.CommandText <-
        "INSERT INTO Students (Id, Name, Department, Age) VALUES ($id, $name, $department, $age)"
    command.Parameters.AddWithValue("$id", student.Id) |> ignore
    command.Parameters.AddWithValue("$name", student.Name) |> ignore
    command.Parameters.AddWithValue("$department", student.Department) |> ignore
    command.Parameters.AddWithValue("$age", student.Age) |> ignore
    command.ExecuteNonQuery() = 1

// READ
let getStudents () : Student list =
    use connection = openConnection ()
    use command = connection.CreateCommand()
    command.CommandText <- "SELECT Id, Name, Department, Age FROM Students ORDER BY Id"
    use reader = command.ExecuteReader()

    let results = ResizeArray<Student>()
    while reader.Read() do
        results.Add(
            {
                Id = reader.GetInt32(0)
                Name = reader.GetString(1)
                Department = reader.GetString(2)
                Age = reader.GetInt32(3)
            }
        )
    results |> Seq.toList

// UPDATE
let updateStudent id name department age =
    use connection = openConnection ()
    use command = connection.CreateCommand()
    command.CommandText <-
        "UPDATE Students SET Name = $name, Department = $department, Age = $age WHERE Id = $id"
    command.Parameters.AddWithValue("$id", id) |> ignore
    command.Parameters.AddWithValue("$name", name) |> ignore
    command.Parameters.AddWithValue("$department", department) |> ignore
    command.Parameters.AddWithValue("$age", age) |> ignore
    command.ExecuteNonQuery() = 1

// DELETE
let removeStudent id =
    use connection = openConnection ()
    use command = connection.CreateCommand()
    command.CommandText <- "DELETE FROM Students WHERE Id = $id"
    command.Parameters.AddWithValue("$id", id) |> ignore
    command.ExecuteNonQuery() = 1

// Higher-order functions: List.map and List.filter.
let studentsInDepartment department students =
    students
    |> List.filter (fun student ->
        student.Department.Equals(department, StringComparison.OrdinalIgnoreCase))

let studentDisplayNames students =
    students
    |> List.map (fun student ->
        sprintf "%s (ID: %d)" student.Name student.Id)

// Pattern matching processes domain data.
let ageCategory age =
    match age with
    | value when value < 18 -> "Minor"
    | value when value <= 25 -> "Young Adult"
    | _ -> "Adult"

let printStudents (students: Student list) =
    if List.isEmpty students then
        printfn "No student records found."
    else
        printfn "\n%-8s %-22s %-18s %-5s %-14s" "ID" "Name" "Department" "Age" "Category"
        printfn "%s" (String.replicate 72 "-")
        students
        |> List.iter (fun student ->
            printfn "%-8d %-22s %-18s %-5d %-14s"
                student.Id student.Name student.Department student.Age (ageCategory student.Age))

// Async workflow: load records and calculate a summary without blocking
// the caller while the workflow is running.
let asyncSummary () =
    async {
        do! Async.Sleep 100
        let students = getStudents ()
        let totalStudents = List.length students
        let averageAge =
            match students with
            | [] -> 0.0
            | values -> values |> List.averageBy (fun student -> float student.Age)
        return totalStudents, averageAge
    }

let readInt prompt =
    printf "%s" prompt
    match Int32.TryParse(Console.ReadLine()) with
    | true, value -> Some value
    | _ ->
        printfn "Invalid input. Please enter a whole number."
        None

let readRequired prompt =
    printf "%s" prompt
    let value = Console.ReadLine()
    if String.IsNullOrWhiteSpace(value) then
        printfn "This field cannot be empty."
        None
    else
        Some(value.Trim())

let addStudentFromInput () =
    match readInt "Enter student ID: ",
          readRequired "Enter student name: ",
          readRequired "Enter department: ",
          readInt "Enter age: " with
    | Some id, Some name, Some department, Some age when age > 0 ->
        let student = { Id = id; Name = name; Department = department; Age = age }
        try
            if addStudent student then
                printfn "Student added successfully."
            else
                printfn "Student could not be added."
        with
        | :? SqliteException as ex -> printfn "Database error: %s" ex.Message
    | Some _, Some _, Some _, Some _ ->
        printfn "Age must be greater than zero."
    | _ -> printfn "Student was not added because the input was invalid."

let updateStudentFromInput () =
    match readInt "Enter ID to update: ",
          readRequired "Enter new name: ",
          readRequired "Enter new department: ",
          readInt "Enter new age: " with
    | Some id, Some name, Some department, Some age when age > 0 ->
        if updateStudent id name department age then
            printfn "Student updated successfully."
        else
            printfn "No student found with that ID."
    | Some _, Some _, Some _, Some _ -> printfn "Age must be greater than zero."
    | _ -> printfn "Update cancelled because the input was invalid."

let removeStudentFromInput () =
    match readInt "Enter ID to remove: " with
    | Some id ->
        if removeStudent id then
            printfn "Student removed successfully."
        else
            printfn "No student found with that ID."
    | None -> ()

let showSummary () =
    let count, averageAge = asyncSummary () |> Async.RunSynchronously
    printfn "\n--- Student Summary ---"
    printfn "Total students: %d" count
    printfn "Average age: %.2f" averageAge

let rec menu () =
    printfn "\n========== F# STUDENT MANAGEMENT =========="
    printfn "1. Add student"
    printfn "2. Display all students"
    printfn "3. Update student"
    printfn "4. Delete student"
    printfn "5. Filter by department"
    printfn "6. Display names using List.map"
    printfn "7. Async summary"
    printfn "8. Exit"

    match readInt "Enter your choice: " with
    | Some 1 ->
        addStudentFromInput ()
        menu ()
    | Some 2 ->
        getStudents () |> printStudents
        menu ()
    | Some 3 ->
        updateStudentFromInput ()
        menu ()
    | Some 4 ->
        removeStudentFromInput ()
        menu ()
    | Some 5 ->
        match readRequired "Enter department: " with
        | Some department ->
            getStudents ()
            |> studentsInDepartment department
            |> printStudents
        | None -> ()
        menu ()
    | Some 6 ->
        getStudents ()
        |> studentDisplayNames
        |> List.iter (printfn "%s")
        menu ()
    | Some 7 ->
        showSummary ()
        menu ()
    | Some 8 -> printfn "Program closed."
    | Some _ ->
        printfn "Please choose a number from 1 to 8."
        menu ()
    | None -> menu ()

[<EntryPoint>]
let main _ =
    try
        initializeDatabase ()
        printfn "Database is ready."
        menu ()
        0
    with
    | :? SqliteException as ex ->
        printfn "Database error: %s" ex.Message
        1
    | ex ->
        printfn "Unexpected error: %s" ex.Message
        1
