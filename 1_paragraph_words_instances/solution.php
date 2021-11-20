<?php

const SEPARATOR = ' ';
const NOT_ALLOWED_CHARS = [
    ',' => true,
    ';' => true,
    '.' => true,
    '!' => true,
    ':' => true,
];

/**
 * @param string $p
 * @return array
 */
function solution(string $p): array
{
    $counter = [];
    $word = '';
    foreach (str_split($p . SEPARATOR) as $char) {
        if ($char !== SEPARATOR) {
            if (!NOT_ALLOWED_CHARS[$char]) {
                $word .= strtolower($char);
            }
            continue;
        }

        if (empty($word)) {
            continue;
        }

        $counter[$word] = $counter[$word] ? $counter[$word] + 1 : 1;
        $word = '';
    }

    return $counter;
}

const TEST_INPUT = "When!      you're down and low, lower than the floor, And you feel like you ain't got a chance. Bom, bom, bom, Don't make a move till you're in the groove And do the Peter Panda Dance:";
print_r(solution(TEST_INPUT));