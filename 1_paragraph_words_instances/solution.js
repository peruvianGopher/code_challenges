const separator = ' '
const notAllowedChars = new Map()
notAllowedChars.set(',', true)
notAllowedChars.set(';', true)
notAllowedChars.set('.', true)
notAllowedChars.set('!', true)
notAllowedChars.set(':', true)

/**
 * @param {string} p paragraph
 *
 * @return {Map}
 */
function solution(p) {
    let counter = new Map()
    p += ' '
    let word = ''
    for (let i = 0; i < p.length; i++) {
        let c = p.charAt(i)
        if (c !== separator) {
            if (!notAllowedChars.get(c)) {
                word += c.toLowerCase()
            }
            continue
        }

        if (word.length === 0) {
            continue
        }

        let wordCounter = counter.get(word)
        counter.set(word, wordCounter ? wordCounter + 1 : 1)
        word = ''
    }

    return counter
}

let testInput = 'When!      you\'re down and low, lower than the floor, And you feel like you ain\'t got a chance. Bom, bom, bom, Don\'t make a move till you\'re in the groove And do the Peter Panda Dance:'
console.log(solution(testInput))