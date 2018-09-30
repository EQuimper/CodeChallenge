let numbers = [1, 3, 5, 7, 9]
let names = ["Bob", "Jon", "Alice"]

func shuffledArray<T>(arr: [T]) -> [T] {
    var mutateArray = arr;
    var shuffledArraySoFar = [T]()

    while mutateArray.count > 0 {
        let removeIndex = Int.random(in: 0..<mutateArray.count)

        let removeElement = mutateArray.remove(at: removeIndex)

        shuffledArraySoFar.append(removeElement)
    }

    return shuffledArraySoFar
}

print(shuffledArray(arr: numbers))
print(shuffledArray(arr: names))
