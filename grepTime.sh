#!/bin/bash

echo "Project 1"

echo "Test 1"

go run ./Project1/multi_run/FristTest/AI_multi-run.go | tail -n 1
go run ./Project1/multi_run/FristTest/MY_multi-run.go | tail -n 1
go run ./Project1/multi_run/FristTest/MY2_multi-run.go | tail -n 1
go run ./Project1/multi_run/FristTest/MY3_multi-run.go | tail -n 1
go run ./Project1/multi_run/FristTest/MY4_multi-run.go | tail -n 1

echo "Test 2"

go run ./Project1/multi_run/SecondTest/AI_multi-run.go | tail -n 1
go run ./Project1/multi_run/SecondTest/MY_multi-run.go | tail -n 1
go run ./Project1/multi_run/SecondTest/MY2_multi-run.go | tail -n 1
go run ./Project1/multi_run/SecondTest/MY3_multi-run.go | tail -n 1
go run ./Project1/multi_run/SecondTest/MY4_multi-run.go | tail -n 1

echo "Test 3"

go run ./Project1/multi_run/ThirdTest/AI_multi-run.go | tail -n 1
go run ./Project1/multi_run/ThirdTest/MY_multi-run.go | tail -n 1
go run ./Project1/multi_run/ThirdTest/MY2_multi-run.go | tail -n 1
go run ./Project1/multi_run/ThirdTest/MY3_multi-run.go | tail -n 1
go run ./Project1/multi_run/ThirdTest/MY4_multi-run.go | tail -n 1

echo "Project 2"

echo "Test 1"

go run ./Project2/multi_run/FristTest/AI_multi-run.go | tail -n 1
go run ./Project2/multi_run/FristTest/MY_multi-run.go | tail -n 1
go run ./Project2/multi_run/FristTest/MY2_multi-run.go | tail -n 1

echo "Test 2"

go run ./Project2/multi_run/SecondTest/AI_multi-run.go | tail -n 1
go run ./Project2/multi_run/SecondTest/MY_multi-run.go | tail -n 1
go run ./Project2/multi_run/SecondTest/MY2_multi-run.go | tail -n 1

echo "Test 3"

go run ./Project2/multi_run/ThirdTest/AI_multi-run.go | tail -n 1
go run ./Project2/multi_run/ThirdTest/MY_multi-run.go | tail -n 1
go run ./Project2/multi_run/ThirdTest/MY2_multi-run.go | tail -n 1

echo "Project 3"

echo "Test 1"

go run ./Project3/multi_run/FristTest/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/multi_run/FristTest/MY_multi-run.go | tail -n 1

echo "Test 1 V2"

go run ./Project3/multi_run/FristTest_V2/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/multi_run/FristTest_V2/MY_multi-run.go | tail -n 1

echo "Test 2 "

go run ./Project3/multi_run/SecondTest/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/multi_run/SecondTest/MY_multi-run.go | tail -n 1

echo "Test 2 V2"

go run ./Project3/multi_run/SecondTest_V2/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/multi_run/SecondTest_V2/MY_multi-run.go | tail -n 1