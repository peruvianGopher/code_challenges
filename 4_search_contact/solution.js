
const testA = [
    ["pim", "pom"],
    ["sander", "amy", "ann", "michael"],
    ["adam", "eva", "leo"]
]

const testB = [
    ["999999999", "777888999"],
    ["123456789", "234567890", "789123456", "123123123"],
    ["121212121", "111111111", "444555666"]
]

const testP = ["88999", "1", "112"]

function main() {
    for (let i = 0; i < testA.length; i++) {
        console.log(solution(testA[i], testB[i], testP[i]))
    }
}

function solution(A, B, P) {
    m =  new Map()
    for (let i = 0; i < A.length; i++) {
        m.set(A[i], B[i])
    }
    A.sort()
    for (let i = 0; i < A.length; i++) {
        if (m.get(A[i]).includes(P)) {
            return A[i]
        }
    }
    return "NO CONTACT"
}

main()