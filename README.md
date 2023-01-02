# BPMN-Flow

A project for automating mapping of system behaviour into BPMN notation.

| :warning:  WARNING            |
|-------------------------------|
| Project currently heavily WIP |


## Motivation

One of the central challenges of code and system maintenance is documentation. Previously, keeping documentation up to date was error-prone and reactive, with little recourse if failures in the docs emerged. Many frameworks now exist to streamline the documentation process, with self-documenting code as the industry standard. Issues still occur, but at far lower frequency than in the past.

However, docstrings do not tell the whole story. In codebases with complex business logic or high degrees of reusability, tracing the exact flow of a program can become a challenge that docstrings only partially assist. Modelling frameworks such as BPMN were developed to help make sense of complex systems and describe the process flow of a system. This is a less-granular view of how a system behaves, but performs the role of a high-level overview that has the advantage of being simpler to understand for non-expert users.

In practice, BMPN diagrams fall into the same hole as documentation did in the past. With the exception of frameworks that are BPMN-first (Camunda/Zeebee etc.) and the simpler DAG charts produced by workflow tools (Airflow etc.), there is a lack of documentation for general-consumption program flow. BPMN diagrams quickly fall out of concordance with the target system unless the team is using BPMN as a fairly strict blueprint, and this constraint smacks of BDUF.

Problems arise when the plan, rather than the system, is treated as the source of truth.

This package steps into this gap and attempts to reverse the dependency. Our systems should not reflect BPMN diagrams, BPMN should reflect our systems. If our system is dynamic, our documentation should reflect that.

### Use Case

This tool is designed to "map" how your process logic actually executes, stripped of non-essential implementation details. It is designed to map the flow from task to task, without having to manually create the maps yourself.

The plan is to use unit and integration tests to call every branch of your logic (you have been writing tests, right?), then aggregate the fragmented views into a unified whole.

### What This Is Not

This is not a tool for every application. This is a tool for "business logic"-centric services that primarily orchestrate other tasks, where orchestration is more important than raw cycle time.

This is not a monitoring tool. This system can be considered an extension of application traces provided by Jaeger/New Relic/Data Dog, but does NOT replace them. There are plans to supplement these features in future.

This is not a call graph. You have call graphs to serve as call graphs. However, if you find call graphs too granular and want to aggretate it up into meaningful chunks, then maybe this is for you.


## Installation

Don't. Not yet.