# Rod Cutting

> Striver SDE Sheet 180 · Day 26 DP II — for each length, choose the first piece; the rest is a smaller rod already solved. Unbounded knapsack in disguise. Go, stdlib only, tested.

## Problem

A rod of length n can be cut into integer pieces, each length having a market price. Maximize total revenue — the CLRS classic (price table `[1,5,8,9,10,17,17,20]`, rod 8 → 22, as 2+6).

## Approach

`dp[l]` is the best revenue for a rod of length l: choose the first piece's size (every cut from 1 to l), bank its price, and add the already-computed best for the remainder — `dp[l] = max(prices[cut−1] + dp[l−cut])`. Pieces repeat freely, so this is the **unbounded** knapsack shape (upward sweep, same family as coin change #178) with weight = length and value = price. Selling the rod whole is just the `cut == l` case — no special handling. The DP's structure also explains greedy's failure here: the best first cut depends on what the remainder can earn, which no local rule sees.

**Complexity:** O(n²) time (or n·len(prices)), O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
