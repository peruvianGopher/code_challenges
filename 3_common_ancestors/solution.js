class TreeNode {
    value = 0
    children = []
    constructor(value, children) {
        this.value = value
        this.children = children
    }
}

class Trace {
    nodeValue = 0
    pathComplete = false
    path = []
    constructor(nodeValue, path) {
        this.nodeValue = nodeValue
        this.path = path
    }

    push(treeNode) {
        this.path.push(treeNode)
        if (treeNode.value === this.nodeValue) {
            this.pathComplete = true
        }
    }
}

class Traces {
    data = []
    constructor(data) {
        this.data = data
    }

    allCompleted() {
        let r = true
        this.data.forEach(el => {
            if (!el.pathComplete) {
                r = false
            }
        })
        return r
    }

    pushPop(pushPop, treeNode) {
        this.data.forEach(el => {
            if (!el.pathComplete) {
                if (pushPop) {
                    el.push(treeNode)
                } else {
                    el.path.pop()
                }
            }
        })
    }
}

function solution(rootNode, nodeA, nodeB) {
    let traceA = new Trace(nodeA, [])
    let traceB = new Trace(nodeB, [])
    let traces = new Traces([traceA, traceB])
    computeTraces(rootNode, traces)

    let r = {}
    traceA.path.every((node, i) => {
        if (node.value !== traceB.path[i].value) {
            r = traceB.path[i - 1]
            return false
        }
        return true
    })

    return r
}

function computeTraces(rootNode, traces) {
    if (traces.allCompleted()) {
        return
    }

    traces.pushPop(true, rootNode)
    if (traces.allCompleted()) {
        return
    }

    if (rootNode.children.length === 0) {
        traces.pushPop(false, null)
    }

    rootNode.children.forEach((c, i) => {
        if (traces.allCompleted()) {
            return
        }
        computeTraces(c, traces)
        if ((rootNode.children.length - 1) === i && !traces.allCompleted()) {
            traces.pushPop(false, null)
        }
    })
}

function main() {
    let tree = loadTestTree()
    let s = solution(tree, 10, 6)
    console.log(s.value)
}

const loadTestTree = () => {
    let childrenOfThree = [new TreeNode(7, [])]
    let childrenOfSix = [
        new TreeNode(8, []),
        new TreeNode(9, [])
    ]
    let childrenOfFive = [new TreeNode(10, [])]
    let childrenOfTwo = [
        new TreeNode(5, childrenOfFive),
        new TreeNode(6, childrenOfSix)
    ]
    let children = [
        new TreeNode(2, childrenOfTwo),
        new TreeNode(3, childrenOfThree),
        new TreeNode(4, [])
    ]

    return new TreeNode(1, children)
}

main()