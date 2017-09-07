#!/bin/bash

rename 's/[^0-3][^0-9]\.sh/-'$(date +%Y%m%d)'.sh/' *.sh