# SparkRegistry Extism Client Plugin

This project is an Extism plugin written in Go to interact with Sulu's registry of L402 APIs. 

## Prerequisites

Before using this plugin, ensure you have the following setup:

- Extism CLI and runtime installed and configured on your system. Visit [Extism's official documentation](https://docs.extism.com/) for installation and setup instructions.
- TinyGo compiler for compiling Go to WebAssembly. Installation instructions can be found at [TinyGo's Getting Started guide](https://tinygo.org/getting-started/install/).

## Installation

1. **Clone the Plugin Repository**

    Clone this repository to get the plugin source code:

    ```sh
    git clone https://your-repository-url-here
    cd sparkreg-extism-plugin
    ```

    Replace `https://your-repository-url-here` with the actual URL of your repository.

2. **Compile the Plugin**

    Use the TinyGo compiler to build the WebAssembly module from the Go source code:

    ```sh
    tinygo build -o sparkreg_plugin.wasm -target wasi main.go
    ```

    This command compiles the Go code in `main.go` to a WebAssembly module named `sparkreg_plugin.wasm`.

## Usage

To use the plugin with the Extism CLI, execute the following command:

```sh
extism call sparkreg_plugin.wasm httpGetSparkRegList --wasi --allow-host='sparkreg.com'
