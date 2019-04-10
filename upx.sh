#!/bin/bash
find ./dist -xdev -maxdepth 3 -type f -name 'AriaNg-go*' | xargs "./upx"
