<?php

function main() {
    $input = [1,4,0,7,null,0,8,9,0,0,4];
    print_r(solution($input));
}

function solution (array $list): array {
    $swapCounter = 0;
    for ($i = 0; $i < count($list); $i++) {
        if ($list[$i] === 0) {
            if ($swapCounter === 0) {
                $swapCounter = $i + 1;
            }

            while (true) {
                if ($swapCounter === sizeof($list)) {
                    break;
                }

                if ($list[$swapCounter] !== 0) {
                    $temp = $list[$i];
                    $list[$i] = $list[$swapCounter];
                    $list[$swapCounter] = $temp;
                    break;
                }
                $swapCounter = $swapCounter + 1;
            }
        }
    }
    return $list;
}

main();