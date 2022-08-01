declare -i max_input_size=250 # max request size in megabytes
declare -i bytes_per_megabyte=1048576 # 1 megabyte 
declare -i output_size=$1*bytes_per_megabyte

if  [[ $# -lt 3 || -z "$3" ]]; then # if the user didn't provide enough arguments or if the placeholder is empty
    echo "Usage: $0 <rand_text_size_in_megabytes> <path_to_function_file> <placeholder_text>"
    exit 1
fi

if [ $1 -gt $max_input_size ]; then
    echo "Input size is too large. Max input size is $max_input_size megabytes."
    exit 1
fi

if [ ! -f $2 ]; then
    echo "Input file does not exist."
    exit 1
fi

if [ $(uname) = Darwin ]; then
sed -i '' -f /dev/stdin $2; else sed -i -f /dev/stdin $2 
fi << EOF 
s/$3/$(base64 /dev/urandom | od -A n -v -t x1 | tr -d ' \n' | head -c $output_size)/g
EOF
