function main () {
    let testA = [1,2,3]
    let testB = [1,3,4]

    for (let i = 0; i < testA.length ; i++) {
        console.log(solution(testA[i], testB[i]))
    }

}

function solution(testA, testB) {
    let c = new Map()
    c.get(1)
    return testA + testB
}

main()