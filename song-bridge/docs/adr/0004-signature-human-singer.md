# Signature human singer, not synthetic voice (at launch)

The Cover Voice is performed by a **Signature Singer** — the channel's own
recurring human vocalist (one person or a small fixed roster) — across every
video. The channel's identity is this one voice, not a match to each Source
Song's artist, and not a synthetic voice (for now).

## Why

Singable adaptations live or die on phrasing and feeling, which a human still
delivers best, and a recurring voice becomes a brand asset people subscribe to.
It also has the cleanest rights of any option: a human signs a simple performer
release, with zero AI-likeness or voice-model licensing questions. This
reinforces the legal-first catalog (ADR-0001) and the no-clone voice rule
(ADR-0002) — the channel trades volume for identity, safety, and trust.

## Considered options

- **Licensed synthetic voicebank** (Synthesizer V / ACE / Vocaloid) — scales to
  volume cheaply and bridges naturally into the Vocaloid fast-follow, but trades
  away emotional delivery and adds per-voicebank licence checks. Kept as the
  deliberate *scale lever* for later, not the launch sound.
- **AI voice conversion** — most public models are real-artist clones, the exact
  thing ADR-0002 rules out. Rejected.
- **Signature human singer (chosen)** — quality, authenticity, cleanest rights;
  becomes the channel's recognizable identity.

## Consequences

- Throughput is bounded by one human; "every song in the world" is off the table
  at launch — accepted, consistent with the legal-first trade.
- The Signature Singer is a key dependency and brand asset; could be the founder,
  a co-founder, or a partner vocalist on revenue-share.
