# KV Datastructure

## Steps To Run

- You can download and install golang using below link.
  `https://golang.org/doc/install`

- Upon executing below command, it will first make build for the main programme and will run all the test cases defined in main_test.go followed by expected code flow.
`$ make run`
- It will display `PASS` if all test cases are passed else `FAIL` with appropriate fails for failed test cases.
- Once tested we can clean the unwanted binaries by using below command.

- `$ make clean`


## Design Choice(s)

- An array is defined of a particular fixed size viz. 
`MAP_SIZE` as a constant.
- Each array element is a structure which contains a key, value and a empty Binary Search Tree (BST).

- Once the array is completely filled or the key value exceeds the size of the array the next element will stored in the `key % MAP_SIZE` index of the array as a node in BST.

- Hence, during fetch if the key falls in the array size then it will be accessed in O(1) time if the key exceeds the size then it will searched in the BST of `key % MAP_SIZE` index of the array in O(log(k)) time where k is the number of nodes in the respective BST.



## How To Run Code with testcase
make runwithtc

## To build the code
make build

## How To Run Code without testcase
make run

## cleans all binary
make clean 




 




