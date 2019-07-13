#! /usr/bin/env python

from livereload import Server

server = Server()
server.watch("helloworld.html")
server.serve(root='./helloworld.html')
