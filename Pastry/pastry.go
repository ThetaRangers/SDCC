package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var targetId string

var m = 2
var b = 4
var base = int(math.Pow(2, float64(b)))
var n = 128
var rows int
var cols int

type metadataStruct struct {
	nodeId    string
	ipAddress string
}

type stateTableStruct struct {
	leafSetLower   []metadataStruct
	leafSetGreater []metadataStruct
	neighbourSet   []metadataStruct
	routingTable   [][]metadataStruct
}

var metadata metadataStruct
var stateTable stateTableStruct

func join(w http.ResponseWriter, req *http.Request) {
	if len(stateTable.neighbourSet) == 0 {
		//Nobody in the network
	}

	//fmt.Fprintf(w, "%s\n", string(nodeId))
	fmt.Fprintf(w, "%s\n", string(metadata.nodeId))
}

func connectToNode(fullAddress string) {
	//Connect to a node
	fmt.Println("Connecting to : ", fullAddress)

	targetIp := fmt.Sprintf(fullAddress)

	//NEED TO SEND JOIN(HASH)

	resp, err := http.Get(fmt.Sprintf("http://%s/join", targetIp))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		targetId = scanner.Text()
	}

	targetMetadata := metadataStruct{nodeId: targetId, ipAddress: targetIp}

	fmt.Printf("Connected to fullAddress %s nodeId %x\n", fullAddress, targetMetadata.nodeId)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func prefixMatch(a, b string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		//Check if char is the same
		if a[i] == b[i] {
			count++
		}
	}

	return count
}

func route(key string) {

	x, _ := strconv.ParseInt(key, base, 0)
	var dest metadataStruct

	if len(stateTable.leafSetLower) == 0 && len(stateTable.leafSetGreater) == 0 {
		//Empty
	}

	//Search in the leaf set
	leafSetLowerNumber, _ := strconv.ParseInt(stateTable.leafSetLower[0].nodeId, base, 0)
	leafSetGreaterNumber, _ := strconv.ParseInt(stateTable.leafSetLower[0].nodeId, base, 0)

	if x > leafSetLowerNumber && x < leafSetGreaterNumber {
		//In the leaf set

	} else {
		//Plaxton Routing
		max := 0
		var currentMax metadataStruct

		for row := 0; row < rows; row++ {
			//TODO Second for not needed just look at the last digit
			for col := 0; col < b; col++ {
				pm := prefixMatch(key, stateTable.routingTable[row][col].nodeId)
				if pm > max {
					//Update max
					max = pm
					currentMax = stateTable.routingTable[row][col]
				}
			}
		}

		dest = currentMax
	}

	fmt.Println("Routed to: ", dest)
}

func generateId(x string) string {
	//nodeId with hash
	h := md5.New()
	h.Write([]byte(x))

	nodeHash := h.Sum(nil)
	//nodeId := ""

	var sb strings.Builder
	fmt.Println(len(nodeHash), nodeHash)

	for i := 0; i < len(nodeHash); i++ {
		c := int(nodeHash[i])

		s := big.NewInt(int64(c)).Text(int(math.Pow(2, float64(b))))

		//Fill missing bits
		/*diff := int(math.Ceil(float64(16/b))) - len(s)
		if diff > 0 {
			var sbTemp strings.Builder
			for j := 0; j < diff; j++ {
				sbTemp.WriteString("0")
			}

			sbTemp.WriteString(s)
			s = sbTemp.String()
		}*/

		sb.WriteString(s)
	}

	return sb.String()
}

/*
Routing table
Leaf set: IP e ID secondo una metrica
Neighbour set: IP e ID di N nodi vicini
*/

/****************************
arg[0] -> IP
arg[1] -> port
arg[2] -> debug connect/server
*****************************/
func main() {
	//remove first one
	argsWithoutProg := os.Args[1:]

	ip := argsWithoutProg[0]
	port := argsWithoutProg[1]

	fullAddress := fmt.Sprintf("%s:%s", ip, port)

	nodeId := generateId(fullAddress)

	//Initialize routing table
	rows = int(math.Log(math.Pow(2, float64(n))) / math.Log(math.Pow(2, float64(b))))
	cols = int(math.Pow(2, float64(b)) - 1)

	fmt.Println("Routing table dimensions ROWS:", rows, " B:", b)
	stateTable.routingTable = make([][]metadataStruct, rows)
	for i := 0; i < rows; i++ {
		stateTable.routingTable[i] = make([]metadataStruct, cols)
	}

	//Initialize leaf set
	l := int(math.Pow(2, float64(b)))
	stateTable.leafSetLower = make([]metadataStruct, l/2)
	stateTable.leafSetGreater = make([]metadataStruct, l/2)
	fmt.Println("Initilized Leaf set with size:", l)

	fmt.Printf("fullAddress %s nodeId %s\n", fullAddress, nodeId)
	fmt.Println("GRR: ", len(nodeId))

	listener, err := net.Listen("tcp", fullAddress)
	if err != nil {
		panic(err)
	}

	metadata.nodeId = nodeId
	metadata.ipAddress = fullAddress

	if len(argsWithoutProg) == 3 {
		connectToNode(argsWithoutProg[2])
		//os.Exit(1)
	}

	//Startup server
	http.HandleFunc("/join", join)

	http.Serve(listener, nil)
}
