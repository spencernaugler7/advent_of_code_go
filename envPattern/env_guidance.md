
So... global variables. All programmers will verbally agree they are bad, yet mysteriously, despite the fact that nobody ever writes one they continue to show up in code bases and libraries. (I blame evil code fairies.) In Go, they are still bad, and there is even less reason than ever to use them. Instead, I create an "environment" object that contains everything I was tempted to put in a global variable, and pass it around.

Suppose I have a logger, and some form of "registry" which will track resources by some string identifier.

