# Hundis programtools
Toolset for creating and testing problems for Hundis locally. 

## Installing programtools
Good luck

## Problem-types
Different types of problems supported by Hundis
### Standard
Code is given input once, gives output. Score is based on output vs answer and grading is based on scoring.

## Commands
### init
```programtools init problem-name problem-type```
Initializes a problem directory and folder structure as well as some functionality.

### verify
```programtools verify problem-name```
Verifies if all solutions get the correct score, checks so that everything compiles.
After verification it runs clean-up to remove all executables it runs.
