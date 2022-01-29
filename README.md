# cosmos_sdk_poc
A Proof of Concept for scaffolding application-driven blockchains on Cosmos SDK

## Proof of Concept Introduction
This is an introduction to a blockchain driven application proof of concept built on top of Cosmos SDK with Starport Tool. 
In this proof of concept we will be testing building a scavenger hunt concept with answering questions for bounties on the blockchain.

## Small introduction of Cosmos SDK
Cosmos SDK is a Go framework that allows to build determinitic state machines on the clou (blockchains) with a single application (with different modules)

## Starport
To make it easier to create a blockchain template, we use the Starport tool to scaffold blockchain and modules with just one command.

## Rule of the Proof Of Concept
Scavenger hunt game with 3 basic mechanics: 
1. Anyone can post a question with an encrypted answer.
2. The question has a bounty of coins in it.
3. Anyone can post an answer on the question and receive the coins if it is correct.

## New Blockchain
First of all, to scaffold a new blockchain use the command `starport scaffold chain github.com/cosmonaut/scavenge --no-module`
This creates a new file with the basic SDK modules `scavenge`. 
Access it via `cd scavenge`

## Scavenge Module
Now we need to implement our first module : Scavenge. This module will be used to send coins to people who answer with the correct question. We can use `starport scaffold module scavenge --dep bank`.
A `bank` dependency is used since it makes it easier to transfer tokens on the blockchain.

## CRUD Messages
After creating our first module we need to create CRUD (create, read, update, delete) messages, which alter the state of our blockchain. 
In our proof of concept, there will only be 3 main messages:
1. Submit question
2. Commit solution
3. Reveal solution

### Submit question 
To create a submit question message: `starport scaffold message submit-scavenge solutionHash description reward`
`SolutionHash`: The hashed solution to the question
`description`: The descrption of the question
`reward`: Amount of coins as bounty for question

### Commit solution
To create a commit solution message: `starport scaffold message commit-solution solutionHash solutionScavengerHash`
`solutionHash`: The hashed solution to the question
`solutionScavengerHash`: The hash of the combination of the solution and the person who solved it. 

### Reveal Solution
To create a reveal solution message: `starport scaffold message reveal-solution solution`
`solution`: The plain text version of the solution

## Types
After creating the messages, we need to create type and methods that can alter the state (remember CRUD).

###  Scavenge
We need to scaffold a map to keep track of all the questions.  
```starport scaffold map scavenge solutionHash solution description reward scavenger --no-message```

### Commit 
Also create a map for the same logic for committing. 
```starport scaffold map commit solutionHash solutionScavengerHash --no-message```

## Functions 
### Submit Scavenge
Submitting a scavenge should do the following: 
1. Check that a scavenge with a given solution hash doesn't exist
2. Send tokens from scavenge creator account to a module account
3. Write the scavenge to the store

Steps: 
1. Follow comments on `x/scavenge/keeper/msg_server_submit_scavenge.go`
2. Follow comments on `x/scavenge/keeper/msg_server_commit_solution.go`
3. Follow comments on `x/scavenge/keeper/msg_server_reveal_solution.go`