package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	
	"../pkg/entryModel"
	"../pkg/handlerFunc"
	)

func main() {

		args := os.Args[1:]
		
		if len(args)>0 {
			//open input file
			
			f,err := os.Open(args[0])
			handlerFunc.Check(err)			
			scanner := bufio.NewScanner(f)
			
			
			//create output file
			outputFile, err := os.Create("./output.txt")
			handlerFunc.Check(err)
			w := bufio.NewWriter(outputFile)
			
			//map to check repeating transaction id
			idCustMap := make(map[string]map[string]int)
			
			for scanner.Scan(){
				data := entryModel.Entry{}
				
				if err := json.Unmarshal(scanner.Bytes(), &data); err != nil {
					panic(err)
				}
				
				if _, ok := idCustMap[data.CustomerID]; ok {
							if _, ok := idCustMap[data.CustomerID][data.ID]; ok {
										continue
							} else {
									idCustMap[data.CustomerID][data.ID] = 1
							}
				} else {
					idCustMap[data.CustomerID] = make(map[string]int)
					idCustMap[data.CustomerID][data.ID] = 1
				}

				//fmt.Println(handlerFunc.CheckIfValidTrans(data))
				
				outputMsg := handlerFunc.CheckIfValidTrans(data) + "\n"
				_, err := w.WriteString(outputMsg)
				handlerFunc.Check(err)
			}
			if scanner.Err() != nil {
				panic(scanner.Err)
			}
			w.Flush()
			f.Close()
		} else{
			fmt.Println("Please add filename as argument")
		}
}