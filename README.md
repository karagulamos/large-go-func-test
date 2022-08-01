# **Large Go Functions on Netlify**

An easy way to generate and test large Go functions on Netlify. Just follow the instructions and link your Netlify site to see the magic happen.

## **Generating Sizes**

To generate large functions, we use a `Makefile` and a custom bash script (`rand_replace_mb.sh`) which takes the following paramaters:

* Size of random text to generate (in `megabytes`)
* Path to the function file
* Placeholder in your function code to replace with the generated text

Note that the final size of the uploaded function may vary due to bundling/compression performed by Netlify (or the CLI) at build time. So, if you wish to play with the functions sizes, you can fork this repo and update the size of the random text generated in the `Makefile`.

Format:

``` bash
    bash scripts/rand_replace_mb.sh <rand_text_size_in_megabytes> <path_to_function_file> <placeholder_text>
```

If you've linked your site and committed changes to your repo, Netlify will run the build script to generate the function before uploading.

To test this locally:

* Ensure you have the latest version of the Netlify CLI installed
* Run the `netlify build` command at the root of the project and you should be able to see random text replace the placeholder in the function.
* You can use the `du` command to check the function size like so:
  * `du -sh .netlify/functions/hello-world` - uncompressed
  * `du -sh .netlify/functions/hello-world.zip` - compressed

## **Final Note**

Although this was created specifically with Go in mind, it should work for JavaScript functions in a manner similar to the Go example in project.
