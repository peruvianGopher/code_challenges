function main() {
    let input = [1, 4, 0, 7, 0, null, 8, 9, 0, 0, 4]
    console.log(solution(input))
}

function solution(list) {
    let swapCounter = 0
    for (let i = 0; i < list.length; i++) {
        if (list[i] === 0) {
            if (swapCounter === 0) {
                swapCounter = i + 1
            }

            while (true) {
                if (swapCounter === list.length) {
                    break
                }

                if (list[swapCounter] !== 0) {
                    let temp = list[i]
                    list[i] = list[swapCounter]
                    list[swapCounter] = temp
                    break
                }
                swapCounter = swapCounter + 1
            }
        }
    }

    return list
}

main()