# Unit 2 (F#) Assessment — Domain Model with Data Persistence

## Requirements covered

1. **Domain model:** `Student` is an F# record with ID, name, department, and age.
2. **Higher-order functions:** `List.map` displays student names and `List.filter` filters students by department.
3. **Pattern matching:** `ageCategory` classifies students using pattern matching.
4. **Persistence and CRUD:** SQLite stores student records. The application supports create, read, update, and delete operations.
5. **Async workflow:** `asyncSummary` uses an F# asynchronous workflow to load records and calculate a summary.

## Prerequisites

- .NET SDK 11
- Internet access for the first NuGet package restore

## Run locally

Open PowerShell in this folder and run:

```powershell
dotnet restore
dotnet build
dotnet run
```

The database file `unit2_students.db` is created automatically in the working directory.

## Suggested manual checks

1. Add a student with ID `101`, name `Aarav`, department `Computer Engineering`, age `20`.
2. Add another student with ID `102`, name `Meera`, department `Computer Engineering`, age `21`.
3. Display all students.
4. Filter by department `Computer Engineering`.
5. Display names using `List.map`.
6. Run the async summary.
7. Update student `101`, then delete student `102`.

The database file is generated at runtime and should not be uploaded to source control.
