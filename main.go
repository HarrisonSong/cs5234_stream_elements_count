package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	N := flag.Int("N", 100000, "total number of elements in stream.")
	M := flag.Int("M", 5000, "range of elements in stream. limited in 2^32")
	B := flag.Int("B", 100, "range of hashing")
	A := flag.Int("A", 100, "number of hash functions")
	streamType := flag.String("stream", "u", "type of the stream")
	alg := flag.Int("alg", 1, "algorithm to be run")
	algoRepeatTime := flag.Int("repeat", 10, "algorithm repeat time")
	totalTrials := flag.Int("times", 10000, "total trial times")
	errorRate := flag.Float64("erate", 0.25, "max acceptable error rate")
	experimentType := flag.String("experiment", "space", "max acceptable error rate")
	flag.Parse()

	if *experimentType == "space" {
		fixSpaceExperiments(*N/10, *N, *M, *alg, *algoRepeatTime, *totalTrials, *errorRate, *streamType)
	} else if *experimentType == "A" {
		fixAExperiments(*N, *M, *A, *alg, *algoRepeatTime, *totalTrials, *errorRate, *streamType)
	} else {
		fixBExperiments(*N, *M, *B, *alg, *algoRepeatTime, *totalTrials, *errorRate, *streamType)
	}
}

func fixAExperiments(N, M, A, alg, algoRepeatTime, totalTrials int, errorRate float64, streamType string) {
	writeToFile("results/stream_"+streamType+"_alg_"+strconv.Itoa(alg)+"_A_"+strconv.Itoa(A)+".csv", fmt.Sprintln(",A,B,accuracy"))
	for space := N / 10; space <= N/2; space += 1000 {
		accuracy := experiment(N, M, space/A, A, alg, algoRepeatTime, totalTrials, errorRate, streamType)
		writeToFile("results/stream_"+streamType+"_alg_"+strconv.Itoa(alg)+"_A_"+strconv.Itoa(A)+".csv", fmt.Sprintf(",%d,%d,%f\n", A, space/A, accuracy))
	}
}

func fixBExperiments(N, M, B, alg, algoRepeatTime, totalTrials int, errorRate float64, streamType string) {
	writeToFile("results/stream_"+streamType+"_alg_"+strconv.Itoa(alg)+"_B_"+strconv.Itoa(B)+".csv", fmt.Sprintln(",A,B,accuracy"))
	for space := N / 10; space <= N/2; space += 1000 {
		accuracy := experiment(N, M, B, space/B, alg, algoRepeatTime, totalTrials, errorRate, streamType)
		writeToFile("results/stream_"+streamType+"_alg_"+strconv.Itoa(alg)+"_B_"+strconv.Itoa(B)+".csv", fmt.Sprintf(",%d,%d,%f\n", space/B, B, accuracy))
	}
}

func fixSpaceExperiments(space, N, M, alg, algoRepeatTime, totalTrials int, errorRate float64, streamType string) {
	writeToFile("results/stream_"+streamType+"_alg_"+strconv.Itoa(alg)+"_space_"+strconv.Itoa(space)+".csv", fmt.Sprintln(",A,B,accuracy"))
	for B := M / 10; B <= M/2; B += 50 {
		accuracy := experiment(N, M, B, space/B, alg, algoRepeatTime, totalTrials, errorRate, streamType)
		writeToFile("results/stream_"+streamType+"_alg_"+strconv.Itoa(alg)+"_space_"+strconv.Itoa(space)+".csv", fmt.Sprintf(",%d,%d,%f\n", space/B, B, accuracy))
	}
}

func experiment(N, M, B, A, alg, algoRepeatTime, totalTrials int, errorRate float64, streamType string) float64 {
	sum := 0.0
	for times := 0; times < algoRepeatTime; times++ {
		sum += RunAlgo(N, M, B, A, alg, algoRepeatTime, totalTrials, errorRate, streamType)
	}

	return sum / float64(algoRepeatTime)
}

func writeToFile(fileName string, content string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		log.Println(err)
	}
}
