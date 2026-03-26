# The Resilient Worker Pool
**The Goal:** Master Go's concurrency primitives (```goroutines```, ```channels```, ```sync.WaitGroup```) and ```context```.

**The Scenario:** You receive a daily batch file of hundreds of scheduled transfers that need to be processed concurrently, but you cannot overwhelm the database.

<br>

## The Implementation:

**Create a Worker Pool:** Write a Go application that spins up a fixed number of workers (e.g., 5). These workers should listen to a single jobs channel.

**Context for Timeouts:** Wrap the entire batch execution in a ```context.WithTimeout```. If the batch takes longer than 10 seconds, the context should cancel, and all workers should gracefully stop what they are doing and exit.

**Graceful Shutdown:** Use a ```sync.WaitGroup``` to ensure your main function waits for all active workers to finish their current task before shutting down the application.

**Self-Assessment:** Did you have any goroutine leaks? If you forcefully stop the program (Ctrl+C), does it stop immediately, or does it finish processing the in-flight transfers first?
