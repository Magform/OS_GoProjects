#!/bin/bash

echo "Project 1"

echo "Test 1"

go run ./Project1/tests/FristTest/AI_multi-run.go | tail -n 1
go run ./Project1/tests/FristTest/MY_multi-run.go | tail -n 1
go run ./Project1/tests/FristTest/MY2_multi-run.go | tail -n 1
go run ./Project1/tests/FristTest/MY3_multi-run.go | tail -n 1
go run ./Project1/tests/FristTest/MY4_multi-run.go | tail -n 1

echo "Test 2"

go run ./Project1/tests/SecondTest/AI_multi-run.go | tail -n 1
go run ./Project1/tests/SecondTest/MY_multi-run.go | tail -n 1
go run ./Project1/tests/SecondTest/MY2_multi-run.go | tail -n 1
go run ./Project1/tests/SecondTest/MY3_multi-run.go | tail -n 1
go run ./Project1/tests/SecondTest/MY4_multi-run.go | tail -n 1

echo "Test 3"

go run ./Project1/tests/ThirdTest/AI_multi-run.go | tail -n 1
go run ./Project1/tests/ThirdTest/MY_multi-run.go | tail -n 1
go run ./Project1/tests/ThirdTest/MY2_multi-run.go | tail -n 1
go run ./Project1/tests/ThirdTest/MY3_multi-run.go | tail -n 1
go run ./Project1/tests/ThirdTest/MY4_multi-run.go | tail -n 1

echo "Project 2"

echo "Test 1"

go run ./Project2/tests/FristTest/AI_multi-run.go | tail -n 1
go run ./Project2/tests/FristTest/MY_multi-run.go | tail -n 1
go run ./Project2/tests/FristTest/MY2_multi-run.go | tail -n 1

echo "Test 2"

go run ./Project2/tests/SecondTest/AI_multi-run.go | tail -n 1
go run ./Project2/tests/SecondTest/MY_multi-run.go | tail -n 1
go run ./Project2/tests/SecondTest/MY2_multi-run.go | tail -n 1

echo "Test 3"

go run ./Project2/tests/ThirdTest/AI_multi-run.go | tail -n 1
go run ./Project2/tests/ThirdTest/MY_multi-run.go | tail -n 1
go run ./Project2/tests/ThirdTest/MY2_multi-run.go | tail -n 1

echo "Project 3"

echo "Test 1"

go run ./Project3/tests/FristTest/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/tests/FristTest/MY_multi-run.go | tail -n 1
go run ./Project3/tests/FristTest/MYV_multi-run.go | tail -n 1

echo "Test 1 V2"

go run ./Project3/tests/FristTest_V2/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/tests/FristTest_V2/MY_multi-run.go | tail -n 1
go run ./Project3/tests/FristTest_V2/MYV_multi-run.go | tail -n 1

echo "Test 2 "

go run ./Project3/tests/SecondTest/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/tests/SecondTest/MY_multi-run.go | tail -n 1
go run ./Project3/tests/SecondTest/MYV_multi-run.go | tail -n 1

echo "Test 2 V2"

go run ./Project3/tests/SecondTest_V2/AI-HumanFixed_multi-run.go | tail -n 1
go run ./Project3/tests/SecondTest_V2/MY_multi-run.go | tail -n 1
go run ./Project3/tests/SecondTest_V2/MYV_multi-run.go | tail -n 1