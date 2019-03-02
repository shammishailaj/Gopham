# Gopham - Static Code Analysis Tool for Go

Gopham is the worlds first static code analysis tool providing a range of various metrics

## Installation

Step 1: Install Docker

Step 2: Clone this Repository and ```cd``` there

Step 3: Create /projects/ and copy the project you wish to analyze under it

Your directory should look like this:
```
/client/
/goservice/
/projects/myProject
```

step 4: Run docker via ```docker-compose up --build```

step 5: Go to localhost:5000 . Type in a name and the name of the folder that you placed in /projects/. 
"myProject" in this case

## Metrics

Currently supported metrics include:

**1. Efferent Couplings**

Tells you for all your files and packages how many other packages they are using. 
For packages only unique imports are considered. 
Importing "fmt" 3 times in a package will only count as a single coupling

**2. Afferent Couplings**

How many times packages are being used by other files. 

**3. Source Lines of Code**

SLoC. Blank lines and comments are ommited.

**4. Number of Functions**

For both files and packages

___

Metrics planned for the future: Complexity (for functions, files and packages) and SLoC for functions,
as well as various visualizations

## License
[GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)