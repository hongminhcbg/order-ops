#!/bin/bash
cat input.txt | parallel 'ab -n 20000 -c 30 {}'