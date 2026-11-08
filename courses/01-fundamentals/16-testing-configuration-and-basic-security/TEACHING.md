# Instructor notes: Testing, configuration, and basic security

[Tier index](../README.md) · [Assessment](../../../instructor/assessment.md) · [Capstone](../CAPSTONE.md)

## Demonstration outline

Open the previous artifact or a known-good copy. State the question in business language before showing syntax. Whiteboard the request/booking flow, then live-code a small slice. Deliberately make one likely beginner error and narrate how to read the first relevant diagnostic. Restore the known-good version and ask a student to predict the next output.

## Explanations and expected reasoning

Tests protect behavior and document expectations. Configuration belongs outside source code. Basic security includes secret handling, input limits, safe SQL, and honest authentication boundaries.

Use the class-booking domain for every example: a student requests a class, the backend validates the request, and stored facts must remain trustworthy. Avoid claiming that the client decides seat availability. Keep code identifiers in English and explain unfamiliar terms aloud.

The lab is the minimum class-time artifact. Pair students at one keyboard; the driver types and the navigator explains. Require both to demonstrate one change. Use the fallback demonstration when setup blocks a student. End with an exit check that asks for reasoning as well as output.

## Support and pacing

If the concept is new to most learners, spend less time on polish and more on prediction and tracing. If learners finish early, extend the booking example without adding unrelated libraries. A student who cannot produce the artifact should receive a targeted clinic task and a clear retry point. Do not treat copying working output as evidence without asking how it was produced.
