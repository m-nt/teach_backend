# Instructor notes: Stateful services, backups, and restore drills

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A replica is not a backup: accidental deletion can replicate. Define recovery point objective as acceptable data loss and recovery time objective as acceptable recovery duration. Choose a backup strategy against those requirements.

2. A PersistentVolumeClaim asks for storage but does not itself provide backup, off-site durability, or proven failover. For a small deployment, compare a managed database or dedicated host against operating a database in Kubernetes.

3. A backup is useful only if restoration is tested. Restore to a separate target, apply necessary schema compatibility checks, compare counts and known records, and run a synthetic business flow before considering recovery complete.

## Demonstration script

Follow resources/backup-drill.md with a synthetic database. Record backup time, create a later booking, restore to a different database, and show exactly which data was recovered.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why is replication not a backup? **Expected reasoning:** Bad changes can replicate too.
- What does RPO measure? **Expected reasoning:** Acceptable data loss, usually expressed as time.
- Does a successful backup command prove recovery? **Expected reasoning:** No; restoration and application validation must be tested.

## Review the lab

- Restored data is validated through SQL and the API.
- Source database remains untouched.
- Observed data loss and recovery duration are compared to targets.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
