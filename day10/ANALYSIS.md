This is a clever solution. Let me walk through how it works.

---

## The Problem (AoC 2025 Day 10, Part 2)

The puzzle involves machines with joltage levels, and buttons that increment specific machines' joltages. You're given a target joltage vector and need to find the **minimum number of button presses** to reach it (starting from zero).

---

## The Core Insight: Divide & Conquer via Even/Odd Parity

The key observation, credited in the code to a Reddit post by `tenthmascot`, is:

> If you apply a combination of buttons whose net joltage effect has the **same even/odd parity** as your current target, you can subtract that effect and then **divide the remainder by 2** — because the difference is guaranteed to be all-even numbers.

This turns a potentially enormous search space into a recursive halving problem, similar to how binary exponentiation works.

---

## Step-by-Step Breakdown

### 1. Parsing
Each input line gives you:
- A target **joltage vector** (e.g., `[4, 6, 2]`) — the goal state for each machine.
- A list of **buttons**, each of which increments a specific subset of machines by 1 when pressed.

### 2. Precomputing Combos by Parity Pattern (`fillPatternsAndCombos`)
Before searching, the code enumerates **every possible subset of buttons** (using bitmask iteration over `2^N` combinations). For each subset it computes:
- The total **joltage delta** from pressing that exact set of buttons once each.
- The **parity pattern** of that delta — a string like `"010"` meaning "machine 0 gets an even bump, machine 1 gets odd, machine 2 gets even."

These are stored in `combosByPattern`, a dictionary keyed by pattern string.

### 3. Recursive Search (`countPresses`)
Given a target joltage vector, the algorithm:

1. **Base cases:** return 0 if all zeros (done), return ∞ if any value is negative (invalid path).
2. Compute the **parity pattern** of the target.
3. Look up all combos whose joltage delta has **the same parity pattern**. This ensures that `target - delta` is all-even, so we can safely divide by 2.
4. For each matching combo, recurse:
   ```
   presses = combo.presses + 2 * countPresses((target - combo.joltage) / 2)
   ```
   The factor of 2 accounts for the fact that the subproblem represents a half-scale version of the remaining work.
5. Take the **minimum** across all matching combos.
6. **Memoize** results by the target vector string to avoid recomputation.

### Why multiply by 2?
When you halve the target and solve it, you're solving for half the remaining joltage. The button presses needed for the full remaining joltage are exactly twice those needed for the half — hence `2 * countPresses(halfTarget)`.

---

## Why This Is Fast

Brute force would mean trying all orderings/repetitions of button presses, which is combinatorially explosive. The optimizations are:

- **Parity filtering** dramatically reduces candidates — instead of testing all `2^N` combos, only a fraction share the same parity pattern as the current target.
- **Recursive halving** means the depth of the search is only `O(log(max_joltage))`.
- **Memoization** ensures each unique subproblem is solved at most once.

Together these bring the runtime to ~100ms even for large inputs.
