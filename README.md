# [Writing An Interpreter In Go](https://interpreterbook.com/)

This is an exercise repo, that has an interpreter implementation for Monkey language.

## Monkey language

Here is how we bind values to names in Monkey:
```
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

Besides integers, booleans and strings, the Monkey interpreter also support arrays and hashes. Here’s what binding an array of integers to a name looks like:
```
let myArray = [1, 2, 3, 4, 5];
```

And here is a hash, where values are associated with keys:
```
let thorsten = {"name": "Thorsten", "age": 28};
```

Accessing the elements in arrays and hashes is done with index expressions:
```
myArray[0]
// => 1
thorsten["name"] // => "Thorsten"
```

The let statements can also be used to bind functions to names:
```
let add = fn(a, b) { return a + b; };
```

But Monkey not only supports return statements. Implicit return values are also possible:
```
let add = fn(a, b) { a + b; };
```

And calling a function is as easy as you’d expect:
```
add(1, 2);
```

A more complex function, such as a fibonacci function that returns the Nth Fibonacci number,
might look like this:
```
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
```

Monkey also supports a special type of functions, called higher order functions. These are
functions that take other functions as arguments. Here is an example:
```
let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};

twice(addTwo, 2); // => 6
```
