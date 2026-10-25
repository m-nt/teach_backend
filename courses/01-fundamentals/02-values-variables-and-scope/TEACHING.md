# Instructor notes: Values, variables, types, and scope

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A type describes the values an expression can hold. Use string for a title, int for seat counts, and bool for an open/closed flag. Start with concrete values rather than explaining every numeric type.

2. A variable is a named value that can change. := declares a local variable; = assigns an existing one. Constants name values that do not change. A variable declared inside a block is unavailable outside that block.

3. A compiler error is feedback, not a failed student. Read the first error, find the named line, make one change, and run again. fmt.Println prints values; fmt.Printf adds formatting, such as %s and %d.

## Demonstration script

Type title := "Go basics", seats := 12, and open := true inside main. Print them. Put a discount variable inside an if block and try to use it outside; repair by choosing the correct scope.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- How do := and = differ? **Expected reasoning:** Declaration with inferred type versus assignment to an existing variable.
- What is the type of capacity - booked if both are int? **Expected reasoning:** int.
- Can a local variable declared inside an if block be used after the block? **Expected reasoning:** No, unless a separate variable was declared in an enclosing scope.

## Review the lab

- Capacity 12 and booked 5 prints remaining 7.
- Student explains why assigning a string to booked fails.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
