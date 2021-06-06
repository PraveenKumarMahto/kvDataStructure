package main

import (
	"fmt"
	"kvDataStructure/bst"
	"kvDataStructure/utility"
	"os"
	"os/signal"
	"syscall"
)

type MyKVMap struct {
	keyValueMap *[MAP_SIZE]*ArrayNode
}

// max zise of map
const MAP_SIZE int64 = 10000

type ArrayNode struct {
	Key   *int64
	Value *string
	Bst   *bst.Node
}

// create array node
func createArrayNode(key int64, value string) *ArrayNode {
	return &ArrayNode{
		Key:   &key,
		Value: &value,
	}
}

// hashFunction return a hash value
func hashFunction(key int64) int64 {
	hashValue := key % MAP_SIZE
	if hashValue < 0 {
		hashValue *= -1
	}
	return hashValue
}

// Insert Key value in map
func InsertOrUpdate(key int64, value string, keyValueMap *[MAP_SIZE]*ArrayNode) {

	hashValue := hashFunction(key)
	arrayNode := keyValueMap[hashValue]

	if arrayNode == nil {
		keyValueMap[hashValue] = createArrayNode(key, value)
		return
	}
	if *arrayNode.Key == key {
		arrayNode.Value = &value
		return
	}
	keyValueMap[hashValue].Bst = keyValueMap[hashValue].Bst.InsertOrUpdateNode(key, value)
}

// find retrive the value of given key if not return nil
func Find(key int64, keyValueMap *[MAP_SIZE]*ArrayNode) (*string, bool) {

	hashValue := hashFunction(key)
	arrayNode := keyValueMap[hashValue]

	if arrayNode == nil {
		return nil, false
	}
	if *arrayNode.Key == key {
		return arrayNode.Value, true
	}

	value := arrayNode.Bst.SearchNode(key)
	return value, true
}

// menu to display
func menu() {
	fmt.Println("1. Insert")
	fmt.Println("2. Retrieve")
	fmt.Println("3. Display current state")
	fmt.Print("Choice (^C to exit): ")
}

func BreakLineWithDash() {
	fmt.Println("\n\n-------------------------------------")
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nBye!")
		os.Exit(0)
	}()
}

func PrintAllMapKeyValue(keyValueMap *[MAP_SIZE]*ArrayNode) {

	var index int64
	for index = 0; index < MAP_SIZE; index++ {
		if keyValueMap[index] == nil {
			continue
		}

		utility.PrintKeyValue(*keyValueMap[index].Key, *keyValueMap[index].Value)

		if keyValueMap[index].Bst != nil {
			keyValueMap[index].Bst.Display()
		}
	}
}

// Program starts from Here
func main() {
	var keyValueMap [MAP_SIZE]*ArrayNode
	SetupCloseHandler()
	var brkLine bool
	for {
		var option int
		if brkLine {
			BreakLineWithDash()
		}
		brkLine = true
		menu()
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			var key int64
			var value string
			fmt.Print("Enter Key: ")
			_, err := fmt.Scanf("%d", &key)
			if err != nil {
				fmt.Println("Invalid Key")
				continue
			}
			fmt.Print("Enter Value: ")
			_, err = fmt.Scanf("%s", &value)
			InsertOrUpdate(key, value, &keyValueMap)
		case 2:
			var key int64
			fmt.Print("Enter Key: ")
			_, err := fmt.Scanf("%d", &key)
			if err != nil {
				fmt.Println("Invalid Key")
				continue
			}
			value, isPresent := Find(key, &keyValueMap)
			if !isPresent {
				fmt.Println("Retrieved value: null")
				continue
			}
			fmt.Println("Retrieved value: ", *value)
		case 3:
			PrintAllMapKeyValue(&keyValueMap)
		default:
			fmt.Println("Invalid option \n\n-------------------------------------")
		}
	}
}
