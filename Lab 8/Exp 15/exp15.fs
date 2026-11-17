
open System
open Microsoft.Data.Sqlite
open MongoDB.Bson
open MongoDB.Driver

// ==========================================
// SQLITE - RELATIONAL DATABASE CRUD
// ==========================================

let showSqliteRecords (connection: SqliteConnection) =
    use cmd = connection.CreateCommand()
    cmd.CommandText <- "SELECT Id, Name, Age FROM Students"

    use reader = cmd.ExecuteReader()

    printfn "\nCurrent SQLite Records:"
    while reader.Read() do
        printfn "ID: %d | Name: %s | Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

let runSqlite () =
    printfn "\n========== SQLITE CRUD =========="

    use connection =
        new SqliteConnection("Data Source=lab8_students.db")

    connection.Open()

    // Create table
    use createCmd = connection.CreateCommand()
    createCmd.CommandText <-
        "CREATE TABLE IF NOT EXISTS Students (
            Id INTEGER PRIMARY KEY,
            Name TEXT NOT NULL,
            Age INTEGER NOT NULL
        )"
    createCmd.ExecuteNonQuery() |> ignore
    printfn "Table created successfully."

    // Clean previous test record for repeatable execution
    use cleanCmd = connection.CreateCommand()
    cleanCmd.CommandText <- "DELETE FROM Students WHERE Id = 101"
    cleanCmd.ExecuteNonQuery() |> ignore

    // CREATE
    use insertCmd = connection.CreateCommand()
    insertCmd.CommandText <-
        "INSERT INTO Students (Id, Name, Age)
         VALUES ($id, $name, $age)"

    insertCmd.Parameters.AddWithValue("$id", 101) |> ignore
    insertCmd.Parameters.AddWithValue("$name", "Student One") |> ignore
    insertCmd.Parameters.AddWithValue("$age", 20) |> ignore
    insertCmd.ExecuteNonQuery() |> ignore

    printfn "\nCREATE: Student inserted."
    showSqliteRecords connection

    // READ
    printfn "\nREAD: Student records retrieved."
    showSqliteRecords connection

    // UPDATE
    use updateCmd = connection.CreateCommand()
    updateCmd.CommandText <-
        "UPDATE Students SET Age = $age WHERE Id = $id"

    updateCmd.Parameters.AddWithValue("$age", 21) |> ignore
    updateCmd.Parameters.AddWithValue("$id", 101) |> ignore
    updateCmd.ExecuteNonQuery() |> ignore

    printfn "\nUPDATE: Student age updated."
    showSqliteRecords connection

    // DELETE
    use deleteCmd = connection.CreateCommand()
    deleteCmd.CommandText <- "DELETE FROM Students WHERE Id = $id"
    deleteCmd.Parameters.AddWithValue("$id", 101) |> ignore
    deleteCmd.ExecuteNonQuery() |> ignore

    printfn "\nDELETE: Student deleted."
    showSqliteRecords connection

// ==========================================
// MONGODB - NOSQL DATABASE CRUD
// ==========================================

let runMongoDB () =
    printfn "\n========== MONGODB CRUD =========="

    let uri = Environment.GetEnvironmentVariable("MONGODB_URI")

    if String.IsNullOrWhiteSpace(uri) then
        failwith "MONGODB_URI environment variable is not set."

    let client = MongoClient(uri)

    let database = client.GetDatabase("Lab8Exp15")

    let collection =
        database.GetCollection<BsonDocument>("Students")

    let filter =
        Builders<BsonDocument>.Filter.Eq(
            "StudentId", BsonInt32(101)
        )

    // Clean previous test record for repeatable execution
    collection.DeleteOne(filter) |> ignore

    // CREATE
    let student = BsonDocument()
    student.Add("StudentId", BsonInt32(101))
    student.Add("Name", BsonString("Student One"))
    student.Add("Age", BsonInt32(20))

    collection.InsertOne(student)

    printfn "\nCREATE: Student inserted."

    // READ
    let records = collection.Find(filter).ToList()

    printfn "\nREAD: Student records retrieved."
    for record in records do
        printfn "%s" (record.ToJson())

    // UPDATE
    let update =
        Builders<BsonDocument>.Update.Set(
            "Age", BsonInt32(21)
        )

    collection.UpdateOne(filter, update) |> ignore

    printfn "\nUPDATE: Student age updated."

    let updatedRecords = collection.Find(filter).ToList()
    for record in updatedRecords do
        printfn "%s" (record.ToJson())

    // DELETE
    collection.DeleteOne(filter) |> ignore

    printfn "\nDELETE: Student deleted."

    let remainingRecords = collection.Find(filter).ToList()

    printfn "Remaining matching records: %d"
        remainingRecords.Count

// ==========================================
// MAIN PROGRAM
// ==========================================

[<EntryPoint>]
let main argv =
    try
        printfn "======================================"
        printfn " LAB 8 - EXPERIMENT 15"
        printfn " DATABASE CRUD OPERATIONS"
        printfn "======================================"

        runSqlite()
        runMongoDB()

        printfn "\n======================================"
        printfn " ALL CRUD OPERATIONS COMPLETED!"
        printfn "======================================"

        0

    with
    | ex ->
        printfn "\nERROR: %s" ex.Message
        1
