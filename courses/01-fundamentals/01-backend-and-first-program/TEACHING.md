# Instructor notes: What a backend does; your first Go program

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A program is a sequence of instructions executed by a computer. Source code is the text we edit; Go compiles that text into a program. A terminal is a way to give commands, and its working directory decides which files a command sees. Demonstrate folders before introducing syntax.

2. A customer asks to book a class. The browser sends a request; the backend checks rules; a database remembers the booking; a response tells the browser what happened. Draw these as four boxes. A database stores facts but does not decide every business rule.

3. A business entity is something the business identifies over time: a Student, Course, ClassSession, or Booking. An ID identifies an entity even when its name changes. Business identity here means domain identity, not company branding. Today students only name entities; tables come later.

## Demonstration script

Open examples/01-hello/main.go. Run go run . from its directory. Change the class name and run again. Remove a quote, read the compiler location aloud, then repair it. Draw the booking request journey beside the terminal.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What is the difference between source code and output? **Expected reasoning:** Source is the instructions we edit; output is what executing them produces.
- Where should the rule “do not overbook” be enforced? **Expected reasoning:** In the backend; later the database also helps enforce consistency.
- Does changing a student name change their identity? **Expected reasoning:** No. A stable ID represents identity independently of the name.

## Review the lab

- A greeting printed from the student’s own file.
- A student can recover from one missing quote and explain the four boxes.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
