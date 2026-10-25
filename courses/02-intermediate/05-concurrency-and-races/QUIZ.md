# Quiz: Goroutines, channels, mutexes, and bounded work

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Does passing -race prove the absence of all races?

2. Who should normally close a channel?

3. Why bound concurrency?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No, it observes only exercised executions and addresses memory races.

2. The sender or owner that knows no more sends will occur.

3. To control resource consumption and backpressure.

</details>
