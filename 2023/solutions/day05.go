package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aoc2023/utils"
)

func Day5() {

  start := time.Now()

  lines, err := utils.ReadLines("input/day05")
  if err != nil {
    fmt.Printf("error reading file %v\n", err)
  }
  _ = lines

  part1 := lowestLocation(lines)
  part2 := lowestLocation2(lines)

  elapsed := time.Since(start)
  fmt.Printf("Day 05 part1 %v\t", part1)
  fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type stuff struct {
  id      int
  number  int
  mapped  bool
}

func lowestLocation(input []string) int {

  var seeds []stuff
  var soils []stuff
  var fertilizers []stuff
  var waters []stuff
  var lights []stuff
  var temperatures []stuff
  var humiditys []stuff
  var locations []stuff

  firstLine := strings.Fields(input[0])
  for i := 1; i < len(firstLine); i++ {
    num, _ := strconv.Atoi(firstLine[i])
    seeds = append(seeds, stuff{id: i, number: num, mapped: false})
  }

  mapCount := 0

  for i := 2; i < len(input); i++ {
    if len(input[i]) > 0 {
      
      line := strings.Fields(input[i])
      if string(line[1]) == "map:" {
        mapCount++
      } else {
        switch mapCount {
        case 1:
          sourceToDestination(&soils, seeds, line)
        case 2:
          finishMapping(&soils, seeds)
          sourceToDestination(&fertilizers, soils, line)
        case 3:
          finishMapping(&fertilizers, soils)
          sourceToDestination(&waters, fertilizers, line)
        case 4:
          finishMapping(&waters, fertilizers)
          sourceToDestination(&lights, waters, line)
        case 5:
          finishMapping(&lights, waters)
          sourceToDestination(&temperatures, lights, line)
        case 6:
          finishMapping(&temperatures, lights)
          sourceToDestination(&humiditys, temperatures, line)
        case 7:
          finishMapping(&humiditys, temperatures)
          sourceToDestination(&locations, humiditys, line)
        }
      }
    }
  }
  finishMapping(&locations, humiditys)

  return getMin(locations)
}

func sourceToDestination(destinations *[]stuff, sources []stuff, line []string) {
  destinationMin, _ := strconv.Atoi(line[0])
  sourceMin, _ := strconv.Atoi(line[1])
  length, _ := strconv.Atoi(line[2])

  for i := range sources {
    if sources[i].number >= sourceMin && sources[i].number < sourceMin + length {
      num := destinationMin + sources[i].number - sourceMin
      *destinations = append(*destinations, stuff{id: sources[i].id, number: num, mapped: false})
      sources[i].mapped = true
    }
  }
}

func finishMapping(destinations *[]stuff, sources []stuff) {
  for i := range sources {
    if !sources[i].mapped {
      sources[i].mapped = true
      num := sources[i].number
      *destinations = append(*destinations, stuff{id: sources[i].id, number: num, mapped: false})
    }
  }
}

func getMin (stuffs []stuff) int {
  min := 100000000000000000

  for _, stuff := range stuffs {
    if min > stuff.number {
      min = stuff.number
    }
  }

  return min
}

type stuff2 struct {
  start   int
  length  int
  source  int
}

func lowestLocation2(input []string) int {

  var seeds []stuff2
  var soilsMap []stuff2
  var fertiMap []stuff2
  var waterMap []stuff2
  var lightMap []stuff2
  var tempMap []stuff2
  var humiMap []stuff2
  var locationMap []stuff2
  
  firstLine := strings.Fields(input[0])
  for i := 1; i < len(firstLine); i += 2 {
    num, _ := strconv.Atoi(firstLine[i])
    length, _ := strconv.Atoi(firstLine[i+1])
    seeds = append(seeds, stuff2{ start: num, length: length })
  }

  whatMap := ""
  _ = whatMap

  for i := 2; i < len(input); i++ {
    if len(input[i]) > 0 {
      line := strings.Fields(input[i])

      if len(strings.Split(line[0], "-")) > 1 {
        whatMap = strings.Split(line[0], "-")[2] 

      } else {
        switch whatMap {
        case "soil":
          fillMap(line, &soilsMap)

        case "fertilizer":
          fillMap(line, &fertiMap)

        case "water":
          fillMap(line, &waterMap)

        case "light":
          fillMap(line, &lightMap)

        case "temperature":
          fillMap(line, &tempMap)

        case "humidity":
          fillMap(line, &humiMap)

        case "location":
          fillMap(line, &locationMap)
        }
      }
    }
  }

  location := 0
  for true {
    humidity := destinationToSource(location, locationMap)
    temperature := destinationToSource(humidity, humiMap)
    light := destinationToSource(temperature, tempMap)
    water := destinationToSource(light, lightMap)
    fertilizer := destinationToSource(water, waterMap)
    soil := destinationToSource(fertilizer, fertiMap)
    seed := destinationToSource(soil, soilsMap)

    if checkSeed(seed, seeds) {
      break
    }

    location++
  }

  return location
}

func fillMap(line []string, stuff *[]stuff2) {
  num, _ := strconv.Atoi(line[0])
  source, _ := strconv.Atoi(line[1])
  length, _ := strconv.Atoi(line[2])

  *stuff = append(*stuff, stuff2{start: num, source: source, length: length})
}

func destinationToSource(destination int, sourceMap []stuff2) int {
  for _, m := range sourceMap {
    delta := m.source - m.start
    if destination >= m.start && destination < m.start + m.length {
      return destination + delta
    }
  }
  return destination
}

func checkSeed(seed int, seeds []stuff2) bool {
  for _, s := range seeds {
    if seed >= s.start && seed < s.start + s.length {
      return true
    }
  }
  return false
}
