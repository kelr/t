#!/bin/bash

platforms=("linux/amd64" "linux/386" "windows/amd64" "windows/386" "darwin/amd64" "darwin/386")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    
    output_name=tl'_'$GOOS'_'$GOARCH
    tar_name=${output_name}'.tar.gz'
    mkdir ${output_name}

    env GOOS=$GOOS GOARCH=$GOARCH go build
    if [ $? -ne 0 ]; then
        echo 'Error in building'
        exit 1
    fi

    if [ $GOOS = "windows" ]; then
        mv tl.exe ${output_name}
    else
    	mv tl ${output_name}
    fi

    tar -czvf ${tar_name} ${output_name}
    rm -rf ${output_name}
done
mkdir release
mv *.tar.gz release