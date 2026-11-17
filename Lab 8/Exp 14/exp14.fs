
open System

// Base class
[<AbstractClass>]
type Shape() =
    abstract member Area : unit -> float

// Derived class: Circle
type Circle(radius: float) =
    inherit Shape()

    override this.Area() =
        Math.PI * radius * radius

// Derived class: Rectangle
type Rectangle(length: float, width: float) =
    inherit Shape()

    override this.Area() =
        length * width

// Main program
[<EntryPoint>]
let main argv =

    printfn "===================================="
    printfn "  EXPERIMENT 14 - OOP IN F#"
    printfn "===================================="

    // Creating objects
    let circle = Circle(5.0)
    let rectangle = Rectangle(10.0, 6.0)

    // Demonstrating polymorphism
    let shapes: Shape list =
        [ circle :> Shape
          rectangle :> Shape ]

    printfn "\n--- Shape Areas ---"

    for shape in shapes do
        printfn "Area: %.2f" (shape.Area())

    printfn "\n--- Individual Objects ---"
    printfn "Circle Area: %.2f" (circle.Area())
    printfn "Rectangle Area: %.2f" (rectangle.Area())

    printfn "\nProgram Completed Successfully!"

    0
