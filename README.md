# Node.js impostor for ProgpJS

## Introduction

This project allows to use ProgpJS as an impostor for Node.js.
It allows you to build an executable, which once renamed "node" allows to replace Node.js by ProgpJS.

## Supported command line arguments

* Param **-v** and **--version**, return node.js version. Here it returns version 50.50 (welcome to the futur!)
* Param **--inspect**, allows to start ProgpJS debugger. It's compatible with Node.js v8 protocol. To start debugging open url chrome://inspect/#devices inside Chrome browser.
* Param **--debug** is the same.
* Param **--inspect-brk** is the same.

## Behaviors

* **node** will use the script named index.js in the current directory. If not found, it will automatically use **index.ts**.
