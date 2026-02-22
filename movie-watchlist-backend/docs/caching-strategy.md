# Caching Strategy

## Overview
This platform utilizes a **Cache-Aside** architecture implemented via Redis to optimize performance and strictly reduce external API calls to TMDB.

## Keys & TTL (Time To Live)

1. **Movie Details** (`movies:details:{movieID}`)
   - **TTL:** 24 Hours
   - **Reasoning:** Movie details (synopsis, crew, streaming providers) rarely change abruptly.

2. **Movie Search Queries** (`movies:search:{query}`)
   - **TTL:** 1 Hour
   - **Reasoning:** Fast response to common searches ("Batman", "Inception").

3. **News Feed**
   - **TTL:** 4 Hours
   - **Reasoning:** News aggregates are moderately static, changing a few times daily.

## Implementation logic (Go)
The Go `MovieService` operates by checking the Redis cache before initiating an HTTP request to the external API client. If there is a cache miss, data is pulled, decoded, sent to the user, and simultaneously serialized back to Redis.
