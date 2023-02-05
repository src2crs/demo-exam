# Src2Crs-Demo: Exams

This repo contains an example of how a programming assignment could look like,
and how Src2Crs or its related tools can be used to process it.
The example consists of the following folders:

* [`task1`](task1):
  * [`task.go`](task1/task.go):  
    Source file with the task description and a solution.  
    The students are be given a version of that file that doesn't contain the solution.
  * [`task_test.go`](task1/task_test.go):  
    Contains tests in `Example` form that are also given to the students.
* [`task1_submission1`](task1_submission1):
  * Contains a solution that has been submitted by a student.  
    The solution in this example is correct and satisfies all tests.
* [`task1_submission2`](task1_submission2):
  * Another student's submission which is syntactically correct, but some tests fail.
* [`task1_submission3`](task1_submission3):
  * The code in this submission doesn't compile.
* [`task1_submission4`](task1_submission4):
  * Another perfect submission.
* [`task1_submission5`](task1_submission5):
  * Another syntactically correct submission with failing tests.
* [`task1_submission6`](task1_submission6):
  * Another syntactically correct submission with crashing tests.

*Note*: The submission folders also contain copies of the tests,
making it easier to check the above assumptions.
For processing the submissions, these test copies are not needed.
