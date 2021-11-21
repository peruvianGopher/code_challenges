<?php
const CLOSE_PAIR = [
    ']' => '[',
    ')' => '(',
    '}' => '{',
];

function solution(string $s): bool
{
    $stack = [];
    foreach (str_split($s) as $c) {
        if (!CLOSE_PAIR[$c]) {
            array_push($stack, $c);
        } else {
            if (array_pop($stack) !== CLOSE_PAIR[$c]) {
                return false;
            }
        }
    }
    return sizeof($stack) === 0;
}

$inputs = [
    "[()]{}{[()()]()}",
	"[(])",
];

foreach ($inputs as $i) {
    echo (solution($i) ? "True" : "False") . PHP_EOL;
}