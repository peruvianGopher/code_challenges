const CLOSE_PAIRS = {
    ']': '[',
    ')': '(',
    '}': '{',
};

/**
 * @param {string} s paragraph
 *
 * @return {boolean}
 */
function solution(s) {
    let stack = []
    for (let i = 0; i < s.length; i++) {
        let c = s.charAt(i)
        if (!CLOSE_PAIRS[c]) {
            stack.push(c)
        } else {
            let last = stack.pop()
            if (CLOSE_PAIRS[c] !== last) {
                return false
            }
        }
    }
    return stack.length === 0
}

let inputs = [
    '[()]{}{[()()]()}',
    '[(])',
]

console.log(solution(inputs[0]))
console.log(solution(inputs[1]))