# Instructor notes: Kubernetes: Pods, Deployments, Services, and configuration

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A Pod is the scheduling unit; a Deployment manages replica rollout and replacement; a Service supplies stable discovery for selected Pods. Containers and Pods are not interchangeable concepts.

2. A namespace groups resources, ConfigMaps carry non-secret configuration, and Secrets carry sensitive values with access controls that must be configured. Resource requests help scheduling and limits constrain consumption.

3. Use a single-node training cluster to learn the control loop. It does not demonstrate high availability. Keep PostgreSQL external to the first stateless deployment or use the explicitly disposable lab database track.

## Demonstration script

Apply infrastructure/k8s/api.yaml after replacing the image and configuration. Delete an API Pod and observe replacement and Service routing. Deploy the worker separately.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What replaces a deleted Deployment-managed Pod? **Expected reasoning:** The Kubernetes controllers reconcile the desired replica count.
- Why use a Service? **Expected reasoning:** It provides stable discovery/routing over changing Pod instances.
- Does a single-node cluster prove high availability? **Expected reasoning:** No.

## Review the lab

- API and worker run as separate Deployments.
- Service reaches replacement Pods.
- No real secret is committed to the manifests.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
