# LeetCode Solutions in Go

This repository contains my solutions to LeetCode problems implemented in Go, organized by algorithmic patterns.

## Structure

```
leetcode-go/
├── problems/                    # Problem solutions organized by algorithmic patterns
│   ├── array/                  # Array manipulation problems
│   ├── hash-table/            # Hash table/map-based solutions
│   ├── two-pointers/          # Two pointer technique
│   ├── sliding-window/        # Sliding window problems
│   ├── binary-search/         # Binary search problems
│   ├── dynamic-programming/   # DP problems
│   ├── backtracking/          # Backtracking problems
│   ├── greedy/                # Greedy algorithm problems
│   ├── graph/                 # Graph traversal problems
│   ├── tree/                  # Tree problems (BST, binary tree, etc.)
│   ├── string/                # String manipulation problems
│   ├── linked-list/           # Linked list problems
│   ├── stack/                 # Stack-based problems
│   ├── queue/                 # Queue-based problems
│   ├── heap/                  # Heap/priority queue problems
│   ├── trie/                  # Trie problems
│   ├── union-find/            # Union-Find/Disjoint Set problems
│   └── bit-manipulation/      # Bit manipulation problems
├── utils/                      # Common utilities and helper functions
└── go.mod                      # Go module file
```

## Algorithmic Patterns

### Hash Table
Problems that use hash maps/sets for O(1) lookups and tracking.
- Example: Two Sum, Group Anagrams, Longest Substring Without Repeating Characters

### Two Pointers
Problems solved using two pointers moving through the array/string.
- Example: Valid Palindrome, Container With Most Water, 3Sum

### Sliding Window
Problems involving contiguous subarrays/substrings with a window.
- Example: Maximum Average Subarray, Minimum Window Substring

### Binary Search
Problems that can be solved using binary search on sorted arrays or search spaces.
- Example: Search in Rotated Sorted Array, Find Peak Element

### Dynamic Programming
Problems that can be broken down into overlapping subproblems.
- Example: Climbing Stairs, Coin Change, Longest Common Subsequence

### Backtracking
Problems that require exploring all possible solutions.
- Example: N-Queens, Generate Parentheses, Word Search

### Greedy
Problems solved by making locally optimal choices.
- Example: Jump Game, Gas Station, Meeting Rooms

### Graph
Problems involving graph traversal (BFS/DFS) and algorithms.
- Example: Number of Islands, Course Schedule, Clone Graph

### Tree
Problems involving tree data structures and tree traversal.
- Example: Maximum Depth of Binary Tree, Validate BST, Invert Binary Tree

## How to Use

1. Identify the algorithmic pattern for your problem
2. Create a new problem directory under the appropriate pattern folder: `problems/[pattern]/[problem-id]-[problem-name]/`
3. Implement the solution in `solution.go`
4. Add tests in `solution_test.go`
5. Run tests: `go test ./problems/[pattern]/[problem-id]-[problem-name]`

## Example

To solve problem "Two Sum" (problem #1) which uses Hash Table pattern:
- Directory: `problems/hash-table/1-two-sum/`
- Files: `solution.go` and `solution_test.go`
- Run: `go test ./problems/hash-table/1-two-sum`
