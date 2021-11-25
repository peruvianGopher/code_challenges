<?php

const TEST_A = [
    ["pim", "pom"],
    ["sander", "amy", "ann", "michael"],
    ["adam", "eva", "leo"]
];

const TEST_B = [
    ["999999999", "777888999"],
    ["123456789", "234567890", "789123456", "123123123"],
    ["121212121", "111111111", "444555666"]
];

const TEST_P = ["88999", "1", "112"];

function main() {
    foreach (TEST_A as $i => $t) {
        echo solution(TEST_A[$i], TEST_B[$i], TEST_P[$i]) . PHP_EOL;
    }
}

function solution(array $A, array $B, string $P): string
{
    $m = [];
    foreach ($A as $i => $s) {
        $m[$s] = $B[$i];
    }
    sort($A);
    foreach ($A as $i => $s) {
        if (strpos($m[$s], $P)) {
            return $s;
        }
    }
    return "NO CONTACT";
}

main();