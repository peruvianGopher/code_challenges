<?php

class TreeNode {
    public $value;
    public $children;

    public function __construct(int $value, array $children)
    {
        $this->value = $value;
        $this->children = $children;
    }
}

class Traces {
    public $data;

    public function __construct(array $data)
    {
        $this->data = $data;
    }

    public function allComplete(): bool
    {
        foreach ($this->data as $trace) {
            if (!$trace->pathComplete) {
                return false;
            }
        }
        return true;
    }

    public function pushPop(bool $pushPop, ?TreeNode $tn)
    {
        foreach ($this->data as $trace) {
            if (!$trace->pathComplete) {
                if ($pushPop) {
                    $trace->push($tn);
                } else {
                    array_pop($trace->path);
                }
            }
        }
    }
}

class Trace {
    public $nodeValue;
    public $pathComplete = false;
    public $path = [];
    public function __construct(int $nodeValue, array $path)
    {
        $this->nodeValue = $nodeValue;
        $this->path = $path;
    }

    public function push(TreeNode $tn)
    {
        array_push($this->path, $tn);
        if ($tn->value == $this->nodeValue) {
            $this->pathComplete = true;
        }
    }
}

function loadTestData(): TreeNode
{
    $childrenOfThree = [new TreeNode(7, [])];
    $childrenOfSix = [
        new TreeNode(8, []),
        new TreeNode(9, [])
    ];
    $childrenOfFive = [new TreeNode(10, [])];
    $childrenOfTwo = [
        new TreeNode(5, $childrenOfFive),
        new TreeNode(6, $childrenOfSix)
    ];
    $children = [
        new TreeNode(2, $childrenOfTwo),
        new TreeNode(3, $childrenOfThree),
        new TreeNode(4, [])
    ];

    return new TreeNode(1, $children);
}

function main() {
    $tree = loadTestdata();
    $s = solution($tree, 8, 9);
    print_r($s->value);
}

function solution(TreeNode $rootNode, int $nodeA, int $nodeB): TreeNode {
    $traceA = new Trace($nodeA, []);
    $traceB = new Trace($nodeB, []);
    $traces = new Traces([$traceA, $traceB]);
    computeTraces($rootNode, $traces);
    foreach ($traceA->path as $i => $node) {
        if ($node->value != $traceB->path[$i]->value) {
            return $traceB->path[$i - 1];
        }
    }

    return new TreeNode(0, []);
}

function computeTraces(TreeNode $root, Traces $traces) {
    if ($traces->allComplete()) {
        return;
    }

    $traces->pushPop(true, $root);
    if ($traces->allComplete()) {
        return;
    }

    if (sizeof($root->children) == 0) {
        $traces->pushPop(false, null);
    }

    foreach ($root->children as $idx => $child) {
        if ($traces->allComplete()) {
            return;
        }
        computeTraces($child, $traces);
        if ((sizeof($root->children) - 1 == $idx) && !$traces->allComplete()) {
            $traces->pushPop(false, null);
        }
    }
}

main();