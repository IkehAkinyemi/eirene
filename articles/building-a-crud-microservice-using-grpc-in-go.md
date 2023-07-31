---
Author: Ikeh Akinyemi
CreatedAt: 2023-07-29
Title: Building a CRUD Microservice using gRPC in Go
Synopsis: Concurrency has been an important topic in computer programming for many years. Concurrent programming is the practice of writing code that can execute multiple tasks simultaneously...
ArticleID: 1
---

## Introduction

Rust is a modern programming language that emphasizes performance, safety, and concurrency. Developed by Mozilla, Rust is designed to be a low-level systems language that can also be used for high-level programming tasks. Rust's unique features, including its memory safety guarantees, ownership and borrowing system, and support for fearless concurrency, make it a powerful language for writing safe and performant concurrent programs.

Concurrency has been an important topic in computer programming for many years. Concurrent programming is the practice of writing code that can execute multiple tasks simultaneously, allowing programs to take full advantage of multi-core processors and other hardware resources.


## Why Rust Excels at Concurrency

Rust was designed with concurrency in mind from the very beginning, and it includes a number of features that make it an excellent language for writing concurrent programs. Here are some of the key features that make Rust excel at concurrency:

- **Memory safety guarantees**: Rust's ownership and borrowing system ensures that code is free from memory-related bugs like null pointer dereferences and data races, which can be major sources of errors in concurrent programs.
- **Ownership and borrowing**: Rust's ownership and borrowing system provides a way to manage memory in a concurrent program that is both safe and efficient. By enforcing ownership rules, Rust ensures that only one thread can access a piece of data at a time, which prevents data races and other concurrency bugs.
- **Fearless concurrency**: Rust's ownership and borrowing system allows programmers to write concurrent code with confidence, knowing that the compiler will catch many common errors before they become bugs. This allows developers to write safe, concurrent code without sacrificing performance or productivity.

In this guide, you will learn about Rust's concurrency patterns for parallel programming. You will learn about the different concurrency patterns available for use in Rust. This will include threads, mutexes and locks, channels, and Arc and atomic reference counting.Finally, we'll explore some advanced and related concepts such as error handling, performance and readability, testing concurrent code, and parallel and async programming in Rust.

## Threads in Rust

One of the most basic primitives for concurrency in Rust is threads. A thread is an independent path of execution within a program that can run concurrently with other threads. Threads allow developers to take full advantage of multi-core processors by dividing a task into smaller sub-tasks that can be executed in parallel.

### Creating and Joining Threads

In Rust, threads can be created using the `std::thread::spawn` function, which takes a closure as an argument. The closure represents the code that will be executed in the new thread. Here's an example of creating a new thread:

```rust
use std::thread;

fn main() {
    let handle = thread::spawn(|| {
        // code to be executed in the new thread
    });
}
```
The `thread::spawn` function returns a `JoinHandle`, which represents the new thread. The `JoinHandle` can be used to wait for the thread to finish executing using the `join` method. Here's an example:

```rust
use std::thread;

fn main() {
    let handle = thread::spawn(|| {
        // code to be executed in the new thread
    });
    
    match handle.join() {
        Ok(result) => {
            // handle success case with result
        },
        Err(_) => {
            // handle error case
        }
    };
}
```
In this above snippet, the `join` method is called on the `JoinHandle` type to wait for the new thread to finish executing. The `match` expression is used to handle the `Result` returned by `handle.join()`. If the variant of `Result` is `Ok`, the result value is accessed and used to handle the success case. If the variant is `Err`, then handle the error case.

### Sharing Data Between Threads

In order for multiple threads to work together, they need to be able to share data. Rust provides several ways to share data between threads, including shared ownership and message passing.

#### Shared Ownership

One way to share data between threads is to use shared ownership with the `Arc` (atomic reference counting) smart pointer. An `Arc` allows multiple threads to share ownership of a value, ensuring that the value is not dropped until all threads have finished using it.


```rust
use std::sync::Arc;
use std::thread;

fn main() {
    let shared_data = Arc::new(42);

    let handle = thread::spawn({
        let shared_data = shared_data.clone();

        move || {
            // use the shared_data value in the new thread
        }
    });

    // do some work in the main thread...

    let result = match handle.join() {
        Ok(result) => {
            // handle success case with result
        },
        Err(_) => {
            // handle error case
        }
    };
}
```
In the above snippet, an `Arc` is used to share the `42` value between the main thread and the new thread. The `clone` method is called on the `Arc` to create a new reference to the shared data that can be passed to the new thread.

#### Message Passing

Another way to share data between threads is to use message passing with channels. Channels allow threads to send messages to each other, which can be used to share data and coordinate tasks.

```rust
use std::sync::mpsc;
use std::thread;

fn main() {
    let (sender, receiver) = mpsc::channel();

    let handle = thread::spawn(move || {
        match receiver.recv() {
            Ok(data) => {
                // use the data value in the new thread
            }
            Err(err) => {
                // handle the error
            }
        }
    });

    // do some work in the main thread...

    let data = 42;
    match sender.send(data) {
        Ok(()) => {}
        Err(_) => {
            // handle error
        }
    };
}
```

In the above snippet, a channel is used to send the `42` value from the main thread to the new thread. Creating a channel returns the sender/receiver halves, used for sending and receiving data across threads. The rest of the code uses the `match` keyword to handle the possible outcomes of calling `recv()` and `send()`. If `recv()` returns `Ok(data)`, then the `data` value is utilised in the new thread. If `recv()` returns an `Err`, the error is handled appropriately. Similarly, if `send()` returns `Ok(())`, this indicates that the data was successfully sent. If `send()` returns an `Err`, the error is handled appropriately as well.

### Thread Synchronization
When multiple threads access shared data, there is a risk of data races and other concurrency bugs. Rust provides several primitives for thread synchronization, like mutexes, locks, and atomic variables. You will learn about these concepts in the next section.

## Concurrency patterns

Concurrency patterns are reusable solutions to common problems that arise in concurrent programming. In Rust, several patterns are available, and we will discuss three of them: mutexes and locks, channels, and atomic reference counting. I'll only cover mutexes and locks in this secction as you have learned the other patterns earlier.

### Mutexes and locks
A mutex (short for mutual exclusion) is a synchronization primitive that allows only one thread to access a shared resource at a time. Mutexes are used to prevent data races, where two or more threads access the same memory location concurrently, and at least one of them modifies it.

In Rust, we use the `Mutex` type from the `std::sync` module to create a mutex.

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut val = counter.lock().unwrap();
            *val += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());
}
```

In the above snippet, you created a mutex called `counter` and wrap it in an `Arc`. Then spawn 10 threads, each of which increments the value of the `counter`. To modify the `counter`'s value, each thread must acquire the lock by calling `counter.lock().unwrap()`. If another thread has already acquired the lock, the calling thread will block until the lock is released. Once the lock is acquired, the thread increments the `counter`'s value by dereferencing the mutex value and adding one to it.

### Channels
A channel serves as a conduit through which data can be sent as a means communication and synchronization between concurrent threads. It ensures that data is sent and received safely and efficiently.

In the following example, you will see the use of channels to coordinate tasks with deadlines in a concurrent setting:

```rust
use std::sync::mpsc;
use std::thread;
use std::time::{Duration, Instant};

fn worker(receiver: mpsc::Receiver<Instant>) {
    loop {
        let deadline = match receiver.recv() {
            Ok(deadline) => deadline,
            Err(_) => break,
        };

        let now = Instant::now();
        if now >= deadline {
            println!("Worker received a task after the deadline!");
        } else {
            let remaining_time = deadline - now;
            println!("Worker received a task. Deadline in {:?}.", remaining_time);
            thread::sleep(remaining_time);
            println!("Task completed!");
        }
    }
}

fn main() {
    let (sender, receiver) = mpsc::channel();

    let worker_handle = thread::spawn(move || {
        worker(receiver);
    });

    // Sending tasks with different deadlines
    sender.send(Instant::now() + Duration::from_secs(3)).unwrap();
    sender.send(Instant::now() + Duration::from_secs(5)).unwrap();
    sender.send(Instant::now() + Duration::from_secs(2)).unwrap();
    sender.send(Instant::now() + Duration::from_secs(4)).unwrap();

    // Signal no more tasks and wait for the worker to finish
    drop(sender);
    worker_handle.join().unwrap();
}
```

In this example, we have a `worker` function running in a separate thread. The worker receives deadlines through a channel (`mpsc::Receiver`) and processes the tasks accordingly. Each task is associated with a deadline represented by an `Instant` value.

In the `main` function, we create a channel (`mpsc::channel`) for communication between the main thread and the worker thread. We then spawn the worker thread, passing in the receiver end of the channel.

Next, we send several tasks with different deadlines to the worker by calling `sender.send`. The deadlines are calculated using `Instant::now` and `Duration` to represent a future point in time. The worker receives these tasks from the channel and processes them accordingly. If a task arrives after its deadline, a message is printed indicating that the task was received late. Otherwise, the worker sleeps for the remaining time until the deadline and then prints a completion message.

After sending all the tasks, you can signal no more tasks by dropping the sender end of the channel. This informs the worker that no further tasks will be arriving. Finally, we wait for the worker thread to finish using `join`.

### Arc (Atomic Reference Counting)

An `Arc` is a smart pointer that provides shared ownership of a value across multiple threads. It uses atomic operations and reference counting to efficiently track the number of references to the shared data. This allows multiple threads to access and modify the shared data concurrently.

Let's consider an example where you have three files, and want to read their contents concurrently using multiple threads. You will use the Arc and Mutex constructs to safely share the read data among the threads.

```rust
use std::sync::{Arc, Mutex};
use std::fs::File;
use std::io::{self, BufRead};
use std::thread;

fn main() {
    // File paths to read
    let file_paths = vec!["file1.txt", "file2.txt", "file3.txt"];
    let shared_data = Arc::new(Mutex::new(Vec::new()));

    let mut handles = vec![];

    // Spawn threads to read files concurrently
    for file_path in file_paths {
        let shared_data = Arc::clone(&shared_data);
        let handle = thread::spawn(move || {
            // Open the file
            let file = File::open(file_path).expect("Failed to open file");

            // Read the contents of the file
            let lines: Vec<String> = io::BufReader::new(file)
                .lines()
                .map(|line| line.expect("Failed to read line"))
                .collect();

            // Lock the shared data
            let mut data = shared_data.lock().unwrap();

            // Append the lines to the shared data
            data.extend(lines);
        });

        handles.push(handle);
    }

    // Wait for all threads to finish
    for handle in handles {
        handle.join().unwrap();
    }

    // Explicitly drop the Arc
    drop(shared_data);

    // Attempting to access the shared data after dropping the Arc would result in a compile-time error
    // Uncomment the following line to see the error:
    // let data = shared_data.lock().unwrap();
}
```

In this snippet, a vector of file paths and a `shared_data` variable of type `Arc<Mutex<Vec<String>>>` are created. The purpose of `Arc` is to allow multiple threads to share ownership of the data, while `Mutex` ensures synchronized access to the shared vector.

Multiple threads are spawned to read each files simultaneously. For each file path, a new `Arc` reference is created by cloning the `shared_data` variable. This allows each thread to possess its own reference to the shared data, ensuring safe concurrent access.

Within each thread's execution, the file is opened, and its contents are read line by line, storing them in a vector of strings.

To access the shared data, the thread locks the associated mutex using the `lock()` method, guaranteeing exclusive access to the shared vector.

The lines read from the file are then appended to the shared vector using the `extend()` method, modifying the shared data in a synchronized manner.

Once a thread finishes processing a file, it releases the lock, enabling other threads to access the shared data concurrently.

The main thread waits for all the spawned threads to complete execution by utilizing the `join()` method.

Finally, the Arc reference is explicitly dropped using the `drop()` function, which deallocates the Arc and its associated data. Afterwards, any attempt to access the shared data will result in a compile-time error, preventing any further use of the shared data.

## Advanced/related concepts
    
### Error handling in concurrent code

Rust's error handling mechanism is designed to be expressive, concise and safe. In concurrent code, Rust provides various ways to handle errors such as using `Result` types, match statements, and error propagation.

```rust
use std::thread;
use std::sync::mpsc::{channel, Sender};

fn main() {
    let (tx, rx) = channel();

    let handle = thread::spawn(move || {
        // Perform some computation...
        let result = 42;

        // Send the result over the channel
        let send_result = tx.send(result);
        match send_result {
            Ok(_) => Ok(()),
            Err(e) => Err(e),
        }
    });

    // Wait for the thread to finish and handle any errors that occur
    let thread_result = handle.join().unwrap(); // Note: unwrap is safe here because we're propagating any errors through the Result type
    match thread_result {
        Ok(_) => {
            // Receive the result from the channel
            let result = rx.recv();
            match result {
                Ok(val) => println!("Result: {}", val),
                Err(e) => println!("Error receiving result: {:?}", e),
            }
        },
        Err(e) => println!("Error sending result: {:?}", e),
    }
}
```

Using `Result` and `match` expressions is an efficient and concise way to handle errors in concurrent Rust code. Wrapping the thread function's return value in a `Result` type ensures that any errors are properly propagated to the calling thread. The match expression then handles potential errors in a readable way, without adding overhead. Using the `unwrap` method further ensures that any errors are caught and handled appropriately, without adding additional error handling code.

### Balancing Performance and Readability
Concurrent programs can be challenging to write because they often require careful consideration of performance trade-offs. Rust provides a variety of tools for optimizing performance, such as low-level concurrency primitives and unsafe code. However, these tools come at the cost of reduced safety and increased complexity. It's important to balance performance considerations with the readability and maintainability of your code. One way to do this is to use high-level concurrency abstractions, such as channels and mutexes, wherever possible and only resort to low-level primitives when necessary.

### Testing concurrent code

When testing concurrent code in Rust, it's important to ensure that your tests are deterministic and do not suffer from race conditions or deadlocks. You can achieve this by using Rust's built-in testing framework and implementing your tests with proper synchronization mechanisms.

One approach is to use a Mutex to ensure that the test is run in a thread-safe manner. You can also use a `Condvar` to allow threads to wait for a signal from another thread.

```rust
use std::sync::{Arc, Mutex, Condvar};
use std::thread;

fn count_to_10(shared_data: Arc<(Mutex<u32>, Condvar)>, thread_num: u32) {
    let &(ref mutex, ref cvar) = &*shared_data;
    let mut count = mutex.lock().unwrap();

    while *count < 10 {
        if *count % 3 == thread_num {
            println!("Thread {} counting: {}", thread_num, *count);
            *count += 1;
            cvar.notify_all();
        } else {
            count = cvar.wait(count).unwrap();
        }
    }
}

#[test]
fn test_count_to_10() {
    let shared_data = Arc::new((Mutex::new(0), Condvar::new()));
    let mut handles = Vec::new();

    for i in 0..3 {
        let shared_data = shared_data.clone();
        let handle = thread::spawn(move || {
            count_to_10(shared_data, i);
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}
```
In this test, you spawn three threads, each of which counts up to 10. The `count_to_10` function takes a shared `Mutex` and `Condvar` as arguments, along with the thread number. The function first locks the `Mutex`, then checks if it is the thread's turn to count (determined by the remainder of the current count divided by 3). If it is, the thread increments the count, prints a message, and signals the `Condvar`. If it is not the thread's turn to count, it waits on the `Condvar` until it is signaled by another thread.

The `test_count_to_10` function sets up the shared data, spawns the threads, waits for them to finish, and asserts that the final count is 10.

By using a `Mutex` and `Condvar`, you ensured that the threads synchronize correctly and the test is deterministic.

### Parallel programming with Rayon

Rayon is a data parallelism library for Rust that allows you to write parallel and concurrent programs more easily. It is designed to work with Rust's ownership and borrowing system, and it can automatically manage thread pools to optimize performance.

To use Rayon, you need to add the rayon crate to your `Cargo.toml` file:

```toml
[dependencies]
rayon = "1.7"
```

In this section, you will explore how to use Rayon to parallelize a merge sort algorithm:

```rust

fn merge_sort_par(arr: &mut [i32]) {
    if arr.len() <= 1 {
        return;
    }

    let mid = arr.len() / 2;
    let (left, right) = arr.split_at_mut(mid);

    rayon::join(|| merge_sort_par(left), || merge_sort_par(right));

    let mut i = 0;
    let mut j = mid;
    let mut temp = Vec::with_capacity(arr.len());

    while i < mid && j < arr.len() {
        if arr[i] < arr[j] {
            temp.push(arr[i]);
            i += 1;
        } else {
            temp.push(arr[j]);
            j += 1;
        }
    }

    while i < mid {
        temp.push(arr[i]);
        i += 1;
    }

    while j < arr.len() {
        temp.push(arr[j]);
        j += 1;
    }

    arr.copy_from_slice(&temp);
}

fn main() {
    let mut arr = vec![8, 2, 5, 9, 1, 3, 7, 6, 4];
    merge_sort_par(&mut arr);
    println!("{:?}", arr);
}
```
In this implementation, the `merge_sort_par` function recursively splits the array in halves until it reaches the base case of having only one element or an empty array. It then uses Rayon's `join` function to spawn two parallel tasks to sort each half of the array. Once the tasks complete, it uses Rayon's `join_slices` function to merge the sorted halves in parallel.

This implementation takes advantage of Rayon's automatic work stealing to parallelize the sorting algorithm efficiently across all available CPU cores.

### Async programming in Rust
Asynchronous programming is an essential technique for writing high-performance I/O-bound applications. Rust has built-in support for asynchronous programming with the `async`/`await` syntax and the `futures` library. In this section, you will learn how to write asynchronous Rust programs using the `tokio` library.

The `tokio` library provides an asynchronous runtime that can run multiple tasks concurrently on a single thread or across multiple threads. It also provides a set of utilities and abstractions for writing asynchronous programs, such as futures, streams, and tasks.

To use `tokio`, you need to add it as a dependency in your `Cargo.toml` file:

```toml
[dependencies]
tokio = { version = "1.11.0", features = ["full"] }
```

Let's start by writing a simple program that downloads the content of a webpage asynchronously. You will use the `reqwest` crate to perform the HTTP request:

```rust
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::net::TcpStream;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut stream = TcpStream::connect("example.com:80").await?;

    let request = "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n";
    stream.write_all(request.as_bytes()).await?;

    let mut response = Vec::new();
    stream.read_to_end(&mut response).await?;

    println!("{}", String::from_utf8_lossy(&response));

    Ok(())
}
```

In this example, you use the `tokio::net::TcpStream` type to open a TCP connection to the `example.com` server. Then send an HTTP GET request and read the response asynchronously using the `read_to_end` method. Finally, print the response to the console.

Notice the use of the `async`/`await` syntax to write asynchronous code that looks like synchronous code. You can also use `await` to wait for a future to complete before proceeding to the next line of code.

We can also use `tokio` to perform parallel computation using `tokio::task::spawn`:

```rust
use tokio::task;

#[tokio::main]
async fn main() {
    let handle1 = task::spawn(async {
        // perform some expensive computation asynchronously
    });

    let handle2 = task::spawn(async {
        // perform some other expensive computation asynchronously
    });

    let (result1, result2) = tokio::join!(handle1, handle2);

    // combine the results of the two tasks
}
```

In this above snippet, you use the `tokio::task::spawn` method to spawn two tasks that perform some expensive computation asynchronously. Then used the `tokio::join!` macro to wait for both tasks to complete and collect their results.

`tokio` also provides a set of abstractions for working with asynchronous streams, such as `tokio::io::AsyncRead` and `tokio::io::AsyncWrite`, and asynchronous channels, such as `tokio::sync::mpsc::channel`. These abstractions make it easy to write high-performance asynchronous network servers, for example.



## Conclusion
Rust's concurrency features make it a powerful language for writing high-performance, concurrent programs. Rust's memory safety guarantees, ownership and borrowing system, and support for fearless concurrency make it a great choice for writing safe and performant concurrent code. In this guide, you learned the basics of Rust's concurrency primitives, as well as more advanced concepts like error handling, performance optimization, testing, parallel programming, and async programming. I hope this guide has been helpful in introducing you to the world of Rust concurrency.

### Links to Further Reading

- Rust Programming Language: https://www.rust-lang.org/
- The Rust Programming Language (Book): https://doc.rust-lang.org/book/
- Rust by Example: https://doc.rust-lang.org/rust-by-example/
- Rust Cookbook: https://rust-lang-nursery.github.io/rust-cookbook/
- The Rayon Parallelism Library: https://github.com/rayon-rs/rayon
- Tokio: https://tokio.rs/
- Async-std: https://async.rs/