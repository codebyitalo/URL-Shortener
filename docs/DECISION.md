URL Shortener — Design Notes

Stack
Redis for anonymous links and as a cache layer. PostgreSQL for registered users and their durable links. Use Upstash and Neon, both free.

Core Idea
Two tiers of links. Anonymous links live only in Redis with a 7-day TTL. Registered links live in Postgres permanently. The user is never forced to register — anonymous is a valid path, just ephemeral.

Short Code
Using base36 that can go to 0 through 2.147.483.647 options. No base36. Globally unique. This is what appears in the URL. The per-user counter is separate and only used for the dashboard listing.

Ownership
Each link belongs to a user (or to no one if anonymous). The link's identity within a user's space is a simple counter. But the redirect path uses the global short code, not the per-user counter. These two concerns are decoupled.

Redirect Flow
Hit the short URL → check Redis → if present, redirect. If not, query Postgres → populate Redis → redirect.

Link Lifecycle

Active — in use.
Soft-deleted — user removed it. Evict Redis immediately. The short code is locked for 7 days, then reusable.
Hard-deleted — account inactivity cleanup. Short code immediately reusable.
Expired — anonymous TTL hit zero. Gone.
Inactivity Cleanup
Zero new links created in 90 days → hard-delete the account and everything in it. No email warning. Document it in the README.

Boundaries

2.147.483.647 at all, counting all the users, that are already stored. → reject with a friendly message.
Destination can 404. Not your problem.
Portfolio Points

Dual-tier: ephemeral vs durable
Composite identity for per-user scoping
Global short code decouples redirect from ownership
TTL as a product decision
Cache-aside pattern
Soft delete with cooldown before code reuse
