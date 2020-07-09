package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sortedList []string = make([]string, 0) //list of all words ordered by usage

func handleRequests() {
	http.HandleFunc("/autocomplete", autoComplete)
	fmt.Println("Service enabled at 'http://localhost:9000/autocomplete' , waiting for requests...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
func autoComplete(w http.ResponseWriter, r *http.Request) {
	//Setting headers for json response
	w.Header().Set("Content-Type", "application/json")
	term := cleanWord(r.URL.Query()["term"][0])
	if len(term) < 1 {
		fmt.Fprintf(w, "not a valid word")
		return
	}
	if _, err := strconv.ParseInt(term, 10, 64); err == nil {
		fmt.Fprintf(w, "not a valid word")
		return
	}

	found := make([]string, 0)
	for _, word := range sortedList {
		if strings.Contains(word, term) {
			found = append(found, word)
		}
		if len(found) >= 25 { //limit to 25
			break
		}
	}
	json.NewEncoder(w).Encode(found)
}

//cleans words for more consistent
func cleanWord(input string) string {
	return strings.Trim(strings.ToLower(input), " ,'.!?\"`/*()[]{}-_:;")
}

//Gets every word, finds the number of instances of them, then sorts words from most used -> least used.
//Data could be exported but it runs quick enough that it's not necessary
func populateDictionary() {
	var dictionary = make(map[string]int)

	file, err := os.Open("./shakespeare-complete.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		cleanedWord := cleanWord(scanner.Text()) //clean each individual word of unnecessary punctuation

		if _, err := strconv.ParseInt(cleanedWord, 10, 64); err == nil { //checks if word is a number
			continue
		}
		_, ok := dictionary[cleanedWord]
		if ok {
			dictionary[cleanedWord]++
		} else {
			dictionary[cleanedWord] = 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k := range dictionary {
		sortedList = append(sortedList, k)
	}

	sort.Slice(sortedList, func(i, j int) bool {
		return dictionary[sortedList[i]] > dictionary[sortedList[j]]
	})
}

func main() {
	fmt.Println("Populating Dictionary...")
	populateDictionary()
	fmt.Println("Dictionary populated; enabling the endpoint...")
	handleRequests()
}
