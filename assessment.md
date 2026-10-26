Technical Assessment: In-Memory Cache with TTL
Time Limit: 60 minutes
Problem Statement
Implement an in-memory cache system that supports automatic expiration of entries based on
Time-To-Live (TTL). Your cache should handle concurrent access safely and provide basic cache
operations.
Core Requirements
Implement a cache with these operations:
1. `set(key, value, ttl_seconds)` - Store a key-value pair with TTL
2. `get(key)` - Retrieve a value by key (return None/null if expired or missing)
3. `delete(key)` - Remove a key-value pair
4. `clear()` - Remove all entries
5. `size()` - Return number of active (non-expired) entries
Basic specifications
• Keys are strings, values can be any type that your chosen language supports
• TTL is in seconds (integer)
• Expired entries should not be returned by `get()`

Bonus Requirements
• Automatic cleanup: Periodically remove expired entries
• Statistics: Track hit/miss ratios, total operations
• Key listing: `keys()` method to list all active keys
• Thread safety: Ensure concurrent access is safe
• Memory limits: Maximum cache size with LRU eviction
• Batch operations: `set_many()`, `get_many()`, `delete_many()`

Submission:
1. Code: Complete implementation in a single file or well-organized module
2. Documentation (i.e. README)
a. How to run your code
b. Which requirements you implemented
c. Any design decisions or trade-offs
d. Time complexity of operations (if relevant)
3. Testing: Include basic test cases or demonstrate functionality