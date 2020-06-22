# Coffee Machine Assignment

The objective here is to implement a Coffee Machine with the given specifications as mentioned in the requirements [here](Assignment%20_%20SYSTEM%20CODING%20.pdf)

## SetUp and Running the program
The following are the guides for setting up and running the program

### Running on Linux
navigate to the `bin` directory containing the script [setup_and_run](bin/setup_and_run), run the script in sudo mode or provide it with the required permissions

execute the script by running the command `./setup_and_run ../resources/input1.json ../resources/input1.txt`

```bash
[root@localhost bin]# ./setup_and_run ../resources/input1.json ../resources/input1.txt
Last metadata expiration check: 1:20:54 ago on Thu 18 Jun 2020 08:54:48 PM UTC.
Package wget-1.19.5-8.el8_1.1.x86_64 is already installed.
Dependencies resolved.
Nothing to do.
Complete!
--2020-06-18 22:15:44--  https://storage.googleapis.com/golang/go1.14.4.linux-amd64.tar.gz
Resolving storage.googleapis.com (storage.googleapis.com)... 74.125.24.128, 2404:6800:4003:c00::80
Connecting to storage.googleapis.com (storage.googleapis.com)|74.125.24.128|:443... connected.
HTTP request sent, awaiting response... 416 Requested range not satisfiable

    The file is already fully retrieved; nothing to do.

green_tea cannot be prepared because green_mixture is not available
hot_tea is prepared
hot_coffee is prepared
black_tea cannot be prepared because item hot_water is not sufficient
```
The script can accept up to two FilePaths as arguments 
1. The machine json file (Mandatory)
2. The command list text file (Non Mandatory)

The script performs the following tasks
1. Installs wget package
2. Script downloads Golang version 1.14 from the Google API.
3. Decompresses and installs Golang.
4. Builds the project. 
5. Runs the project.

### Running on Mac OS
The same script [setup_and_run](bin/setup_and_run) can be used to run the project on Mac OS also as it can auto detect the Operating system

1. The installation requires homebrew as a pre requisite, please download and install [homebrew](https://brew.sh/) 

Note :- The script was only tested in a Linux environment incase it's not functioning contact me at my [email](mailto:gauravk.it@nsit.net.in) `gauravk.it@nsit.net.in`

## Project Characteristics
The program can execute commands in two modes 
1. File mode with concurrently serving drinks with outlet number (`count_n`) of drinks served in a concurrent fashion
2. Console mode where you can supply the beverage name to the program until you provide the `exit` command

The `resources` directory contains the various input test cases with two files per test case `inputX.json` and `inputX.txt`

you can create your own test cases and provide them as arguments to the program 