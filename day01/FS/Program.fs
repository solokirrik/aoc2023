open System.IO

let readLines (path: string) : string[] =
    try
        let lines: string[] = File.ReadAllLines(path = path)
        lines
    with :? IOException as e ->
        printfn "The file could not be read: %s" e.Message
        [||]

let task1 (lines: string[]) : int =
    let mutable sum: int = 0

    for line: string in lines do
        let foundDigits: ResizeArray<int> = ResizeArray<int>()

        for i = 0 to line.Length - 1 do
            let c: char = line.[i]

            if c >= '1' && c <= '9' then
                foundDigits.Add(int (c) - 48)

        if foundDigits.Count > 0 then
            let first: int = foundDigits.[0] * 10
            let last: int = foundDigits.[foundDigits.Count - 1]

            sum <- sum + first + last

    sum

[<EntryPoint>]
let main (_: string array) : int =
    let path: string = "../input"
    let result: int = task1 (readLines path)
    printfn "Result1: %d" result
    0
