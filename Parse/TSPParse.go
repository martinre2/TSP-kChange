package Parse

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const TOKEN_NAME string = "INAME:"
const TOKEN_TYPE string = "TYPE:"
const TOKEN_COMMENT string = "COMMENT:"
const TOKEN_DIMENSION string = "DIMENSION:"
const TOKEN_BEST_KNOWN string = "BEST_KNOWN:"
const TOKEN_EDGE_WEIGHT_TYPE string = "EDGE_WEIGHT_TYPE:"
const TOKEN_EDGE_WEIGHT_FORMAT string = "EDGE_WEIGHT_FORMAT:"
const TOKEN_EDGE_WEIGHT_SECTION string = "EDGE_WEIGHT_SECTION"
const TOKEN_EOF string = "EOF"

type Ftv struct {
	Name             string
	Type             string
	Comment          string
	Dimension        string
	Size             int
	BestKnow         int
	EdgeWeightType   string
	EdgeWeightFormat string
	Weights          [][]float32
}

type TSPParser struct {
	Data Ftv
}

func (p *TSPParser) ParseFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, TOKEN_NAME) {
			p.Data.Name = strings.TrimSpace(strings.Replace(line, TOKEN_NAME, "", -1))
		} else if strings.HasPrefix(line, TOKEN_TYPE) {
			p.Data.Type = strings.TrimSpace(strings.Replace(line, TOKEN_TYPE, "", -1))
		} else if strings.HasPrefix(line, TOKEN_COMMENT) {
			p.Data.Name = strings.TrimSpace(strings.Replace(line, TOKEN_COMMENT, "", -1))
		} else if strings.HasPrefix(line, TOKEN_DIMENSION) {
			p.Data.Dimension = strings.TrimSpace(strings.Replace(line, TOKEN_DIMENSION, "", -1))
			p.Data.Size, _ = strconv.Atoi(p.Data.Dimension)
		} else if strings.HasPrefix(line, TOKEN_BEST_KNOWN) {
			bK := strings.TrimSpace(strings.Replace(line, TOKEN_BEST_KNOWN, "", -1))
			p.Data.BestKnow, _ = strconv.Atoi(bK)
		} else if strings.HasPrefix(line, TOKEN_EDGE_WEIGHT_TYPE) {
			p.Data.EdgeWeightType = strings.TrimSpace(strings.Replace(line, TOKEN_EDGE_WEIGHT_TYPE, "", -1))
		} else if strings.HasPrefix(line, TOKEN_EDGE_WEIGHT_FORMAT) {
			p.Data.EdgeWeightFormat = strings.TrimSpace(strings.Replace(line, TOKEN_EDGE_WEIGHT_FORMAT, "", -1))
		} else if strings.HasPrefix(line, TOKEN_EDGE_WEIGHT_SECTION) {

			p.Data.Weights = make([][]float32, p.Data.Size)
			p.Data.Weights[0] = make([]float32, p.Data.Size)
			var f, c int
			f, c = 0, 0

			for scanner.Scan() {
				inner_line := scanner.Text()
				if strings.HasPrefix(inner_line, TOKEN_EOF) {
					break
				} else {
					weights := strings.Fields(inner_line)
					for _, w := range weights {
						if f == c {
							p.Data.Weights[f][c] = 0.0
						} else {
							pF, _ := strconv.ParseFloat(w, 64)
							p.Data.Weights[f][c] = float32(pF)
						}

						if c == p.Data.Size-1 {
							f++
							if f < p.Data.Size {
								p.Data.Weights[f] = make([]float32, p.Data.Size)
								c = 0
							}
						} else {
							c++
						}

					}

				}

			}

		} else {
			fmt.Println(" ")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
